// Code generated by protoc-gen-goext. DO NOT EDIT.

package opensearch

import (
	config "github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/opensearch/v1/config"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

func (m *Cluster) SetSecurityGroupIds(v []string) {
	m.SecurityGroupIds = v
}

func (m *Cluster) SetServiceAccountId(v string) {
	m.ServiceAccountId = v
}

func (m *Cluster) SetDeletionProtection(v bool) {
	m.DeletionProtection = v
}

func (m *Cluster) SetMaintenanceWindow(v *MaintenanceWindow) {
	m.MaintenanceWindow = v
}

func (m *Cluster) SetPlannedOperation(v *MaintenanceOperation) {
	m.PlannedOperation = v
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

func (m *ClusterConfig) SetVersion(v string) {
	m.Version = v
}

func (m *ClusterConfig) SetOpensearch(v *OpenSearch) {
	m.Opensearch = v
}

func (m *ClusterConfig) SetDashboards(v *Dashboards) {
	m.Dashboards = v
}

func (m *ClusterConfig) SetAccess(v *Access) {
	m.Access = v
}

type OpenSearch_Config = isOpenSearch_Config

func (m *OpenSearch) SetConfig(v OpenSearch_Config) {
	m.Config = v
}

func (m *OpenSearch) SetPlugins(v []string) {
	m.Plugins = v
}

func (m *OpenSearch) SetNodeGroups(v []*OpenSearch_NodeGroup) {
	m.NodeGroups = v
}

func (m *OpenSearch) SetOpensearchConfigSet_2(v *config.OpenSearchConfigSet2) {
	m.Config = &OpenSearch_OpensearchConfigSet_2{
		OpensearchConfigSet_2: v,
	}
}

func (m *OpenSearch_NodeGroup) SetName(v string) {
	m.Name = v
}

func (m *OpenSearch_NodeGroup) SetResources(v *Resources) {
	m.Resources = v
}

func (m *OpenSearch_NodeGroup) SetHostsCount(v int64) {
	m.HostsCount = v
}

func (m *OpenSearch_NodeGroup) SetZoneIds(v []string) {
	m.ZoneIds = v
}

func (m *OpenSearch_NodeGroup) SetSubnetIds(v []string) {
	m.SubnetIds = v
}

func (m *OpenSearch_NodeGroup) SetAssignPublicIp(v bool) {
	m.AssignPublicIp = v
}

func (m *OpenSearch_NodeGroup) SetRoles(v []OpenSearch_GroupRole) {
	m.Roles = v
}

func (m *Dashboards) SetNodeGroups(v []*Dashboards_NodeGroup) {
	m.NodeGroups = v
}

func (m *Dashboards_NodeGroup) SetName(v string) {
	m.Name = v
}

func (m *Dashboards_NodeGroup) SetResources(v *Resources) {
	m.Resources = v
}

func (m *Dashboards_NodeGroup) SetHostsCount(v int64) {
	m.HostsCount = v
}

func (m *Dashboards_NodeGroup) SetZoneIds(v []string) {
	m.ZoneIds = v
}

func (m *Dashboards_NodeGroup) SetSubnetIds(v []string) {
	m.SubnetIds = v
}

func (m *Dashboards_NodeGroup) SetAssignPublicIp(v bool) {
	m.AssignPublicIp = v
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

func (m *Host) SetType(v Host_Type) {
	m.Type = v
}

func (m *Host) SetHealth(v Host_Health) {
	m.Health = v
}

func (m *Host) SetSubnetId(v string) {
	m.SubnetId = v
}

func (m *Host) SetAssignPublicIp(v bool) {
	m.AssignPublicIp = v
}

func (m *Host) SetSystem(v *Host_SystemMetrics) {
	m.System = v
}

func (m *Host) SetNodeGroup(v string) {
	m.NodeGroup = v
}

func (m *Host) SetRoles(v []OpenSearch_GroupRole) {
	m.Roles = v
}

func (m *Host_CPUMetric) SetTimestamp(v int64) {
	m.Timestamp = v
}

func (m *Host_CPUMetric) SetUsed(v float64) {
	m.Used = v
}

func (m *Host_MemoryMetric) SetTimestamp(v int64) {
	m.Timestamp = v
}

func (m *Host_MemoryMetric) SetUsed(v int64) {
	m.Used = v
}

func (m *Host_MemoryMetric) SetTotal(v int64) {
	m.Total = v
}

func (m *Host_DiskMetric) SetTimestamp(v int64) {
	m.Timestamp = v
}

func (m *Host_DiskMetric) SetUsed(v int64) {
	m.Used = v
}

func (m *Host_DiskMetric) SetTotal(v int64) {
	m.Total = v
}

func (m *Host_SystemMetrics) SetCpu(v *Host_CPUMetric) {
	m.Cpu = v
}

func (m *Host_SystemMetrics) SetMemory(v *Host_MemoryMetric) {
	m.Memory = v
}

func (m *Host_SystemMetrics) SetDisk(v *Host_DiskMetric) {
	m.Disk = v
}

func (m *Access) SetDataTransfer(v bool) {
	m.DataTransfer = v
}

func (m *Access) SetServerless(v bool) {
	m.Serverless = v
}