// Code generated by protoc-gen-goext. DO NOT EDIT.

package mongodb

import (
	config "github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/mongodb/v1/config"
	timeofday "google.golang.org/genproto/googleapis/type/timeofday"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

func (m *Cluster) SetId(v string) {
	m.Id = v
}

func (m *Cluster) SetFolderId(v string) {
	m.FolderId = v
}

func (m *Cluster) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}

func (m *Cluster) SetName(v string) {
	m.Name = v
}

func (m *Cluster) SetDescription(v string) {
	m.Description = v
}

func (m *Cluster) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *Cluster) SetEnvironment(v Cluster_Environment) {
	m.Environment = v
}

func (m *Cluster) SetMonitoring(v []*Monitoring) {
	m.Monitoring = v
}

func (m *Cluster) SetConfig(v *ClusterConfig) {
	m.Config = v
}

func (m *Cluster) SetNetworkId(v string) {
	m.NetworkId = v
}

func (m *Cluster) SetHealth(v Cluster_Health) {
	m.Health = v
}

func (m *Cluster) SetStatus(v Cluster_Status) {
	m.Status = v
}

func (m *Cluster) SetSharded(v bool) {
	m.Sharded = v
}

func (m *Cluster) SetMaintenanceWindow(v *MaintenanceWindow) {
	m.MaintenanceWindow = v
}

func (m *Cluster) SetPlannedOperation(v *MaintenanceOperation) {
	m.PlannedOperation = v
}

func (m *Cluster) SetSecurityGroupIds(v []string) {
	m.SecurityGroupIds = v
}

func (m *Cluster) SetDeletionProtection(v bool) {
	m.DeletionProtection = v
}

func (m *Monitoring) SetName(v string) {
	m.Name = v
}

func (m *Monitoring) SetDescription(v string) {
	m.Description = v
}

func (m *Monitoring) SetLink(v string) {
	m.Link = v
}

type ClusterConfig_Mongodb = isClusterConfig_Mongodb

func (m *ClusterConfig) SetMongodb(v ClusterConfig_Mongodb) {
	m.Mongodb = v
}

func (m *ClusterConfig) SetVersion(v string) {
	m.Version = v
}

func (m *ClusterConfig) SetFeatureCompatibilityVersion(v string) {
	m.FeatureCompatibilityVersion = v
}

func (m *ClusterConfig) SetMongodb_3_6(v *Mongodb3_6) {
	m.Mongodb = &ClusterConfig_Mongodb_3_6{
		Mongodb_3_6: v,
	}
}

func (m *ClusterConfig) SetMongodb_4_0(v *Mongodb4_0) {
	m.Mongodb = &ClusterConfig_Mongodb_4_0{
		Mongodb_4_0: v,
	}
}

func (m *ClusterConfig) SetMongodb_4_2(v *Mongodb4_2) {
	m.Mongodb = &ClusterConfig_Mongodb_4_2{
		Mongodb_4_2: v,
	}
}

func (m *ClusterConfig) SetMongodb_4_4(v *Mongodb4_4) {
	m.Mongodb = &ClusterConfig_Mongodb_4_4{
		Mongodb_4_4: v,
	}
}

func (m *ClusterConfig) SetMongodb_5_0(v *Mongodb5_0) {
	m.Mongodb = &ClusterConfig_Mongodb_5_0{
		Mongodb_5_0: v,
	}
}

func (m *ClusterConfig) SetMongodb_6_0(v *Mongodb6_0) {
	m.Mongodb = &ClusterConfig_Mongodb_6_0{
		Mongodb_6_0: v,
	}
}

func (m *ClusterConfig) SetMongodb_4_4Enterprise(v *Mongodb4_4Enterprise) {
	m.Mongodb = &ClusterConfig_Mongodb_4_4Enterprise{
		Mongodb_4_4Enterprise: v,
	}
}

func (m *ClusterConfig) SetMongodb_5_0Enterprise(v *Mongodb5_0Enterprise) {
	m.Mongodb = &ClusterConfig_Mongodb_5_0Enterprise{
		Mongodb_5_0Enterprise: v,
	}
}

func (m *ClusterConfig) SetMongodb_6_0Enterprise(v *Mongodb6_0Enterprise) {
	m.Mongodb = &ClusterConfig_Mongodb_6_0Enterprise{
		Mongodb_6_0Enterprise: v,
	}
}

func (m *ClusterConfig) SetBackupWindowStart(v *timeofday.TimeOfDay) {
	m.BackupWindowStart = v
}

func (m *ClusterConfig) SetBackupRetainPeriodDays(v *wrapperspb.Int64Value) {
	m.BackupRetainPeriodDays = v
}

func (m *ClusterConfig) SetPerformanceDiagnostics(v *PerformanceDiagnosticsConfig) {
	m.PerformanceDiagnostics = v
}

func (m *ClusterConfig) SetAccess(v *Access) {
	m.Access = v
}

func (m *ClusterConfig) SetMongodbConfig(v *Mongodb) {
	m.MongodbConfig = v
}

func (m *Mongodb3_6) SetMongod(v *Mongodb3_6_Mongod) {
	m.Mongod = v
}

func (m *Mongodb3_6) SetMongocfg(v *Mongodb3_6_MongoCfg) {
	m.Mongocfg = v
}

func (m *Mongodb3_6) SetMongos(v *Mongodb3_6_Mongos) {
	m.Mongos = v
}

func (m *Mongodb3_6) SetMongoinfra(v *Mongodb3_6_MongoInfra) {
	m.Mongoinfra = v
}

func (m *Mongodb3_6_Mongod) SetConfig(v *config.MongodConfigSet3_6) {
	m.Config = v
}

func (m *Mongodb3_6_Mongod) SetResources(v *Resources) {
	m.Resources = v
}

func (m *Mongodb3_6_Mongod) SetDiskSizeAutoscaling(v *DiskSizeAutoscaling) {
	m.DiskSizeAutoscaling = v
}

func (m *Mongodb3_6_MongoCfg) SetConfig(v *config.MongoCfgConfigSet3_6) {
	m.Config = v
}

func (m *Mongodb3_6_MongoCfg) SetResources(v *Resources) {
	m.Resources = v
}

func (m *Mongodb3_6_MongoCfg) SetDiskSizeAutoscaling(v *DiskSizeAutoscaling) {
	m.DiskSizeAutoscaling = v
}

func (m *Mongodb3_6_Mongos) SetConfig(v *config.MongosConfigSet3_6) {
	m.Config = v
}

func (m *Mongodb3_6_Mongos) SetResources(v *Resources) {
	m.Resources = v
}

func (m *Mongodb3_6_Mongos) SetDiskSizeAutoscaling(v *DiskSizeAutoscaling) {
	m.DiskSizeAutoscaling = v
}

func (m *Mongodb3_6_MongoInfra) SetConfigMongos(v *config.MongosConfigSet3_6) {
	m.ConfigMongos = v
}

func (m *Mongodb3_6_MongoInfra) SetConfigMongocfg(v *config.MongoCfgConfigSet3_6) {
	m.ConfigMongocfg = v
}

func (m *Mongodb3_6_MongoInfra) SetResources(v *Resources) {
	m.Resources = v
}

func (m *Mongodb3_6_MongoInfra) SetDiskSizeAutoscaling(v *DiskSizeAutoscaling) {
	m.DiskSizeAutoscaling = v
}

func (m *Mongodb4_0) SetMongod(v *Mongodb4_0_Mongod) {
	m.Mongod = v
}

func (m *Mongodb4_0) SetMongocfg(v *Mongodb4_0_MongoCfg) {
	m.Mongocfg = v
}

func (m *Mongodb4_0) SetMongos(v *Mongodb4_0_Mongos) {
	m.Mongos = v
}

func (m *Mongodb4_0) SetMongoinfra(v *Mongodb4_0_MongoInfra) {
	m.Mongoinfra = v
}

func (m *Mongodb4_0_Mongod) SetConfig(v *config.MongodConfigSet4_0) {
	m.Config = v
}

func (m *Mongodb4_0_Mongod) SetResources(v *Resources) {
	m.Resources = v
}

func (m *Mongodb4_0_Mongod) SetDiskSizeAutoscaling(v *DiskSizeAutoscaling) {
	m.DiskSizeAutoscaling = v
}

func (m *Mongodb4_0_MongoCfg) SetConfig(v *config.MongoCfgConfigSet4_0) {
	m.Config = v
}

func (m *Mongodb4_0_MongoCfg) SetResources(v *Resources) {
	m.Resources = v
}

func (m *Mongodb4_0_MongoCfg) SetDiskSizeAutoscaling(v *DiskSizeAutoscaling) {
	m.DiskSizeAutoscaling = v
}

func (m *Mongodb4_0_Mongos) SetConfig(v *config.MongosConfigSet4_0) {
	m.Config = v
}

func (m *Mongodb4_0_Mongos) SetResources(v *Resources) {
	m.Resources = v
}

func (m *Mongodb4_0_Mongos) SetDiskSizeAutoscaling(v *DiskSizeAutoscaling) {
	m.DiskSizeAutoscaling = v
}

func (m *Mongodb4_0_MongoInfra) SetConfigMongos(v *config.MongosConfigSet4_0) {
	m.ConfigMongos = v
}

func (m *Mongodb4_0_MongoInfra) SetConfigMongocfg(v *config.MongoCfgConfigSet4_0) {
	m.ConfigMongocfg = v
}

func (m *Mongodb4_0_MongoInfra) SetResources(v *Resources) {
	m.Resources = v
}

func (m *Mongodb4_0_MongoInfra) SetDiskSizeAutoscaling(v *DiskSizeAutoscaling) {
	m.DiskSizeAutoscaling = v
}

func (m *Mongodb4_2) SetMongod(v *Mongodb4_2_Mongod) {
	m.Mongod = v
}

func (m *Mongodb4_2) SetMongocfg(v *Mongodb4_2_MongoCfg) {
	m.Mongocfg = v
}

func (m *Mongodb4_2) SetMongos(v *Mongodb4_2_Mongos) {
	m.Mongos = v
}

func (m *Mongodb4_2) SetMongoinfra(v *Mongodb4_2_MongoInfra) {
	m.Mongoinfra = v
}

func (m *Mongodb4_2_Mongod) SetConfig(v *config.MongodConfigSet4_2) {
	m.Config = v
}

func (m *Mongodb4_2_Mongod) SetResources(v *Resources) {
	m.Resources = v
}

func (m *Mongodb4_2_Mongod) SetDiskSizeAutoscaling(v *DiskSizeAutoscaling) {
	m.DiskSizeAutoscaling = v
}

func (m *Mongodb4_2_MongoCfg) SetConfig(v *config.MongoCfgConfigSet4_2) {
	m.Config = v
}

func (m *Mongodb4_2_MongoCfg) SetResources(v *Resources) {
	m.Resources = v
}

func (m *Mongodb4_2_MongoCfg) SetDiskSizeAutoscaling(v *DiskSizeAutoscaling) {
	m.DiskSizeAutoscaling = v
}

func (m *Mongodb4_2_Mongos) SetConfig(v *config.MongosConfigSet4_2) {
	m.Config = v
}

func (m *Mongodb4_2_Mongos) SetResources(v *Resources) {
	m.Resources = v
}

func (m *Mongodb4_2_Mongos) SetDiskSizeAutoscaling(v *DiskSizeAutoscaling) {
	m.DiskSizeAutoscaling = v
}

func (m *Mongodb4_2_MongoInfra) SetConfigMongos(v *config.MongosConfigSet4_2) {
	m.ConfigMongos = v
}

func (m *Mongodb4_2_MongoInfra) SetConfigMongocfg(v *config.MongoCfgConfigSet4_2) {
	m.ConfigMongocfg = v
}

func (m *Mongodb4_2_MongoInfra) SetResources(v *Resources) {
	m.Resources = v
}

func (m *Mongodb4_2_MongoInfra) SetDiskSizeAutoscaling(v *DiskSizeAutoscaling) {
	m.DiskSizeAutoscaling = v
}

func (m *Mongodb4_4) SetMongod(v *Mongodb4_4_Mongod) {
	m.Mongod = v
}

func (m *Mongodb4_4) SetMongocfg(v *Mongodb4_4_MongoCfg) {
	m.Mongocfg = v
}

func (m *Mongodb4_4) SetMongos(v *Mongodb4_4_Mongos) {
	m.Mongos = v
}

func (m *Mongodb4_4) SetMongoinfra(v *Mongodb4_4_MongoInfra) {
	m.Mongoinfra = v
}

func (m *Mongodb4_4_Mongod) SetConfig(v *config.MongodConfigSet4_4) {
	m.Config = v
}

func (m *Mongodb4_4_Mongod) SetResources(v *Resources) {
	m.Resources = v
}

func (m *Mongodb4_4_Mongod) SetDiskSizeAutoscaling(v *DiskSizeAutoscaling) {
	m.DiskSizeAutoscaling = v
}

func (m *Mongodb4_4_MongoCfg) SetConfig(v *config.MongoCfgConfigSet4_4) {
	m.Config = v
}

func (m *Mongodb4_4_MongoCfg) SetResources(v *Resources) {
	m.Resources = v
}

func (m *Mongodb4_4_MongoCfg) SetDiskSizeAutoscaling(v *DiskSizeAutoscaling) {
	m.DiskSizeAutoscaling = v
}

func (m *Mongodb4_4_Mongos) SetConfig(v *config.MongosConfigSet4_4) {
	m.Config = v
}

func (m *Mongodb4_4_Mongos) SetResources(v *Resources) {
	m.Resources = v
}

func (m *Mongodb4_4_Mongos) SetDiskSizeAutoscaling(v *DiskSizeAutoscaling) {
	m.DiskSizeAutoscaling = v
}

func (m *Mongodb4_4_MongoInfra) SetConfigMongos(v *config.MongosConfigSet4_4) {
	m.ConfigMongos = v
}

func (m *Mongodb4_4_MongoInfra) SetConfigMongocfg(v *config.MongoCfgConfigSet4_4) {
	m.ConfigMongocfg = v
}

func (m *Mongodb4_4_MongoInfra) SetResources(v *Resources) {
	m.Resources = v
}

func (m *Mongodb4_4_MongoInfra) SetDiskSizeAutoscaling(v *DiskSizeAutoscaling) {
	m.DiskSizeAutoscaling = v
}

func (m *Mongodb4_4Enterprise) SetMongod(v *Mongodb4_4Enterprise_Mongod) {
	m.Mongod = v
}

func (m *Mongodb4_4Enterprise) SetMongocfg(v *Mongodb4_4Enterprise_MongoCfg) {
	m.Mongocfg = v
}

func (m *Mongodb4_4Enterprise) SetMongos(v *Mongodb4_4Enterprise_Mongos) {
	m.Mongos = v
}

func (m *Mongodb4_4Enterprise) SetMongoinfra(v *Mongodb4_4Enterprise_MongoInfra) {
	m.Mongoinfra = v
}

func (m *Mongodb4_4Enterprise_Mongod) SetConfig(v *config.MongodConfigSet4_4Enterprise) {
	m.Config = v
}

func (m *Mongodb4_4Enterprise_Mongod) SetResources(v *Resources) {
	m.Resources = v
}

func (m *Mongodb4_4Enterprise_Mongod) SetDiskSizeAutoscaling(v *DiskSizeAutoscaling) {
	m.DiskSizeAutoscaling = v
}

func (m *Mongodb4_4Enterprise_MongoCfg) SetConfig(v *config.MongoCfgConfigSet4_4Enterprise) {
	m.Config = v
}

func (m *Mongodb4_4Enterprise_MongoCfg) SetResources(v *Resources) {
	m.Resources = v
}

func (m *Mongodb4_4Enterprise_MongoCfg) SetDiskSizeAutoscaling(v *DiskSizeAutoscaling) {
	m.DiskSizeAutoscaling = v
}

func (m *Mongodb4_4Enterprise_Mongos) SetConfig(v *config.MongosConfigSet4_4Enterprise) {
	m.Config = v
}

func (m *Mongodb4_4Enterprise_Mongos) SetResources(v *Resources) {
	m.Resources = v
}

func (m *Mongodb4_4Enterprise_Mongos) SetDiskSizeAutoscaling(v *DiskSizeAutoscaling) {
	m.DiskSizeAutoscaling = v
}

func (m *Mongodb4_4Enterprise_MongoInfra) SetConfigMongos(v *config.MongosConfigSet4_4Enterprise) {
	m.ConfigMongos = v
}

func (m *Mongodb4_4Enterprise_MongoInfra) SetConfigMongocfg(v *config.MongoCfgConfigSet4_4Enterprise) {
	m.ConfigMongocfg = v
}

func (m *Mongodb4_4Enterprise_MongoInfra) SetResources(v *Resources) {
	m.Resources = v
}

func (m *Mongodb4_4Enterprise_MongoInfra) SetDiskSizeAutoscaling(v *DiskSizeAutoscaling) {
	m.DiskSizeAutoscaling = v
}

func (m *Mongodb5_0) SetMongod(v *Mongodb5_0_Mongod) {
	m.Mongod = v
}

func (m *Mongodb5_0) SetMongocfg(v *Mongodb5_0_MongoCfg) {
	m.Mongocfg = v
}

func (m *Mongodb5_0) SetMongos(v *Mongodb5_0_Mongos) {
	m.Mongos = v
}

func (m *Mongodb5_0) SetMongoinfra(v *Mongodb5_0_MongoInfra) {
	m.Mongoinfra = v
}

func (m *Mongodb5_0_Mongod) SetConfig(v *config.MongodConfigSet5_0) {
	m.Config = v
}

func (m *Mongodb5_0_Mongod) SetResources(v *Resources) {
	m.Resources = v
}

func (m *Mongodb5_0_Mongod) SetDiskSizeAutoscaling(v *DiskSizeAutoscaling) {
	m.DiskSizeAutoscaling = v
}

func (m *Mongodb5_0_MongoCfg) SetConfig(v *config.MongoCfgConfigSet5_0) {
	m.Config = v
}

func (m *Mongodb5_0_MongoCfg) SetResources(v *Resources) {
	m.Resources = v
}

func (m *Mongodb5_0_MongoCfg) SetDiskSizeAutoscaling(v *DiskSizeAutoscaling) {
	m.DiskSizeAutoscaling = v
}

func (m *Mongodb5_0_Mongos) SetConfig(v *config.MongosConfigSet5_0) {
	m.Config = v
}

func (m *Mongodb5_0_Mongos) SetResources(v *Resources) {
	m.Resources = v
}

func (m *Mongodb5_0_Mongos) SetDiskSizeAutoscaling(v *DiskSizeAutoscaling) {
	m.DiskSizeAutoscaling = v
}

func (m *Mongodb5_0_MongoInfra) SetConfigMongos(v *config.MongosConfigSet5_0) {
	m.ConfigMongos = v
}

func (m *Mongodb5_0_MongoInfra) SetConfigMongocfg(v *config.MongoCfgConfigSet5_0) {
	m.ConfigMongocfg = v
}

func (m *Mongodb5_0_MongoInfra) SetResources(v *Resources) {
	m.Resources = v
}

func (m *Mongodb5_0_MongoInfra) SetDiskSizeAutoscaling(v *DiskSizeAutoscaling) {
	m.DiskSizeAutoscaling = v
}

func (m *Mongodb5_0Enterprise) SetMongod(v *Mongodb5_0Enterprise_Mongod) {
	m.Mongod = v
}

func (m *Mongodb5_0Enterprise) SetMongocfg(v *Mongodb5_0Enterprise_MongoCfg) {
	m.Mongocfg = v
}

func (m *Mongodb5_0Enterprise) SetMongos(v *Mongodb5_0Enterprise_Mongos) {
	m.Mongos = v
}

func (m *Mongodb5_0Enterprise) SetMongoinfra(v *Mongodb5_0Enterprise_MongoInfra) {
	m.Mongoinfra = v
}

func (m *Mongodb5_0Enterprise_Mongod) SetConfig(v *config.MongodConfigSet5_0Enterprise) {
	m.Config = v
}

func (m *Mongodb5_0Enterprise_Mongod) SetResources(v *Resources) {
	m.Resources = v
}

func (m *Mongodb5_0Enterprise_Mongod) SetDiskSizeAutoscaling(v *DiskSizeAutoscaling) {
	m.DiskSizeAutoscaling = v
}

func (m *Mongodb5_0Enterprise_MongoCfg) SetConfig(v *config.MongoCfgConfigSet5_0Enterprise) {
	m.Config = v
}

func (m *Mongodb5_0Enterprise_MongoCfg) SetResources(v *Resources) {
	m.Resources = v
}

func (m *Mongodb5_0Enterprise_MongoCfg) SetDiskSizeAutoscaling(v *DiskSizeAutoscaling) {
	m.DiskSizeAutoscaling = v
}

func (m *Mongodb5_0Enterprise_Mongos) SetConfig(v *config.MongosConfigSet5_0Enterprise) {
	m.Config = v
}

func (m *Mongodb5_0Enterprise_Mongos) SetResources(v *Resources) {
	m.Resources = v
}

func (m *Mongodb5_0Enterprise_Mongos) SetDiskSizeAutoscaling(v *DiskSizeAutoscaling) {
	m.DiskSizeAutoscaling = v
}

func (m *Mongodb5_0Enterprise_MongoInfra) SetConfigMongos(v *config.MongosConfigSet5_0Enterprise) {
	m.ConfigMongos = v
}

func (m *Mongodb5_0Enterprise_MongoInfra) SetConfigMongocfg(v *config.MongoCfgConfigSet5_0Enterprise) {
	m.ConfigMongocfg = v
}

func (m *Mongodb5_0Enterprise_MongoInfra) SetResources(v *Resources) {
	m.Resources = v
}

func (m *Mongodb5_0Enterprise_MongoInfra) SetDiskSizeAutoscaling(v *DiskSizeAutoscaling) {
	m.DiskSizeAutoscaling = v
}

func (m *Mongodb6_0) SetMongod(v *Mongodb6_0_Mongod) {
	m.Mongod = v
}

func (m *Mongodb6_0) SetMongocfg(v *Mongodb6_0_MongoCfg) {
	m.Mongocfg = v
}

func (m *Mongodb6_0) SetMongos(v *Mongodb6_0_Mongos) {
	m.Mongos = v
}

func (m *Mongodb6_0) SetMongoinfra(v *Mongodb6_0_MongoInfra) {
	m.Mongoinfra = v
}

func (m *Mongodb6_0_Mongod) SetConfig(v *config.MongodConfigSet6_0) {
	m.Config = v
}

func (m *Mongodb6_0_Mongod) SetResources(v *Resources) {
	m.Resources = v
}

func (m *Mongodb6_0_Mongod) SetDiskSizeAutoscaling(v *DiskSizeAutoscaling) {
	m.DiskSizeAutoscaling = v
}

func (m *Mongodb6_0_MongoCfg) SetConfig(v *config.MongoCfgConfigSet6_0) {
	m.Config = v
}

func (m *Mongodb6_0_MongoCfg) SetResources(v *Resources) {
	m.Resources = v
}

func (m *Mongodb6_0_MongoCfg) SetDiskSizeAutoscaling(v *DiskSizeAutoscaling) {
	m.DiskSizeAutoscaling = v
}

func (m *Mongodb6_0_Mongos) SetConfig(v *config.MongosConfigSet6_0) {
	m.Config = v
}

func (m *Mongodb6_0_Mongos) SetResources(v *Resources) {
	m.Resources = v
}

func (m *Mongodb6_0_Mongos) SetDiskSizeAutoscaling(v *DiskSizeAutoscaling) {
	m.DiskSizeAutoscaling = v
}

func (m *Mongodb6_0_MongoInfra) SetConfigMongos(v *config.MongosConfigSet6_0) {
	m.ConfigMongos = v
}

func (m *Mongodb6_0_MongoInfra) SetConfigMongocfg(v *config.MongoCfgConfigSet6_0) {
	m.ConfigMongocfg = v
}

func (m *Mongodb6_0_MongoInfra) SetResources(v *Resources) {
	m.Resources = v
}

func (m *Mongodb6_0_MongoInfra) SetDiskSizeAutoscaling(v *DiskSizeAutoscaling) {
	m.DiskSizeAutoscaling = v
}

func (m *Mongodb6_0Enterprise) SetMongod(v *Mongodb6_0Enterprise_Mongod) {
	m.Mongod = v
}

func (m *Mongodb6_0Enterprise) SetMongocfg(v *Mongodb6_0Enterprise_MongoCfg) {
	m.Mongocfg = v
}

func (m *Mongodb6_0Enterprise) SetMongos(v *Mongodb6_0Enterprise_Mongos) {
	m.Mongos = v
}

func (m *Mongodb6_0Enterprise) SetMongoinfra(v *Mongodb6_0Enterprise_MongoInfra) {
	m.Mongoinfra = v
}

func (m *Mongodb6_0Enterprise_Mongod) SetConfig(v *config.MongodConfigSet6_0Enterprise) {
	m.Config = v
}

func (m *Mongodb6_0Enterprise_Mongod) SetResources(v *Resources) {
	m.Resources = v
}

func (m *Mongodb6_0Enterprise_Mongod) SetDiskSizeAutoscaling(v *DiskSizeAutoscaling) {
	m.DiskSizeAutoscaling = v
}

func (m *Mongodb6_0Enterprise_MongoCfg) SetConfig(v *config.MongoCfgConfigSet6_0Enterprise) {
	m.Config = v
}

func (m *Mongodb6_0Enterprise_MongoCfg) SetResources(v *Resources) {
	m.Resources = v
}

func (m *Mongodb6_0Enterprise_MongoCfg) SetDiskSizeAutoscaling(v *DiskSizeAutoscaling) {
	m.DiskSizeAutoscaling = v
}

func (m *Mongodb6_0Enterprise_Mongos) SetConfig(v *config.MongosConfigSet6_0Enterprise) {
	m.Config = v
}

func (m *Mongodb6_0Enterprise_Mongos) SetResources(v *Resources) {
	m.Resources = v
}

func (m *Mongodb6_0Enterprise_Mongos) SetDiskSizeAutoscaling(v *DiskSizeAutoscaling) {
	m.DiskSizeAutoscaling = v
}

func (m *Mongodb6_0Enterprise_MongoInfra) SetConfigMongos(v *config.MongosConfigSet6_0Enterprise) {
	m.ConfigMongos = v
}

func (m *Mongodb6_0Enterprise_MongoInfra) SetConfigMongocfg(v *config.MongoCfgConfigSet6_0Enterprise) {
	m.ConfigMongocfg = v
}

func (m *Mongodb6_0Enterprise_MongoInfra) SetResources(v *Resources) {
	m.Resources = v
}

func (m *Mongodb6_0Enterprise_MongoInfra) SetDiskSizeAutoscaling(v *DiskSizeAutoscaling) {
	m.DiskSizeAutoscaling = v
}

func (m *Mongodb) SetMongod(v *Mongodb_Mongod) {
	m.Mongod = v
}

func (m *Mongodb) SetMongocfg(v *Mongodb_MongoCfg) {
	m.Mongocfg = v
}

func (m *Mongodb) SetMongos(v *Mongodb_Mongos) {
	m.Mongos = v
}

func (m *Mongodb) SetMongoinfra(v *Mongodb_MongoInfra) {
	m.Mongoinfra = v
}

func (m *Mongodb_Mongod) SetConfig(v *config.MongodConfigSet) {
	m.Config = v
}

func (m *Mongodb_Mongod) SetResources(v *Resources) {
	m.Resources = v
}

func (m *Mongodb_Mongod) SetDiskSizeAutoscaling(v *DiskSizeAutoscaling) {
	m.DiskSizeAutoscaling = v
}

func (m *Mongodb_MongoCfg) SetConfig(v *config.MongoCfgConfigSet) {
	m.Config = v
}

func (m *Mongodb_MongoCfg) SetResources(v *Resources) {
	m.Resources = v
}

func (m *Mongodb_MongoCfg) SetDiskSizeAutoscaling(v *DiskSizeAutoscaling) {
	m.DiskSizeAutoscaling = v
}

func (m *Mongodb_Mongos) SetConfig(v *config.MongosConfigSet) {
	m.Config = v
}

func (m *Mongodb_Mongos) SetResources(v *Resources) {
	m.Resources = v
}

func (m *Mongodb_Mongos) SetDiskSizeAutoscaling(v *DiskSizeAutoscaling) {
	m.DiskSizeAutoscaling = v
}

func (m *Mongodb_MongoInfra) SetConfigMongos(v *config.MongosConfigSet) {
	m.ConfigMongos = v
}

func (m *Mongodb_MongoInfra) SetConfigMongocfg(v *config.MongoCfgConfigSet) {
	m.ConfigMongocfg = v
}

func (m *Mongodb_MongoInfra) SetResources(v *Resources) {
	m.Resources = v
}

func (m *Mongodb_MongoInfra) SetDiskSizeAutoscaling(v *DiskSizeAutoscaling) {
	m.DiskSizeAutoscaling = v
}

func (m *Shard) SetName(v string) {
	m.Name = v
}

func (m *Shard) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *Host) SetName(v string) {
	m.Name = v
}

func (m *Host) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *Host) SetZoneId(v string) {
	m.ZoneId = v
}

func (m *Host) SetResources(v *Resources) {
	m.Resources = v
}

func (m *Host) SetRole(v Host_Role) {
	m.Role = v
}

func (m *Host) SetHealth(v Host_Health) {
	m.Health = v
}

func (m *Host) SetServices(v []*Service) {
	m.Services = v
}

func (m *Host) SetSubnetId(v string) {
	m.SubnetId = v
}

func (m *Host) SetAssignPublicIp(v bool) {
	m.AssignPublicIp = v
}

func (m *Host) SetShardName(v string) {
	m.ShardName = v
}

func (m *Host) SetType(v Host_Type) {
	m.Type = v
}

func (m *Host) SetHostParameters(v *Host_HostParameters) {
	m.HostParameters = v
}

func (m *Host_HostParameters) SetHidden(v bool) {
	m.Hidden = v
}

func (m *Host_HostParameters) SetSecondaryDelaySecs(v int64) {
	m.SecondaryDelaySecs = v
}

func (m *Host_HostParameters) SetPriority(v float64) {
	m.Priority = v
}

func (m *Host_HostParameters) SetTags(v map[string]string) {
	m.Tags = v
}

func (m *Service) SetType(v Service_Type) {
	m.Type = v
}

func (m *Service) SetHealth(v Service_Health) {
	m.Health = v
}

func (m *Resources) SetResourcePresetId(v string) {
	m.ResourcePresetId = v
}

func (m *Resources) SetDiskSize(v int64) {
	m.DiskSize = v
}

func (m *Resources) SetDiskTypeId(v string) {
	m.DiskTypeId = v
}

func (m *Access) SetDataLens(v bool) {
	m.DataLens = v
}

func (m *Access) SetWebSql(v bool) {
	m.WebSql = v
}

func (m *Access) SetDataTransfer(v bool) {
	m.DataTransfer = v
}

func (m *PerformanceDiagnosticsConfig) SetProfilingEnabled(v bool) {
	m.ProfilingEnabled = v
}

func (m *DiskSizeAutoscaling) SetPlannedUsageThreshold(v *wrapperspb.Int64Value) {
	m.PlannedUsageThreshold = v
}

func (m *DiskSizeAutoscaling) SetEmergencyUsageThreshold(v *wrapperspb.Int64Value) {
	m.EmergencyUsageThreshold = v
}

func (m *DiskSizeAutoscaling) SetDiskSizeLimit(v *wrapperspb.Int64Value) {
	m.DiskSizeLimit = v
}
