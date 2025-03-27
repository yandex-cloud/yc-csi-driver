// Code generated by protoc-gen-goext. DO NOT EDIT.

package greenplum

import (
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (m *Backup) SetId(v string) {
	m.Id = v
}

func (m *Backup) SetFolderId(v string) {
	m.FolderId = v
}

func (m *Backup) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}

func (m *Backup) SetSourceClusterId(v string) {
	m.SourceClusterId = v
}

func (m *Backup) SetStartedAt(v *timestamppb.Timestamp) {
	m.StartedAt = v
}

func (m *Backup) SetSize(v int64) {
	m.Size = v
}

func (m *Backup) SetType(v Backup_BackupCreationType) {
	m.Type = v
}

func (m *Backup) SetMethod(v Backup_BackupMethod) {
	m.Method = v
}

func (m *Backup) SetJournalSize(v int64) {
	m.JournalSize = v
}
