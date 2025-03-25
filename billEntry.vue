<template>
  <div>
    <dash />
    <div>
      <v-app-bar color="pink" dense dark>
        <v-icon @click="back()">mdi-arrow-left</v-icon>
        <v-toolbar-title class="ml-16">Bill Entry</v-toolbar-title>
        <v-spacer></v-spacer>

        <v-icon class="black--text" @click="userDetailShown()">
          mdi-chevron-down

        </v-icon>

      </v-app-bar>
      <div class="d-flex justify-center">
        <v-form class="ml-16 mt-16" v-show="userDetail">
          <v-container>
            <v-row>
              <v-col cols="12" md="4"> <v-select v-model="medicine" :items="getstock" item-text="medicineName"
                  label="medicine Name" outlined></v-select>
              </v-col>
              <v-col cols="12" md="4">
                <v-text-field type="number" v-model.number="Qty" :counter="10" label="Quantity" hide-details required
                  outlined></v-text-field>
              </v-col>


              <v-col cols="12" md="4">
                <div class="my-2">
                  <v-btn color="warning" dark @click="add()" v-show="show">
                    Add
                  </v-btn>
                </div>

              </v-col>
            </v-row>
          </v-container>
        </v-form>
      </div>
    </div>
    <div class="d-flex justify-center mt-16">
      <v-main>
        <v-row align="center" justify="space-around">
          <v-btn color="primary" dark @click.stop="dialog = true">
            preiview
          </v-btn>
          <v-btn depressed color="error" @click="save()">
            save
          </v-btn>
          <pre>BillNo:{{ this.billno }}</pre>
          <pre>Date:{{ this.date }}</pre>
          <pre>Total:{{ this.totalCost1 }}</pre>
          <pre>GST:{{ this.gst1}}</pre>
          <pre>NetPayable:{{ this.netprice1}}</pre>

          <v-dialog v-model="dialog" width="500">
            <v-app-bar color="blue" dense dark>
              <v-icon class="black--text" @click="dialog = false">
                mdi-minus


              </v-icon>

              <v-toolbar-title class="ml-5">Preview</v-toolbar-title>
              <v-spacer></v-spacer>


            </v-app-bar>
            <v-card>
              <v-card-title class="text-h5">
                MEDICAL SHOP NAME
              </v-card-title>

              <v-card-text>
                <v-data-table :headers="medHead" :items="desserts" hide-default-footer></v-data-table>
                <pre></pre>
                <!-- <pre>Qty:{{ this.Qty }}</pre> -->
                <!-- <pre>Amount: {{ this.price }}</pre> -->
                <v-row align="center" justify="space-around">
                  <pre class="font-weight-bold black--text">total:{{ this.totalCost1 }}</pre>
                  <pre class="font-weight-bold black--text">gst:{{ this.gst1 }}</pre>
                  <pre class="font-weight-bold black--text">netPrice:{{ this.netprice1 }}</pre>
                </v-row>
              </v-card-text>

              <v-card-actions>
                <v-spacer></v-spacer>

                <!-- <v-btn
            color="green darken-1"
            text
            @click="dialog = false"
          >
            Disagree
          </v-btn> -->

                <v-btn color="green darken-1" text @click="dialog = false">
                  print
                </v-btn>
              </v-card-actions>
            </v-card>
          </v-dialog>
        </v-row>
      </v-main>
    </div>
    <div class=" mt-16 text-right">
      <v-main>
        <v-btn depressed color="primary" @click="download()">
          Download
        </v-btn>

      </v-main>
    </div>
    <div>
      <v-card>
        <v-card-title color="primary" height="100"></v-card-title>
        <v-data-table :headers="headers" :items="desserts"  class="elevation-15">
          <template v-slot:item.actions="{ item }">
            <v-icon small @click="deleteItem(item)">
              mdi-delete
            </v-icon>
          </template>
        </v-data-table>
      </v-card>
    </div>
    <template>
      <div class="text-center">


        <v-snackbar v-model="snackbar" color="primary">
          empty field

          <template v-slot:action="{ attrs }">
            <v-btn color="" text v-bind="attrs" @click="snackbar = false">
              Close
            </v-btn>
          </template>
        </v-snackbar>
      </div>
    </template>
    <template>
      <div class="text-center">


        <v-snackbar v-model="snackbar3" color="primary">
          empty field

          <template v-slot:action="{ attrs }">
            <v-btn color="" text v-bind="attrs" @click="snackbar3 = false">
              Close
            </v-btn>
          </template>
        </v-snackbar>
      </div>
    </template>

    <template>
      <div class="text-center">


        <v-snackbar v-model="snackbar1" color="primary">
          invalid quantity

          <template v-slot:action="{ attrs }">
            <v-btn color="" text v-bind="attrs" @click="snackbar1 = false">
              Close
            </v-btn>
          </template>
        </v-snackbar>
      </div>
    </template>
    <template>
      <div class="text-center">


        <v-snackbar v-model="snackbar2" color="primary">
          Quantity is less
          <template v-slot:action="{ attrs }">
            <v-btn color="" text v-bind="attrs" @click="snackbar2 = false">
              Close
            </v-btn>
          </template>
        </v-snackbar>
      </div>
    </template>
  </div>
</template>
<script>
import dash from '../components/dash.vue';
import EventServices from "@/Service/EventServices";
export default {
  name: "",
  data() {
    return {
      medicine1: [],
      medicine: '',
      Qty: 0,
      getstock: [],
      item1: [],
      userDetail: false,
      dialog: false,
      totalCost: 0,
      price: 0,
      billno: 0,
      netprice: 0,
      date: '',
      billmaster: {
        billmaster: [],
      },
      billdetail: {
        billdetail: [],
      },
      gst: 0,
      gst1: 0,
      totalCost1: 0,
      netprice1: 0,
      user: '',
      userId: '',
      med: [],
      med1: [],
      brandDum: [],
      qty: [],
      medHead: [{ text: 'medicineName', value: "medicineName" },
      { text: 'quantity', value: "quantity" }, { text: 'price', value: "price" }
      ],
      show: true,
      headers: [

        { text: 'billNo', value: "billNo" },
        { text: 'medicineName', value: "medicineName" },
        { text: 'quantity', value: "quantity" },
        { text: 'amount', value: "amount" },
        {text:'delete',value:'actions'}

      ],
      desserts: [],
      snackbar: false,
      snackbar1: false,
      snackbar2: false,
      snackbar3:false,
    }
  },
  components: {
    dash
  },
  watch: {
    medicine() {
      if (this.medicine) {
        this.show = true
      }
    }
  },
  mounted() {
    console.log("stockview")
    EventServices.StockView()
      .then((response) => {
        console.log(response)
        this.item1 = response.data.stockviewArr
        this.med1 = this.item1
        console.log(this.item1)
      })
      .catch((error) => {
        console.log(error)
        console.log("hjhk")
      })

  },
  methods: {

    deleteItem (item) {
        let value = this.desserts.indexOf(item)
        if (value>-1){
          alert(item.quantity)
          this.desserts.splice(value,1)
          this.totalCost1 -=item.amount
          this.gst1 =this.totalCost1*0.18
          this.netprice1=this.gst1+this.totalCost1
          for(var o=0;o<this.med1.length;o++){
            if(this.med1[o].medicinename==item.medicineName){
              this.med1[o].quantity+=item.quantity
            }
          }
        }
      },
    userDetailShown() {
      this.userDetail = !this.userDetail;
    },
    add() {
      var flag1 = false
      for (var i = 0; i < this.med1.length; i++) {
        console.log("first")
        if ((this.medicine == this.med1[i].medicinename) && (this.Qty != 0 && this.Qty > 0)) {
          flag1 = true
          if (this.Qty % 1 == 0) {
            this.price = this.med1[i].unitprice;
            if (this.med1[i].quantity >= this.Qty) {
              this.med1[i].quantity -= this.Qty;
              this.totalCost = this.price * this.Qty;
              this.totalCost1 += this.totalCost
              this.gst = this.totalCost * (18 / 100);
              this.gst = parseFloat(this.gst.toFixed(2))
              this. gst1 += this.gst;
              this.gst1=parseFloat(this.gst1.toFixed(2))
              var net = this.totalCost + this.gst;
              this.netprice = parseFloat(net.toFixed(2))
              this.netprice1 += this.netprice
              this.netprice1=parseFloat(this.netprice1.toFixed(2))
              console.log("fourth")
              if (this.desserts.length == 0) {
                let MedArr = { billNo: this.billno, medicineName: this.medicine, quantity: this.Qty, amount: this.totalCost ,price:this.price, gst: this.gst}
                this.desserts.push(MedArr)
                console.log(this.desserts)
              } else {
                var index = 1;
                var flag = false;

                for (var l = 0; l < this.desserts.length; l++) {
                  if (this.desserts[l].medicineName == this.medicine) {
                    index = l;
                    flag = true;
                  }
                }
                if (flag == true) {
                  this.desserts[index].quantity += this.Qty;
                  this.desserts[index].amount += this.totalCost;
                } else {
                  let MedArr = { billNo: this.billno, medicineName: this.medicine, quantity: this.Qty, amount: this.totalCost ,price:this.price, gst: this.gst}
                  this.desserts.push(MedArr)

                }

              }

            }
            else {
              this.snackbar2 = true
              break;
            }

          }else{
            this.snackbar1=true
          }
        }
      }
      if (flag1 == false) {
        this.snackbar3 = true
      }
      
     
    },
    save() {
      if ((this.medicine != "") && (this.Qty != 0)) {
          for (let j = 0; j < this.desserts.length; j++) {
            console.log('loop');
              this.billdetail.billdetail.push({ "billno": this.billno, "medicinename": this.desserts[j].medicineName, "quantity": this.desserts[j].quantity, "unitprice": this.desserts[j].price, "amount": this.desserts[j].amount })
              console.log("billdetails", this.billdetail)
             
            }
            EventServices.BillDetails(this.billdetail)
                .then((response) => {
                  console.log(response)
                  // this.desserts = response.data
                })
                .catch((error) => {
                  console.log(error)
                  console.log("hjhk")
                })
        
        this.billmaster.billmaster.push({ "billno": this.billno, "billdate": this.date, "billamount": this.totalCost1, "billgst":this.gst1, "netprice": this.netprice1 })
        EventServices.BillMaster(this.billmaster)
                .then((response) => {
                  console.log(response)
                  // this.desserts = response.data
                })
                .catch((error) => {
                  console.log(error)
                  console.log("hjhk")
                })
        this.billno = this.billno + 1;
        this.totalCost1 = 0;
        this.gst1 = 0;
        this.netprice1 = 0;
        this.desserts = [];
        this.medicine = "";
        this.Qty = 0;
        this.med = [];
      }
      
      else {
        this.snackbar = true
      }

    },
    back() {
      this.$router.push("/dashboard")
    }
  },
  created() {
    this.date = new Date().toISOString().split('T')[0]
    console.log("billno")
    EventServices.BillNo()
      .then((response) => {
        console.log(response)
        this.billno = (response.data.billno) + 1
        console.log(this.billno)
      })
      .catch((error) => {
        console.log(error)
        console.log("hjhk")
      })
    console.log("Getmedicine")
    EventServices.GetMedicine()
      .then((response) => {
        console.log(response)
        var desserts = response.data.medicine
        for (var i = 0; i < desserts.length; i++) {
          this.getstock.push(desserts[i].medicinename)
        }
        // this.desserts=response.data
      })
      .catch((error) => {
        console.log(error)
        console.log("hjhk1")
      })
  },
}
</script>