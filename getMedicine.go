package getmedicine

import (
	"log"
	dbconnection "medapp_go/DbConnection"
)


func getMedicine() (MedicineRespStruct, error) {


	var MedicineRec MedicineStruct   //variable declaration
	var resp MedicineRespStruct


	db, err := dbconnection.LocalDBConnect()   //db connection
	if err != nil {
		log.Println(err)
		resp.Errmsg = "MM01"+err.Error()
		resp.Status = "E"
		return resp, err
	}
	defer db.Close()
	//TO GET MEDICINE AND BRAND
	coreString := `select medicine_name,brand from medapp_medicine_master`
	rows, err := db.Query(coreString)
	if err != nil {
		log.Println(err)
		resp.Errmsg = "MM02"+err.Error()
		resp.Status = "E"
		return resp, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&MedicineRec.MedicineName, &MedicineRec.Brand)
		if err != nil {
			log.Println(err)
			resp.Errmsg = "MM03"+err.Error()
			resp.Status = "E"
			return resp, err
		}
		resp.Medicine = append(resp.Medicine, MedicineRec)
	}
	return resp, nil
}