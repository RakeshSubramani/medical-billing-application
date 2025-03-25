package salesreport

import (
	"database/sql"
	"log"
	dbconnection "medapp_go/DbConnection"
)

func SalesReport(salesReq SalesReqStruct) (SalesReportRespStruct, error) {
	var salesReport SalesReportStruct
	var salesResp SalesReportRespStruct
	log.Println(salesReq)
	db, err := dbconnection.LocalDBConnect()
	if err != nil {
		log.Println("LoginValidation: Error connecting to database:", err)
		salesResp.Status = "E"
		salesResp.Errmsg = "Error connecting to database: SRM01" + err.Error()
	} else {
		defer db.Close()
		// TO GET SALES REPORT
		sqlstring := `Select bill.bill_no,master.bill_date,med.medicine_name,bill.quantity,bill.amount
					from medapp_bill_details  bill inner join medapp_bill_master  master
					on bill.bill_no = master.bill_no  inner join medapp_medicine_master med
					on bill.medicine_master_id=med.medicine_master_id and
					 master.bill_date between ? And ?`
		rows, err := db.Query(sqlstring, salesReq.From_date, salesReq.To_date)
		if err != nil {
			if err == sql.ErrNoRows {

				log.Println("LoginValidation: No matching user found")
				salesResp.Status = "E"
				salesResp.Errmsg = "No matching user found"
			} else {
				log.Println("LoginValidation: Error querying database:", err)
				salesResp.Status = "E"
				salesResp.Errmsg = "Error querying database: SRM02" + err.Error()
			}
		} else {
			for rows.Next() {
				err := rows.Scan(&salesReport.Bill_no, &salesReport.Bill_date, &salesReport.Medicine_name, &salesReport.Qty, &salesReport.Amount)
				if err != nil {
					log.Println(err)
					salesResp.Errmsg = err.Error()
					salesResp.Status = "E"
				} else {
					salesResp.SalesArr = append(salesResp.SalesArr, salesReport)
				}
			}

		}
	}
	return salesResp, nil
}
