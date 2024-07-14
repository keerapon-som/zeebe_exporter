package service

import (
	"fmt"

	"google.golang.org/protobuf/proto"
)

var mng = NewManager()

var (
	recordDeployment                    = make(chan DeploymentRecord, 1000)
	recordDeploymentDistribution        = make(chan DeploymentDistributionRecord, 1000)
	recordError                         = make(chan ErrorRecord, 1000)
	recordIncident                      = make(chan IncidentRecord, 1000)
	recordJobs                          = make(chan JobRecord, 1000)
	recordJobBatch                      = make(chan JobBatchRecord, 1000)
	recordMessage                       = make(chan MessageRecord, 1000)
	recordMessageSubscription           = make(chan MessageSubscriptionRecord, 1000)
	recordMessageStartEventSubscription = make(chan MessageStartEventSubscriptionRecord, 1000)
	recordTimer                         = make(chan TimerRecord, 1000)
	recordVariable                      = make(chan VariableRecord, 1000)
	recordVariableDocument              = make(chan VariableDocumentRecord, 1000)
	recordProcessInstance               = make(chan ProcessInstanceRecord, 1000)
	recordProcessInstanceCreation       = make(chan ProcessInstanceCreationRecord, 1000)
	recordProcessMessageSubscription    = make(chan ProcessMessageSubscriptionRecord, 1000)
	recordProcess                       = make(chan ProcessRecord, 1000)
	recordProcessEvent                  = make(chan ProcessEventRecord, 1000)
	recordDecision                      = make(chan DecisionRecord, 1000)
	recordDecisionRequirementsMetadata  = make(chan DecisionRequirementsMetadata, 1000)
	recordDecisionRequirements          = make(chan DecisionRequirementsRecord, 1000)
	recordDecisionEvaluation            = make(chan DecisionEvaluationRecord, 1000)
	recordProcessInstanceModification   = make(chan ProcessInstanceModificationRecord, 1000)
	// recordCheckpoint                    = make(chan CheckpointRecord, 1000)
	// recordSignal                        = make(chan SignalRecord, 1000)
	// recordSignalSubscription            = make(chan SignalSubscriptionRecord, 1000)
	// recordForm                          = make(chan FormRecord, 1000)
	// recordResourceDeletion              = make(chan ResourceDeletionRecord, 1000)
	// recordUserTask                                                                   = make(chan []Record, 1000)
	// recordCompensationSubscription                                                   = make(chan []Record, 1000)
	// recordEscalation                                                                 = make(chan []Record, 1000)
	// recordDeploymentResource                                                         = make(chan []Record, 1000)
	// recordDeploymentProcessMetadata                                                  = make(chan []Record, 1000)
	// recordDeploymentDecisionMetadata                                                 = make(chan []Record, 1000)
	// recordDeploymentFormMetadata                                                     = make(chan []Record, 1000)
	// recordDecisionEvaluationEvaluatedDecision                                        = make(chan []Record, 1000)
	// recordDecisionEvaluationMatchedRule                                              = make(chan []Record, 1000)
	// recordDecisionEvaluationEvaluatedInput                                           = make(chan []Record, 1000)
	// recordDecisionEvaluationEvaluatedOutput                                          = make(chan []Record, 1000)
	// recordProcessInstanceModificationProcessInstanceModificationTerminateInstruction = make(chan []Record, 1000)
)

func TypeClasify(c []byte) {
	//หมุนข้อมูลตรงนี้
	//แต่ละประเภทข้อมูลส่งเข้าท่อของแต่ละประเภท ตามที่กำหนดไว้ ใน recordJobs, recordIncidents, recordDeployments
	var record Record

	err := proto.Unmarshal(c, &record)
	if err != nil {
		fmt.Println("Error unmarshalling message")
	}

	if record.Record.TypeUrl == "type.googleapis.com/exporter_protocol.DeploymentRecord" {
		var Deployment DeploymentRecord // Change to slice
		err = proto.Unmarshal(record.Record.Value, &Deployment)
		recordDeployment <- Deployment // Use append to add elements to the slice
	} else if record.Record.TypeUrl == "type.googleapis.com/exporter_protocol.DeploymentDistributionRecord" {
		var DeploymentDistribution DeploymentDistributionRecord // Change to slice
		err = proto.Unmarshal(record.Record.Value, &DeploymentDistribution)
		recordDeploymentDistribution <- DeploymentDistribution // Use append to add elements to the slice
	} else if record.Record.TypeUrl == "type.googleapis.com/exporter_protocol.ErrorRecord" {
		var Error ErrorRecord // Change to slice
		err = proto.Unmarshal(record.Record.Value, &Error)
		recordError <- Error // Use append to add elements to the slice
	} else if record.Record.TypeUrl == "type.googleapis.com/exporter_protocol.IncidentRecord" {
		var Incident IncidentRecord // Change to slice
		err = proto.Unmarshal(record.Record.Value, &Incident)
		recordIncident <- Incident // Use append to add elements to the slice
	} else if record.Record.TypeUrl == "type.googleapis.com/exporter_protocol.JobRecord" {
		var Job JobRecord // Change to slice
		err = proto.Unmarshal(record.Record.Value, &Job)
		recordJobs <- Job // Use append to add elements to the slice
	} else if record.Record.TypeUrl == "type.googleapis.com/exporter_protocol.JobBatchRecord" {
		var JobBatch JobBatchRecord // Change to slice
		err = proto.Unmarshal(record.Record.Value, &JobBatch)
		recordJobBatch <- JobBatch // Use append to add elements to the slice
	} else if record.Record.TypeUrl == "type.googleapis.com/exporter_protocol.MessageRecord" {
		var Message MessageRecord // Change to slice
		err = proto.Unmarshal(record.Record.Value, &Message)
		recordMessage <- Message // Use append to add elements to the slice
	} else if record.Record.TypeUrl == "type.googleapis.com/exporter_protocol.MessageSubscriptionRecord" {
		var MessageSubscription MessageSubscriptionRecord // Change to slice
		err = proto.Unmarshal(record.Record.Value, &MessageSubscription)
		recordMessageSubscription <- MessageSubscription // Use append to add elements to the slice
	} else if record.Record.TypeUrl == "type.googleapis.com/exporter_protocol.MessageStartEventSubscriptionRecord" {
		var MessageStartEventSubscription MessageStartEventSubscriptionRecord // Change to slice
		err = proto.Unmarshal(record.Record.Value, &MessageStartEventSubscription)
		recordMessageStartEventSubscription <- MessageStartEventSubscription // Use append to add elements to the slice
	} else if record.Record.TypeUrl == "type.googleapis.com/exporter_protocol.TimerRecord" {
		var Timer TimerRecord // Change to slice
		err = proto.Unmarshal(record.Record.Value, &Timer)
		recordTimer <- Timer // Use append to add elements to the slice
	} else if record.Record.TypeUrl == "type.googleapis.com/exporter_protocol.VariableRecord" {
		var Variable VariableRecord // Change to slice
		err = proto.Unmarshal(record.Record.Value, &Variable)
		recordVariable <- Variable // Use append to add elements to the slice
	} else if record.Record.TypeUrl == "type.googleapis.com/exporter_protocol.VariableDocumentRecord" {
		var VariableDocument VariableDocumentRecord // Change to slice
		err = proto.Unmarshal(record.Record.Value, &VariableDocument)
		recordVariableDocument <- VariableDocument // Use append to add elements to the slice
	} else if record.Record.TypeUrl == "type.googleapis.com/exporter_protocol.ProcessInstanceRecord" {
		var ProcessInstance ProcessInstanceRecord // Change to slice
		err = proto.Unmarshal(record.Record.Value, &ProcessInstance)
		recordProcessInstance <- ProcessInstance // Use append to add elements to the slice
	} else if record.Record.TypeUrl == "type.googleapis.com/exporter_protocol.ProcessInstanceCreationRecord" {
		var ProcessInstanceCreation ProcessInstanceCreationRecord // Change to slice
		err = proto.Unmarshal(record.Record.Value, &ProcessInstanceCreation)
		recordProcessInstanceCreation <- ProcessInstanceCreation // Use append to add elements to the slice
	} else if record.Record.TypeUrl == "type.googleapis.com/exporter_protocol.ProcessMessageSubscriptionRecord" {
		var ProcessMessageSubscription ProcessMessageSubscriptionRecord // Change to slice
		err = proto.Unmarshal(record.Record.Value, &ProcessMessageSubscription)
		recordProcessMessageSubscription <- ProcessMessageSubscription // Use append to add elements to the slice
	} else if record.Record.TypeUrl == "type.googleapis.com/exporter_protocol.ProcessRecord" {
		var Process ProcessRecord // Change to slice
		err = proto.Unmarshal(record.Record.Value, &Process)
		recordProcess <- Process // Use append to add elements to the slice
	} else if record.Record.TypeUrl == "type.googleapis.com/exporter_protocol.ProcessEventRecord" {
		var ProcessEvent ProcessEventRecord // Change to slice
		err = proto.Unmarshal(record.Record.Value, &ProcessEvent)
		recordProcessEvent <- ProcessEvent // Use append to add elements to the slice
	} else if record.Record.TypeUrl == "type.googleapis.com/exporter_protocol.DecisionRecord" {
		var Decision DecisionRecord // Change to slice
		err = proto.Unmarshal(record.Record.Value, &Decision)
		recordDecision <- Decision // Use append to add elements to the slice
	} else if record.Record.TypeUrl == "type.googleapis.com/exporter_protocol.DecisionRequirementsMetadata" {
		var DecisionRequirementsMetadata DecisionRequirementsMetadata // Change to slice
		err = proto.Unmarshal(record.Record.Value, &DecisionRequirementsMetadata)
		recordDecisionRequirementsMetadata <- DecisionRequirementsMetadata // Use append to add elements to the slice
	} else if record.Record.TypeUrl == "type.googleapis.com/exporter_protocol.DecisionRequirementsRecord" {
		var DecisionRequirements DecisionRequirementsRecord // Change to slice
		err = proto.Unmarshal(record.Record.Value, &DecisionRequirements)
		recordDecisionRequirements <- DecisionRequirements // Use append to add elements to the slice
	} else if record.Record.TypeUrl == "type.googleapis.com/exporter_protocol.DecisionEvaluationRecord" {
		var DecisionEvaluation DecisionEvaluationRecord // Change to slice
		err = proto.Unmarshal(record.Record.Value, &DecisionEvaluation)
		recordDecisionEvaluation <- DecisionEvaluation // Use append to add elements to the slice
	} else if record.Record.TypeUrl == "type.googleapis.com/exporter_protocol.ProcessInstanceModificationRecord" {
		var ProcessInstanceModification ProcessInstanceModificationRecord // Change to slice
		err = proto.Unmarshal(record.Record.Value, &ProcessInstanceModification)
		recordProcessInstanceModification <- ProcessInstanceModification // Use append to add elements to the slice
	} else {
		fmt.Println("No record type found")
	}
	if err != nil {
		fmt.Println("Error unmarshalling message")
	}
}

func PerformBatchRecord() {

	fmt.Println("PerformBatchRecord")

	if len(recordJobs) != 0 {
		go JobsToDB(recordJobs)
	}

	if len(recordVariable) != 0 {
		go VariablesToDB(recordVariable)
	}
}
