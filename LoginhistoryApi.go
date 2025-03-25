package loginhistory

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type LoginHistoryStruct struct {
	User_id     string `json:"user_id"`
	Login_date  string `json:"logindate"`
	Login_time  string `json:"logintime"`
	Logout_date string `json:"logoutdate"`
	Logout_time string `json:"logouttime"`
}
type LoginHistoryRespStruct struct {
	Errmsg          string               `json:"errmsg"`
	Status          string               `json:"status"`
	LoginHistoryArr []LoginHistoryStruct `json:"loginhistory_arr"`
}

func LoginHistoryApi(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "GET,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "USER,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")
	log.Println("getData(+)")
	if r.Method == "GET" {
		var resp LoginHistoryRespStruct
		resp.Status = "S"
		resp, err := LoginHistory()
		if err != nil {
			log.Println(err)
			resp.Errmsg = err.Error()
			resp.Status = "E"
		} else {
			data, err := json.Marshal(resp)
			if err != nil {
				fmt.Fprintf(w, "error taking data LHA00"+err.Error())
			} else {
				fmt.Fprintf(w, string(data))
			}
			log.Println("getdata(-)")
		}

	}

}
