package getmedicine	
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)
type MedicineStruct struct{
	MedicineName string `json:"medicinename"`
	Brand string `json:"brand"`
}
type MedicineRespStruct struct{
Errmsg string `json:"errmsg"`
Status string `json:"status"`
Medicine []MedicineStruct `json:"medicine"`

}

func GetMedicineApi(w http.ResponseWriter, r *http.Request){
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "USER,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")
	log.Println("getData(+)")
	if r.Method == "PUT" {
		log.Println("inside tghe api")
		var resp MedicineRespStruct
		resp.Status = "S"
		resp,err:=getMedicine()
		if err!=nil{
			log.Println(err)
			resp.Errmsg = "MA01"+err.Error()
			resp.Status = "E"
		}else{

			data, err := json.Marshal(resp)
			if err != nil {
				fmt.Fprintf(w, "error taking data"+err.Error())
			} else {
				fmt.Fprintf(w, string(data))
			}
			log.Println("getdata(-)")
		}
	}

}

