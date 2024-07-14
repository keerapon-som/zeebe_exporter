package io.zeebe.redis.exporter;

public class Config {
    private boolean job;
    private boolean deployment;
    private boolean processInstance;
    private boolean incident;
    private boolean message;
    private boolean messageSubscription;
    private boolean processMessageSubscription;
    private boolean jobBatch;
    private boolean timer;
    private boolean messageStartEventSubscription;
    private boolean variable;
    private boolean variableDocument;
    private boolean processInstanceCreation;
    private boolean error;
    private boolean processInstanceResult;
    private boolean process;
    private boolean deploymentDistribution;
    private boolean processEvent;
    private boolean decision;
    private boolean decisionRequirements;
    private boolean decisionEvaluation;
    private boolean processInstanceModification;
    private boolean escalation;
    private boolean signalSubscription;
    private boolean signal;
    private boolean resourceDeletion;
    private boolean commandDistribution;
    private boolean processInstanceBatch;
    private boolean messageBatch;
    private boolean form;
    private boolean userTask;
    private boolean processInstanceMigration;
    private boolean compensationSubscription;
    public String redisPositionKeyName;
    public String redisHost;
    public int redisPort;
    public String redisPassword;
    public long periodExportPosition;
    public String streamName;

    public Config() {
        job = Boolean.parseBoolean(System.getenv().getOrDefault("JOB", "false"));
        deployment = Boolean.parseBoolean(System.getenv().getOrDefault("DEPLOYMENT", "false"));
        processInstance = Boolean.parseBoolean(System.getenv().getOrDefault("PROCESS_INSTANCE", "false"));
        incident = Boolean.parseBoolean(System.getenv().getOrDefault("INCIDENT", "false"));
        message = Boolean.parseBoolean(System.getenv().getOrDefault("MESSAGE", "false"));
        messageSubscription = Boolean.parseBoolean(System.getenv().getOrDefault("MESSAGE_SUBSCRIPTION", "false"));
        processMessageSubscription = Boolean.parseBoolean(System.getenv().getOrDefault("PROCESS_MESSAGE_SUBSCRIPTION", "false"));
        jobBatch = Boolean.parseBoolean(System.getenv().getOrDefault("JOB_BATCH", "false"));
        timer = Boolean.parseBoolean(System.getenv().getOrDefault("TIMER", "false"));
        messageStartEventSubscription = Boolean.parseBoolean(System.getenv().getOrDefault("MESSAGE_START_EVENT_SUBSCRIPTION", "false"));
        variable = Boolean.parseBoolean(System.getenv().getOrDefault("VARIABLE", "false"));
        variableDocument = Boolean.parseBoolean(System.getenv().getOrDefault("VARIABLE_DOCUMENT", "false"));
        processInstanceCreation = Boolean.parseBoolean(System.getenv().getOrDefault("PROCESS_INSTANCE_CREATION", "false"));
        error = Boolean.parseBoolean(System.getenv().getOrDefault("ERROR", "false"));
        processInstanceResult = Boolean.parseBoolean(System.getenv().getOrDefault("PROCESS_INSTANCE_RESULT", "false"));
        process = Boolean.parseBoolean(System.getenv().getOrDefault("PROCESS", "false"));
        deploymentDistribution = Boolean.parseBoolean(System.getenv().getOrDefault("DEPLOYMENT_DISTRIBUTION", "false"));
        processEvent = Boolean.parseBoolean(System.getenv().getOrDefault("PROCESS_EVENT", "false"));
        decision = Boolean.parseBoolean(System.getenv().getOrDefault("DECISION", "false"));
        decisionRequirements = Boolean.parseBoolean(System.getenv().getOrDefault("DECISION_REQUIREMENTS", "false"));
        decisionEvaluation = Boolean.parseBoolean(System.getenv().getOrDefault("DECISION_EVALUATION", "false"));
        processInstanceModification = Boolean.parseBoolean(System.getenv().getOrDefault("PROCESS_INSTANCE_MODIFICATION", "false"));
        escalation = Boolean.parseBoolean(System.getenv().getOrDefault("ESCALATION", "false"));
        signalSubscription = Boolean.parseBoolean(System.getenv().getOrDefault("SIGNAL_SUBSCRIPTION", "false"));
        signal = Boolean.parseBoolean(System.getenv().getOrDefault("SIGNAL", "false"));
        resourceDeletion = Boolean.parseBoolean(System.getenv().getOrDefault("RESOURCE_DELETION", "false"));
        commandDistribution = Boolean.parseBoolean(System.getenv().getOrDefault("COMMAND_DISTRIBUTION", "false"));
        processInstanceBatch = Boolean.parseBoolean(System.getenv().getOrDefault("PROCESS_INSTANCE_BATCH", "false"));
        messageBatch = Boolean.parseBoolean(System.getenv().getOrDefault("MESSAGE_BATCH", "false"));
        form = Boolean.parseBoolean(System.getenv().getOrDefault("FORM", "false"));
        userTask = Boolean.parseBoolean(System.getenv().getOrDefault("USER_TASK", "false"));
        processInstanceMigration = Boolean.parseBoolean(System.getenv().getOrDefault("PROCESS_INSTANCE_MIGRATION", "false"));
        compensationSubscription = Boolean.parseBoolean(System.getenv().getOrDefault("COMPENSATION_SUBSCRIPTION", "false"));
        redisPositionKeyName = System.getenv().getOrDefault("RECORD_POSITION_KEY_NAME", "");
        redisHost = System.getenv().getOrDefault("REDIS_HOST", "");
        redisPort = Integer.parseInt(System.getenv().getOrDefault("REDIS_PORT", ""));
        redisPassword = System.getenv().getOrDefault("REDIS_PASSWORD", "");
        periodExportPosition = Long.parseLong(System.getenv().getOrDefault("PERIOD_EXPORT_POSITION", ""));
        streamName = System.getenv().getOrDefault("STREAM_NAME", "");
    }

    

    // Example of a getter
    public boolean IsOpen(String valueType) {
        switch (valueType) {
            case "JOB":
                return job;
            case "DEPLOYMENT":
                return deployment;
            case "PROCESS_INSTANCE":
                return processInstance;
            case "INCIDENT":
                return incident;
            case "MESSAGE":
                return message;
            case "MESSAGE_SUBSCRIPTION":
                return messageSubscription;
            case "PROCESS_MESSAGE_SUBSCRIPTION":
                return processMessageSubscription;
            case "JOB_BATCH":
                return jobBatch;
            case "TIMER":
                return timer;
            case "MESSAGE_START_EVENT_SUBSCRIPTION":
                return messageStartEventSubscription;
            case "VARIABLE":
                return variable;
            case "VARIABLE_DOCUMENT":
                return variableDocument;
            case "PROCESS_INSTANCE_CREATION":
                return processInstanceCreation;
            case "ERROR":
                return error;
            case "PROCESS_INSTANCE_RESULT":
                return processInstanceResult;
            case "PROCESS":
                return process;
            case "DEPLOYMENT_DISTRIBUTION":
                return deploymentDistribution;
            case "PROCESS_EVENT":
                return processEvent;
            case "DECISION":
                return decision;
            case "DECISION_REQUIREMENTS":
                return decisionRequirements;
            case "DECISION_EVALUATION":
                return decisionEvaluation;
            case "PROCESS_INSTANCE_MODIFICATION":
                return processInstanceModification;
            case "ESCALATION":
                return escalation;
            case "SIGNAL_SUBSCRIPTION":
                return signalSubscription;
            case "SIGNAL":
                return signal;
            case "RESOURCE_DELETION":
                return resourceDeletion;
            case "COMMAND_DISTRIBUTION":
                return commandDistribution;
            case "PROCESS_INSTANCE_BATCH":
                return processInstanceBatch;
            case "MESSAGE_BATCH":
                return messageBatch;
            case "FORM":
                return form;
            case "USER_TASK":
                return userTask;
            case "PROCESS_INSTANCE_MIGRATION":
                return processInstanceMigration;
            case "COMPENSATION_SUBSCRIPTION":
                return compensationSubscription;
            default:
                return false;
        }
    }
    // Getters and potentially setters
}