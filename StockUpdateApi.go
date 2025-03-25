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

type StockUpdateReqStruct struct {
	Medicine_Name      string  `json:"medicinename"`
	User_id            string  `json:"userid"`
	Quantity           int     `json:"quantity"`
	Unit_price         float64 `json:"unitprice"`
	Medicine_master_id int     `json:"medicinemasterid"`
}
type StockUpdateRespStruct struct {
	ErrMsg string `json:"errmsg"`
	Status string `json:"status"`
	Msg    string `json:"msg"`
}

func StockUpdateApi(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "USER,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")
	log.Println("AddStock: Request received")

	if r.Method == "PUT" {
		var UpdateRec StockUpdateReqStruct
		var Resp StockUpdateRespStruct
		Resp.Status = "S"
		Resp.ErrMsg = "Validation done"

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("AddStock: Error reading request body:", err)
			Resp.Status = "E"
			Resp.ErrMsg = "Error reading request body: SUA01" + err.Error()
		} else {
			err = json.Unmarshal(body, &UpdateRec)
			if err != nil {
				log.Println("AddStock: Error unmarshalling request body:", err)
				Resp.Status = "E"
				Resp.ErrMsg = "Error unmarshalling request body: SUA02" + err.Error()
			} else {
				Resp,err=StockUpdate(UpdateRec)
				if err!=nil{
					Resp.ErrMsg = "Error unmarshalling request body: SUA03" + err.Error()

				}
				
			}
		}

		data, err := json.Marshal(Resp)
		if err != nil {
			log.Println("AddStock: Error marshalling response:", err)
		} else {
			fmt.Fprintf(w, string(data))
			log.Println("AddStock: Response sent")
		}
	}
}
