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
	"fmt"

	"github.com/yandex-cloud/yc-csi-driver/pkg/diskapi"
	"github.com/yandex-cloud/yc-csi-driver/pkg/inflight"
	"github.com/yandex-cloud/yc-csi-driver/pkg/server"
	"github.com/yandex-cloud/yc-csi-driver/pkg/services"
	"github.com/yandex-cloud/yc-csi-driver/pkg/services/controller"
	"github.com/yandex-cloud/yc-csi-driver/pkg/services/identity"
	"github.com/yandex-cloud/yc-csi-driver/pkg/util"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"google.golang.org/grpc"
)

type Options struct {
	Endpoint     string
	DriverName   string
	DiskAPI      diskapi.DiskAPI
	Capabilities []*csi.ControllerServiceCapability
}

func New(opts Options) (*server.Server, error) {
	scheme, addr, err := util.ParseEndpoint(opts.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("endpoint parse: %w", err)
	}

	srv := grpc.NewServer()

	csi.RegisterIdentityServer(srv, &identity.Identity{
		DriverName:   opts.DriverName,
		Capabilities: services.ProductionPluginCapabilities,
	})

	csi.RegisterControllerServer(srv,
		controller.New(
			opts.DiskAPI,
			inflight.NewWithTTL(),
			opts.Capabilities,
		),
	)

	return &server.Server{Srv: srv, Scheme: scheme, Addr: addr}, nil
}
