package repo

import (
	"readq/internal/config"
	"readq/internal/utils/postgresql"
	"testing"
)

// ExecSQL executes a given SQL command
func ExecSQL(command string) error {
	db, _ := postgresql.Open()
	_, err := db.Exec(command)
	return err
}

func TestCreateTable(t *testing.T) {

	config := config.GetConfig()
	postgresql.InitDatabase(config.PostgresDB.Postgres_connectionstring)

	command := `
		CREATE TABLE IF NOT EXISTS public.tasklist_variables
		(
			id character varying NOT NULL,
			tenantId character varying,
			key bigint, -- Changed from integer to bigint
			partitionId bigint, -- Changed from integer to bigint
			name character varying,
			value character varying,
			fullValue character varying,
			isPreview boolean,
			scopeFlowNodeId character varying,
			processInstanceId character varying,
			position bigint, -- Changed from integer to bigint
			PRIMARY KEY (id)
		)`
	ExecSQL(command)
}

func TestCreateTasklistTaskTable(t *testing.T) {
	config := config.GetConfig()
	postgresql.InitDatabase(config.PostgresDB.Postgres_connectionstring)

	command := `
		CREATE TABLE IF NOT EXISTS public.tasklist_tasks
		(
			id character varying NOT NULL,
			tenantId character varying,
			key bigint,
			partitionId integer,
			bpmnProcessId character varying,
			processDefinitionId character varying,
			flowNodeBpmnId character varying,
			flowNodeInstanceId character varying,
			processInstanceId character varying,
			creationTime bigint,
			completionTime bigint,
			state character varying,
			assignee character varying,
			candidateGroups text[],
			candidateUsers text[],
			formKey character varying,
			formId character varying,
			formVersion integer,
			isFormEmbedded boolean,
			followUpDate character varying,
			dueDate character varying,
			position bigint,
			PRIMARY KEY (id)
		)`
	ExecSQL(command)
}
