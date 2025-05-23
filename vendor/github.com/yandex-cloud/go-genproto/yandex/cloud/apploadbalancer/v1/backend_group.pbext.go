// Code generated by protoc-gen-goext. DO NOT EDIT.

package apploadbalancer

import (
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

type BackendGroup_Backend = isBackendGroup_Backend

func (m *BackendGroup) SetBackend(v BackendGroup_Backend) {
	m.Backend = v
}

func (m *BackendGroup) SetId(v string) {
	m.Id = v
}

func (m *BackendGroup) SetName(v string) {
	m.Name = v
}

func (m *BackendGroup) SetDescription(v string) {
	m.Description = v
}

func (m *BackendGroup) SetFolderId(v string) {
	m.FolderId = v
}

func (m *BackendGroup) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *BackendGroup) SetHttp(v *HttpBackendGroup) {
	m.Backend = &BackendGroup_Http{
		Http: v,
	}
}

func (m *BackendGroup) SetGrpc(v *GrpcBackendGroup) {
	m.Backend = &BackendGroup_Grpc{
		Grpc: v,
	}
}

func (m *BackendGroup) SetStream(v *StreamBackendGroup) {
	m.Backend = &BackendGroup_Stream{
		Stream: v,
	}
}

func (m *BackendGroup) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}

type StreamBackendGroup_SessionAffinity = isStreamBackendGroup_SessionAffinity

func (m *StreamBackendGroup) SetSessionAffinity(v StreamBackendGroup_SessionAffinity) {
	m.SessionAffinity = v
}

func (m *StreamBackendGroup) SetBackends(v []*StreamBackend) {
	m.Backends = v
}

func (m *StreamBackendGroup) SetConnection(v *ConnectionSessionAffinity) {
	m.SessionAffinity = &StreamBackendGroup_Connection{
		Connection: v,
	}
}

type HttpBackendGroup_SessionAffinity = isHttpBackendGroup_SessionAffinity

func (m *HttpBackendGroup) SetSessionAffinity(v HttpBackendGroup_SessionAffinity) {
	m.SessionAffinity = v
}

func (m *HttpBackendGroup) SetBackends(v []*HttpBackend) {
	m.Backends = v
}

func (m *HttpBackendGroup) SetConnection(v *ConnectionSessionAffinity) {
	m.SessionAffinity = &HttpBackendGroup_Connection{
		Connection: v,
	}
}

func (m *HttpBackendGroup) SetHeader(v *HeaderSessionAffinity) {
	m.SessionAffinity = &HttpBackendGroup_Header{
		Header: v,
	}
}

func (m *HttpBackendGroup) SetCookie(v *CookieSessionAffinity) {
	m.SessionAffinity = &HttpBackendGroup_Cookie{
		Cookie: v,
	}
}

type GrpcBackendGroup_SessionAffinity = isGrpcBackendGroup_SessionAffinity

func (m *GrpcBackendGroup) SetSessionAffinity(v GrpcBackendGroup_SessionAffinity) {
	m.SessionAffinity = v
}

func (m *GrpcBackendGroup) SetBackends(v []*GrpcBackend) {
	m.Backends = v
}

func (m *GrpcBackendGroup) SetConnection(v *ConnectionSessionAffinity) {
	m.SessionAffinity = &GrpcBackendGroup_Connection{
		Connection: v,
	}
}

func (m *GrpcBackendGroup) SetHeader(v *HeaderSessionAffinity) {
	m.SessionAffinity = &GrpcBackendGroup_Header{
		Header: v,
	}
}

func (m *GrpcBackendGroup) SetCookie(v *CookieSessionAffinity) {
	m.SessionAffinity = &GrpcBackendGroup_Cookie{
		Cookie: v,
	}
}

func (m *HeaderSessionAffinity) SetHeaderName(v string) {
	m.HeaderName = v
}

func (m *CookieSessionAffinity) SetName(v string) {
	m.Name = v
}

func (m *CookieSessionAffinity) SetTtl(v *durationpb.Duration) {
	m.Ttl = v
}

func (m *ConnectionSessionAffinity) SetSourceIp(v bool) {
	m.SourceIp = v
}

func (m *LoadBalancingConfig) SetPanicThreshold(v int64) {
	m.PanicThreshold = v
}

func (m *LoadBalancingConfig) SetLocalityAwareRoutingPercent(v int64) {
	m.LocalityAwareRoutingPercent = v
}

func (m *LoadBalancingConfig) SetStrictLocality(v bool) {
	m.StrictLocality = v
}

func (m *LoadBalancingConfig) SetMode(v LoadBalancingMode) {
	m.Mode = v
}

type StreamBackend_BackendType = isStreamBackend_BackendType

func (m *StreamBackend) SetBackendType(v StreamBackend_BackendType) {
	m.BackendType = v
}

func (m *StreamBackend) SetName(v string) {
	m.Name = v
}

func (m *StreamBackend) SetBackendWeight(v *wrapperspb.Int64Value) {
	m.BackendWeight = v
}

func (m *StreamBackend) SetLoadBalancingConfig(v *LoadBalancingConfig) {
	m.LoadBalancingConfig = v
}

func (m *StreamBackend) SetPort(v int64) {
	m.Port = v
}

func (m *StreamBackend) SetTargetGroups(v *TargetGroupsBackend) {
	m.BackendType = &StreamBackend_TargetGroups{
		TargetGroups: v,
	}
}

func (m *StreamBackend) SetHealthchecks(v []*HealthCheck) {
	m.Healthchecks = v
}

func (m *StreamBackend) SetTls(v *BackendTls) {
	m.Tls = v
}

func (m *StreamBackend) SetEnableProxyProtocol(v bool) {
	m.EnableProxyProtocol = v
}

func (m *StreamBackend) SetKeepConnectionsOnHostHealthFailure(v bool) {
	m.KeepConnectionsOnHostHealthFailure = v
}

type HttpBackend_BackendType = isHttpBackend_BackendType

func (m *HttpBackend) SetBackendType(v HttpBackend_BackendType) {
	m.BackendType = v
}

func (m *HttpBackend) SetName(v string) {
	m.Name = v
}

func (m *HttpBackend) SetBackendWeight(v *wrapperspb.Int64Value) {
	m.BackendWeight = v
}

func (m *HttpBackend) SetLoadBalancingConfig(v *LoadBalancingConfig) {
	m.LoadBalancingConfig = v
}

func (m *HttpBackend) SetPort(v int64) {
	m.Port = v
}

func (m *HttpBackend) SetTargetGroups(v *TargetGroupsBackend) {
	m.BackendType = &HttpBackend_TargetGroups{
		TargetGroups: v,
	}
}

func (m *HttpBackend) SetStorageBucket(v *StorageBucketBackend) {
	m.BackendType = &HttpBackend_StorageBucket{
		StorageBucket: v,
	}
}

func (m *HttpBackend) SetHealthchecks(v []*HealthCheck) {
	m.Healthchecks = v
}

func (m *HttpBackend) SetTls(v *BackendTls) {
	m.Tls = v
}

func (m *HttpBackend) SetUseHttp2(v bool) {
	m.UseHttp2 = v
}

type GrpcBackend_BackendType = isGrpcBackend_BackendType

func (m *GrpcBackend) SetBackendType(v GrpcBackend_BackendType) {
	m.BackendType = v
}

func (m *GrpcBackend) SetName(v string) {
	m.Name = v
}

func (m *GrpcBackend) SetBackendWeight(v *wrapperspb.Int64Value) {
	m.BackendWeight = v
}

func (m *GrpcBackend) SetLoadBalancingConfig(v *LoadBalancingConfig) {
	m.LoadBalancingConfig = v
}

func (m *GrpcBackend) SetPort(v int64) {
	m.Port = v
}

func (m *GrpcBackend) SetTargetGroups(v *TargetGroupsBackend) {
	m.BackendType = &GrpcBackend_TargetGroups{
		TargetGroups: v,
	}
}

func (m *GrpcBackend) SetHealthchecks(v []*HealthCheck) {
	m.Healthchecks = v
}

func (m *GrpcBackend) SetTls(v *BackendTls) {
	m.Tls = v
}

func (m *TargetGroupsBackend) SetTargetGroupIds(v []string) {
	m.TargetGroupIds = v
}

func (m *SecureTransportSettings) SetSni(v string) {
	m.Sni = v
}

func (m *SecureTransportSettings) SetValidationContext(v *ValidationContext) {
	m.ValidationContext = v
}

func (m *BackendTls) SetSni(v string) {
	m.Sni = v
}

func (m *BackendTls) SetValidationContext(v *ValidationContext) {
	m.ValidationContext = v
}

func (m *StorageBucketBackend) SetBucket(v string) {
	m.Bucket = v
}

type HealthCheck_Healthcheck = isHealthCheck_Healthcheck

func (m *HealthCheck) SetHealthcheck(v HealthCheck_Healthcheck) {
	m.Healthcheck = v
}

type HealthCheck_TransportSettings = isHealthCheck_TransportSettings

func (m *HealthCheck) SetTransportSettings(v HealthCheck_TransportSettings) {
	m.TransportSettings = v
}

func (m *HealthCheck) SetTimeout(v *durationpb.Duration) {
	m.Timeout = v
}

func (m *HealthCheck) SetInterval(v *durationpb.Duration) {
	m.Interval = v
}

func (m *HealthCheck) SetIntervalJitterPercent(v float64) {
	m.IntervalJitterPercent = v
}

func (m *HealthCheck) SetHealthyThreshold(v int64) {
	m.HealthyThreshold = v
}

func (m *HealthCheck) SetUnhealthyThreshold(v int64) {
	m.UnhealthyThreshold = v
}

func (m *HealthCheck) SetHealthcheckPort(v int64) {
	m.HealthcheckPort = v
}

func (m *HealthCheck) SetStream(v *HealthCheck_StreamHealthCheck) {
	m.Healthcheck = &HealthCheck_Stream{
		Stream: v,
	}
}

func (m *HealthCheck) SetHttp(v *HealthCheck_HttpHealthCheck) {
	m.Healthcheck = &HealthCheck_Http{
		Http: v,
	}
}

func (m *HealthCheck) SetGrpc(v *HealthCheck_GrpcHealthCheck) {
	m.Healthcheck = &HealthCheck_Grpc{
		Grpc: v,
	}
}

func (m *HealthCheck) SetPlaintext(v *PlaintextTransportSettings) {
	m.TransportSettings = &HealthCheck_Plaintext{
		Plaintext: v,
	}
}

func (m *HealthCheck) SetTls(v *SecureTransportSettings) {
	m.TransportSettings = &HealthCheck_Tls{
		Tls: v,
	}
}

func (m *HealthCheck_StreamHealthCheck) SetSend(v *Payload) {
	m.Send = v
}

func (m *HealthCheck_StreamHealthCheck) SetReceive(v *Payload) {
	m.Receive = v
}

func (m *HealthCheck_HttpHealthCheck) SetHost(v string) {
	m.Host = v
}

func (m *HealthCheck_HttpHealthCheck) SetPath(v string) {
	m.Path = v
}

func (m *HealthCheck_HttpHealthCheck) SetUseHttp2(v bool) {
	m.UseHttp2 = v
}

func (m *HealthCheck_HttpHealthCheck) SetExpectedStatuses(v []int64) {
	m.ExpectedStatuses = v
}

func (m *HealthCheck_GrpcHealthCheck) SetServiceName(v string) {
	m.ServiceName = v
}
