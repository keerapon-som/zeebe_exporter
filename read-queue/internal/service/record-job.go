package service

import (
	"encoding/json"
	"fmt"
	"readq/internal/data"
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
	var CreationTime string
	var CompletionTime string

	for _, job := range jobrecords {
		if job.Type == "io.camunda.zeebe:userTask" {
			CustomHeaders := job.CustomHeaders.AsMap()
			var candidateGroups []string
			var candidateUsers []string

			// Assuming CustomHeaders is a map[string]interface{} with your data
			candidateGroupsStr, ok := CustomHeaders["io.camunda.zeebe:candidateGroups"].(string)
			if ok && candidateGroupsStr != "" {
				err := json.Unmarshal([]byte(candidateGroupsStr), &candidateGroups)
				if err != nil {
					fmt.Printf("Error unmarshalling candidateGroups: %v\n", err)
				}
			}

			candidateUsersStr, ok := CustomHeaders["io.camunda.zeebe:candidateUsers"].(string)
			if ok && candidateUsersStr != "" {
				err := json.Unmarshal([]byte(candidateUsersStr), &candidateUsers)
				if err != nil {
					fmt.Printf("Error unmarshalling candidateUsers: %v\n", err)
				}
			}
			if job.Metadata.Intent == "CREATED" {
				CreationTime = string(job.Metadata.Timestamp)
			} else if job.Metadata.Intent == "COMPLETED" {
				CompletionTime = string(job.Metadata.Timestamp)
			}
			task := data.TasklistTask{
				ID:                  string(job.Metadata.Key),
				TenantID:            job.TenantId,
				Key:                 job.Metadata.Key,
				PartitionID:         int(job.Metadata.PartitionId),
				BPMNProcessID:       job.BpmnProcessId,
				ProcessDefinitionID: fmt.Sprintf("%d", job.ProcessDefinitionKey),
				FlowNodeBPMNID:      job.ElementId,
				FlowNodeInstanceId:  string(job.ElementInstanceKey),
				ProcessInstanceID:   fmt.Sprintf("%d", job.ProcessInstanceKey),
				CreationTime:        CreationTime,
				CompletionTime:      CompletionTime,
				State:               job.Metadata.Intent,
				Assignee:            CustomHeaders["io.camunda.zeebe:assignee"].(string),
				CandidateGroups:     candidateGroups,
				CandidateUsers:      candidateUsers,
				FormKey:             "",
				FormID:              "",
				FormVersion:         0,
				IsFormEmbedded:      true,
				FollowupDate:        "",
				DueDate:             "",
			}
			listToTasklistTask = append(listToTasklistTask, task)
		}
		fmt.Println(listToTasklistTask)
	}
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
			fmt.Println("-----Perform History Job Table-----")
			go mng.JobManager.TohistoryTable(batchjobRecords)
			return
		}
	}

}
