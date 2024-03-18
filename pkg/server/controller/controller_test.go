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

package controller

import (
	"context"
	"net"
	"os"
	"sync"
	"testing"

	diskapi "github.com/yandex-cloud/yc-csi-driver/pkg/diskapi/fake"
	"github.com/yandex-cloud/yc-csi-driver/pkg/services"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-csi/csi-test/v4/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestController(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "csi-controller-test")
	defer os.RemoveAll(tmpDir)
	require.NoError(t, err)

	pubSrv, err := New(
		Options{
			Endpoint:     "unix://" + tmpDir + "/csi-pub.sock",
			DriverName:   "yc.disk.test",
			DiskAPI:      diskapi.NewFakeDiskAPI("test", "test"),
			Capabilities: services.ProductionCapsForPublicAPIController,
		})

	require.NoError(t, err)

	pubListener, err := net.Listen(pubSrv.Scheme, pubSrv.Addr)
	require.NoError(t, err)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		require.NoError(t, pubSrv.Srv.Serve(pubListener))
		defer wg.Done()
	}()

	defer func() {
		pubSrv.Srv.Stop()
		wg.Wait()
	}()

	testCaps(t, pubSrv.Addr, services.ProductionCapsForPublicAPIController)
}

func testCaps(t *testing.T, addr string, caps []*csi.ControllerServiceCapability) {
	conn, err := utils.Connect(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	require.NoError(t, err)
	require.NotNil(t, conn)
	cl := csi.NewControllerClient(conn)
	i := csi.NewIdentityClient(conn)
	for _, c := range caps {
		assert.NotEqual(t, csi.ControllerServiceCapability_RPC_UNKNOWN, c.GetRpc().GetType())
		assert.True(t, isControllerCapabilitySupported(t, cl, c.GetRpc().GetType()))
	}
	for _, c := range services.ProductionPluginCapabilities {
		if c.GetService().GetType() == csi.PluginCapability_Service_UNKNOWN {
			continue
		}
		assert.True(t, isPluginCapabilitySupported(t, i, c.GetService().GetType()))
	}
	for _, c := range services.ProductionPluginCapabilities {
		if c.GetVolumeExpansion().GetType() == csi.PluginCapability_VolumeExpansion_UNKNOWN {
			continue
		}
		assert.True(t, isPluginVolumeExpansionSupported(t, i, c.GetVolumeExpansion().GetType()))
	}
}

func isControllerCapabilitySupported(
	t *testing.T,
	c csi.ControllerClient,
	capType csi.ControllerServiceCapability_RPC_Type,
) bool {
	caps, err := c.ControllerGetCapabilities(
		context.Background(),
		&csi.ControllerGetCapabilitiesRequest{})
	require.NoError(t, err)
	for _, c := range caps.GetCapabilities() {
		require.NotNil(t, c.GetRpc())
		if c.GetRpc().GetType() == capType {
			return true
		}
	}
	return false
}

func isPluginCapabilitySupported(
	t *testing.T,
	i csi.IdentityClient,
	capType csi.PluginCapability_Service_Type,
) bool {
	caps, err := i.GetPluginCapabilities(
		context.Background(),
		&csi.GetPluginCapabilitiesRequest{})
	require.NoError(t, err)
	for _, c := range caps.GetCapabilities() {
		require.NotNil(t, c.GetType())
		if c.GetService().GetType() == capType {
			return true
		}
	}
	return false
}

func isPluginVolumeExpansionSupported(
	t *testing.T,
	i csi.IdentityClient,
	capType csi.PluginCapability_VolumeExpansion_Type,
) bool {
	caps, err := i.GetPluginCapabilities(
		context.Background(),
		&csi.GetPluginCapabilitiesRequest{})
	require.NoError(t, err)
	for _, c := range caps.GetCapabilities() {
		require.NotNil(t, c.GetType())
		if c.GetVolumeExpansion().GetType() == capType {
			return true
		}
	}
	return false
}
