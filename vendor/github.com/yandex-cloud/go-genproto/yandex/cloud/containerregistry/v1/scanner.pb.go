// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/containerregistry/v1/scanner.proto

package containerregistry

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ScanResult_Status int32

const (
	ScanResult_STATUS_UNSPECIFIED ScanResult_Status = 0
	// Image scan is in progress.
	ScanResult_RUNNING ScanResult_Status = 1
	// Image has been scanned and result is ready.
	ScanResult_READY ScanResult_Status = 2
	// Image scan is failed.
	ScanResult_ERROR ScanResult_Status = 3
)

// Enum value maps for ScanResult_Status.
var (
	ScanResult_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "RUNNING",
		2: "READY",
		3: "ERROR",
	}
	ScanResult_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"RUNNING":            1,
		"READY":              2,
		"ERROR":              3,
	}
)

func (x ScanResult_Status) Enum() *ScanResult_Status {
	p := new(ScanResult_Status)
	*p = x
	return p
}

func (x ScanResult_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ScanResult_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_containerregistry_v1_scanner_proto_enumTypes[0].Descriptor()
}

func (ScanResult_Status) Type() protoreflect.EnumType {
	return &file_yandex_cloud_containerregistry_v1_scanner_proto_enumTypes[0]
}

func (x ScanResult_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ScanResult_Status.Descriptor instead.
func (ScanResult_Status) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_containerregistry_v1_scanner_proto_rawDescGZIP(), []int{0, 0}
}

type Vulnerability_Severity int32

const (
	Vulnerability_SEVERITY_UNSPECIFIED Vulnerability_Severity = 0
	// Critical severity is a world-burning problem, exploitable for nearly all users.
	// Includes remote root privilege escalations, or massive data loss.
	Vulnerability_CRITICAL Vulnerability_Severity = 1
	// High severity is a real problem, exploitable for many users in a default installation.
	// Includes serious remote denial of services, local root privilege escalations, or data loss.
	Vulnerability_HIGH Vulnerability_Severity = 2
	// Medium severity is a real security problem, and is exploitable for many users.
	// Includes network daemon denial of service attacks, cross-site scripting, and gaining user privileges.
	// Updates should be made soon for this priority of issue.
	Vulnerability_MEDIUM Vulnerability_Severity = 3
	// Low severity is a security problem, but is hard to exploit due to environment, requires a user-assisted attack,
	// a small install base, or does very little damage. These tend to be included in security updates only when
	// higher priority issues require an update, or if many low priority issues have built up.
	Vulnerability_LOW Vulnerability_Severity = 4
	// Negligible severity is technically a security problem, but is only theoretical in nature, requires a very special situation,
	// has almost no install base, or does no real damage. These tend not to get backport from upstream,
	// and will likely not be included in security updates unless there is an easy fix and some other issue causes an update.
	Vulnerability_NEGLIGIBLE Vulnerability_Severity = 5
	// Unknown severity is either a security problem that has not been assigned to a priority yet or
	// a priority that our system did not recognize.
	Vulnerability_UNDEFINED Vulnerability_Severity = 6
)

// Enum value maps for Vulnerability_Severity.
var (
	Vulnerability_Severity_name = map[int32]string{
		0: "SEVERITY_UNSPECIFIED",
		1: "CRITICAL",
		2: "HIGH",
		3: "MEDIUM",
		4: "LOW",
		5: "NEGLIGIBLE",
		6: "UNDEFINED",
	}
	Vulnerability_Severity_value = map[string]int32{
		"SEVERITY_UNSPECIFIED": 0,
		"CRITICAL":             1,
		"HIGH":                 2,
		"MEDIUM":               3,
		"LOW":                  4,
		"NEGLIGIBLE":           5,
		"UNDEFINED":            6,
	}
)

func (x Vulnerability_Severity) Enum() *Vulnerability_Severity {
	p := new(Vulnerability_Severity)
	*p = x
	return p
}

func (x Vulnerability_Severity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Vulnerability_Severity) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_containerregistry_v1_scanner_proto_enumTypes[1].Descriptor()
}

func (Vulnerability_Severity) Type() protoreflect.EnumType {
	return &file_yandex_cloud_containerregistry_v1_scanner_proto_enumTypes[1]
}

func (x Vulnerability_Severity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Vulnerability_Severity.Descriptor instead.
func (Vulnerability_Severity) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_containerregistry_v1_scanner_proto_rawDescGZIP(), []int{2, 0}
}

// A ScanResult resource.
type ScanResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. ID of the ScanResult.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Output only. ID of the Image that the ScanResult belongs to.
	ImageId string `protobuf:"bytes,2,opt,name=image_id,json=imageId,proto3" json:"image_id,omitempty"`
	// Output only. The timestamp in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format when the scan been finished.
	ScannedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=scanned_at,json=scannedAt,proto3" json:"scanned_at,omitempty"`
	// Output only. The status of the ScanResult.
	Status ScanResult_Status `protobuf:"varint,4,opt,name=status,proto3,enum=yandex.cloud.containerregistry.v1.ScanResult_Status" json:"status,omitempty"`
	// Output only. Summary information about vulnerabilities found.
	Vulnerabilities *VulnerabilityStats `protobuf:"bytes,5,opt,name=vulnerabilities,proto3" json:"vulnerabilities,omitempty"`
}

func (x *ScanResult) Reset() {
	*x = ScanResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_containerregistry_v1_scanner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScanResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScanResult) ProtoMessage() {}

func (x *ScanResult) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_containerregistry_v1_scanner_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScanResult.ProtoReflect.Descriptor instead.
func (*ScanResult) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_containerregistry_v1_scanner_proto_rawDescGZIP(), []int{0}
}

func (x *ScanResult) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ScanResult) GetImageId() string {
	if x != nil {
		return x.ImageId
	}
	return ""
}

func (x *ScanResult) GetScannedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ScannedAt
	}
	return nil
}

func (x *ScanResult) GetStatus() ScanResult_Status {
	if x != nil {
		return x.Status
	}
	return ScanResult_STATUS_UNSPECIFIED
}

func (x *ScanResult) GetVulnerabilities() *VulnerabilityStats {
	if x != nil {
		return x.Vulnerabilities
	}
	return nil
}

// A VulnerabilityStats resource.
type VulnerabilityStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Count of CRITICAL vulnerabilities.
	Critical int64 `protobuf:"varint,1,opt,name=critical,proto3" json:"critical,omitempty"`
	// Count of HIGH vulnerabilities.
	High int64 `protobuf:"varint,2,opt,name=high,proto3" json:"high,omitempty"`
	// Count of MEDIUM vulnerabilities.
	Medium int64 `protobuf:"varint,3,opt,name=medium,proto3" json:"medium,omitempty"`
	// Count of LOW vulnerabilities.
	Low int64 `protobuf:"varint,4,opt,name=low,proto3" json:"low,omitempty"`
	// Count of NEGLIGIBLE vulnerabilities.
	Negligible int64 `protobuf:"varint,5,opt,name=negligible,proto3" json:"negligible,omitempty"`
	// Count of other vulnerabilities.
	Undefined int64 `protobuf:"varint,6,opt,name=undefined,proto3" json:"undefined,omitempty"`
}

func (x *VulnerabilityStats) Reset() {
	*x = VulnerabilityStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_containerregistry_v1_scanner_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VulnerabilityStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VulnerabilityStats) ProtoMessage() {}

func (x *VulnerabilityStats) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_containerregistry_v1_scanner_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VulnerabilityStats.ProtoReflect.Descriptor instead.
func (*VulnerabilityStats) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_containerregistry_v1_scanner_proto_rawDescGZIP(), []int{1}
}

func (x *VulnerabilityStats) GetCritical() int64 {
	if x != nil {
		return x.Critical
	}
	return 0
}

func (x *VulnerabilityStats) GetHigh() int64 {
	if x != nil {
		return x.High
	}
	return 0
}

func (x *VulnerabilityStats) GetMedium() int64 {
	if x != nil {
		return x.Medium
	}
	return 0
}

func (x *VulnerabilityStats) GetLow() int64 {
	if x != nil {
		return x.Low
	}
	return 0
}

func (x *VulnerabilityStats) GetNegligible() int64 {
	if x != nil {
		return x.Negligible
	}
	return 0
}

func (x *VulnerabilityStats) GetUndefined() int64 {
	if x != nil {
		return x.Undefined
	}
	return 0
}

// A Vulnerability resource.
type Vulnerability struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. Severity of the Vulnerability.
	Severity Vulnerability_Severity `protobuf:"varint,1,opt,name=severity,proto3,enum=yandex.cloud.containerregistry.v1.Vulnerability_Severity" json:"severity,omitempty"`
	// Details of vulnerability depending on type. Only `package` vulnerability is supported at the moment.
	//
	// Types that are assignable to Vulnerability:
	//
	//	*Vulnerability_Package
	Vulnerability isVulnerability_Vulnerability `protobuf_oneof:"vulnerability"`
}

func (x *Vulnerability) Reset() {
	*x = Vulnerability{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_containerregistry_v1_scanner_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vulnerability) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vulnerability) ProtoMessage() {}

func (x *Vulnerability) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_containerregistry_v1_scanner_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vulnerability.ProtoReflect.Descriptor instead.
func (*Vulnerability) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_containerregistry_v1_scanner_proto_rawDescGZIP(), []int{2}
}

func (x *Vulnerability) GetSeverity() Vulnerability_Severity {
	if x != nil {
		return x.Severity
	}
	return Vulnerability_SEVERITY_UNSPECIFIED
}

func (m *Vulnerability) GetVulnerability() isVulnerability_Vulnerability {
	if m != nil {
		return m.Vulnerability
	}
	return nil
}

func (x *Vulnerability) GetPackage() *PackageVulnerability {
	if x, ok := x.GetVulnerability().(*Vulnerability_Package); ok {
		return x.Package
	}
	return nil
}

type isVulnerability_Vulnerability interface {
	isVulnerability_Vulnerability()
}

type Vulnerability_Package struct {
	Package *PackageVulnerability `protobuf:"bytes,2,opt,name=package,proto3,oneof"`
}

func (*Vulnerability_Package) isVulnerability_Vulnerability() {}

// A PackageVulnerability resource.
type PackageVulnerability struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of vulnerability in CVE database.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// URL to the page with description of vulnerability.
	Link string `protobuf:"bytes,2,opt,name=link,proto3" json:"link,omitempty"`
	// The package name where vulnerability has been found.
	Package string `protobuf:"bytes,3,opt,name=package,proto3" json:"package,omitempty"`
	// The package manager name. Ex.: yum, rpm, dpkg.
	Source string `protobuf:"bytes,4,opt,name=source,proto3" json:"source,omitempty"`
	// The version of the package where vulnerability has been found.
	Version string `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
	// The version of the package where vulnerability has been fixed.
	FixedBy string `protobuf:"bytes,6,opt,name=fixed_by,json=fixedBy,proto3" json:"fixed_by,omitempty"`
	// The place where vulnerability is originated (OS, lang package, etc.)
	Origin string `protobuf:"bytes,7,opt,name=origin,proto3" json:"origin,omitempty"`
	// The type of vulnerability origin - name of OS if origin="os" or package type (jar, gobinary, etc.) if origin="lang"
	Type string `protobuf:"bytes,8,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *PackageVulnerability) Reset() {
	*x = PackageVulnerability{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_containerregistry_v1_scanner_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PackageVulnerability) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PackageVulnerability) ProtoMessage() {}

func (x *PackageVulnerability) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_containerregistry_v1_scanner_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PackageVulnerability.ProtoReflect.Descriptor instead.
func (*PackageVulnerability) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_containerregistry_v1_scanner_proto_rawDescGZIP(), []int{3}
}

func (x *PackageVulnerability) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PackageVulnerability) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *PackageVulnerability) GetPackage() string {
	if x != nil {
		return x.Package
	}
	return ""
}

func (x *PackageVulnerability) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *PackageVulnerability) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *PackageVulnerability) GetFixedBy() string {
	if x != nil {
		return x.FixedBy
	}
	return ""
}

func (x *PackageVulnerability) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *PackageVulnerability) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

var File_yandex_cloud_containerregistry_v1_scanner_proto protoreflect.FileDescriptor

var file_yandex_cloud_containerregistry_v1_scanner_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x21, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe6, 0x02, 0x0a, 0x0a, 0x53, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x39,
	0x0a, 0x0a, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x64, 0x41, 0x74, 0x12, 0x4c, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x34, 0x2e, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63,
	0x61, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x5f, 0x0a, 0x0f, 0x76, 0x75, 0x6c, 0x6e, 0x65,
	0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x35, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x0f, 0x76, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x22, 0x43, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55,
	0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x45, 0x41, 0x44, 0x59,
	0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x03, 0x22, 0xac, 0x01,
	0x0a, 0x12, 0x56, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c,
	0x12, 0x12, 0x0a, 0x04, 0x68, 0x69, 0x67, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x68, 0x69, 0x67, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x12, 0x10, 0x0a, 0x03,
	0x6c, 0x6f, 0x77, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6c, 0x6f, 0x77, 0x12, 0x1e,
	0x0a, 0x0a, 0x6e, 0x65, 0x67, 0x6c, 0x69, 0x67, 0x69, 0x62, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x6e, 0x65, 0x67, 0x6c, 0x69, 0x67, 0x69, 0x62, 0x6c, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x75, 0x6e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x75, 0x6e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x22, 0xc4, 0x02, 0x0a,
	0x0d, 0x56, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x55,
	0x0a, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x39, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x2e, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x52, 0x08, 0x73, 0x65, 0x76,
	0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x53, 0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x56, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x48,
	0x00, 0x52, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x22, 0x70, 0x0a, 0x08, 0x53, 0x65,
	0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x14, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49,
	0x54, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x0c, 0x0a, 0x08, 0x43, 0x52, 0x49, 0x54, 0x49, 0x43, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x08,
	0x0a, 0x04, 0x48, 0x49, 0x47, 0x48, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x45, 0x44, 0x49,
	0x55, 0x4d, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x4c, 0x4f, 0x57, 0x10, 0x04, 0x12, 0x0e, 0x0a,
	0x0a, 0x4e, 0x45, 0x47, 0x4c, 0x49, 0x47, 0x49, 0x42, 0x4c, 0x45, 0x10, 0x05, 0x12, 0x0d, 0x0a,
	0x09, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x06, 0x42, 0x15, 0x0a, 0x0d,
	0x76, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x04, 0xc0,
	0xc1, 0x31, 0x01, 0x22, 0xd1, 0x01, 0x0a, 0x14, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x56,
	0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x19, 0x0a, 0x08, 0x66, 0x69, 0x78, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x66, 0x69, 0x78, 0x65, 0x64, 0x42, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x42, 0x80, 0x01, 0x0a, 0x25, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76,
	0x31, 0x5a, 0x57, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_yandex_cloud_containerregistry_v1_scanner_proto_rawDescOnce sync.Once
	file_yandex_cloud_containerregistry_v1_scanner_proto_rawDescData = file_yandex_cloud_containerregistry_v1_scanner_proto_rawDesc
)

func file_yandex_cloud_containerregistry_v1_scanner_proto_rawDescGZIP() []byte {
	file_yandex_cloud_containerregistry_v1_scanner_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_containerregistry_v1_scanner_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_containerregistry_v1_scanner_proto_rawDescData)
	})
	return file_yandex_cloud_containerregistry_v1_scanner_proto_rawDescData
}

var file_yandex_cloud_containerregistry_v1_scanner_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_yandex_cloud_containerregistry_v1_scanner_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_yandex_cloud_containerregistry_v1_scanner_proto_goTypes = []any{
	(ScanResult_Status)(0),        // 0: yandex.cloud.containerregistry.v1.ScanResult.Status
	(Vulnerability_Severity)(0),   // 1: yandex.cloud.containerregistry.v1.Vulnerability.Severity
	(*ScanResult)(nil),            // 2: yandex.cloud.containerregistry.v1.ScanResult
	(*VulnerabilityStats)(nil),    // 3: yandex.cloud.containerregistry.v1.VulnerabilityStats
	(*Vulnerability)(nil),         // 4: yandex.cloud.containerregistry.v1.Vulnerability
	(*PackageVulnerability)(nil),  // 5: yandex.cloud.containerregistry.v1.PackageVulnerability
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_yandex_cloud_containerregistry_v1_scanner_proto_depIdxs = []int32{
	6, // 0: yandex.cloud.containerregistry.v1.ScanResult.scanned_at:type_name -> google.protobuf.Timestamp
	0, // 1: yandex.cloud.containerregistry.v1.ScanResult.status:type_name -> yandex.cloud.containerregistry.v1.ScanResult.Status
	3, // 2: yandex.cloud.containerregistry.v1.ScanResult.vulnerabilities:type_name -> yandex.cloud.containerregistry.v1.VulnerabilityStats
	1, // 3: yandex.cloud.containerregistry.v1.Vulnerability.severity:type_name -> yandex.cloud.containerregistry.v1.Vulnerability.Severity
	5, // 4: yandex.cloud.containerregistry.v1.Vulnerability.package:type_name -> yandex.cloud.containerregistry.v1.PackageVulnerability
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_yandex_cloud_containerregistry_v1_scanner_proto_init() }
func file_yandex_cloud_containerregistry_v1_scanner_proto_init() {
	if File_yandex_cloud_containerregistry_v1_scanner_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_containerregistry_v1_scanner_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ScanResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_yandex_cloud_containerregistry_v1_scanner_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*VulnerabilityStats); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_yandex_cloud_containerregistry_v1_scanner_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Vulnerability); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_yandex_cloud_containerregistry_v1_scanner_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*PackageVulnerability); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_yandex_cloud_containerregistry_v1_scanner_proto_msgTypes[2].OneofWrappers = []any{
		(*Vulnerability_Package)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_containerregistry_v1_scanner_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_containerregistry_v1_scanner_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_containerregistry_v1_scanner_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_containerregistry_v1_scanner_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_containerregistry_v1_scanner_proto_msgTypes,
	}.Build()
	File_yandex_cloud_containerregistry_v1_scanner_proto = out.File
	file_yandex_cloud_containerregistry_v1_scanner_proto_rawDesc = nil
	file_yandex_cloud_containerregistry_v1_scanner_proto_goTypes = nil
	file_yandex_cloud_containerregistry_v1_scanner_proto_depIdxs = nil
}
