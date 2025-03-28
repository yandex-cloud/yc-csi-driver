// Code generated by protoc-gen-goext. DO NOT EDIT.

package redis

import (
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

func (m *RedisConfig) SetMaxmemoryPolicy(v RedisConfig_MaxmemoryPolicy) {
	m.MaxmemoryPolicy = v
}

func (m *RedisConfig) SetTimeout(v *wrapperspb.Int64Value) {
	m.Timeout = v
}

func (m *RedisConfig) SetPassword(v string) {
	m.Password = v
}

func (m *RedisConfig) SetDatabases(v *wrapperspb.Int64Value) {
	m.Databases = v
}

func (m *RedisConfig) SetSlowlogLogSlowerThan(v *wrapperspb.Int64Value) {
	m.SlowlogLogSlowerThan = v
}

func (m *RedisConfig) SetSlowlogMaxLen(v *wrapperspb.Int64Value) {
	m.SlowlogMaxLen = v
}

func (m *RedisConfig) SetNotifyKeyspaceEvents(v string) {
	m.NotifyKeyspaceEvents = v
}

func (m *RedisConfig) SetClientOutputBufferLimitPubsub(v *RedisConfig_ClientOutputBufferLimit) {
	m.ClientOutputBufferLimitPubsub = v
}

func (m *RedisConfig) SetClientOutputBufferLimitNormal(v *RedisConfig_ClientOutputBufferLimit) {
	m.ClientOutputBufferLimitNormal = v
}

func (m *RedisConfig) SetMaxmemoryPercent(v *wrapperspb.Int64Value) {
	m.MaxmemoryPercent = v
}

func (m *RedisConfig) SetLuaTimeLimit(v *wrapperspb.Int64Value) {
	m.LuaTimeLimit = v
}

func (m *RedisConfig) SetReplBacklogSizePercent(v *wrapperspb.Int64Value) {
	m.ReplBacklogSizePercent = v
}

func (m *RedisConfig) SetClusterRequireFullCoverage(v *wrapperspb.BoolValue) {
	m.ClusterRequireFullCoverage = v
}

func (m *RedisConfig) SetClusterAllowReadsWhenDown(v *wrapperspb.BoolValue) {
	m.ClusterAllowReadsWhenDown = v
}

func (m *RedisConfig) SetClusterAllowPubsubshardWhenDown(v *wrapperspb.BoolValue) {
	m.ClusterAllowPubsubshardWhenDown = v
}

func (m *RedisConfig) SetLfuDecayTime(v *wrapperspb.Int64Value) {
	m.LfuDecayTime = v
}

func (m *RedisConfig) SetLfuLogFactor(v *wrapperspb.Int64Value) {
	m.LfuLogFactor = v
}

func (m *RedisConfig) SetTurnBeforeSwitchover(v *wrapperspb.BoolValue) {
	m.TurnBeforeSwitchover = v
}

func (m *RedisConfig) SetAllowDataLoss(v *wrapperspb.BoolValue) {
	m.AllowDataLoss = v
}

func (m *RedisConfig) SetUseLuajit(v *wrapperspb.BoolValue) {
	m.UseLuajit = v
}

func (m *RedisConfig) SetIoThreadsAllowed(v *wrapperspb.BoolValue) {
	m.IoThreadsAllowed = v
}

func (m *RedisConfig) SetZsetMaxListpackEntries(v *wrapperspb.Int64Value) {
	m.ZsetMaxListpackEntries = v
}

func (m *RedisConfig_ClientOutputBufferLimit) SetHardLimit(v *wrapperspb.Int64Value) {
	m.HardLimit = v
}

func (m *RedisConfig_ClientOutputBufferLimit) SetSoftLimit(v *wrapperspb.Int64Value) {
	m.SoftLimit = v
}

func (m *RedisConfig_ClientOutputBufferLimit) SetSoftSeconds(v *wrapperspb.Int64Value) {
	m.SoftSeconds = v
}

func (m *RedisConfigSet) SetEffectiveConfig(v *RedisConfig) {
	m.EffectiveConfig = v
}

func (m *RedisConfigSet) SetUserConfig(v *RedisConfig) {
	m.UserConfig = v
}

func (m *RedisConfigSet) SetDefaultConfig(v *RedisConfig) {
	m.DefaultConfig = v
}
