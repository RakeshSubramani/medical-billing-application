import axios from 'axios'
const baseApiClient = axios.create({
    baseURL: 'http://localhost:29014',
    withCredentials: false,
    headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json'
    },
})
export default {
    LoginValidation(body) {
        return baseApiClient.put("/loginValidation", body)
    },
    LoginHistory() {
        return baseApiClient.get("/loginHistory")

    },
    StockView() {
        return baseApiClient.get("/stockView")
    },
    SalesReport(body) {
        return baseApiClient.put("/salesReport", body)
    },
    AddMedicine(body) {
        return baseApiClient.put("/addMedicine", body)
    },
    AddUser(body) {
        return baseApiClient.put("/addUser", body)
    },
    LogOut() {
        return baseApiClient.put("/logOut")
    },
    GetMedicine(){
        return baseApiClient.put("/getMedicine")   
    },
    UpdateMedicine(body){
        return baseApiClient.put("/updateStock",body)   

    },
    Dash(){
        return baseApiClient.get("/dash")   

    },
    BillerDash(){
        return baseApiClient.get("/billerDash")   

    },
    ManagerDash(){
        return baseApiClient.get("/managerDash")   

    },
    BillNo(){
        return baseApiClient.get("/billno")   
    },
    BillDetails(body){
        return baseApiClient.put("/billDetails",body)   

    },
    BillMaster(body){
        return baseApiClient.put("/billMaster",body)   

    },
    GetUser(){
        return baseApiClient.put("/getUser")   

    }

}