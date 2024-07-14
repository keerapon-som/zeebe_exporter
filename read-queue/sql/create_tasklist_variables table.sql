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
)