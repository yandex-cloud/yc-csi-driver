// Code generated by protoc-gen-goext. DO NOT EDIT.

package postgresql

import (
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

func (m *PostgresqlConfig12) SetMaxConnections(v *wrapperspb.Int64Value) {
	m.MaxConnections = v
}

func (m *PostgresqlConfig12) SetSharedBuffers(v *wrapperspb.Int64Value) {
	m.SharedBuffers = v
}

func (m *PostgresqlConfig12) SetTempBuffers(v *wrapperspb.Int64Value) {
	m.TempBuffers = v
}

func (m *PostgresqlConfig12) SetMaxPreparedTransactions(v *wrapperspb.Int64Value) {
	m.MaxPreparedTransactions = v
}

func (m *PostgresqlConfig12) SetWorkMem(v *wrapperspb.Int64Value) {
	m.WorkMem = v
}

func (m *PostgresqlConfig12) SetMaintenanceWorkMem(v *wrapperspb.Int64Value) {
	m.MaintenanceWorkMem = v
}

func (m *PostgresqlConfig12) SetAutovacuumWorkMem(v *wrapperspb.Int64Value) {
	m.AutovacuumWorkMem = v
}

func (m *PostgresqlConfig12) SetTempFileLimit(v *wrapperspb.Int64Value) {
	m.TempFileLimit = v
}

func (m *PostgresqlConfig12) SetVacuumCostDelay(v *wrapperspb.Int64Value) {
	m.VacuumCostDelay = v
}

func (m *PostgresqlConfig12) SetVacuumCostPageHit(v *wrapperspb.Int64Value) {
	m.VacuumCostPageHit = v
}

func (m *PostgresqlConfig12) SetVacuumCostPageMiss(v *wrapperspb.Int64Value) {
	m.VacuumCostPageMiss = v
}

func (m *PostgresqlConfig12) SetVacuumCostPageDirty(v *wrapperspb.Int64Value) {
	m.VacuumCostPageDirty = v
}

func (m *PostgresqlConfig12) SetVacuumCostLimit(v *wrapperspb.Int64Value) {
	m.VacuumCostLimit = v
}

func (m *PostgresqlConfig12) SetBgwriterDelay(v *wrapperspb.Int64Value) {
	m.BgwriterDelay = v
}

func (m *PostgresqlConfig12) SetBgwriterLruMaxpages(v *wrapperspb.Int64Value) {
	m.BgwriterLruMaxpages = v
}

func (m *PostgresqlConfig12) SetBgwriterLruMultiplier(v *wrapperspb.DoubleValue) {
	m.BgwriterLruMultiplier = v
}

func (m *PostgresqlConfig12) SetBgwriterFlushAfter(v *wrapperspb.Int64Value) {
	m.BgwriterFlushAfter = v
}

func (m *PostgresqlConfig12) SetBackendFlushAfter(v *wrapperspb.Int64Value) {
	m.BackendFlushAfter = v
}

func (m *PostgresqlConfig12) SetOldSnapshotThreshold(v *wrapperspb.Int64Value) {
	m.OldSnapshotThreshold = v
}

func (m *PostgresqlConfig12) SetWalLevel(v PostgresqlConfig12_WalLevel) {
	m.WalLevel = v
}

func (m *PostgresqlConfig12) SetSynchronousCommit(v PostgresqlConfig12_SynchronousCommit) {
	m.SynchronousCommit = v
}

func (m *PostgresqlConfig12) SetCheckpointTimeout(v *wrapperspb.Int64Value) {
	m.CheckpointTimeout = v
}

func (m *PostgresqlConfig12) SetCheckpointCompletionTarget(v *wrapperspb.DoubleValue) {
	m.CheckpointCompletionTarget = v
}

func (m *PostgresqlConfig12) SetCheckpointFlushAfter(v *wrapperspb.Int64Value) {
	m.CheckpointFlushAfter = v
}

func (m *PostgresqlConfig12) SetMaxWalSize(v *wrapperspb.Int64Value) {
	m.MaxWalSize = v
}

func (m *PostgresqlConfig12) SetMinWalSize(v *wrapperspb.Int64Value) {
	m.MinWalSize = v
}

func (m *PostgresqlConfig12) SetMaxStandbyStreamingDelay(v *wrapperspb.Int64Value) {
	m.MaxStandbyStreamingDelay = v
}

func (m *PostgresqlConfig12) SetDefaultStatisticsTarget(v *wrapperspb.Int64Value) {
	m.DefaultStatisticsTarget = v
}

func (m *PostgresqlConfig12) SetConstraintExclusion(v PostgresqlConfig12_ConstraintExclusion) {
	m.ConstraintExclusion = v
}

func (m *PostgresqlConfig12) SetCursorTupleFraction(v *wrapperspb.DoubleValue) {
	m.CursorTupleFraction = v
}

func (m *PostgresqlConfig12) SetFromCollapseLimit(v *wrapperspb.Int64Value) {
	m.FromCollapseLimit = v
}

func (m *PostgresqlConfig12) SetJoinCollapseLimit(v *wrapperspb.Int64Value) {
	m.JoinCollapseLimit = v
}

func (m *PostgresqlConfig12) SetForceParallelMode(v PostgresqlConfig12_ForceParallelMode) {
	m.ForceParallelMode = v
}

func (m *PostgresqlConfig12) SetClientMinMessages(v PostgresqlConfig12_LogLevel) {
	m.ClientMinMessages = v
}

func (m *PostgresqlConfig12) SetLogMinMessages(v PostgresqlConfig12_LogLevel) {
	m.LogMinMessages = v
}

func (m *PostgresqlConfig12) SetLogMinErrorStatement(v PostgresqlConfig12_LogLevel) {
	m.LogMinErrorStatement = v
}

func (m *PostgresqlConfig12) SetLogMinDurationStatement(v *wrapperspb.Int64Value) {
	m.LogMinDurationStatement = v
}

func (m *PostgresqlConfig12) SetLogCheckpoints(v *wrapperspb.BoolValue) {
	m.LogCheckpoints = v
}

func (m *PostgresqlConfig12) SetLogConnections(v *wrapperspb.BoolValue) {
	m.LogConnections = v
}

func (m *PostgresqlConfig12) SetLogDisconnections(v *wrapperspb.BoolValue) {
	m.LogDisconnections = v
}

func (m *PostgresqlConfig12) SetLogDuration(v *wrapperspb.BoolValue) {
	m.LogDuration = v
}

func (m *PostgresqlConfig12) SetLogErrorVerbosity(v PostgresqlConfig12_LogErrorVerbosity) {
	m.LogErrorVerbosity = v
}

func (m *PostgresqlConfig12) SetLogLockWaits(v *wrapperspb.BoolValue) {
	m.LogLockWaits = v
}

func (m *PostgresqlConfig12) SetLogStatement(v PostgresqlConfig12_LogStatement) {
	m.LogStatement = v
}

func (m *PostgresqlConfig12) SetLogTempFiles(v *wrapperspb.Int64Value) {
	m.LogTempFiles = v
}

func (m *PostgresqlConfig12) SetSearchPath(v string) {
	m.SearchPath = v
}

func (m *PostgresqlConfig12) SetRowSecurity(v *wrapperspb.BoolValue) {
	m.RowSecurity = v
}

func (m *PostgresqlConfig12) SetDefaultTransactionIsolation(v PostgresqlConfig12_TransactionIsolation) {
	m.DefaultTransactionIsolation = v
}

func (m *PostgresqlConfig12) SetStatementTimeout(v *wrapperspb.Int64Value) {
	m.StatementTimeout = v
}

func (m *PostgresqlConfig12) SetLockTimeout(v *wrapperspb.Int64Value) {
	m.LockTimeout = v
}

func (m *PostgresqlConfig12) SetIdleInTransactionSessionTimeout(v *wrapperspb.Int64Value) {
	m.IdleInTransactionSessionTimeout = v
}

func (m *PostgresqlConfig12) SetByteaOutput(v PostgresqlConfig12_ByteaOutput) {
	m.ByteaOutput = v
}

func (m *PostgresqlConfig12) SetXmlbinary(v PostgresqlConfig12_XmlBinary) {
	m.Xmlbinary = v
}

func (m *PostgresqlConfig12) SetXmloption(v PostgresqlConfig12_XmlOption) {
	m.Xmloption = v
}

func (m *PostgresqlConfig12) SetGinPendingListLimit(v *wrapperspb.Int64Value) {
	m.GinPendingListLimit = v
}

func (m *PostgresqlConfig12) SetDeadlockTimeout(v *wrapperspb.Int64Value) {
	m.DeadlockTimeout = v
}

func (m *PostgresqlConfig12) SetMaxLocksPerTransaction(v *wrapperspb.Int64Value) {
	m.MaxLocksPerTransaction = v
}

func (m *PostgresqlConfig12) SetMaxPredLocksPerTransaction(v *wrapperspb.Int64Value) {
	m.MaxPredLocksPerTransaction = v
}

func (m *PostgresqlConfig12) SetArrayNulls(v *wrapperspb.BoolValue) {
	m.ArrayNulls = v
}

func (m *PostgresqlConfig12) SetBackslashQuote(v PostgresqlConfig12_BackslashQuote) {
	m.BackslashQuote = v
}

func (m *PostgresqlConfig12) SetDefaultWithOids(v *wrapperspb.BoolValue) {
	m.DefaultWithOids = v
}

func (m *PostgresqlConfig12) SetEscapeStringWarning(v *wrapperspb.BoolValue) {
	m.EscapeStringWarning = v
}

func (m *PostgresqlConfig12) SetLoCompatPrivileges(v *wrapperspb.BoolValue) {
	m.LoCompatPrivileges = v
}

func (m *PostgresqlConfig12) SetOperatorPrecedenceWarning(v *wrapperspb.BoolValue) {
	m.OperatorPrecedenceWarning = v
}

func (m *PostgresqlConfig12) SetQuoteAllIdentifiers(v *wrapperspb.BoolValue) {
	m.QuoteAllIdentifiers = v
}

func (m *PostgresqlConfig12) SetStandardConformingStrings(v *wrapperspb.BoolValue) {
	m.StandardConformingStrings = v
}

func (m *PostgresqlConfig12) SetSynchronizeSeqscans(v *wrapperspb.BoolValue) {
	m.SynchronizeSeqscans = v
}

func (m *PostgresqlConfig12) SetTransformNullEquals(v *wrapperspb.BoolValue) {
	m.TransformNullEquals = v
}

func (m *PostgresqlConfig12) SetExitOnError(v *wrapperspb.BoolValue) {
	m.ExitOnError = v
}

func (m *PostgresqlConfig12) SetSeqPageCost(v *wrapperspb.DoubleValue) {
	m.SeqPageCost = v
}

func (m *PostgresqlConfig12) SetRandomPageCost(v *wrapperspb.DoubleValue) {
	m.RandomPageCost = v
}

func (m *PostgresqlConfig12) SetAutovacuumMaxWorkers(v *wrapperspb.Int64Value) {
	m.AutovacuumMaxWorkers = v
}

func (m *PostgresqlConfig12) SetAutovacuumVacuumCostDelay(v *wrapperspb.Int64Value) {
	m.AutovacuumVacuumCostDelay = v
}

func (m *PostgresqlConfig12) SetAutovacuumVacuumCostLimit(v *wrapperspb.Int64Value) {
	m.AutovacuumVacuumCostLimit = v
}

func (m *PostgresqlConfig12) SetAutovacuumNaptime(v *wrapperspb.Int64Value) {
	m.AutovacuumNaptime = v
}

func (m *PostgresqlConfig12) SetArchiveTimeout(v *wrapperspb.Int64Value) {
	m.ArchiveTimeout = v
}

func (m *PostgresqlConfig12) SetTrackActivityQuerySize(v *wrapperspb.Int64Value) {
	m.TrackActivityQuerySize = v
}

func (m *PostgresqlConfig12) SetEnableBitmapscan(v *wrapperspb.BoolValue) {
	m.EnableBitmapscan = v
}

func (m *PostgresqlConfig12) SetEnableHashagg(v *wrapperspb.BoolValue) {
	m.EnableHashagg = v
}

func (m *PostgresqlConfig12) SetEnableHashjoin(v *wrapperspb.BoolValue) {
	m.EnableHashjoin = v
}

func (m *PostgresqlConfig12) SetEnableIndexscan(v *wrapperspb.BoolValue) {
	m.EnableIndexscan = v
}

func (m *PostgresqlConfig12) SetEnableIndexonlyscan(v *wrapperspb.BoolValue) {
	m.EnableIndexonlyscan = v
}

func (m *PostgresqlConfig12) SetEnableMaterial(v *wrapperspb.BoolValue) {
	m.EnableMaterial = v
}

func (m *PostgresqlConfig12) SetEnableMergejoin(v *wrapperspb.BoolValue) {
	m.EnableMergejoin = v
}

func (m *PostgresqlConfig12) SetEnableNestloop(v *wrapperspb.BoolValue) {
	m.EnableNestloop = v
}

func (m *PostgresqlConfig12) SetEnableSeqscan(v *wrapperspb.BoolValue) {
	m.EnableSeqscan = v
}

func (m *PostgresqlConfig12) SetEnableSort(v *wrapperspb.BoolValue) {
	m.EnableSort = v
}

func (m *PostgresqlConfig12) SetEnableTidscan(v *wrapperspb.BoolValue) {
	m.EnableTidscan = v
}

func (m *PostgresqlConfig12) SetMaxWorkerProcesses(v *wrapperspb.Int64Value) {
	m.MaxWorkerProcesses = v
}

func (m *PostgresqlConfig12) SetMaxParallelWorkers(v *wrapperspb.Int64Value) {
	m.MaxParallelWorkers = v
}

func (m *PostgresqlConfig12) SetMaxParallelWorkersPerGather(v *wrapperspb.Int64Value) {
	m.MaxParallelWorkersPerGather = v
}

func (m *PostgresqlConfig12) SetAutovacuumVacuumScaleFactor(v *wrapperspb.DoubleValue) {
	m.AutovacuumVacuumScaleFactor = v
}

func (m *PostgresqlConfig12) SetAutovacuumAnalyzeScaleFactor(v *wrapperspb.DoubleValue) {
	m.AutovacuumAnalyzeScaleFactor = v
}

func (m *PostgresqlConfig12) SetDefaultTransactionReadOnly(v *wrapperspb.BoolValue) {
	m.DefaultTransactionReadOnly = v
}

func (m *PostgresqlConfig12) SetTimezone(v string) {
	m.Timezone = v
}

func (m *PostgresqlConfig12) SetEnableParallelAppend(v *wrapperspb.BoolValue) {
	m.EnableParallelAppend = v
}

func (m *PostgresqlConfig12) SetEnableParallelHash(v *wrapperspb.BoolValue) {
	m.EnableParallelHash = v
}

func (m *PostgresqlConfig12) SetEnablePartitionPruning(v *wrapperspb.BoolValue) {
	m.EnablePartitionPruning = v
}

func (m *PostgresqlConfig12) SetEnablePartitionwiseAggregate(v *wrapperspb.BoolValue) {
	m.EnablePartitionwiseAggregate = v
}

func (m *PostgresqlConfig12) SetEnablePartitionwiseJoin(v *wrapperspb.BoolValue) {
	m.EnablePartitionwiseJoin = v
}

func (m *PostgresqlConfig12) SetJit(v *wrapperspb.BoolValue) {
	m.Jit = v
}

func (m *PostgresqlConfig12) SetMaxParallelMaintenanceWorkers(v *wrapperspb.Int64Value) {
	m.MaxParallelMaintenanceWorkers = v
}

func (m *PostgresqlConfig12) SetParallelLeaderParticipation(v *wrapperspb.BoolValue) {
	m.ParallelLeaderParticipation = v
}

func (m *PostgresqlConfig12) SetVacuumCleanupIndexScaleFactor(v *wrapperspb.DoubleValue) {
	m.VacuumCleanupIndexScaleFactor = v
}

func (m *PostgresqlConfig12) SetLogTransactionSampleRate(v *wrapperspb.DoubleValue) {
	m.LogTransactionSampleRate = v
}

func (m *PostgresqlConfig12) SetPlanCacheMode(v PostgresqlConfig12_PlanCacheMode) {
	m.PlanCacheMode = v
}

func (m *PostgresqlConfig12) SetEffectiveIoConcurrency(v *wrapperspb.Int64Value) {
	m.EffectiveIoConcurrency = v
}

func (m *PostgresqlConfig12) SetEffectiveCacheSize(v *wrapperspb.Int64Value) {
	m.EffectiveCacheSize = v
}

func (m *PostgresqlConfig12) SetSharedPreloadLibraries(v []PostgresqlConfig12_SharedPreloadLibraries) {
	m.SharedPreloadLibraries = v
}

func (m *PostgresqlConfig12) SetAutoExplainLogMinDuration(v *wrapperspb.Int64Value) {
	m.AutoExplainLogMinDuration = v
}

func (m *PostgresqlConfig12) SetAutoExplainLogAnalyze(v *wrapperspb.BoolValue) {
	m.AutoExplainLogAnalyze = v
}

func (m *PostgresqlConfig12) SetAutoExplainLogBuffers(v *wrapperspb.BoolValue) {
	m.AutoExplainLogBuffers = v
}

func (m *PostgresqlConfig12) SetAutoExplainLogTiming(v *wrapperspb.BoolValue) {
	m.AutoExplainLogTiming = v
}

func (m *PostgresqlConfig12) SetAutoExplainLogTriggers(v *wrapperspb.BoolValue) {
	m.AutoExplainLogTriggers = v
}

func (m *PostgresqlConfig12) SetAutoExplainLogVerbose(v *wrapperspb.BoolValue) {
	m.AutoExplainLogVerbose = v
}

func (m *PostgresqlConfig12) SetAutoExplainLogNestedStatements(v *wrapperspb.BoolValue) {
	m.AutoExplainLogNestedStatements = v
}

func (m *PostgresqlConfig12) SetAutoExplainSampleRate(v *wrapperspb.DoubleValue) {
	m.AutoExplainSampleRate = v
}

func (m *PostgresqlConfig12) SetPgHintPlanEnableHint(v *wrapperspb.BoolValue) {
	m.PgHintPlanEnableHint = v
}

func (m *PostgresqlConfig12) SetPgHintPlanEnableHintTable(v *wrapperspb.BoolValue) {
	m.PgHintPlanEnableHintTable = v
}

func (m *PostgresqlConfig12) SetPgHintPlanDebugPrint(v PostgresqlConfig12_PgHintPlanDebugPrint) {
	m.PgHintPlanDebugPrint = v
}

func (m *PostgresqlConfig12) SetPgHintPlanMessageLevel(v PostgresqlConfig12_LogLevel) {
	m.PgHintPlanMessageLevel = v
}

func (m *PostgresqlConfig12) SetPgQualstatsEnabled(v *wrapperspb.BoolValue) {
	m.PgQualstatsEnabled = v
}

func (m *PostgresqlConfig12) SetPgQualstatsTrackConstants(v *wrapperspb.BoolValue) {
	m.PgQualstatsTrackConstants = v
}

func (m *PostgresqlConfig12) SetPgQualstatsMax(v *wrapperspb.Int64Value) {
	m.PgQualstatsMax = v
}

func (m *PostgresqlConfig12) SetPgQualstatsResolveOids(v *wrapperspb.BoolValue) {
	m.PgQualstatsResolveOids = v
}

func (m *PostgresqlConfig12) SetPgQualstatsSampleRate(v *wrapperspb.DoubleValue) {
	m.PgQualstatsSampleRate = v
}

func (m *PostgresqlConfig12) SetMaxStackDepth(v *wrapperspb.Int64Value) {
	m.MaxStackDepth = v
}

func (m *PostgresqlConfig12) SetGeqo(v *wrapperspb.BoolValue) {
	m.Geqo = v
}

func (m *PostgresqlConfig12) SetGeqoThreshold(v *wrapperspb.Int64Value) {
	m.GeqoThreshold = v
}

func (m *PostgresqlConfig12) SetGeqoEffort(v *wrapperspb.Int64Value) {
	m.GeqoEffort = v
}

func (m *PostgresqlConfig12) SetGeqoPoolSize(v *wrapperspb.Int64Value) {
	m.GeqoPoolSize = v
}

func (m *PostgresqlConfig12) SetGeqoGenerations(v *wrapperspb.Int64Value) {
	m.GeqoGenerations = v
}

func (m *PostgresqlConfig12) SetGeqoSelectionBias(v *wrapperspb.DoubleValue) {
	m.GeqoSelectionBias = v
}

func (m *PostgresqlConfig12) SetGeqoSeed(v *wrapperspb.DoubleValue) {
	m.GeqoSeed = v
}

func (m *PostgresqlConfig12) SetPgTrgmSimilarityThreshold(v *wrapperspb.DoubleValue) {
	m.PgTrgmSimilarityThreshold = v
}

func (m *PostgresqlConfig12) SetPgTrgmWordSimilarityThreshold(v *wrapperspb.DoubleValue) {
	m.PgTrgmWordSimilarityThreshold = v
}

func (m *PostgresqlConfig12) SetPgTrgmStrictWordSimilarityThreshold(v *wrapperspb.DoubleValue) {
	m.PgTrgmStrictWordSimilarityThreshold = v
}

func (m *PostgresqlConfig12) SetMaxStandbyArchiveDelay(v *wrapperspb.Int64Value) {
	m.MaxStandbyArchiveDelay = v
}

func (m *PostgresqlConfig12) SetSessionDurationTimeout(v *wrapperspb.Int64Value) {
	m.SessionDurationTimeout = v
}

func (m *PostgresqlConfig12) SetLogReplicationCommands(v *wrapperspb.BoolValue) {
	m.LogReplicationCommands = v
}

func (m *PostgresqlConfig12) SetLogAutovacuumMinDuration(v *wrapperspb.Int64Value) {
	m.LogAutovacuumMinDuration = v
}

func (m *PostgresqlConfigSet12) SetEffectiveConfig(v *PostgresqlConfig12) {
	m.EffectiveConfig = v
}

func (m *PostgresqlConfigSet12) SetUserConfig(v *PostgresqlConfig12) {
	m.UserConfig = v
}

func (m *PostgresqlConfigSet12) SetDefaultConfig(v *PostgresqlConfig12) {
	m.DefaultConfig = v
}
