package userid
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)
type UserStruct struct{
	UserId string `json:"UserId"`
	// Brand string `json:"brand"`
}
type UserRespStruct struct{
Errmsg string `json:"errmsg"`
Status string `json:"status"`
User []UserStruct`json:"user"`

}

func GetUserApi(w http.ResponseWriter, r *http.Request){
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "USER,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")
	log.Println("getData(+)")
	if r.Method == "PUT" {
		log.Println("inside tghe api")
		var resp UserRespStruct
		resp.Status = "S"
		resp,err:=getUser()
		if err!=nil{
			log.Println(err)
			resp.Errmsg = "UA01"+err.Error()
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

