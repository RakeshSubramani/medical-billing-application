package dashboard

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ManagerDashRespStruct struct {
	Errmsg             string  `json:"errmsg"`
	Status             string  `json:"status"`
	Manager_today_sale float64 `json:"managertodaysale"`
	Inventary_value    float64 `json:"inventaryvalue"`
}

func ManagerDashBoardApi(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "GET,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "USER,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")
	log.Println("getData(+)")
	if r.Method == "GET" {
		var resp ManagerDashRespStruct
		resp.Status = "S"
		resp, err := ManagerDashBoard()
		if err != nil {
			log.Println(err)
			resp.Errmsg = "MIA01"+err.Error()
			resp.Status = "E"
		}
		data, err := json.Marshal(resp)
		if err != nil {
			fmt.Fprintf(w, "error taking data"+err.Error())
		} else {
			fmt.Fprintf(w, string(data))
		}
		log.Println("getdata(-)")
	}

}
