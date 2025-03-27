// Code generated by protoc-gen-goext. DO NOT EDIT.

package smartwebsecurity

import (
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (m *SecurityProfile) SetId(v string) {
	m.Id = v
}

func (m *SecurityProfile) SetFolderId(v string) {
	m.FolderId = v
}

func (m *SecurityProfile) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *SecurityProfile) SetName(v string) {
	m.Name = v
}

func (m *SecurityProfile) SetDescription(v string) {
	m.Description = v
}

func (m *SecurityProfile) SetDefaultAction(v SecurityProfile_DefaultAction) {
	m.DefaultAction = v
}

func (m *SecurityProfile) SetSecurityRules(v []*SecurityRule) {
	m.SecurityRules = v
}

func (m *SecurityProfile) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}

func (m *SecurityProfile) SetCloudId(v string) {
	m.CloudId = v
}

func (m *SecurityProfile) SetCaptchaId(v string) {
	m.CaptchaId = v
}

func (m *SecurityProfile) SetAdvancedRateLimiterProfileId(v string) {
	m.AdvancedRateLimiterProfileId = v
}

func (m *SecurityProfile) SetAnalyzeRequestBody(v *SecurityProfile_AnalyzeRequestBody) {
	m.AnalyzeRequestBody = v
}

func (m *SecurityProfile_AnalyzeRequestBody) SetSizeLimit(v int64) {
	m.SizeLimit = v
}

func (m *SecurityProfile_AnalyzeRequestBody) SetSizeLimitAction(v SecurityProfile_AnalyzeRequestBody_Action) {
	m.SizeLimitAction = v
}

type SecurityRule_RuleSpecifier = isSecurityRule_RuleSpecifier

func (m *SecurityRule) SetRuleSpecifier(v SecurityRule_RuleSpecifier) {
	m.RuleSpecifier = v
}

func (m *SecurityRule) SetName(v string) {
	m.Name = v
}

func (m *SecurityRule) SetPriority(v int64) {
	m.Priority = v
}

func (m *SecurityRule) SetDryRun(v bool) {
	m.DryRun = v
}

func (m *SecurityRule) SetRuleCondition(v *SecurityRule_RuleCondition) {
	m.RuleSpecifier = &SecurityRule_RuleCondition_{
		RuleCondition: v,
	}
}

func (m *SecurityRule) SetSmartProtection(v *SecurityRule_SmartProtection) {
	m.RuleSpecifier = &SecurityRule_SmartProtection_{
		SmartProtection: v,
	}
}

func (m *SecurityRule) SetWaf(v *SecurityRule_Waf) {
	m.RuleSpecifier = &SecurityRule_Waf_{
		Waf: v,
	}
}

func (m *SecurityRule) SetDescription(v string) {
	m.Description = v
}

func (m *SecurityRule_RuleCondition) SetAction(v SecurityRule_RuleCondition_Action) {
	m.Action = v
}

func (m *SecurityRule_RuleCondition) SetCondition(v *Condition) {
	m.Condition = v
}

func (m *SecurityRule_SmartProtection) SetMode(v SecurityRule_SmartProtection_Mode) {
	m.Mode = v
}

func (m *SecurityRule_SmartProtection) SetCondition(v *Condition) {
	m.Condition = v
}

func (m *SecurityRule_Waf) SetMode(v SecurityRule_Waf_Mode) {
	m.Mode = v
}

func (m *SecurityRule_Waf) SetCondition(v *Condition) {
	m.Condition = v
}

func (m *SecurityRule_Waf) SetWafProfileId(v string) {
	m.WafProfileId = v
}

func (m *Condition) SetAuthority(v *Condition_AuthorityMatcher) {
	m.Authority = v
}

func (m *Condition) SetHttpMethod(v *Condition_HttpMethodMatcher) {
	m.HttpMethod = v
}

func (m *Condition) SetRequestUri(v *Condition_RequestUriMatcher) {
	m.RequestUri = v
}

func (m *Condition) SetHeaders(v []*Condition_HeaderMatcher) {
	m.Headers = v
}

func (m *Condition) SetSourceIp(v *Condition_IpMatcher) {
	m.SourceIp = v
}

type Condition_StringMatcher_Match = isCondition_StringMatcher_Match

func (m *Condition_StringMatcher) SetMatch(v Condition_StringMatcher_Match) {
	m.Match = v
}

func (m *Condition_StringMatcher) SetExactMatch(v string) {
	m.Match = &Condition_StringMatcher_ExactMatch{
		ExactMatch: v,
	}
}

func (m *Condition_StringMatcher) SetExactNotMatch(v string) {
	m.Match = &Condition_StringMatcher_ExactNotMatch{
		ExactNotMatch: v,
	}
}

func (m *Condition_StringMatcher) SetPrefixMatch(v string) {
	m.Match = &Condition_StringMatcher_PrefixMatch{
		PrefixMatch: v,
	}
}

func (m *Condition_StringMatcher) SetPrefixNotMatch(v string) {
	m.Match = &Condition_StringMatcher_PrefixNotMatch{
		PrefixNotMatch: v,
	}
}

func (m *Condition_StringMatcher) SetPireRegexMatch(v string) {
	m.Match = &Condition_StringMatcher_PireRegexMatch{
		PireRegexMatch: v,
	}
}

func (m *Condition_StringMatcher) SetPireRegexNotMatch(v string) {
	m.Match = &Condition_StringMatcher_PireRegexNotMatch{
		PireRegexNotMatch: v,
	}
}

func (m *Condition_HttpMethodMatcher) SetHttpMethods(v []*Condition_StringMatcher) {
	m.HttpMethods = v
}

func (m *Condition_AuthorityMatcher) SetAuthorities(v []*Condition_StringMatcher) {
	m.Authorities = v
}

func (m *Condition_RequestUriMatcher) SetPath(v *Condition_StringMatcher) {
	m.Path = v
}

func (m *Condition_RequestUriMatcher) SetQueries(v []*Condition_QueryMatcher) {
	m.Queries = v
}

func (m *Condition_QueryMatcher) SetKey(v string) {
	m.Key = v
}

func (m *Condition_QueryMatcher) SetValue(v *Condition_StringMatcher) {
	m.Value = v
}

func (m *Condition_HeaderMatcher) SetName(v string) {
	m.Name = v
}

func (m *Condition_HeaderMatcher) SetValue(v *Condition_StringMatcher) {
	m.Value = v
}

func (m *Condition_IpMatcher) SetIpRangesMatch(v *Condition_IpRangesMatcher) {
	m.IpRangesMatch = v
}

func (m *Condition_IpMatcher) SetIpRangesNotMatch(v *Condition_IpRangesMatcher) {
	m.IpRangesNotMatch = v
}

func (m *Condition_IpMatcher) SetGeoIpMatch(v *Condition_GeoIpMatcher) {
	m.GeoIpMatch = v
}

func (m *Condition_IpMatcher) SetGeoIpNotMatch(v *Condition_GeoIpMatcher) {
	m.GeoIpNotMatch = v
}

func (m *Condition_IpRangesMatcher) SetIpRanges(v []string) {
	m.IpRanges = v
}

func (m *Condition_GeoIpMatcher) SetLocations(v []string) {
	m.Locations = v
}
