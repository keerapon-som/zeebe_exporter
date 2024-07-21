package repo

import (
	"fmt"
	"readq/internal/data"
	"readq/internal/utils/postgresql"

	"github.com/lib/pq"
)

type TasklistTask interface {
	InsertAndUpdate(records []data.TasklistTask) (err error)
	// GetTypeList() (res []data.MasterTypeListForCommon, err error)
	// GetServiceList() (res []data.ServiceListForCommon, err error)
	// GetCategoryTypeList() (res []data.ParameterForCommon, err error)
	// GetHardwareTypeList() (res []data.ParameterForCommon, err error)
}

type tasklistTask struct {
}

func NewTasksRepo() TasklistTask {
	return &tasklistTask{}
}

func (r tasklistTask) InsertAndUpdate(records []data.TasklistTask) (err error) {
	db, err := postgresql.Open()
	if err != nil {
		fmt.Println("error in open db")
		return err
	}
	// defer postgresql.Close()

	tx, err := db.Begin()
	if err != nil {
		fmt.Println("error in begin tx")
		return err
	}

	stmt, err := tx.Prepare(`
		INSERT INTO public.tasklist_tasks 
		(id, tenantid, key, partitionid, bpmnprocessid, processdefinitionid, flownodebpmnid, flownodeinstanceid, processinstanceid, creationtime, completiontime, state, assignee, candidategroups, candidateusers, formkey, formid, formversion, isformembedded, followupdate, duedate, position) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22)
		ON CONFLICT (id) DO UPDATE SET
		state = EXCLUDED.state,
		creationtime = CASE WHEN EXCLUDED.state = 'CREATED' THEN EXCLUDED.creationtime ELSE public.tasklist_tasks.creationtime END,
		completiontime = CASE WHEN EXCLUDED.state = 'COMPLETED' THEN EXCLUDED.completiontime ELSE public.tasklist_tasks.completiontime END,
		tenantid = CASE WHEN EXCLUDED.state = 'COMPLETED' THEN EXCLUDED.tenantid ELSE public.tasklist_tasks.tenantid END,
		key = CASE WHEN EXCLUDED.state = 'COMPLETED' THEN EXCLUDED.key ELSE public.tasklist_tasks.key END,
		partitionid = CASE WHEN EXCLUDED.state = 'COMPLETED' THEN EXCLUDED.partitionid ELSE public.tasklist_tasks.partitionid END,
		bpmnprocessid = CASE WHEN EXCLUDED.state = 'COMPLETED' THEN EXCLUDED.bpmnprocessid ELSE public.tasklist_tasks.bpmnprocessid END,
		processdefinitionid = CASE WHEN EXCLUDED.state = 'COMPLETED' THEN EXCLUDED.processdefinitionid ELSE public.tasklist_tasks.processdefinitionid END,
		flownodebpmnid = CASE WHEN EXCLUDED.state = 'COMPLETED' THEN EXCLUDED.flownodebpmnid ELSE public.tasklist_tasks.flownodebpmnid END,
		flownodeinstanceid = CASE WHEN EXCLUDED.state = 'COMPLETED' THEN EXCLUDED.flownodeinstanceid ELSE public.tasklist_tasks.flownodeinstanceid END,
		processinstanceid = CASE WHEN EXCLUDED.state = 'COMPLETED' THEN EXCLUDED.processinstanceid ELSE public.tasklist_tasks.processinstanceid END,
		assignee = CASE WHEN EXCLUDED.state = 'COMPLETED' THEN EXCLUDED.assignee ELSE public.tasklist_tasks.assignee END,
		candidategroups = CASE WHEN EXCLUDED.state = 'COMPLETED' THEN EXCLUDED.candidategroups ELSE public.tasklist_tasks.candidategroups END,
		candidateusers = CASE WHEN EXCLUDED.state = 'COMPLETED' THEN EXCLUDED.candidateusers ELSE public.tasklist_tasks.candidateusers END,
		formkey = CASE WHEN EXCLUDED.state = 'COMPLETED' THEN EXCLUDED.formkey ELSE public.tasklist_tasks.formkey END,
		formid = CASE WHEN EXCLUDED.state = 'COMPLETED' THEN EXCLUDED.formid ELSE public.tasklist_tasks.formid END,
		formversion = CASE WHEN EXCLUDED.state = 'COMPLETED' THEN EXCLUDED.formversion ELSE public.tasklist_tasks.formversion END,
		isformembedded = CASE WHEN EXCLUDED.state = 'COMPLETED' THEN EXCLUDED.isformembedded ELSE public.tasklist_tasks.isformembedded END,
		followupdate = CASE WHEN EXCLUDED.state = 'COMPLETED' THEN EXCLUDED.followupdate ELSE public.tasklist_tasks.followupdate END,
		duedate = CASE WHEN EXCLUDED.state = 'COMPLETED' THEN EXCLUDED.duedate ELSE public.tasklist_tasks.duedate END,
		position = CASE WHEN EXCLUDED.state = 'COMPLETED' THEN EXCLUDED.position ELSE public.tasklist_tasks.position END
		WHERE public.tasklist_tasks.position < EXCLUDED.position;
	`)
	if err != nil {
		fmt.Println("error preparing statement:", err)
		return err // Make sure to handle the error appropriately, possibly rolling back the transaction
	}
	var creationTime int64
	var completionTime int64
	for _, record := range records {
		if record.State == "COMPLETED" {
			creationTime = 0
			completionTime = record.CompletionTime
		} else if record.State == "CREATED" {
			creationTime = record.CreationTime
			completionTime = 0
		}
		_, err = stmt.Exec(record.ID, record.TenantID, record.Key, record.PartitionID, record.BPMNProcessID, record.ProcessDefinitionID, record.FlowNodeBPMNID, record.FlowNodeInstanceId, record.ProcessInstanceID, creationTime, completionTime, record.State, record.Assignee, pq.Array(record.CandidateGroups), pq.Array(record.CandidateUsers), record.FormKey, record.FormID, record.FormVersion, record.IsFormEmbedded, record.FollowupDate, record.DueDate, record.Position)
		if err != nil {
			fmt.Printf("Error executing statement: %v\n", err)
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println("error in commit tx")
		return err
	}

	return err

}

// func (r tasklistTask) GetTypeList() (res []data.MasterTypeListForCommon, err error) {
// 	db, err := db.Open()
// 	if err != nil {
// 		fmt.Println("error in open db")
// 		return res, err
// 	}
// 	defer db.Close()

// 	var result []data.MasterTypeListForCommon

// 	//TODO: JOIN service_manager <> user to get name
// 	query := `
// 	SELECT category, code, name, owner, hardware_type
// 	FROM cmdb.master_type
// 	WHERE flag = true
// `
// 	fmt.Println(query)
// 	err = db.Select(&result, query)
// 	if utils.IsSQLReallyError(err) {
// 		fmt.Println("error in select db")
// 		return result, err
// 	}

// 	return result, err

// }
