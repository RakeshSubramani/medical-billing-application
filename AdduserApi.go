package adduser

import (
	_"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
)

type AdduserReqStruct struct {
	User_id         string `json:"userid"`
	Password        string `json:"password"`
	Role            string `json:"role"`
	Created_User_id string `json:"createduserid"`
}

type AdduserRespStruct struct {
	ErrMsg string `json:"errmsg"`
	Status string `json:"status"`
	Msg    string `json:"msg"`
}

func AddUserApi(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "USER,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")
	log.Println("AddUser: Request received")

	if r.Method == "PUT" {
		var AddRec AdduserReqStruct
		var Resp AdduserRespStruct
		Resp.Status = "S"
		Resp.ErrMsg = "Validation done"

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("AddUser: Error reading request body:", err)
			Resp.Status = "E"
			Resp.ErrMsg = "Error reading request body: AUA01" + err.Error()
		} else {
			err = json.Unmarshal(body, &AddRec)
			if err != nil {
				log.Println("AddUser: Error unmarshalling request body:", err)
				Resp.Status = "E"
				Resp.ErrMsg = "Error unmarshalling request body: AUA01" + err.Error()
			} else {
				log.Println("AddUser: Request body:", AddRec)
				Resp,err=AddUser(AddRec)
			}
		}

		data, err := json.Marshal(Resp)
		if err != nil {
			log.Println("AddUser: Error marshalling response:", err)
		} else {
			fmt.Fprintf(w, string(data))
			log.Println("AddUser: Response sent")
		}
	}
}
