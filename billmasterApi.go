package save

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	

	_ "github.com/go-sql-driver/mysql"
)

type BillMasterResp struct {
	ErrMsg string `json:"errmsg"`
	Status string `json:"status"`
	Msg    string `json:"msg"`
}

type BillMasterReqStruct struct {
	Bill_no     int     `json:"billno"`
	Bill_date   string  `json:"billdate"`
	Bill_amount float64 `json:"billamount"`
	Bill_gst    float64 `json:"billgst"`
	Net_price   float64 `json:"netprice"`
	Login_id    int     `json:"loginid"`
	User_id     string  `json:"userid"`
}

type BillMasterReq struct {
	BillMaster []BillMasterReqStruct `json:"billmaster"`
}

func BillMasterApi(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "USER,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")
	log.Println("BillMaster: Request received")

	if r.Method == "PUT" {
		var masterarr BillMasterReq
		var resp BillMasterResp
		resp.Status = "S"
		resp.ErrMsg = "validation done"

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("BillMaster: Error reading request body:", err)
			resp.Status = "E"
			resp.ErrMsg = "Error reading request body: BMA01" + err.Error()
		} else {
			err = json.Unmarshal(body, &masterarr)
			log.Println(masterarr)
			if err != nil {
				log.Println("BillMaster: Error unmarshalling request body:", err)
				resp.Status = "E"
				resp.ErrMsg = "Error unmarshalling request body: BMA02" + err.Error()
			} else {

				for _, masterRec := range masterarr.BillMaster {
					err := insertBillMaster(masterRec)
					if err != nil {
						log.Println("BillMaster: Error inserting bill master:", err)
						resp.Status = "E"
						resp.ErrMsg = "Error inserting bill master:BMA03 " + err.Error()
						break
					}
				}
			}
		}

		data, err := json.Marshal(resp)
		if err != nil {
			log.Println("BillMaster: Error marshalling response:BMA04", err)
			
		}else{
			fmt.Fprintf(w, string(data))
		}
	}
}
