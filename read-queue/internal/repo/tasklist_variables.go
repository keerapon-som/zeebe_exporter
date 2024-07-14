package repo

import (
	"fmt"
	"readq/internal/data"
	"readq/internal/utils/postgresql"
)

type TasklistVariable interface {
	InsertAndUpdate(records []data.TasklistVariables) (err error)
	// GetTypeList() (res []data.MasterTypeListForCommon, err error)
	// GetServiceList() (res []data.ServiceListForCommon, err error)
	// GetCategoryTypeList() (res []data.ParameterForCommon, err error)
	// GetHardwareTypeList() (res []data.ParameterForCommon, err error)
}

type tasklistVariable struct {
}

func NewCommonRepo() TasklistVariable {
	return &tasklistVariable{}
}

func (r tasklistVariable) InsertAndUpdate(records []data.TasklistVariables) (err error) {
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
        INSERT INTO public.tasklist_variables 
        (id, tenantid, key, partitionid, name, value, fullValue, isPreview, scopeFlowNodeId, processInstanceId, position) 
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
        ON CONFLICT (id) DO UPDATE SET
        tenantid = EXCLUDED.tenantid,
        key = EXCLUDED.key,
        partitionid = EXCLUDED.partitionid,
        name = EXCLUDED.name,
        value = EXCLUDED.value,
        fullValue = EXCLUDED.fullValue,
        isPreview = EXCLUDED.isPreview,
        scopeFlowNodeId = EXCLUDED.scopeFlowNodeId,
        processInstanceId = EXCLUDED.processInstanceId,
        position = EXCLUDED.position
        WHERE EXCLUDED.position > public.tasklist_variables.position
    `)

	for _, record := range records {
		_, err = stmt.Exec(record.ID, record.TenantID, record.Key, record.PartitionID, record.Name, record.Value, record.FullValue, record.IsPreview, record.ScopeFlowNodeID, record.ProcessInstanceID, record.Position)
		if err != nil {
			fmt.Println("error in exec stmt")
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

// func (r tasklistVariable) GetTypeList() (res []data.MasterTypeListForCommon, err error) {
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
