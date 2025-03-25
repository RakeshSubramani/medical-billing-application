package Login

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"net/http"
)

type LoginReqStruct struct {
	User_id  string `json:"user_id"`
	Password string `json:"password"`
}

type LoginStruct struct {
	User_id  string `json:"user_id"`
	Role     string `json:"role"`
	Login_id int    `json:"login_id"`
}

type LoginRespStruct struct {
	ErrMsg string `json:"errmsg"`
	Status string `json:"status"`
	Role   string `json:"role"`
}

func LoginValidationApi(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "USER,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")
	log.Println("LoginValidation: Request received")

	if r.Method == "PUT" {
		var loginRec LoginReqStruct
		var loginrespRec LoginRespStruct
		var loginStructRec LoginStruct
		loginrespRec.Status = "S"
		loginrespRec.ErrMsg = "validation done"
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("LoginValidation: Error reading request body:", err)
			loginrespRec.Status = "E"
			loginrespRec.ErrMsg = "Error reading request body:LA01 " + err.Error()
		} else {
			err = json.Unmarshal(body, &loginRec)
			if err != nil {
				log.Println("LoginValidation: Error unmarshalling request body:", err)
				loginrespRec.Status = "E"
				loginrespRec.ErrMsg = "Error unmarshalling request body: LA02" + err.Error()
			} else {
				log.Println(loginRec)

				loginrespRec, err = LoginValidation(loginRec, loginStructRec)
				if err != nil {
					loginrespRec.Status = "E"
				}
			}
		}

		data, err := json.Marshal(loginrespRec)
		if err != nil {
			log.Println("LoginValidation: Error marshalling response: LAO3", err)
		} else {
			fmt.Fprintf(w, string(data))
			log.Println("LoginValidation: Response sent")
		}

	}
}
