// Code generated by protoc-gen-goext. DO NOT EDIT.

package agent

func (m *Agent) SetId(v string) {
	m.Id = v
}

func (m *Agent) SetFolderId(v string) {
	m.FolderId = v
}

func (m *Agent) SetName(v string) {
	m.Name = v
}

func (m *Agent) SetDescription(v string) {
	m.Description = v
}

func (m *Agent) SetComputeInstanceId(v string) {
	m.ComputeInstanceId = v
}

func (m *Agent) SetStatus(v Status) {
	m.Status = v
}

func (m *Agent) SetErrors(v []string) {
	m.Errors = v
}

func (m *Agent) SetCurrentJobId(v string) {
	m.CurrentJobId = v
}

func (m *Agent) SetAgentVersionId(v string) {
	m.AgentVersionId = v
}
