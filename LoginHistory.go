package loginhistory

import (
	"log"
	dbconnection "medapp_go/DbConnection"
)

func LoginHistory() (LoginHistoryRespStruct, error) {
	var resp LoginHistoryRespStruct
	var historyrec LoginHistoryStruct
	db, err := dbconnection.LocalDBConnect()
	if err != nil {
		log.Println(err)
		resp.Errmsg = err.Error()
		resp.Status = "E"

	} else {
		defer db.Close()
		//TO GET LOGINHISTORY
		coreString := `SELECT 
		u.user_id, 
		l.login_date, 
		l.login_time, 
		IFNULL(l.logout_date, 'working') AS logout_date, 
		IFNULL(l.logout_time, 'working') AS logout_time
	FROM 
		medapp_login_history AS l
	INNER JOIN 
		medapp_login AS u 
	ON 
		l.login_id = u.login_id
	WHERE 
		l.login_date = CURDATE()
	ORDER BY 
		l.login_time;
	`
		rows, err := db.Query(coreString)
		if err != nil {
			log.Println(err)
			resp.Errmsg = "LHMO1" + err.Error()
			resp.Status = "E"
		} else {
			for rows.Next() {
				err := rows.Scan(&historyrec.User_id, &historyrec.Login_date, &historyrec.Login_time, &historyrec.Logout_date, &historyrec.Logout_time)
				if err != nil {
					log.Println(err)
					resp.Errmsg = err.Error()
					resp.Status = "E"
				} else {
					resp.LoginHistoryArr = append(resp.LoginHistoryArr, historyrec)
				}
			}
		}
	}
	return resp, nil
}
