package dashboard

import (
	"log"
	dbconnection "medapp_go/DbConnection"
)

func BillerDashBoard() (BillerDashRespStruct, error) {
	var resp BillerDashRespStruct
	var req BillerReqStruct

	db, err := dbconnection.LocalDBConnect()
	if err != nil {
		log.Println(err)
		resp.Errmsg = "BIM01"+err.Error()
		resp.Status = "E"
		return resp, err
	}
	defer db.Close()

	// To get the latest login ID
	sqlstring1 := `SELECT Login_id FROM medapp_login_history ORDER BY login_history_id DESC LIMIT 1`
	err = db.QueryRow(sqlstring1).Scan(&req.Login_id)
	if err != nil {
		log.Println(err)
		resp.Errmsg = "BIM02"+err.Error()
		resp.Status = "E"
		return resp, err
	}

	// To get today's sales
	coreString := `SELECT IFNULL(SUM(bill_amount), 0) FROM medapp_bill_master WHERE DATE(bill_date) = CURDATE() AND login_id = ?`
	err = db.QueryRow(coreString, req.Login_id).Scan(&resp.Biller_today_sale)
	if err != nil {
		log.Println(err)
		resp.Errmsg = "BIM03"+err.Error()
		resp.Status = "E"
		return resp, err
	}

	// To get yesterday's sales
	coreString1 := `SELECT IFNULL(SUM(bill_amount), 0) FROM medapp_bill_master WHERE DATE(bill_date) = DATE_SUB(CURDATE(), INTERVAL 1 DAY) AND login_id = ?`
	err = db.QueryRow(coreString1, req.Login_id).Scan(&resp.Biller_Yesterday_sale)
	if err != nil {
		log.Println(err)
		resp.Errmsg = "BIM04"+err.Error()
		resp.Status = "E"
		return resp, err
	}

	// resp.Status = "S"
	return resp, nil
}
