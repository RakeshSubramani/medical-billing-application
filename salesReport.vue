<template>
  <div>
    <dash />
    <v-icon @click="back()">mdi-arrow-left</v-icon>
    <div class="d-flex justify-center text-h5 blue--text font-weight-bold mt-2">
      Sales Report
    </div>
    <div class="d-flex justify-center mt-16">
      <v-row class="d-flex justify-center">

        <v-col cols="6" sm="4" md="4">
          <v-text-field v-model="stockentry.fromdate" label="From Date" type="date" :max="maxDate"></v-text-field>

        </v-col>
        <v-col cols="6" sm="4" md="4" color="primary">

          <v-text-field v-model="stockentry.todate" label="To Date" type="date" :max="maxDate"></v-text-field>
        </v-col>
        <v-btn @click="search()" class="mt-10 ml-10" color="primary">search</v-btn>
      </v-row>
    </div>

    <v-container class=" mt-16">
      <v-data-table :headers="headers" :items="desserts" class="elevation-19 font-weight-bold"></v-data-table>
    </v-container>
    <template>
      <div class="text-center">


        <v-snackbar v-model="snackbar" color="primary">
          give valid date

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


        <v-snackbar v-model="snackbar1" color="primary">
          between two dates there is no sales

          <template v-slot:action="{ attrs }">
            <v-btn color="" text v-bind="attrs" @click="snackbar1 = false">
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
  data() {
    return {
      headers: [
        {
          text: 'Billno',
          align: 'start',
          // sortable: false,
          value: 'billno',
        },
        { text: 'BillDate', value: "billdate" },
        { text: 'MedicineName', value: "medicinename" },
        { text: 'Qty', value: "qty" },
        { text: 'Amount', value: "amount" },

      ],
      desserts: [],
      // date: (new Date(Date.now() - (new Date()).getTimezoneOffset() * 60000)).toISOString().substr(0, 10),?
      // date1: (new Date(Date.now() - (new Date()).getTimezoneOffset() * 60000)).toISOString().substr(0, 10),
      menu: false,

      modal: false,
      menu2: false,
      menu1: false,
      billNo: '',
      date2: '',
      quantity: '',
      Amount: 0,
      userId: '',
      medicineName: '',
      stockentry: {
        fromdate: '',
        todate: '',
      },
      snackbar: false,
      snackbar1: false,

      // show:true,
      // v1:true
    }
  },
  components: {
    dash
  },
  methods: {
    back() {
      this.$router.push("/dashboard")
    },
    search() {
      if (this.stockentry.fromdate != "" && this.stockentry.todate != "") {
        if (this.stockentry.todate >= this.stockentry.fromdate) {
          console.log("SalesReport")
          EventServices.SalesReport(this.stockentry)
            .then((response) => {
              console.log(response)
              if (response.data.salesarr != null) {
                this.desserts = response.data.salesarr
                console.log("dfgdf")
                console.log(this.desserts)
              } else {
                this.snackbar1 = true
              }

            })
            .catch((error) => {
              console.log(error)
              console.log("hjhk")
            })
        }

      } else {
        this.snackbar = true
      }


    },
  },


  computed: {
    maxDate() {
      const today = new Date().toISOString().split('T')[0];
      return today;
    }

  },
}
</script>