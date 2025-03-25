package logout

import (
	"database/sql"
	"log"
	dbconnection "medapp_go/DbConnection"
)
func Logout(logRec LogOutReqStruct)(LogOutRespStruct,error){
	var Resp LogOutRespStruct
	db, err := dbconnection.LocalDBConnect()
		if err != nil {
			log.Println("Logout: Error connecting to database:", err)
			Resp.Status = "E"
			Resp.ErrMsg = "Error connecting to database: " + err.Error()
		} else {
			defer db.Close()
			//TO GET USERID AND LOGINHISTORYID
			sqlstring1 := `SELECT created_by,login_history_id FROM medapp_login_history ORDER BY login_history_id DESC LIMIT 1`
			err := db.QueryRow(sqlstring1).Scan(&logRec.User_id, &logRec.Login_History_Id)
			if err != nil {
				if err == sql.ErrNoRows {
					log.Println("Logout: No matching user found")
					Resp.Status = "E"
					Resp.ErrMsg = "No matching user found"
				} else {
					log.Println("Logout: Error querying database:", err)
					Resp.Status = "E"
					Resp.ErrMsg = "Error querying database:LOM01 " + err.Error()
				}
			} else {
				log.Println(logRec.User_id)
				//TO UPDATE THE LOGINHISTORY 
				sqlstring := `Update medapp_login_history set logout_date=now(),logout_time=now(),updated_by=?,updated_date=now()
						where login_history_id=?`
				_, err := db.Exec(sqlstring, logRec.User_id, logRec.Login_History_Id)
				if err != nil {
					log.Println("Logout: Error inserting into database:", err)
					Resp.Status = "E"
					Resp.ErrMsg = "Error inserting into database: LOM02" + err.Error()
				}
			}
		}
return Resp,nil
	}