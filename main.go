package main

import (
	adduser "medapp_go/Adduser"
	dash "medapp_go/Dash"
	getmedicine "medapp_go/GetMedicine"
	"medapp_go/Login"
	loginhistory "medapp_go/LoginHistory"
	logout "medapp_go/Logout"
	salesreport "medapp_go/SalesReport"
	save "medapp_go/Save"
	stockenrty "medapp_go/StockEnrty"
	userid "medapp_go/Userid"
	"medapp_go/dashboard"
	stockview "medapp_go/stockView"
	"net/http"
)

func main() {
	http.HandleFunc("/loginValidation", Login.LoginValidationApi)
	http.HandleFunc("/loginHistory", loginhistory.LoginHistoryApi)
	http.HandleFunc("/stockView", stockview.StockViewApi)
	http.HandleFunc("/salesReport", salesreport.SalesReportApi)
	http.HandleFunc("/addMedicine", stockenrty.AddStockApi)
	http.HandleFunc("/addUser", adduser.AddUserApi)
	http.HandleFunc("/logOut", logout.LogOutApi)
	http.HandleFunc("/getMedicine", getmedicine.GetMedicineApi)
	http.HandleFunc("/updateStock", stockenrty.StockUpdateApi)
	http.HandleFunc("/dash", dash.Dash)
	http.HandleFunc("/billerDash",dashboard.BillerDashBoardApi)
	http.HandleFunc("/managerDash",dashboard.ManagerDashBoardApi)
	http.HandleFunc("/billDetails",save.BillDetailsApi)
	http.HandleFunc("/billMaster",save.BillMasterApi)
	http.HandleFunc("/billno",dash.BillNo)
	http.HandleFunc("/getUser",userid.GetUserApi)
	http.ListenAndServe(":29014", nil)
}