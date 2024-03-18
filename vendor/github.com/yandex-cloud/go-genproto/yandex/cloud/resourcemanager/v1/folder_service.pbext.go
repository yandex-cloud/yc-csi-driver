// Code generated by protoc-gen-goext. DO NOT EDIT.

package resourcemanager

import (
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (m *GetFolderRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *ListFoldersRequest) SetCloudId(v string) {
	m.CloudId = v
}

func (m *ListFoldersRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListFoldersRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListFoldersRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListFoldersResponse) SetFolders(v []*Folder) {
	m.Folders = v
}

func (m *ListFoldersResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *CreateFolderRequest) SetCloudId(v string) {
	m.CloudId = v
}

func (m *CreateFolderRequest) SetName(v string) {
	m.Name = v
}

func (m *CreateFolderRequest) SetDescription(v string) {
	m.Description = v
}

func (m *CreateFolderRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *CreateFolderMetadata) SetFolderId(v string) {
	m.FolderId = v
}

func (m *UpdateFolderRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *UpdateFolderRequest) SetUpdateMask(v *fieldmaskpb.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateFolderRequest) SetName(v string) {
	m.Name = v
}

func (m *UpdateFolderRequest) SetDescription(v string) {
	m.Description = v
}

func (m *UpdateFolderRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *UpdateFolderMetadata) SetFolderId(v string) {
	m.FolderId = v
}

func (m *DeleteFolderRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *DeleteFolderRequest) SetDeleteAfter(v *timestamppb.Timestamp) {
	m.DeleteAfter = v
}

func (m *DeleteFolderMetadata) SetFolderId(v string) {
	m.FolderId = v
}

func (m *DeleteFolderMetadata) SetDeleteAfter(v *timestamppb.Timestamp) {
	m.DeleteAfter = v
}

func (m *DeleteFolderMetadata) SetCancelledBy(v string) {
	m.CancelledBy = v
}

func (m *DeleteFolderMetadata) SetCancelledAt(v *timestamppb.Timestamp) {
	m.CancelledAt = v
}

func (m *ListFolderOperationsRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *ListFolderOperationsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListFolderOperationsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListFolderOperationsResponse) SetOperations(v []*operation.Operation) {
	m.Operations = v
}

func (m *ListFolderOperationsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}