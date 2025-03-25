package Login

import (
	"database/sql"
	"log"
	dbconnection "medapp_go/DbConnection"
)

func LoginValidation(loginRec LoginReqStruct, loginStructRec LoginStruct) (LoginRespStruct, error) {
	var loginrespRec LoginRespStruct
	db, err := dbconnection.LocalDBConnect()
	if err != nil {
		log.Println("LoginValidation: Error connecting to database:", err)
		loginrespRec.Status = "E"
		loginrespRec.ErrMsg = "Error connecting to database:LM01" + err.Error()
	} else {
		defer db.Close()
    //TO GET LOGINID,USERID,ROLE
		sqlstring := `SELECT login_id, user_id, role FROM medapp_login WHERE user_id=? AND password=?`
		err := db.QueryRow(sqlstring, loginRec.User_id, loginRec.Password).Scan(&loginStructRec.Login_id, &loginStructRec.User_id, &loginStructRec.Role)
		if err != nil {
			if err == sql.ErrNoRows {

				log.Println("LoginValidation: No matching user found")
				loginrespRec.Status = "E"
				loginrespRec.ErrMsg = "No matching user found"
			} else {

				log.Println("LoginValidation: Error querying database:", err)
				loginrespRec.Status = "E"
				loginrespRec.ErrMsg = "Error querying database: LMO2" + err.Error()
			}
		} else {
			loginrespRec.Role = loginStructRec.Role
			//TO INSERT INTO LOGINHISTORY
			coreString := `INSERT INTO medapp_login_history(login_id, login_date, login_time, created_by, created_date,user_id,password,role) VALUES (?, CURDATE(), CURTIME(), ?, NOW(),?,?,?)`
			_, err := db.Exec(coreString, loginStructRec.Login_id, loginStructRec.User_id, loginStructRec.User_id, loginRec.Password, loginStructRec.Role)
			if err != nil {
				log.Println("LoginValidation: Error inserting into login history:", err)
				loginrespRec.Status = "E"
				loginrespRec.ErrMsg = "Error inserting into login history: " + err.Error()
			}
		}
	}
	return loginrespRec,nil
}
