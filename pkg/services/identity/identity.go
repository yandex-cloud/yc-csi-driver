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

package identity

import (
	"context"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"k8s.io/klog"

	"github.com/yandex-cloud/yc-csi-driver/pkg/services"
	"github.com/yandex-cloud/yc-csi-driver/pkg/version"
)

type Identity struct {
	DriverName   string
	Capabilities []*csi.PluginCapability
}

func (i *Identity) GetPluginInfo(ctx context.Context, req *csi.GetPluginInfoRequest) (*csi.GetPluginInfoResponse, error) {
	name := services.DriverName

	if i.DriverName != "" {
		name = i.DriverName
	}

	return &csi.GetPluginInfoResponse{
		Name:          name,
		VendorVersion: version.GitRevision(),
	}, nil
}

func (i *Identity) GetPluginCapabilities(ctx context.Context, req *csi.GetPluginCapabilitiesRequest) (*csi.GetPluginCapabilitiesResponse, error) {
	klog.Infof("GetPluginCapabilities")
	return &csi.GetPluginCapabilitiesResponse{Capabilities: i.Capabilities}, nil
}

func (*Identity) Probe(context.Context, *csi.ProbeRequest) (*csi.ProbeResponse, error) {
	return &csi.ProbeResponse{
		Ready: &wrapperspb.BoolValue{
			Value: true,
		},
	}, nil
}
