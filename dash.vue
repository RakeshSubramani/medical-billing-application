<template>
  <div>
    <v-app-bar app color="#01579B" dark class="">
      <p class="mt-6"><span class="font-weight-medium text-h4 pink--text">MED</span> <span
          class="orange--text text-h5">APP</span></p>
      <!-- <v-spacer></v-spacer> -->

      <v-btn text @click="stockViewEntry()" v-if="job == 'Manager'" class="ml-16">Stock View</v-btn>
      <v-btn text @click="stockEntryEntry()" v-if="job == 'Manager'" class="ml-16">Stock Entry</v-btn>
      <v-btn text @click="salesReportEntry()" v-if="job == 'Manager'" class="ml-16">Sales Report</v-btn>


      <v-btn text @click="stockViewEntry()" v-if="job == 'Biller'" class="ml-16">Stock View</v-btn>
      <v-btn text @click="billEntry()" v-if="job == 'Biller'" class="ml-16">Bill entry</v-btn>




      <v-btn text @click="addUserEntry()" v-if="job == 'SystemAdmin'" class="ml-16">Add user</v-btn>
      <v-btn text @click="loginHistoryEntry()" v-if="job == 'SystemAdmin'" class="ml-16">Login History</v-btn>


      <v-btn text @click="stockViewEntry()" v-if="job == 'Inventry'" class="ml-16">Stock View</v-btn>
      <v-btn text @click="stockEntryEntry()" v-if="job == 'Inventry'" class="ml-16">Stock Entry</v-btn>


      <v-spacer></v-spacer>


      <!-- <v-btn @click="logOut()" text >logOut</v-btn> -->
      <v-dialog v-model="dialog" persistent max-width="290">
        <template v-slot:activator="{ on, attrs }">
          <v-btn color="primary" dark v-bind="attrs" v-on="on" >
            logOut
          </v-btn>
        </template>
        <v-card>
          <v-card-title class="text-h5">
            Logout
          </v-card-title>
          <v-card-text>Do you want to logout </v-card-text>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn color="green darken-1" text @click="dialog = false">
              No
            </v-btn>
            <v-btn color="green darken-1" text @click="logOut()">
              Yes
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>

    </v-app-bar>
    <template>


    </template>
  </div>
</template>
<script>
import EventServices from "@/Service/EventServices";
export default {
  name: "dash",
  data() {
    return {
      job: "",
      userId: '',
      login: '',
      loginHistory: {},
      user: "",
      dialog: false,
    }
  },
  methods: {
    stockEntryEntry() {
      this.$router.push("/stockEntry")
    },
    stockViewEntry() {
      this.$router.push("/stockView")
    },
    loginHistoryEntry() {
      this.$router.push("/loginHistory")
    },
    addUserEntry() {
      this.$router.push("/addUser")
    },
    billEntry() {
      this.$router.push("/billEntry")
    },
    salesReportEntry() {
      this.$router.push("/salesReport")
    },
    logOut() {
      console.log("logout")
      EventServices.LogOut()
        .then((response) => {
          console.log(response)
          var desserts = response.data
          if (desserts.status == '' || desserts.status=='S') {

            this.$router.push("/")


          }
        })
        .catch((error) => {
          console.log(error)
          console.log("hjhk")
        })
    }
  },
  created() {
    console.log("Dash3")
    EventServices.Dash()
      .then((response) => {
        console.log(response)
        this.job = response.data.role
      })
      .catch((error) => {
        console.log(error)
        console.log("hjhk")
      })
  },
  mounted() {

  }

};
</script>