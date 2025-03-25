package stockenrty

import (
	"database/sql"
	"log"
	dbconnection "medapp_go/DbConnection"
)

func AddStock(stockRec StockEntryReqStruct)(StockEntryRespStruct,error){
	var stockResp StockEntryRespStruct
	log.Println(stockRec)
	db, err := dbconnection.LocalDBConnect()
	if err != nil {
		log.Println("AddStock: Error connecting to database:", err)
		stockResp.Status = "E"
		stockResp.ErrMsg = "Error connecting to database:AUM01 " + err.Error()
	} else {
		defer db.Close()
		// To Get Userid
		var userID string
		err := db.QueryRow("SELECT created_by FROM medapp_login_history ORDER BY login_history_id DESC LIMIT 1").Scan(&userID)
		if err != nil {
			if err == sql.ErrNoRows {
				log.Println("AddStock: No matching user found")
				stockResp.Status = "E"
				stockResp.ErrMsg = "No matching user found"
			} else {
				log.Println("AddStock: Error querying database:", err)
				stockResp.Status = "E"
				stockResp.ErrMsg = "Error querying database: AUM02" + err.Error()
			}
		} else {
			// Insert medicine master 
			_, err := db.Exec("INSERT INTO medapp_medicine_master (medicine_name, brand, created_by, created_date) SELECT ?, ?, ?, NOW() WHERE NOT EXISTS (SELECT medicine_master_id FROM medapp_medicine_master WHERE medicine_name = ? AND brand = ?)",
				stockRec.MedicineName, stockRec.Brand, userID, stockRec.MedicineName, stockRec.Brand)
			if err != nil {
				log.Println("AddStock: Error inserting into medapp_medicine_master:", err)
				stockResp.Status = "E"
				stockResp.ErrMsg = "Error inserting into medapp_medicine_master: AUM03" + err.Error()
			} else {
				log.Println("AddStock: Medicine master inserted successfully")
			}

			// Get  medicine master ID
			err = db.QueryRow("SELECT medicine_master_id FROM medapp_medicine_master ORDER BY medicine_master_id DESC LIMIT 1").Scan(&stockRec.MedicineMasterID)
			if err != nil {
				if err == sql.ErrNoRows {
					log.Println("AddStock: No matching medicine master found")
					stockResp.Status = "E"
					stockResp.ErrMsg = "No matching medicine master found"
				} else {
					log.Println("AddStock: Error querying database for medicine master ID:", err)
					stockResp.Status = "E"
					stockResp.ErrMsg = "Error querying database for medicine master ID: AUM04" + err.Error()
				}
			} else {
				log.Println("AddStock: Medicine master ID retrieved successfully")
			}

			// Insert  into stock entry
			result, err := db.Exec("INSERT INTO medapp_stock (medicine_master_id, quantity, unit_price, created_by, created_date, updated_by, updated_date) SELECT ?, ?, ?, ?, NOW(), ?, NOW() FROM DUAL WHERE EXISTS (SELECT medicine_master_id FROM medapp_medicine_master WHERE medicine_master_id = ?)",
				stockRec.MedicineMasterID, stockRec.Quantity, stockRec.UnitPrice, userID, userID, stockRec.MedicineMasterID)
			if err != nil {
				log.Println("AddStock: Error inserting into medapp_stock:", err)
				stockResp.Status = "E"
				stockResp.ErrMsg = "Error inserting into medapp_stock: AUM05" + err.Error()
			} else {
				rowsAffected, _ := result.RowsAffected()
				if rowsAffected == 0 {
					log.Println("AddStock: No rows affected, medicine master not found")
					stockResp.Status = "E"
					stockResp.ErrMsg = "No rows affected, medicine master not found"
				} else {
					log.Println("AddStock: Stock entry inserted successfully")
				}
			}
		}
	}
	return stockResp,nil
}