package dash

import (
	"encoding/json"
	"fmt"
	"log"
	dbconnection "medapp_go/DbConnection"
	"net/http"
)

type DashResp struct {
	Role   string `json:"role"`
	Status string `json:"status"`
	Errmsg string `json:"errmsg"`
}

func Dash(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "GET,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "USER,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")
	log.Println("getData(+)")
	if r.Method == "GET" {
		var resp DashResp
		resp.Status = "S"
		resp.Errmsg="done"
		db, err := dbconnection.LocalDBConnect()
		if err != nil {
			log.Println("Logout: Error connecting to database:", err)
			resp.Status = "E"
			resp.Errmsg = "Error connecting to database: DA01" + err.Error()
		} else {
			defer db.Close()
			//TO GET ROLE
			sqlstring1 := `SELECT role FROM medapp_login_history ORDER BY login_history_id DESC LIMIT 1`
			err := db.QueryRow(sqlstring1).Scan(&resp.Role)
			if err != nil {
				log.Println("Logout: Error connecting to database:", err)
				resp.Status = "E"
				resp.Errmsg = "Error connecting to database:DA02 " + err.Error()
			}
			data, err := json.Marshal(resp)
			if err != nil {
				fmt.Fprintf(w, "error taking data DA04"+err.Error())
			} else {
				fmt.Fprintf(w, string(data))
			}
			log.Println("getdata(-)")
		}

	}
}
