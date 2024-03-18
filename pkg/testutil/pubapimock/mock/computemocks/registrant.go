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

package computemocks

import (
	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1"
)

func (_m *DiskServiceServer) GRPCTestRegister(srv *grpc.Server) {
	compute.RegisterDiskServiceServer(srv, _m)
}

func (_m *InstanceServiceServer) GRPCTestRegister(srv *grpc.Server) {
	compute.RegisterInstanceServiceServer(srv, _m)
}

func (_m *ZoneServiceServer) GRPCTestRegister(srv *grpc.Server) {
	compute.RegisterZoneServiceServer(srv, _m)
}
