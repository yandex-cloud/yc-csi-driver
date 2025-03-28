// Code generated by protoc-gen-goext. DO NOT EDIT.

package cdn

import (
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

func (m *GetResourceRequest) SetResourceId(v string) {
	m.ResourceId = v
}

func (m *ListResourcesRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *ListResourcesRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListResourcesRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListResourcesResponse) SetResources(v []*Resource) {
	m.Resources = v
}

func (m *ListResourcesResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *CreateResourceRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *CreateResourceRequest) SetCname(v string) {
	m.Cname = v
}

func (m *CreateResourceRequest) SetOrigin(v *CreateResourceRequest_Origin) {
	m.Origin = v
}

func (m *CreateResourceRequest) SetSecondaryHostnames(v *SecondaryHostnames) {
	m.SecondaryHostnames = v
}

func (m *CreateResourceRequest) SetOriginProtocol(v OriginProtocol) {
	m.OriginProtocol = v
}

func (m *CreateResourceRequest) SetActive(v *wrapperspb.BoolValue) {
	m.Active = v
}

func (m *CreateResourceRequest) SetOptions(v *ResourceOptions) {
	m.Options = v
}

func (m *CreateResourceRequest) SetSslCertificate(v *SSLTargetCertificate) {
	m.SslCertificate = v
}

func (m *CreateResourceRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

type CreateResourceRequest_Origin_OriginVariant = isCreateResourceRequest_Origin_OriginVariant

func (m *CreateResourceRequest_Origin) SetOriginVariant(v CreateResourceRequest_Origin_OriginVariant) {
	m.OriginVariant = v
}

func (m *CreateResourceRequest_Origin) SetOriginGroupId(v int64) {
	m.OriginVariant = &CreateResourceRequest_Origin_OriginGroupId{
		OriginGroupId: v,
	}
}

func (m *CreateResourceRequest_Origin) SetOriginSource(v string) {
	m.OriginVariant = &CreateResourceRequest_Origin_OriginSource{
		OriginSource: v,
	}
}

func (m *CreateResourceRequest_Origin) SetOriginSourceParams(v *ResourceOriginParams) {
	m.OriginVariant = &CreateResourceRequest_Origin_OriginSourceParams{
		OriginSourceParams: v,
	}
}

func (m *ResourceOriginParams) SetSource(v string) {
	m.Source = v
}

func (m *ResourceOriginParams) SetMeta(v *OriginMeta) {
	m.Meta = v
}

func (m *CreateResourceMetadata) SetResourceId(v string) {
	m.ResourceId = v
}

func (m *UpdateResourceRequest) SetResourceId(v string) {
	m.ResourceId = v
}

func (m *UpdateResourceRequest) SetOriginGroupId(v *wrapperspb.Int64Value) {
	m.OriginGroupId = v
}

func (m *UpdateResourceRequest) SetSecondaryHostnames(v *SecondaryHostnames) {
	m.SecondaryHostnames = v
}

func (m *UpdateResourceRequest) SetOptions(v *ResourceOptions) {
	m.Options = v
}

func (m *UpdateResourceRequest) SetOriginProtocol(v OriginProtocol) {
	m.OriginProtocol = v
}

func (m *UpdateResourceRequest) SetActive(v *wrapperspb.BoolValue) {
	m.Active = v
}

func (m *UpdateResourceRequest) SetSslCertificate(v *SSLTargetCertificate) {
	m.SslCertificate = v
}

func (m *UpdateResourceRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *UpdateResourceRequest) SetRemoveLabels(v bool) {
	m.RemoveLabels = v
}

func (m *UpdateResourceMetadata) SetResourceId(v string) {
	m.ResourceId = v
}

func (m *DeleteResourceRequest) SetResourceId(v string) {
	m.ResourceId = v
}

func (m *DeleteResourceMetadata) SetResourceId(v string) {
	m.ResourceId = v
}

func (m *GetProviderCNameRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *GetProviderCNameResponse) SetCname(v string) {
	m.Cname = v
}

func (m *GetProviderCNameResponse) SetFolderId(v string) {
	m.FolderId = v
}
