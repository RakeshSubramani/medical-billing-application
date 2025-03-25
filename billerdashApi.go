package dashboard

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)
type BillerReqStruct struct{
	Login_id int `json:"loginid"`
}
type BillerDashRespStruct struct {
	Errmsg                string  `json:"errmsg"`
	Status                string  `json:"status"`
	Biller_today_sale     float64 `json:"billertodaysale"`
	Biller_Yesterday_sale float64 `json:"yesterdaysale"`
}

func BillerDashBoardApi(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "GET,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "USER,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")
	log.Println("getData(+)")
	if r.Method == "GET" {
		// var historyrec LoginHistoryStruct
		var resp BillerDashRespStruct
		resp.Status = "S"
        resp,err:=BillerDashBoard()
		if err!=nil{
			log.Println(err)
			resp.Errmsg = "BI01"+err.Error()
			resp.Status = "E"
		}
		data, err := json.Marshal(resp)
		if err != nil {
			fmt.Fprintf(w, "error taking data"+err.Error())
		} else {
			fmt.Fprintf(w, string(data))
		}
		log.Println("getdata(-)")
	}

}
