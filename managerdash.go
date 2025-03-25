package dashboard

import (
	"log"
	dbconnection "medapp_go/DbConnection"
)

func ManagerDashBoard() (ManagerDashRespStruct, error) {
	var resp ManagerDashRespStruct
	db, err := dbconnection.LocalDBConnect()
	if err != nil {
		log.Println(err)
		resp.Errmsg = err.Error()
		resp.Status = "E"

	} else {
		defer db.Close()
		//TO GET TODAY SALE
		coreString := `SELECT IFNULL(SUM(bill_amount), 0) FROM medapp_bill_master WHERE DATE(bill_date) = CURDATE()`
		err := db.QueryRow(coreString).Scan(&resp.Manager_today_sale)
		if err != nil {
			log.Println(err)
			resp.Errmsg = "MIM01"+err.Error()
			resp.Status = "E"
		} else {
			//TO GET SUM OF UNITPRICE
			coreString1 := `Select sum(unit_price*quantity)
				from medapp_stock`
			err := db.QueryRow(coreString1).Scan(&resp.Inventary_value)
			if err != nil {
				log.Println(err)
				resp.Errmsg = "MIM02"+err.Error()
				resp.Status = "E"
			}
		}
	}
	return resp, nil
}
