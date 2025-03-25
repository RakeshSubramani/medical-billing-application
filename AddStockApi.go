package stockenrty

import (
	_"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type StockEntryReqStruct struct {
	MedicineName      string  `json:"medicineName"`
	Brand             string  `json:"brand"`
	UserID            string  `json:"userID"`
	Quantity          int     `json:"quantity"`
	UnitPrice         float64 `json:"unitPrice"`
	MedicineMasterID  int     `json:"medicineMasterID"`
}

type StockStruct struct {
	MedicineName string `json:"medicineName"`
	Brand        string `json:"brand"`
}

type StockEntryRespStruct struct {
	ErrMsg   string       `json:"errMsg"`
	Status   string       `json:"status"`
	StockArr []StockStruct `json:"stockArr"`
}

func AddStockApi(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "USER,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")
	log.Println("AddStock: Request received")

	if r.Method == "PUT" {
		var stockRec StockEntryReqStruct
		var stockResp StockEntryRespStruct
		stockResp.Status = "S"
		stockResp.ErrMsg = "Validation done"

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("AddStock: Error reading request body:", err)
			stockResp.Status = "E"
			stockResp.ErrMsg = "Error reading request body: AUA01" + err.Error()
		} else {
			err = json.Unmarshal(body, &stockRec)
			if err != nil {
				log.Println("AddStock: Error unmarshalling request body:", err)
				stockResp.Status = "E"
				stockResp.ErrMsg = "Error unmarshalling request body: AUA02" + err.Error()
			} else {
				stockResp,err=AddStock(stockRec);
				if err!=nil{
				stockResp.ErrMsg = "Error unmarshalling request body: AUA03" + err.Error()
					
				}
			}

			data, err := json.Marshal(stockResp)
			if err != nil {
				log.Println("AddStock: Error marshalling response:", err)
			} else {
				fmt.Fprintf(w, string(data))
				log.Println("AddStock: Response sent")
			}
		}
	}
}
