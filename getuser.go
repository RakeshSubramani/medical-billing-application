package userid

import (
	"log"
	dbconnection "medapp_go/DbConnection"
)


func getUser() (UserRespStruct, error) {
	var MedicineRec UserStruct
	var resp UserRespStruct
	db, err := dbconnection.LocalDBConnect()
	if err != nil {
		log.Println(err)
		resp.Errmsg = "UM01"+err.Error()
		resp.Status = "E"
		return resp, err
	}
	defer db.Close()
	//TO GET USER ID
	coreString := `select user_id from medapp_login`
	rows, err := db.Query(coreString)
	if err != nil {
		log.Println(err)
		resp.Errmsg = "UM02"+err.Error()
		resp.Status = "E"
		return resp, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&MedicineRec.UserId)
		if err != nil {
			log.Println(err)
			resp.Errmsg = "UM03"+err.Error()
			resp.Status = "E"
			return resp, err
		}
		resp.User = append(resp.User, MedicineRec)
	}
	return resp, nil
}