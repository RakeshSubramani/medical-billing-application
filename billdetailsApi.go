// Please adjust the package name accordingly
package save

import (
	_"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	
	_ "github.com/go-sql-driver/mysql"
)

type BillDetailResp struct {
	ErrMsg string `json:"errmsg"`
	Status string `json:"status"`
	Msg    string `json:"msg"`
}

type BillDetailReqStruct struct {
	Bill_no          int     `json:"billno"`
	Quantity         int     `json:"quantity"`
	Unit_price       float64 `json:"unitprice"`
	Amount           float64 `json:"amount"`
	LoginId          int     `json:"loginid"`
	UserId           string  `json:"userid"`
	MedicineMasterId int     `json:"medicinemasterid"`
	MedicineName     string  `json:"medicinename"`
}

type BillDetailReq struct {
	BillDetail []BillDetailReqStruct `json:"billdetail"`
}

func BillDetailsApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "USER,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")
	log.Println("BillDetails: Request received")

	if r.Method == "PUT" {
		var Detailarr BillDetailReq
		var resp BillDetailResp
		resp.Status = "S"
		resp.ErrMsg = "validation done"

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("BillDetails: Error reading request body:", err)
			resp.Status = "E"
			resp.ErrMsg = "Error reading request body: BD01" + err.Error()
		} else {
			err = json.Unmarshal(body, &Detailarr)
			log.Println(Detailarr)
			if err != nil {
				log.Println("BillDetails: Error unmarshalling request body:", err)
				resp.Status = "E"
				resp.ErrMsg = "Error unmarshalling request body: BD02" + err.Error()
			} else {

				for _, DetailRec := range Detailarr.BillDetail {
					resp,err := BillDetail(DetailRec)
					if err != nil {
						log.Println("BillDetails: Error inserting bill detail:", err)
						resp.Status = "E"
						resp.ErrMsg = "Error inserting bill detail:BD03 " + err.Error()
						break
					}
				}

			}
		}

		data, err := json.Marshal(resp)
		if err != nil {
			log.Println("BillDetails: Error marshalling response:BD04", err.Error())
		} else {
			fmt.Fprintf(w, string(data))
		}

	}
}
