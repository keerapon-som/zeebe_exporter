-- Table: public.job

-- DROP TABLE IF EXISTS public.job;

CREATE TABLE IF NOT EXISTS public.tasklist_task
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
    creationTime character varying,
    completionTime character varying,
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
    PRIMARY KEY (id)
)