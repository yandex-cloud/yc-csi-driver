// Code generated by protoc-gen-goext. DO NOT EDIT.

package endpoint

func (m *YdbSource) SetDatabase(v string) {
	m.Database = v
}

func (m *YdbSource) SetInstance(v string) {
	m.Instance = v
}

func (m *YdbSource) SetPaths(v []string) {
	m.Paths = v
}

func (m *YdbSource) SetServiceAccountId(v string) {
	m.ServiceAccountId = v
}

func (m *YdbSource) SetSubnetId(v string) {
	m.SubnetId = v
}

func (m *YdbSource) SetSaKeyContent(v string) {
	m.SaKeyContent = v
}

func (m *YdbSource) SetSecurityGroups(v []string) {
	m.SecurityGroups = v
}

func (m *YdbSource) SetChangefeedCustomName(v string) {
	m.ChangefeedCustomName = v
}

func (m *YdbSource) SetChangefeedCustomConsumerName(v string) {
	m.ChangefeedCustomConsumerName = v
}

func (m *YdbTarget) SetDatabase(v string) {
	m.Database = v
}

func (m *YdbTarget) SetInstance(v string) {
	m.Instance = v
}

func (m *YdbTarget) SetPath(v string) {
	m.Path = v
}

func (m *YdbTarget) SetServiceAccountId(v string) {
	m.ServiceAccountId = v
}

func (m *YdbTarget) SetCleanupPolicy(v YdbCleanupPolicy) {
	m.CleanupPolicy = v
}

func (m *YdbTarget) SetSubnetId(v string) {
	m.SubnetId = v
}

func (m *YdbTarget) SetSaKeyContent(v string) {
	m.SaKeyContent = v
}

func (m *YdbTarget) SetSecurityGroups(v []string) {
	m.SecurityGroups = v
}

func (m *YdbTarget) SetIsTableColumnOriented(v bool) {
	m.IsTableColumnOriented = v
}

func (m *YdbTarget) SetDefaultCompression(v YdbDefaultCompression) {
	m.DefaultCompression = v
}
