// Code generated by protoc-gen-goext. DO NOT EDIT.

package k8s_marketplace

import (
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (m *HelmRelease) SetId(v string) {
	m.Id = v
}

func (m *HelmRelease) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *HelmRelease) SetAppName(v string) {
	m.AppName = v
}

func (m *HelmRelease) SetAppNamespace(v string) {
	m.AppNamespace = v
}

func (m *HelmRelease) SetProductId(v string) {
	m.ProductId = v
}

func (m *HelmRelease) SetProductName(v string) {
	m.ProductName = v
}

func (m *HelmRelease) SetProductVersion(v string) {
	m.ProductVersion = v
}

func (m *HelmRelease) SetStatus(v HelmRelease_Status) {
	m.Status = v
}

func (m *HelmRelease) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}
