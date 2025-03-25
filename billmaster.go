package save

import (
	"database/sql"
	"log"
	dbconnection "medapp_go/DbConnection"
)

func insertBillMaster(masterRec BillMasterReqStruct) error {
	db, err := dbconnection.LocalDBConnect()
	if err != nil {
		log.Println("insertBillMaster: Error connecting to database:", err)
		return err
	}
	defer db.Close()
	//TO GET USERID AND LOGINID
	sqlstring1 := `SELECT created_by, login_id FROM medapp_login_history ORDER BY login_history_id DESC LIMIT 1`
	err = db.QueryRow(sqlstring1).Scan(&masterRec.User_id, &masterRec.Login_id)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("insertBillDetail: No matching login history found")
		} else {
			log.Println("insertBillDetail: Error querying database:", err)
		}
		return err
	}
	//TO INSERT INTO BILLMASTER
	sqlstring := `INSERT INTO medapp_bill_master (bill_no, bill_date, bill_amount, bill_gst, net_price, login_id, created_by, created_date, updated_by, updated_date)
				VALUES (?, NOW(), ?, ?, ?, ?, ?, NOW(), ?, NOW())`
	_, err = db.Exec(sqlstring, masterRec.Bill_no, masterRec.Bill_amount, masterRec.Bill_gst, masterRec.Net_price, masterRec.Login_id, masterRec.User_id, masterRec.User_id)
	if err != nil {
		log.Println("insertBillMaster: Error inserting into bill master:", err)
		return err
	}

	return nil
}
