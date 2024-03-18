/*
Copyright 2024 YANDEX LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package sanitytests

import (
	"os"
	"testing"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-csi/csi-test/v4/pkg/sanity"
	"github.com/onsi/ginkgo"
	ginkgo_config "github.com/onsi/ginkgo/config"
	"github.com/onsi/ginkgo/reporters"
	"github.com/onsi/gomega"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"k8s.io/kubernetes/pkg/volume/util/hostutil"

	diskapi "github.com/yandex-cloud/yc-csi-driver/pkg/diskapi/fake"
	"github.com/yandex-cloud/yc-csi-driver/pkg/inflight"
	"github.com/yandex-cloud/yc-csi-driver/pkg/server"
	"github.com/yandex-cloud/yc-csi-driver/pkg/services"
	"github.com/yandex-cloud/yc-csi-driver/pkg/services/controller"
	"github.com/yandex-cloud/yc-csi-driver/pkg/services/identity"
	"github.com/yandex-cloud/yc-csi-driver/pkg/services/node"
	"github.com/yandex-cloud/yc-csi-driver/pkg/util"
)

type fakeFileXattrs struct {
}

func (f fakeFileXattrs) SetImmutable(target string) error {
	return nil
}

func (f fakeFileXattrs) UnsetImmutable(target string) error {
	return nil
}

type nodeOptions struct {
	MaxVolumesPerNode int
}

type options struct {
	nodeOptions
	Endpoint   string
	DriverName string
}

func defaultOptions(endpoint string) options {
	return options{
		Endpoint:    endpoint,
		DriverName:  "yc.disk.test",
		nodeOptions: nodeOptions{MaxVolumesPerNode: 5},
	}
}

func newServer(opts options) (*server.Server, error) {
	scheme, addr, err := util.ParseEndpoint(opts.Endpoint)
	if err != nil {
		return nil, err
	}

	srv := grpc.NewServer()
	csi.RegisterIdentityServer(srv, &identity.Identity{
		DriverName:   services.DriverName,
		Capabilities: services.SanityPluginCapabilities,
	})

	nodeSvc := node.NewTestNodeService(
		mounter,
		hostutil.NewFakeHostUtil(nil),
		fakeFileXattrs{},
		node.NewFakeMetadataGetter(diskapi.FakeInstID, diskapi.FakeZoneID),
		services.SanityNodeCaps,
	)

	controllerSvc := controller.New(
		diskapi.NewFakeDiskAPI("test", "test"),
		inflight.NewWithTTL(),
		services.SanityControllerCaps,
	)
	csi.RegisterNodeServer(srv, nodeSvc)
	csi.RegisterControllerServer(srv, controllerSvc)

	return &server.Server{Srv: srv, Scheme: scheme, Addr: addr}, nil
}

func createConfig(tmpDir, srvAddr string) sanity.TestConfig {
	config := sanity.NewTestConfig()
	config.Address = srvAddr
	config.TargetPath = tmpDir + "/csi-mount"
	config.StagingPath = tmpDir + "/csi-staging"
	return config
}

func TestCSIDriver(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "sanity")
	defer os.RemoveAll(tmpDir)
	require.NoError(t, err)
	opts := defaultOptions("unix://" + tmpDir + "/csi.sock")
	srv, err := newServer(opts)
	require.NoError(t, err)
	go func() {
		if err := srv.Run(); err != nil {
			require.NoError(t, err)
		}
	}()
	defer srv.Srv.Stop()

	ginkgo_config.DefaultReporterConfig.Verbose = true
	test(t, createConfig(tmpDir, srv.Addr))
}

func test(t ginkgo.GinkgoTestingT, config sanity.TestConfig) {
	sc := sanity.GinkgoTest(&config)
	gomega.RegisterFailHandler(ginkgo.Fail)
	teamcityReporter := reporters.NewTeamCityReporter(ginkgo.GinkgoWriter)
	teamcityReporter.ReporterConfig.Verbose = true
	ginkgo.RunSpecsWithCustomReporters(t, "CSI Driver Test Suite", []ginkgo.Reporter{teamcityReporter})
	sc.Finalize()
}
