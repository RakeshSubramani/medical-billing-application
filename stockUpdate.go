package stockenrty

import (
	"database/sql"
	"log"
	dbconnection "medapp_go/DbConnection"
)

func StockUpdate(UpdateRec StockUpdateReqStruct) (StockUpdateRespStruct, error) {
	var Resp StockUpdateRespStruct
	log.Println(UpdateRec)
	db, err := dbconnection.LocalDBConnect()
	if err != nil {
		log.Println("StockUpdate: Error connecting to database:", err)
		Resp.Status = "E"
		Resp.ErrMsg = "Error unmarshalling request body: SUA02" + err.Error()
		Resp.ErrMsg = "Error connecting to database: SUM01" + err.Error()
		return Resp, err
	}
	defer db.Close()

	sqlstring1 := `SELECT created_by FROM medapp_login_history ORDER BY login_history_id DESC LIMIT 1`
	err = db.QueryRow(sqlstring1).Scan(&UpdateRec.User_id)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("StockUpdate: No matching user found")
			Resp.Status = "E"
			Resp.ErrMsg = "No matching user found"
			return Resp, nil
		}
		log.Println("StockUpdate: Error querying database:", err)
		Resp.Status = "E"
		Resp.ErrMsg = "Error querying database:SUM02 " + err.Error()
		return Resp, err
	}

	sqlstring2 := `SELECT medicine_master_id FROM medapp_medicine_master WHERE medicine_name=?`
	err = db.QueryRow(sqlstring2, UpdateRec.Medicine_Name).Scan(&UpdateRec.Medicine_master_id)
	if err != nil {
		log.Println(err)
		if err == sql.ErrNoRows {
			log.Println("StockUpdate: No matching medicine found")
			Resp.Status = "E"
			Resp.ErrMsg = "No matching medicine found 2"
			return Resp, nil
		}
		log.Println("StockUpdate: Error querying database:", err)
		Resp.Status = "E"
		Resp.ErrMsg = "Error querying database: SUM03" + err.Error()
		return Resp, err
	}

	sqlstring := `UPDATE medapp_stock SET quantity = quantity+?, unit_price=?, updated_by=?, updated_date=NOW() WHERE medicine_master_id=?`
	_, err = db.Exec(sqlstring, UpdateRec.Quantity, UpdateRec.Unit_price, UpdateRec.User_id, UpdateRec.Medicine_master_id)
	if err != nil {
		log.Println("StockUpdate: Error updating database:", err)
		Resp.Status = "E"
		Resp.ErrMsg = "Error updating database:SUM04 " + err.Error()
		return Resp, err
	}

	Resp.Status = "S"
	Resp.ErrMsg = "Stock updated successfully"
	return Resp, nil
}
