package logout

import (
	"encoding/json"
	"fmt"
	_"io/ioutil"
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
)

type LogOutReqStruct struct {
	User_id          string `json:"userid"`
	Login_History_Id int    `json:"loginhistoryid"`
}

type LogOutRespStruct struct {
	ErrMsg string `json:"errmsg"`
	Status string `json:"status"`
	Msg    string `json:"msg"`
}

func LogOutApi(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "USER,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")
	log.Println("AddUser: Request received")

	if r.Method == "PUT" {
		var logRec LogOutReqStruct
		var Resp LogOutRespStruct
		Resp.Status = "S"
		Resp.ErrMsg = "Validation done"
		Resp,err:=Logout(logRec);
		if err!=nil{
			Resp.ErrMsg="e"
		}

		data, err := json.Marshal(Resp)
		if err != nil {
			log.Println("Logout: Error marshalling response:LOA01",err.Error())
		} else {
			fmt.Fprintf(w, string(data))
			log.Println("Logout: Response sent")
		}
	}
} 
