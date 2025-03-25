package stockview

import (
	"log"
	dbconnection "medapp_go/DbConnection"
)

func StockView() (StockViewRespStruct, error) {
	var stockRec StockViewStruct
	var resp StockViewRespStruct
	db, err := dbconnection.LocalDBConnect()
	if err != nil {
		log.Println(err)
		resp.Errmsg = "SVM01"+err.Error()
		resp.Status = "E"

	} else {
		defer db.Close()
		//TO GET STOCK VIEW 
		coreString := `Select med.medicine_name,med.brand,stock.quantity,stock.unit_price
			from medapp_medicine_master  med inner join medapp_stock stock
			on med.medicine_master_id= stock.medicine_master_id`
		rows, err := db.Query(coreString)
		if err != nil {
			log.Println(err)
			resp.Errmsg = "SVM02"+err.Error()
			resp.Status = "E"
		} else {
			for rows.Next() {
				err := rows.Scan(&stockRec.Medicine_name, &stockRec.Brand, &stockRec.Quantity, &stockRec.Unit_price)
				if err != nil {
					log.Println(err)
					resp.Errmsg = "SVM03"+err.Error()
					resp.Status = "E"
				} else {
					resp.StockviewArr = append(resp.StockviewArr, stockRec)
				}
			}
		}
	}
	return resp, nil
}
