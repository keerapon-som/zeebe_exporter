package service

import (
	"encoding/json"
	"fmt"
	"readq/internal/data"
	"readq/internal/repo"
)

type variableManager struct {
}

type VariableManager interface {
	TohistoryTable(variableRecords []VariableRecord)
	ToTasklistVariablesTable(variableRecords []VariableRecord)
}

func (m *variableManager) TohistoryTable(variableRecords []VariableRecord) {
	ListJobJson, err := json.MarshalIndent(variableRecords, "", "  ")
	if err != nil {
		fmt.Println("Error converting record to JSON:", err)
		return
	}
	fmt.Println(string(ListJobJson))
}

func (m *variableManager) ToTasklistVariablesTable(variableRecords []VariableRecord) {
	repo := repo.NewCommonRepo()

	var records []data.TasklistVariables

	for _, record := range variableRecords {

		records = append(records, data.TasklistVariables{
			ID:                fmt.Sprintf("%d-%s", record.ScopeKey, record.Name),
			TenantID:          record.TenantId,
			Key:               record.Metadata.Key,
			PartitionID:       int(record.Metadata.PartitionId),
			Name:              record.Name,
			Value:             record.Value,
			FullValue:         record.Value,
			IsPreview:         false, // notfoundinevent
			ScopeFlowNodeID:   fmt.Sprintf("%d", record.ScopeKey),
			ProcessInstanceID: fmt.Sprintf("%d", record.ProcessInstanceKey),
			Position:          record.Metadata.Position,
		})
	}
	fmt.Println("Insert ------- Data")
	fmt.Println(records)

	err := repo.InsertAndUpdate(records)
	if err != nil {
		fmt.Println("ERROR TO APPEND TO POSTGRESQL", err)
		return
	}
	fmt.Println("SUCCESS TO APPEND TO POSTGRESQL tasklist_variables")

}

func VariablesToDB(pipe chan VariableRecord) {
	var batchVariableRecords []VariableRecord

	for {
		select {
		case variable := <-pipe:
			batchVariableRecords = append(batchVariableRecords, variable)
		default:
			// jobmng.ToTasklistTaskTable(batchjobRecords)
			// fmt.Println("-----Perform History Variable Table-----")
			// mng.VariableManager.TohistoryTable(batchVariableRecords)
			// mng.VariableManager.ToTasklistVariablesTable(batchVariableRecords)
			return
		}
	}

}
