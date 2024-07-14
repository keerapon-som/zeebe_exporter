package service

import (
	"encoding/json"
	"fmt"
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
	fmt.Println("TasklistTaskTable")
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
