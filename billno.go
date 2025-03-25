package dash

import (
	"encoding/json"
	"fmt"
	"log"
	dbconnection "medapp_go/DbConnection"
	"net/http"
)

type BillResp struct {
	BillNo int `json:"billno"`
	Status string `json:"status"`
	Errmsg string `json:"errmsg"`
}
func BillNo(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "GET,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "USER,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")
	log.Println("getData(+)")
	if r.Method == "GET" {
		var resp1 BillResp
		resp1.Status = "S"
		resp1.Errmsg="done"
		db, err := dbconnection.LocalDBConnect()
		if err != nil {
			log.Println("Logout: Error connecting to database:", err)
			resp1.Status = "E"
			resp1.Errmsg = "Error connecting to database: BA01" + err.Error()
		} else {
			defer db.Close()
			//TO GET BILLNO
			sqlstring := `SELECT bill_no FROM medapp_bill_master ORDER BY bill_master_id DESC LIMIT 1`
			log.Println(sqlstring)
			err := db.QueryRow(sqlstring).Scan(&resp1.BillNo)
			log.Println(resp1.BillNo)
			if err != nil {
				log.Println("Logout: Error connecting to database:", err)
				resp1.Status = "E"
				resp1.Errmsg = "Error connecting to database:BA02 " + err.Error()
			}
			data, err := json.Marshal(resp1)
			if err != nil {
				fmt.Fprintf(w, "error taking data BA03"+err.Error())
			} else {
				fmt.Fprintf(w, string(data))
			}
			log.Println("getdata(-)")
		}

	}
}