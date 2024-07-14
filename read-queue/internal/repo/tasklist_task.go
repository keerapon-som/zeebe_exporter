// package exportertasklist_repo
package repo

// import (
// 	"fmt"
// )

// type ExporterTasklistRepo interface {
// 	// GetTypeList() (res []data.MasterTypeListForCommon, err error)
// 	// GetServiceList() (res []data.ServiceListForCommon, err error)
// 	// GetCategoryTypeList() (res []data.ParameterForCommon, err error)
// 	// GetHardwareTypeList() (res []data.ParameterForCommon, err error)
// }

// type exporterTastlistrRepo struct {
// }

// func NewCommonRepo() ExporterTasklistRepo {
// 	return &exporterTastlistrRepo{}
// }

// func (r exporterTastlistrRepo) GetTypeList() (res []data.MasterTypeListForCommon, err error) {
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
