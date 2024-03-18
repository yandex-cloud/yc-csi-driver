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

package node

import (
	"github.com/container-storage-interface/spec/lib/go/csi"
	"google.golang.org/grpc"
	"k8s.io/kubernetes/pkg/volume/util/hostutil"

	"github.com/yandex-cloud/yc-csi-driver/pkg/instancemeta"
	"github.com/yandex-cloud/yc-csi-driver/pkg/server"
	"github.com/yandex-cloud/yc-csi-driver/pkg/services"
	"github.com/yandex-cloud/yc-csi-driver/pkg/services/identity"
	"github.com/yandex-cloud/yc-csi-driver/pkg/services/node"
	"github.com/yandex-cloud/yc-csi-driver/pkg/util"
)

type Options struct {
	node.Options
	Endpoint   string
	DriverName string
}

func DefaultOptions() Options {
	return Options{
		Endpoint:   "unix://tmp/csi.sock",
		DriverName: "",
		Options:    node.DefaultOptions(),
	}
}

func New(opts Options) (*server.Server, error) {
	scheme, addr, err := util.ParseEndpoint(opts.Endpoint)
	if err != nil {
		return nil, err
	}

	srv := grpc.NewServer()

	csi.RegisterIdentityServer(srv, &identity.Identity{
		DriverName:   opts.DriverName,
		Capabilities: services.ProductionPluginCapabilities,
	})

	csi.RegisterNodeServer(srv, node.New(
		node.NewSafeMounter(),
		hostutil.NewHostUtil(),
		new(VolumeStats),
		new(FileXattrs),
		instancemeta.New(),
		services.ProductionNodeCaps,
		opts.Options,
	))

	return &server.Server{Srv: srv, Scheme: scheme, Addr: addr}, nil
}
