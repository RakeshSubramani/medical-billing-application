package dbconnection

import (
	"database/sql"
	"fmt"
	"log"
)

func LocalDBConnect() (*sql.DB, error) {
	log.Println("localDBConnected +")
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", "root", "root", "192.168.2.5", 3306, "rakesh")
	db, err := sql.Open("mysql", connString)
	if err != nil {
		log.Println("open connection failed:", err.Error())
		return db, err
	} else {
		log.Println("local conenected-")
		return db, nil
	}

}