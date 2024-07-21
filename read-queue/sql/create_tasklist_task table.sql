-- Table: public.job

-- DROP TABLE IF EXISTS public.job;

CREATE TABLE IF NOT EXISTS public.tasklist_tasks
(
	id character varying NOT NULL,
    tenantId character varying,
    key integer,
    partitionId integer,
    bpmnProcessId character varying,
    processDefinitionId character varying,
    flowNodeBpmnId character varying,
    flowNodeInstanceId character varying,
    processInstanceId character varying,
    creationTime integer,
    completionTime integer,
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
    position integer,
    PRIMARY KEY (id)
)