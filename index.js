import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
  //   login:[{userId:"mi2",password:"mi2@",role:"biller"}
  //           ,{userId:"mi1",password:"mi1@",role:"manager"},
  //              {userId:"mi3",password:"mi3@",role:"systemAdmin"},
  //                 {userId:"mi4",password:"mi4@",role:"inventry"}],
  //   loginHistory1:[],
  //   medicineMaster:[{medicineName:"crosin",brand:"cerosine"},
  //                      {medicineName:"dolo",brand:"dologs"}],
  //   stock:[{medicineName:"dolo",quantity:20,amount:1000},{medicineName:"crosin",quantity:20,amount:2000}],
  //   billMaster:[{ billNo:99, date:"2024-03-03", billAmount:2000, billGst: 289, netPrice: 2289,userId: "mi2" }],
  //   billDetails:[{ billNo:99, medicineName:"dolo", quantity:4, unitPrice: 2000, amount: 8000}],
    roll:""
  },
  mutations: {
    
    MUTATE_UPLOAD_DATE1(state,value){
      state.loginHistory1.push(value);
    },
    UPDATE_ROLE(state,value){
      state.roll=value
    },
    MUTATE_ADD_USER(state,value){
      state.login.push(value);
    },
    MUTATE_UPDATE_STOCK(state,value){
      state.stock.push(value);
    },
    MUTATE_UPDATE_MEDICINE(state,value){
      state.medicineMaster.push(value)
    },
    MUTATE_UPDATE_BILLDETAIL(state,value){
      state.billDetails.push(value);
    },
    MUTATE_UPDATE_BILLMASTER(state,value){
      state.billMaster.push(value)
    }
  },
  actions: {},
  modules: {},
});
