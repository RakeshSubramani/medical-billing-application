package adduser

import (
	"database/sql"
	"log"
	dbconnection "medapp_go/DbConnection"
)
func AddUser(AddRec AdduserReqStruct)(AdduserRespStruct,error){
	var Resp AdduserRespStruct
	db, err := dbconnection.LocalDBConnect()
				if err != nil {
					log.Println("AddUser: Error connecting to database:", err)
					Resp.Status = "E"
					Resp.ErrMsg = "Error connecting to database: AUM01" + err.Error()
				} else {
					defer db.Close()
					//TO GET USERID
					sqlstring1 := `SELECT created_by FROM medapp_login_history ORDER BY login_history_id DESC LIMIT 1`
					err := db.QueryRow(sqlstring1).Scan(&AddRec.Created_User_id)
					if err != nil {
						if err == sql.ErrNoRows {
							log.Println("AddUser: No matching user found")
							Resp.Status = "E"
							Resp.ErrMsg = "No matching user found"
						} else {
							log.Println("AddUser: Error querying database:", err)
							Resp.Status = "E"
							Resp.ErrMsg = "Error querying database:AUM03 " + err.Error()
						}
					} else {
                        log.Println(AddRec.User_id)
						//TO INSERT THE NEW USER
						sqlstring := `INSERT INTO medapp_login(user_id, password, role, created_by, created_date, updated_by, updated_date) 
                                      SELECT ?, ?, ?, ?, NOW(), ?, NOW() 
                                      FROM dual 
                                      WHERE NOT EXISTS (SELECT login_id FROM medapp_login WHERE user_id = ?)`
						_, err := db.Exec(sqlstring, AddRec.User_id, AddRec.Password, AddRec.Role, AddRec.Created_User_id, AddRec.Created_User_id, AddRec.User_id)
						if err != nil {
							log.Println("AddUser: Error inserting into database:", err)
							Resp.Status = "E"
							Resp.ErrMsg = "Error inserting into database: " + err.Error()
						}
					}
				}
				return Resp,nil
}