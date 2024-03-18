// Code generated by protoc-gen-goext. DO NOT EDIT.

package mysql

import (
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

func (m *MysqlConfig5_7) SetInnodbBufferPoolSize(v *wrapperspb.Int64Value) {
	m.InnodbBufferPoolSize = v
}

func (m *MysqlConfig5_7) SetMaxConnections(v *wrapperspb.Int64Value) {
	m.MaxConnections = v
}

func (m *MysqlConfig5_7) SetLongQueryTime(v *wrapperspb.DoubleValue) {
	m.LongQueryTime = v
}

func (m *MysqlConfig5_7) SetGeneralLog(v *wrapperspb.BoolValue) {
	m.GeneralLog = v
}

func (m *MysqlConfig5_7) SetAuditLog(v *wrapperspb.BoolValue) {
	m.AuditLog = v
}

func (m *MysqlConfig5_7) SetSqlMode(v []MysqlConfig5_7_SQLMode) {
	m.SqlMode = v
}

func (m *MysqlConfig5_7) SetMaxAllowedPacket(v *wrapperspb.Int64Value) {
	m.MaxAllowedPacket = v
}

func (m *MysqlConfig5_7) SetDefaultAuthenticationPlugin(v MysqlConfig5_7_AuthPlugin) {
	m.DefaultAuthenticationPlugin = v
}

func (m *MysqlConfig5_7) SetInnodbFlushLogAtTrxCommit(v *wrapperspb.Int64Value) {
	m.InnodbFlushLogAtTrxCommit = v
}

func (m *MysqlConfig5_7) SetInnodbLockWaitTimeout(v *wrapperspb.Int64Value) {
	m.InnodbLockWaitTimeout = v
}

func (m *MysqlConfig5_7) SetTransactionIsolation(v MysqlConfig5_7_TransactionIsolation) {
	m.TransactionIsolation = v
}

func (m *MysqlConfig5_7) SetInnodbPrintAllDeadlocks(v *wrapperspb.BoolValue) {
	m.InnodbPrintAllDeadlocks = v
}

func (m *MysqlConfig5_7) SetNetReadTimeout(v *wrapperspb.Int64Value) {
	m.NetReadTimeout = v
}

func (m *MysqlConfig5_7) SetNetWriteTimeout(v *wrapperspb.Int64Value) {
	m.NetWriteTimeout = v
}

func (m *MysqlConfig5_7) SetGroupConcatMaxLen(v *wrapperspb.Int64Value) {
	m.GroupConcatMaxLen = v
}

func (m *MysqlConfig5_7) SetTmpTableSize(v *wrapperspb.Int64Value) {
	m.TmpTableSize = v
}

func (m *MysqlConfig5_7) SetMaxHeapTableSize(v *wrapperspb.Int64Value) {
	m.MaxHeapTableSize = v
}

func (m *MysqlConfig5_7) SetDefaultTimeZone(v string) {
	m.DefaultTimeZone = v
}

func (m *MysqlConfig5_7) SetCharacterSetServer(v string) {
	m.CharacterSetServer = v
}

func (m *MysqlConfig5_7) SetCollationServer(v string) {
	m.CollationServer = v
}

func (m *MysqlConfig5_7) SetInnodbAdaptiveHashIndex(v *wrapperspb.BoolValue) {
	m.InnodbAdaptiveHashIndex = v
}

func (m *MysqlConfig5_7) SetInnodbNumaInterleave(v *wrapperspb.BoolValue) {
	m.InnodbNumaInterleave = v
}

func (m *MysqlConfig5_7) SetInnodbLogBufferSize(v *wrapperspb.Int64Value) {
	m.InnodbLogBufferSize = v
}

func (m *MysqlConfig5_7) SetInnodbLogFileSize(v *wrapperspb.Int64Value) {
	m.InnodbLogFileSize = v
}

func (m *MysqlConfig5_7) SetInnodbIoCapacity(v *wrapperspb.Int64Value) {
	m.InnodbIoCapacity = v
}

func (m *MysqlConfig5_7) SetInnodbIoCapacityMax(v *wrapperspb.Int64Value) {
	m.InnodbIoCapacityMax = v
}

func (m *MysqlConfig5_7) SetInnodbReadIoThreads(v *wrapperspb.Int64Value) {
	m.InnodbReadIoThreads = v
}

func (m *MysqlConfig5_7) SetInnodbWriteIoThreads(v *wrapperspb.Int64Value) {
	m.InnodbWriteIoThreads = v
}

func (m *MysqlConfig5_7) SetInnodbPurgeThreads(v *wrapperspb.Int64Value) {
	m.InnodbPurgeThreads = v
}

func (m *MysqlConfig5_7) SetInnodbThreadConcurrency(v *wrapperspb.Int64Value) {
	m.InnodbThreadConcurrency = v
}

func (m *MysqlConfig5_7) SetInnodbTempDataFileMaxSize(v *wrapperspb.Int64Value) {
	m.InnodbTempDataFileMaxSize = v
}

func (m *MysqlConfig5_7) SetThreadCacheSize(v *wrapperspb.Int64Value) {
	m.ThreadCacheSize = v
}

func (m *MysqlConfig5_7) SetThreadStack(v *wrapperspb.Int64Value) {
	m.ThreadStack = v
}

func (m *MysqlConfig5_7) SetJoinBufferSize(v *wrapperspb.Int64Value) {
	m.JoinBufferSize = v
}

func (m *MysqlConfig5_7) SetSortBufferSize(v *wrapperspb.Int64Value) {
	m.SortBufferSize = v
}

func (m *MysqlConfig5_7) SetTableDefinitionCache(v *wrapperspb.Int64Value) {
	m.TableDefinitionCache = v
}

func (m *MysqlConfig5_7) SetTableOpenCache(v *wrapperspb.Int64Value) {
	m.TableOpenCache = v
}

func (m *MysqlConfig5_7) SetTableOpenCacheInstances(v *wrapperspb.Int64Value) {
	m.TableOpenCacheInstances = v
}

func (m *MysqlConfig5_7) SetExplicitDefaultsForTimestamp(v *wrapperspb.BoolValue) {
	m.ExplicitDefaultsForTimestamp = v
}

func (m *MysqlConfig5_7) SetAutoIncrementIncrement(v *wrapperspb.Int64Value) {
	m.AutoIncrementIncrement = v
}

func (m *MysqlConfig5_7) SetAutoIncrementOffset(v *wrapperspb.Int64Value) {
	m.AutoIncrementOffset = v
}

func (m *MysqlConfig5_7) SetSyncBinlog(v *wrapperspb.Int64Value) {
	m.SyncBinlog = v
}

func (m *MysqlConfig5_7) SetBinlogCacheSize(v *wrapperspb.Int64Value) {
	m.BinlogCacheSize = v
}

func (m *MysqlConfig5_7) SetBinlogGroupCommitSyncDelay(v *wrapperspb.Int64Value) {
	m.BinlogGroupCommitSyncDelay = v
}

func (m *MysqlConfig5_7) SetBinlogRowImage(v MysqlConfig5_7_BinlogRowImage) {
	m.BinlogRowImage = v
}

func (m *MysqlConfig5_7) SetBinlogRowsQueryLogEvents(v *wrapperspb.BoolValue) {
	m.BinlogRowsQueryLogEvents = v
}

func (m *MysqlConfig5_7) SetRplSemiSyncMasterWaitForSlaveCount(v *wrapperspb.Int64Value) {
	m.RplSemiSyncMasterWaitForSlaveCount = v
}

func (m *MysqlConfig5_7) SetSlaveParallelType(v MysqlConfig5_7_SlaveParallelType) {
	m.SlaveParallelType = v
}

func (m *MysqlConfig5_7) SetSlaveParallelWorkers(v *wrapperspb.Int64Value) {
	m.SlaveParallelWorkers = v
}

func (m *MysqlConfig5_7) SetMdbPreserveBinlogBytes(v *wrapperspb.Int64Value) {
	m.MdbPreserveBinlogBytes = v
}

func (m *MysqlConfig5_7) SetInteractiveTimeout(v *wrapperspb.Int64Value) {
	m.InteractiveTimeout = v
}

func (m *MysqlConfig5_7) SetWaitTimeout(v *wrapperspb.Int64Value) {
	m.WaitTimeout = v
}

func (m *MysqlConfig5_7) SetMdbOfflineModeEnableLag(v *wrapperspb.Int64Value) {
	m.MdbOfflineModeEnableLag = v
}

func (m *MysqlConfig5_7) SetMdbOfflineModeDisableLag(v *wrapperspb.Int64Value) {
	m.MdbOfflineModeDisableLag = v
}

func (m *MysqlConfig5_7) SetRangeOptimizerMaxMemSize(v *wrapperspb.Int64Value) {
	m.RangeOptimizerMaxMemSize = v
}

func (m *MysqlConfig5_7) SetSlowQueryLog(v *wrapperspb.BoolValue) {
	m.SlowQueryLog = v
}

func (m *MysqlConfig5_7) SetSlowQueryLogAlwaysWriteTime(v *wrapperspb.DoubleValue) {
	m.SlowQueryLogAlwaysWriteTime = v
}

func (m *MysqlConfig5_7) SetLogSlowRateType(v MysqlConfig5_7_LogSlowRateType) {
	m.LogSlowRateType = v
}

func (m *MysqlConfig5_7) SetLogSlowRateLimit(v *wrapperspb.Int64Value) {
	m.LogSlowRateLimit = v
}

func (m *MysqlConfig5_7) SetLogSlowSpStatements(v *wrapperspb.BoolValue) {
	m.LogSlowSpStatements = v
}

func (m *MysqlConfig5_7) SetLogSlowFilter(v []MysqlConfig5_7_LogSlowFilterType) {
	m.LogSlowFilter = v
}

func (m *MysqlConfig5_7) SetMdbPriorityChoiceMaxLag(v *wrapperspb.Int64Value) {
	m.MdbPriorityChoiceMaxLag = v
}

func (m *MysqlConfig5_7) SetInnodbPageSize(v *wrapperspb.Int64Value) {
	m.InnodbPageSize = v
}

func (m *MysqlConfig5_7) SetInnodbOnlineAlterLogMaxSize(v *wrapperspb.Int64Value) {
	m.InnodbOnlineAlterLogMaxSize = v
}

func (m *MysqlConfig5_7) SetInnodbFtMinTokenSize(v *wrapperspb.Int64Value) {
	m.InnodbFtMinTokenSize = v
}

func (m *MysqlConfig5_7) SetInnodbFtMaxTokenSize(v *wrapperspb.Int64Value) {
	m.InnodbFtMaxTokenSize = v
}

func (m *MysqlConfig5_7) SetLowerCaseTableNames(v *wrapperspb.Int64Value) {
	m.LowerCaseTableNames = v
}

func (m *MysqlConfig5_7) SetShowCompatibility_56(v *wrapperspb.BoolValue) {
	m.ShowCompatibility_56 = v
}

func (m *MysqlConfig5_7) SetMaxSpRecursionDepth(v *wrapperspb.Int64Value) {
	m.MaxSpRecursionDepth = v
}

func (m *MysqlConfig5_7) SetInnodbCompressionLevel(v *wrapperspb.Int64Value) {
	m.InnodbCompressionLevel = v
}

func (m *MysqlConfig5_7) SetBinlogTransactionDependencyTracking(v MysqlConfig5_7_BinlogTransactionDependencyTracking) {
	m.BinlogTransactionDependencyTracking = v
}

func (m *MysqlConfig5_7) SetAutocommit(v *wrapperspb.BoolValue) {
	m.Autocommit = v
}

func (m *MysqlConfig5_7) SetInnodbStatusOutput(v *wrapperspb.BoolValue) {
	m.InnodbStatusOutput = v
}

func (m *MysqlConfig5_7) SetInnodbStrictMode(v *wrapperspb.BoolValue) {
	m.InnodbStrictMode = v
}

func (m *MysqlConfig5_7) SetInnodbPrintLockWaitTimeoutInfo(v *wrapperspb.BoolValue) {
	m.InnodbPrintLockWaitTimeoutInfo = v
}

func (m *MysqlConfig5_7) SetLogErrorVerbosity(v *wrapperspb.Int64Value) {
	m.LogErrorVerbosity = v
}

func (m *MysqlConfig5_7) SetMaxDigestLength(v *wrapperspb.Int64Value) {
	m.MaxDigestLength = v
}

func (m *MysqlConfig5_7) SetQueryCacheLimit(v *wrapperspb.Int64Value) {
	m.QueryCacheLimit = v
}

func (m *MysqlConfig5_7) SetQueryCacheSize(v *wrapperspb.Int64Value) {
	m.QueryCacheSize = v
}

func (m *MysqlConfig5_7) SetQueryCacheType(v *wrapperspb.Int64Value) {
	m.QueryCacheType = v
}

func (m *MysqlConfig5_7) SetLockWaitTimeout(v *wrapperspb.Int64Value) {
	m.LockWaitTimeout = v
}

func (m *MysqlConfig5_7) SetMaxPreparedStmtCount(v *wrapperspb.Int64Value) {
	m.MaxPreparedStmtCount = v
}

func (m *MysqlConfig5_7) SetOptimizerSwitch(v string) {
	m.OptimizerSwitch = v
}

func (m *MysqlConfig5_7) SetOptimizerSearchDepth(v *wrapperspb.Int64Value) {
	m.OptimizerSearchDepth = v
}

func (m *MysqlConfig5_7) SetQueryResponseTimeStats(v *wrapperspb.BoolValue) {
	m.QueryResponseTimeStats = v
}

func (m *MysqlConfig5_7) SetUserstat(v *wrapperspb.BoolValue) {
	m.Userstat = v
}

func (m *MysqlConfig5_7) SetMaxExecutionTime(v *wrapperspb.Int64Value) {
	m.MaxExecutionTime = v
}

func (m *MysqlConfigSet5_7) SetEffectiveConfig(v *MysqlConfig5_7) {
	m.EffectiveConfig = v
}

func (m *MysqlConfigSet5_7) SetUserConfig(v *MysqlConfig5_7) {
	m.UserConfig = v
}

func (m *MysqlConfigSet5_7) SetDefaultConfig(v *MysqlConfig5_7) {
	m.DefaultConfig = v
}
