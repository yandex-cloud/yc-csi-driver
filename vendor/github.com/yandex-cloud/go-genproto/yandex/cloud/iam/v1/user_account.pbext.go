// Code generated by protoc-gen-goext. DO NOT EDIT.

package iam

import (
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type UserAccount_UserAccount = isUserAccount_UserAccount

func (m *UserAccount) SetUserAccount(v UserAccount_UserAccount) {
	m.UserAccount = v
}

func (m *UserAccount) SetId(v string) {
	m.Id = v
}

func (m *UserAccount) SetYandexPassportUserAccount(v *YandexPassportUserAccount) {
	m.UserAccount = &UserAccount_YandexPassportUserAccount{
		YandexPassportUserAccount: v,
	}
}

func (m *UserAccount) SetSamlUserAccount(v *SamlUserAccount) {
	m.UserAccount = &UserAccount_SamlUserAccount{
		SamlUserAccount: v,
	}
}

func (m *UserAccount) SetLastAuthenticatedAt(v *timestamppb.Timestamp) {
	m.LastAuthenticatedAt = v
}

func (m *YandexPassportUserAccount) SetLogin(v string) {
	m.Login = v
}

func (m *YandexPassportUserAccount) SetDefaultEmail(v string) {
	m.DefaultEmail = v
}

func (m *SamlUserAccount) SetFederationId(v string) {
	m.FederationId = v
}

func (m *SamlUserAccount) SetNameId(v string) {
	m.NameId = v
}

func (m *SamlUserAccount) SetAttributes(v map[string]*SamlUserAccount_Attribute) {
	m.Attributes = v
}

func (m *SamlUserAccount_Attribute) SetValue(v []string) {
	m.Value = v
}
