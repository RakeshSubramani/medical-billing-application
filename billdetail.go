package save

import (
	"database/sql"
	"log"
	dbconnection "medapp_go/DbConnection"
)

func BillDetail(detail BillDetailReqStruct) (BillDetailResp, error) {
	var resp BillDetailResp
	db, err := dbconnection.LocalDBConnect()
	if err != nil {
		log.Println("insertBillDetail: Error connecting to database:", err)
		resp.ErrMsg = "insertingbill details:BDM01" + err.Error()
	}
	defer db.Close()
	//TO GET MEDICINE MASTER ID
	sqlstring := `SELECT medicine_master_id FROM medapp_medicine_master WHERE medicine_name=?`
	err = db.QueryRow(sqlstring, detail.MedicineName).Scan(&detail.MedicineMasterId)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("insertBillDetail: No matching medicine found")
		} else {
			log.Println("insertBillDetail: Error querying database:", err)
		}
		resp.ErrMsg = "insertingbill details:BDM02" + err.Error()
	}
	// TO GET LOGINID , USERID
	sqlstring1 := `SELECT created_by, login_id FROM medapp_login_history ORDER BY login_history_id DESC LIMIT 1`
	err = db.QueryRow(sqlstring1).Scan(&detail.UserId, &detail.LoginId)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("insertBillDetail: No matching login history found")
		} else {
			log.Println("insertBillDetail: Error querying database:", err)
		}
		resp.ErrMsg = "insertingbill details:BDM03" + err.Error()
	}
	_, err = db.Exec(`
	UPDATE medapp_stock
	SET quantity = quantity - ?
	WHERE medicine_master_id = ?
`, detail.Quantity, detail.MedicineMasterId)
	if err != nil {
		log.Println("insertBillDetail: Error updating stock quantity:", err)
		resp.ErrMsg = "insertingbill details:BDM04" + err.Error()
	}
	coreString := `INSERT INTO medapp_bill_details (bill_no, medicine_master_id, quantity, unit_price, amount, created_by, created_date, updated_by, updated_date)
					VALUES (?, ?, ?, ?, ?, ?, NOW(), ?, NOW())`
	_, err = db.Exec(coreString, detail.Bill_no, detail.MedicineMasterId, detail.Quantity, detail.Unit_price, detail.Amount, detail.UserId, detail.UserId)
	if err != nil {
		log.Println("insertBillDetail: Error inserting into bill details:", err)
		resp.ErrMsg = "insertingbill details:BDM05" + err.Error()
	}

	return resp, nil
}
