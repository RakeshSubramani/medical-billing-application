package salesreport

import (
	_ "database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type SalesReqStruct struct {
	From_date string `json:"fromdate"`
	To_date   string `json:"todate"`
}

type SalesReportStruct struct {
	Bill_no       int     `json:"billno"`
	Bill_date     string  `json:"billdate"`
	Medicine_name string  `json:"medicinename"`
	Qty           int     `json:"qty"`
	Amount        float64 `json:"amount"`
}
type SalesReportRespStruct struct {
	Errmsg   string              `json:"errmsg"`
	Status   string              `json:"status"`
	SalesArr []SalesReportStruct `json:"salesarr"`
}

func SalesReportApi(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "USER,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")
	log.Println("LoginValidation: Request received")

	if r.Method == "PUT" {
		var salesReq SalesReqStruct
		var salesResp SalesReportRespStruct
		salesResp.Status = "S"
		salesResp.Errmsg = "validation done"

		body, err := ioutil.ReadAll(r.Body)
		// log.Println(body)
		if err != nil {
			log.Println("LoginValidation: Error reading request body:", err)
			salesResp.Status = "E"
			salesResp.Errmsg = "Error reading request body: " + err.Error()
		} else {
			err = json.Unmarshal(body, &salesReq)
			if err != nil {
				log.Println("LoginValidation: Error unmarshalling request body:", err)
				salesResp.Status = "E"
				salesResp.Errmsg = "Error unmarshalling request body:SRA01 " + err.Error()
			} else {
				salesResp, err = SalesReport(salesReq)
				if err != nil {
					salesResp.Status = "E"
				}
			}
		}
		data, err := json.Marshal(salesResp)
		if err != nil {
			log.Println("LoginValidation: Error marshalling response:SRA02", err.Error())
		} else {
			fmt.Fprintf(w, string(data))
			log.Println("LoginValidation: Response sent")
		}

	}

}
