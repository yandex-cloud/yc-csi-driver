// Code generated by protoc-gen-goext. DO NOT EDIT.

package greenplum

import (
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

func (m *ResourceGroup) SetName(v string) {
	m.Name = v
}

func (m *ResourceGroup) SetIsUserDefined(v *wrapperspb.BoolValue) {
	m.IsUserDefined = v
}

func (m *ResourceGroup) SetConcurrency(v *wrapperspb.Int64Value) {
	m.Concurrency = v
}

func (m *ResourceGroup) SetCpuRateLimit(v *wrapperspb.Int64Value) {
	m.CpuRateLimit = v
}

func (m *ResourceGroup) SetMemoryLimit(v *wrapperspb.Int64Value) {
	m.MemoryLimit = v
}

func (m *ResourceGroup) SetMemorySharedQuota(v *wrapperspb.Int64Value) {
	m.MemorySharedQuota = v
}

func (m *ResourceGroup) SetMemorySpillRatio(v *wrapperspb.Int64Value) {
	m.MemorySpillRatio = v
}
