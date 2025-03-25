package stockview

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type StockViewStruct struct {
	Medicine_name string  `json:"medicinename"`
	Brand         string  `json:"brand"`
	Quantity      int     `json:"quantity"`
	Unit_price    float64 `json:"unitprice"`
}
type StockViewRespStruct struct {
	Errmsg       string            `json:"errmsg"`
	Status       string            `json:"status"`
	StockviewArr []StockViewStruct `json:"stockviewArr"`
}

func StockViewApi(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "GET,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "USER,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")
	log.Println("getData(+)")
	if r.Method == "GET" {
		var resp StockViewRespStruct
		resp.Status = "S"
		resp, err := StockView()
		if err != nil {
			log.Println(err)
			resp.Errmsg = "SVA01"+err.Error()
			resp.Status = "E"
		}
		data, err := json.Marshal(resp)
		if err != nil {
			fmt.Fprintf(w, "error taking data:SVA02"+err.Error())
		} else {
			fmt.Fprintf(w, string(data))
		}
		log.Println("getdata(-)")
	}

}
