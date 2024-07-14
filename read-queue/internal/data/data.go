package data

type TasklistTask struct {
	ID                  string
	TenantID            string
	Key                 int64
	PartitionID         int
	BPMNProcessID       string
	ProcessDefinitionID string
	FlowNodeBPMNID      string
	FlowNodeInstanceId  string
	ProcessInstanceID   string
	CreationTime        string
	CompletionTime      string
	State               string
	Assignee            string
	CandidateGroups     []string
	CandidateUsers      []string
	FormKey             string
	FormID              string
	FormVersion         int
	IsFormEmbedded      bool
	FollowupDate        string
	DueDate             string
}

type TasklistVariables struct {
	ID                string
	TenantID          string
	Key               int64
	PartitionID       int
	Name              string
	Value             string
	FullValue         string
	IsPreview         bool
	ScopeFlowNodeID   string
	ProcessInstanceID string
	Position          int64
}
