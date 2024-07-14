package repodata

import "database/sql"

type Tasklisttask struct {
	ID                  sql.NullString `db:"id"`
	TenantId            sql.NullString `db:"tenantid"`
	Key                 int64          `db:"key"`
	PartitionId         int            `db:"partitionid"`
	BPMNProcessID       sql.NullString `db:"bpmnprocessid"`
	ProcessDefinitionId sql.NullString `db:"processdefinitionid"`
	FlowNodeBPMNId      sql.NullString `db:"flownodebpmnid"`
	ProcessInstanceId   sql.NullString `db:"processinstanceid"`
	CreationTime        sql.NullString `db:"creationtime"`
	CompletionTime      sql.NullString `db:"completiontime"`
	Assignee            sql.NullString `db:"assignee"`
	CandidateGroups     []string       `db:"candidategroups"`
	CandidateUsers      []string       `db:"candidateusers"`
	FormKey             sql.NullString `db:"formkey"`
	FormId              sql.NullString `db:"formid"`
	FormVersion         int            `db:"formversion"`
	IsFormEmbedded      bool           `db:"isformembedded"`
	FollowupDate        sql.NullString `db:"followupdate"`
	DueDate             sql.NullString `db:"duedate"` // Changed to sql.NullString
}

type Tasklistvariables struct {
	ID                string         `db:"id"`
	TenantId          string         `db:"tenantid"`
	Key               int64          `db:"key"`
	PartitionId       int            `db:"partitionid"`
	Name              string         `db:"name"`
	Value             sql.NullString `db:"value"`
	FullValue         sql.NullString `db:"fullValue"`
	IsPreview         bool           `db:"isPreview"`
	ScopeFlowNodeId   sql.NullString `db:"scopeFlowNodeId"`
	ProcessInstanceId sql.NullString `db:"processInstanceId"`
	Position          int64          `db:"position"`
}
