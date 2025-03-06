// Code generated by protoc-gen-goext. DO NOT EDIT.

package mongodb

import (
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

func (m *MongodConfig) SetStorage(v *MongodConfig_Storage) {
	m.Storage = v
}

func (m *MongodConfig) SetOperationProfiling(v *MongodConfig_OperationProfiling) {
	m.OperationProfiling = v
}

func (m *MongodConfig) SetNet(v *MongodConfig_Network) {
	m.Net = v
}

func (m *MongodConfig) SetSecurity(v *MongodConfig_Security) {
	m.Security = v
}

func (m *MongodConfig) SetAuditLog(v *MongodConfig_AuditLog) {
	m.AuditLog = v
}

func (m *MongodConfig) SetSetParameter(v *MongodConfig_SetParameter) {
	m.SetParameter = v
}

func (m *MongodConfig_Storage) SetWiredTiger(v *MongodConfig_Storage_WiredTiger) {
	m.WiredTiger = v
}

func (m *MongodConfig_Storage) SetJournal(v *MongodConfig_Storage_Journal) {
	m.Journal = v
}

func (m *MongodConfig_Storage_WiredTiger) SetEngineConfig(v *MongodConfig_Storage_WiredTiger_EngineConfig) {
	m.EngineConfig = v
}

func (m *MongodConfig_Storage_WiredTiger) SetCollectionConfig(v *MongodConfig_Storage_WiredTiger_CollectionConfig) {
	m.CollectionConfig = v
}

func (m *MongodConfig_Storage_WiredTiger) SetIndexConfig(v *MongodConfig_Storage_WiredTiger_IndexConfig) {
	m.IndexConfig = v
}

func (m *MongodConfig_Storage_WiredTiger_EngineConfig) SetCacheSizeGb(v *wrapperspb.DoubleValue) {
	m.CacheSizeGb = v
}

func (m *MongodConfig_Storage_WiredTiger_CollectionConfig) SetBlockCompressor(v MongodConfig_Storage_WiredTiger_CollectionConfig_Compressor) {
	m.BlockCompressor = v
}

func (m *MongodConfig_Storage_WiredTiger_IndexConfig) SetPrefixCompression(v *wrapperspb.BoolValue) {
	m.PrefixCompression = v
}

func (m *MongodConfig_Storage_Journal) SetCommitInterval(v *wrapperspb.Int64Value) {
	m.CommitInterval = v
}

func (m *MongodConfig_OperationProfiling) SetMode(v MongodConfig_OperationProfiling_Mode) {
	m.Mode = v
}

func (m *MongodConfig_OperationProfiling) SetSlowOpThreshold(v *wrapperspb.Int64Value) {
	m.SlowOpThreshold = v
}

func (m *MongodConfig_OperationProfiling) SetSlowOpSampleRate(v *wrapperspb.DoubleValue) {
	m.SlowOpSampleRate = v
}

func (m *MongodConfig_Network) SetMaxIncomingConnections(v *wrapperspb.Int64Value) {
	m.MaxIncomingConnections = v
}

func (m *MongodConfig_Network) SetCompression(v *MongodConfig_Network_Compression) {
	m.Compression = v
}

func (m *MongodConfig_Network_Compression) SetCompressors(v []MongodConfig_Network_Compression_Compressor) {
	m.Compressors = v
}

func (m *MongodConfig_Security) SetEnableEncryption(v *wrapperspb.BoolValue) {
	m.EnableEncryption = v
}

func (m *MongodConfig_Security) SetKmip(v *MongodConfig_Security_KMIP) {
	m.Kmip = v
}

func (m *MongodConfig_Security_KMIP) SetServerName(v string) {
	m.ServerName = v
}

func (m *MongodConfig_Security_KMIP) SetPort(v *wrapperspb.Int64Value) {
	m.Port = v
}

func (m *MongodConfig_Security_KMIP) SetServerCa(v string) {
	m.ServerCa = v
}

func (m *MongodConfig_Security_KMIP) SetClientCertificate(v string) {
	m.ClientCertificate = v
}

func (m *MongodConfig_Security_KMIP) SetKeyIdentifier(v string) {
	m.KeyIdentifier = v
}

func (m *MongodConfig_AuditLog) SetFilter(v string) {
	m.Filter = v
}

func (m *MongodConfig_AuditLog) SetRuntimeConfiguration(v *wrapperspb.BoolValue) {
	m.RuntimeConfiguration = v
}

func (m *MongodConfig_SetParameter) SetAuditAuthorizationSuccess(v *wrapperspb.BoolValue) {
	m.AuditAuthorizationSuccess = v
}

func (m *MongodConfig_SetParameter) SetEnableFlowControl(v *wrapperspb.BoolValue) {
	m.EnableFlowControl = v
}

func (m *MongodConfig_SetParameter) SetMinSnapshotHistoryWindowInSeconds(v *wrapperspb.Int64Value) {
	m.MinSnapshotHistoryWindowInSeconds = v
}

func (m *MongoCfgConfig) SetStorage(v *MongoCfgConfig_Storage) {
	m.Storage = v
}

func (m *MongoCfgConfig) SetOperationProfiling(v *MongoCfgConfig_OperationProfiling) {
	m.OperationProfiling = v
}

func (m *MongoCfgConfig) SetNet(v *MongoCfgConfig_Network) {
	m.Net = v
}

func (m *MongoCfgConfig) SetSetParameter(v *MongoCfgConfig_SetParameter) {
	m.SetParameter = v
}

func (m *MongoCfgConfig) SetAuditLog(v *MongoCfgConfig_AuditLog) {
	m.AuditLog = v
}

func (m *MongoCfgConfig_Storage) SetWiredTiger(v *MongoCfgConfig_Storage_WiredTiger) {
	m.WiredTiger = v
}

func (m *MongoCfgConfig_Storage_WiredTiger) SetEngineConfig(v *MongoCfgConfig_Storage_WiredTiger_EngineConfig) {
	m.EngineConfig = v
}

func (m *MongoCfgConfig_Storage_WiredTiger) SetIndexConfig(v *MongoCfgConfig_Storage_WiredTiger_IndexConfig) {
	m.IndexConfig = v
}

func (m *MongoCfgConfig_Storage_WiredTiger_EngineConfig) SetCacheSizeGb(v *wrapperspb.DoubleValue) {
	m.CacheSizeGb = v
}

func (m *MongoCfgConfig_Storage_WiredTiger_IndexConfig) SetPrefixCompression(v *wrapperspb.BoolValue) {
	m.PrefixCompression = v
}

func (m *MongoCfgConfig_OperationProfiling) SetMode(v MongoCfgConfig_OperationProfiling_Mode) {
	m.Mode = v
}

func (m *MongoCfgConfig_OperationProfiling) SetSlowOpThreshold(v *wrapperspb.Int64Value) {
	m.SlowOpThreshold = v
}

func (m *MongoCfgConfig_Network) SetMaxIncomingConnections(v *wrapperspb.Int64Value) {
	m.MaxIncomingConnections = v
}

func (m *MongoCfgConfig_Network) SetCompression(v *MongoCfgConfig_Network_Compression) {
	m.Compression = v
}

func (m *MongoCfgConfig_Network_Compression) SetCompressors(v []MongoCfgConfig_Network_Compression_Compressor) {
	m.Compressors = v
}

func (m *MongoCfgConfig_SetParameter) SetEnableFlowControl(v *wrapperspb.BoolValue) {
	m.EnableFlowControl = v
}

func (m *MongoCfgConfig_SetParameter) SetAuditAuthorizationSuccess(v *wrapperspb.BoolValue) {
	m.AuditAuthorizationSuccess = v
}

func (m *MongoCfgConfig_AuditLog) SetFilter(v string) {
	m.Filter = v
}

func (m *MongosConfig) SetNet(v *MongosConfig_Network) {
	m.Net = v
}

func (m *MongosConfig) SetSetParameter(v *MongosConfig_SetParameter) {
	m.SetParameter = v
}

func (m *MongosConfig) SetAuditLog(v *MongosConfig_AuditLog) {
	m.AuditLog = v
}

func (m *MongosConfig_Network) SetMaxIncomingConnections(v *wrapperspb.Int64Value) {
	m.MaxIncomingConnections = v
}

func (m *MongosConfig_Network) SetCompression(v *MongosConfig_Network_Compression) {
	m.Compression = v
}

func (m *MongosConfig_Network_Compression) SetCompressors(v []MongosConfig_Network_Compression_Compressor) {
	m.Compressors = v
}

func (m *MongosConfig_SetParameter) SetAuditAuthorizationSuccess(v *wrapperspb.BoolValue) {
	m.AuditAuthorizationSuccess = v
}

func (m *MongosConfig_AuditLog) SetFilter(v string) {
	m.Filter = v
}

func (m *MongodConfigSet) SetEffectiveConfig(v *MongodConfig) {
	m.EffectiveConfig = v
}

func (m *MongodConfigSet) SetUserConfig(v *MongodConfig) {
	m.UserConfig = v
}

func (m *MongodConfigSet) SetDefaultConfig(v *MongodConfig) {
	m.DefaultConfig = v
}

func (m *MongoCfgConfigSet) SetEffectiveConfig(v *MongoCfgConfig) {
	m.EffectiveConfig = v
}

func (m *MongoCfgConfigSet) SetUserConfig(v *MongoCfgConfig) {
	m.UserConfig = v
}

func (m *MongoCfgConfigSet) SetDefaultConfig(v *MongoCfgConfig) {
	m.DefaultConfig = v
}

func (m *MongosConfigSet) SetEffectiveConfig(v *MongosConfig) {
	m.EffectiveConfig = v
}

func (m *MongosConfigSet) SetUserConfig(v *MongosConfig) {
	m.UserConfig = v
}

func (m *MongosConfigSet) SetDefaultConfig(v *MongosConfig) {
	m.DefaultConfig = v
}
