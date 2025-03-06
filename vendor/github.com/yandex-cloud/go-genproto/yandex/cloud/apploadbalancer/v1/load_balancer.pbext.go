// Code generated by protoc-gen-goext. DO NOT EDIT.

package apploadbalancer

import (
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (m *LoadBalancer) SetId(v string) {
	m.Id = v
}

func (m *LoadBalancer) SetName(v string) {
	m.Name = v
}

func (m *LoadBalancer) SetDescription(v string) {
	m.Description = v
}

func (m *LoadBalancer) SetFolderId(v string) {
	m.FolderId = v
}

func (m *LoadBalancer) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *LoadBalancer) SetStatus(v LoadBalancer_Status) {
	m.Status = v
}

func (m *LoadBalancer) SetRegionId(v string) {
	m.RegionId = v
}

func (m *LoadBalancer) SetNetworkId(v string) {
	m.NetworkId = v
}

func (m *LoadBalancer) SetListeners(v []*Listener) {
	m.Listeners = v
}

func (m *LoadBalancer) SetAllocationPolicy(v *AllocationPolicy) {
	m.AllocationPolicy = v
}

func (m *LoadBalancer) SetLogGroupId(v string) {
	m.LogGroupId = v
}

func (m *LoadBalancer) SetSecurityGroupIds(v []string) {
	m.SecurityGroupIds = v
}

func (m *LoadBalancer) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}

func (m *LoadBalancer) SetAutoScalePolicy(v *AutoScalePolicy) {
	m.AutoScalePolicy = v
}

func (m *LoadBalancer) SetLogOptions(v *LogOptions) {
	m.LogOptions = v
}

type Address_Address = isAddress_Address

func (m *Address) SetAddress(v Address_Address) {
	m.Address = v
}

func (m *Address) SetExternalIpv4Address(v *ExternalIpv4Address) {
	m.Address = &Address_ExternalIpv4Address{
		ExternalIpv4Address: v,
	}
}

func (m *Address) SetInternalIpv4Address(v *InternalIpv4Address) {
	m.Address = &Address_InternalIpv4Address{
		InternalIpv4Address: v,
	}
}

func (m *Address) SetExternalIpv6Address(v *ExternalIpv6Address) {
	m.Address = &Address_ExternalIpv6Address{
		ExternalIpv6Address: v,
	}
}

func (m *ExternalIpv4Address) SetAddress(v string) {
	m.Address = v
}

func (m *InternalIpv4Address) SetAddress(v string) {
	m.Address = v
}

func (m *InternalIpv4Address) SetSubnetId(v string) {
	m.SubnetId = v
}

func (m *ExternalIpv6Address) SetAddress(v string) {
	m.Address = v
}

func (m *Location) SetZoneId(v string) {
	m.ZoneId = v
}

func (m *Location) SetSubnetId(v string) {
	m.SubnetId = v
}

func (m *Location) SetDisableTraffic(v bool) {
	m.DisableTraffic = v
}

func (m *AllocationPolicy) SetLocations(v []*Location) {
	m.Locations = v
}

type Listener_Listener = isListener_Listener

func (m *Listener) SetListener(v Listener_Listener) {
	m.Listener = v
}

func (m *Listener) SetName(v string) {
	m.Name = v
}

func (m *Listener) SetEndpoints(v []*Endpoint) {
	m.Endpoints = v
}

func (m *Listener) SetHttp(v *HttpListener) {
	m.Listener = &Listener_Http{
		Http: v,
	}
}

func (m *Listener) SetTls(v *TlsListener) {
	m.Listener = &Listener_Tls{
		Tls: v,
	}
}

func (m *Listener) SetStream(v *StreamListener) {
	m.Listener = &Listener_Stream{
		Stream: v,
	}
}

func (m *Endpoint) SetAddresses(v []*Address) {
	m.Addresses = v
}

func (m *Endpoint) SetPorts(v []int64) {
	m.Ports = v
}

func (m *HttpListener) SetHandler(v *HttpHandler) {
	m.Handler = v
}

func (m *HttpListener) SetRedirects(v *Redirects) {
	m.Redirects = v
}

func (m *TlsListener) SetDefaultHandler(v *TlsHandler) {
	m.DefaultHandler = v
}

func (m *TlsListener) SetSniHandlers(v []*SniMatch) {
	m.SniHandlers = v
}

func (m *StreamListener) SetHandler(v *StreamHandler) {
	m.Handler = v
}

func (m *Http2Options) SetMaxConcurrentStreams(v int64) {
	m.MaxConcurrentStreams = v
}

func (m *StreamHandler) SetBackendGroupId(v string) {
	m.BackendGroupId = v
}

func (m *StreamHandler) SetIdleTimeout(v *durationpb.Duration) {
	m.IdleTimeout = v
}

type HttpHandler_ProtocolSettings = isHttpHandler_ProtocolSettings

func (m *HttpHandler) SetProtocolSettings(v HttpHandler_ProtocolSettings) {
	m.ProtocolSettings = v
}

func (m *HttpHandler) SetHttpRouterId(v string) {
	m.HttpRouterId = v
}

func (m *HttpHandler) SetHttp2Options(v *Http2Options) {
	m.ProtocolSettings = &HttpHandler_Http2Options{
		Http2Options: v,
	}
}

func (m *HttpHandler) SetAllowHttp10(v bool) {
	m.ProtocolSettings = &HttpHandler_AllowHttp10{
		AllowHttp10: v,
	}
}

func (m *HttpHandler) SetRewriteRequestId(v bool) {
	m.RewriteRequestId = v
}

func (m *Redirects) SetHttpToHttps(v bool) {
	m.HttpToHttps = v
}

func (m *SniMatch) SetName(v string) {
	m.Name = v
}

func (m *SniMatch) SetServerNames(v []string) {
	m.ServerNames = v
}

func (m *SniMatch) SetHandler(v *TlsHandler) {
	m.Handler = v
}

type TlsHandler_Handler = isTlsHandler_Handler

func (m *TlsHandler) SetHandler(v TlsHandler_Handler) {
	m.Handler = v
}

func (m *TlsHandler) SetHttpHandler(v *HttpHandler) {
	m.Handler = &TlsHandler_HttpHandler{
		HttpHandler: v,
	}
}

func (m *TlsHandler) SetStreamHandler(v *StreamHandler) {
	m.Handler = &TlsHandler_StreamHandler{
		StreamHandler: v,
	}
}

func (m *TlsHandler) SetCertificateIds(v []string) {
	m.CertificateIds = v
}

func (m *TargetState) SetStatus(v *TargetState_HealthcheckStatus) {
	m.Status = v
}

func (m *TargetState) SetTarget(v *Target) {
	m.Target = v
}

func (m *TargetState_HealthcheckStatus) SetZoneStatuses(v []*TargetState_ZoneHealthcheckStatus) {
	m.ZoneStatuses = v
}

func (m *TargetState_ZoneHealthcheckStatus) SetZoneId(v string) {
	m.ZoneId = v
}

func (m *TargetState_ZoneHealthcheckStatus) SetStatus(v TargetState_Status) {
	m.Status = v
}

func (m *TargetState_ZoneHealthcheckStatus) SetFailedActiveHc(v bool) {
	m.FailedActiveHc = v
}

func (m *AutoScalePolicy) SetMinZoneSize(v int64) {
	m.MinZoneSize = v
}

func (m *AutoScalePolicy) SetMaxSize(v int64) {
	m.MaxSize = v
}
