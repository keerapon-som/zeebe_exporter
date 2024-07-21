package service

import (
	"encoding/json"
	"fmt"
	"readq/internal/data"
	"readq/internal/repo"
)

type jobManager struct {
}

type JobManager interface {
	TohistoryTable(jobrecords []JobRecord)
	ToTasklistTaskTable(jobrecords []JobRecord)
}

func (m *jobManager) TohistoryTable(jobrecords []JobRecord) {
	ListJobJson, err := json.MarshalIndent(jobrecords, "", "  ")
	if err != nil {
		fmt.Println("Error converting record to JSON:", err)
		return
	}
	fmt.Println(string(ListJobJson))
}

func (m *jobManager) ToTasklistTaskTable(jobrecords []JobRecord) {
	var listToTasklistTask []data.TasklistTask
	repo := repo.NewTasksRepo()

	for _, job := range jobrecords {
		if job.Type == "io.camunda.zeebe:userTask" {
			CustomHeaders := job.CustomHeaders.AsMap()
			var candidateGroups []string
			var candidateUsers []string

			// Safely handle candidateGroups
			if val, ok := CustomHeaders["io.camunda.zeebe:candidateGroups"]; ok && val != nil {
				strVal, ok := val.(string)
				if !ok {
					// Handle error: value is not a string
				} else {
					err := json.Unmarshal([]byte(strVal), &candidateGroups)
					if err != nil {
						// Handle error: JSON unmarshaling failed
					}
				}
			}

			// Safely handle candidateUsers
			if val, ok := CustomHeaders["io.camunda.zeebe:candidateUsers"]; ok && val != nil {
				strVal, ok := val.(string)
				if !ok {
					// Handle error: value is not a string
				} else {
					err := json.Unmarshal([]byte(strVal), &candidateUsers)
					if err != nil {
						// Handle error: JSON unmarshaling failed
					}
				}
			}

			var isFormEmbedded bool
			var formKey string
			// Safely handle formKey
			if val, ok := CustomHeaders["io.camunda.zeebe:formKey"]; ok && val != nil {
				formKey, ok := val.(string)
				if !ok {
					// Handle error: value is not a string
				} else {
					fmt.Println(formKey)
					if len(formKey) >= 18 && formKey[0:18] == "camunda-forms:bpmn" {
						isFormEmbedded = true
					} else {
						isFormEmbedded = false
					}
				}
			}

			task := data.TasklistTask{
				ID:                  fmt.Sprintf("%d", job.Metadata.Key),
				TenantID:            job.TenantId,
				Key:                 job.Metadata.Key,
				PartitionID:         int(job.Metadata.PartitionId),
				BPMNProcessID:       job.BpmnProcessId,
				ProcessDefinitionID: fmt.Sprintf("%d", job.ProcessDefinitionKey),
				FlowNodeBPMNID:      job.ElementId,
				FlowNodeInstanceId:  fmt.Sprintf("%d", job.ElementInstanceKey),
				ProcessInstanceID:   fmt.Sprintf("%d", job.ProcessInstanceKey),
				CreationTime:        job.Metadata.Timestamp,
				CompletionTime:      job.Metadata.Timestamp,
				State:               job.Metadata.Intent,
				Assignee:            CustomHeaders["io.camunda.zeebe:assignee"].(string),
				CandidateGroups:     candidateGroups,
				CandidateUsers:      candidateUsers,
				FormKey:             formKey,
				FormID:              "",
				FormVersion:         0,
				IsFormEmbedded:      isFormEmbedded,
				FollowupDate:        CustomHeaders["io.camunda.zeebe:followUpDate"].(string),
				DueDate:             CustomHeaders["io.camunda.zeebe:dueDate"].(string),
				Position:            job.Metadata.Position,
			}
			listToTasklistTask = append(listToTasklistTask, task)
		}
		// fmt.Println(listToTasklistTask)
	}
	prettyJsonFormat, _ := json.MarshalIndent(listToTasklistTask, "", "  ")
	fmt.Println(string(prettyJsonFormat))
	repo.InsertAndUpdate(listToTasklistTask)
}

func JobsToDB(pipe chan JobRecord) {
	var batchjobRecords []JobRecord

	for {
		select {
		case job := <-pipe:
			batchjobRecords = append(batchjobRecords, job)
		default:
			fmt.Println("-----Perform Tasklist Task Table-----")
			go mng.JobManager.ToTasklistTaskTable(batchjobRecords)
			// fmt.Println("-----Perform History Job Table-----")
			// go mng.JobManager.TohistoryTable(batchjobRecords)
			return
		}
	}

}
