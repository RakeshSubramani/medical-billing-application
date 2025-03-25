<template>
  <div>
    <dash />
    <v-app-bar color="pink" dense dark>
      <v-icon @click="back()">mdi-arrow-left</v-icon>


      <v-toolbar-title class="ml-16">User details</v-toolbar-title>
      <v-spacer></v-spacer>

      <v-icon class="black--text" @click="userDetailShown()">
        mdi-chevron-down

      </v-icon>

    </v-app-bar>
    <template>
      <div class="text-center">
        <v-dialog v-model="dialog" width="500">

          <v-card>
            <v-card-title class="text-h5 blue lighten-2">
              Add User
            </v-card-title>

            <v-card-text class="text-h5  red--text mt-5">
              Empty Field </v-card-text>

            <v-divider></v-divider>

            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn color="primary" text @click="dialog = false">
                accept
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-dialog>
      </div>
    </template>
    <template>
      <div class="text-center">
        <v-dialog v-model="dialog1" width="500">

          <v-card>
            <v-card-title class="text-h5 blue lighten-2">
              Add User
            </v-card-title>

            <v-card-text class="text-h5  red--text mt-5">
              User Already exist </v-card-text>

            <v-divider></v-divider>

            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn color="primary" text @click="dialog1 = false">
                accept
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-dialog>
      </div>
    </template>

    <template>
      <div class="text-center">
        <v-dialog v-model="dialog2" width="500">

          <v-card>
            <v-card-title class="text-h5 blue lighten-2">
              Add User
            </v-card-title>

            <v-card-text class="text-h5  green--text mt-5">
              user add successfully </v-card-text>

            <v-divider></v-divider>

            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn color="primary" text @click="dialog2 = false">
                ok
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-dialog>
      </div>
    </template>

    <div class="d-flex justify-center mt-16">
      <v-card height="255" width="900" color="#3949AB" v-show="userDetail" class="">
      <v-card height="210" width="820" class="ml-10 mt-1" color="">
      <v-form class=" ml-16" >
        <v-container class="mt-5">
          <v-row class="mt-10">
            <v-col cols="12" md="3">
              <v-text-field v-model="AddUser.userid" :counter="10" label="User Id" required outlined
                hide-details></v-text-field>
            </v-col>

            <v-col cols="12" md="3">
              <v-text-field v-model="AddUser.password" :counter="10" label="password" type="password" hide-details
                required outlined></v-text-field>
            </v-col>

            <v-col cols="12" md="3"> <v-select v-model="AddUser.role" :items="items" label="Role" outlined></v-select>
            </v-col>
            <v-col cols="12" md="3">
              <div class="my-2">
                <v-btn color="warning" dark @click="add()">
                  Add
                </v-btn>
              </div>

            </v-col>
          </v-row>
        </v-container>
      </v-form>
    </v-card>
  </v-card>
    </div>
  </div>
</template>
<script>
import dash from '../components/dash.vue';
import EventServices from "@/Service/EventServices";
export default {
  name: "",
  data() {
    return {
      firstname: "",
      password: "",
      role: "",
      AddUser: {
        userid: '',
        password: '',
        role: '',
      },
      items: ['Biller', 'Manager', 'SystemAdmin', 'Inventry'],
      userDetail: false,
      desserts: [],
      dialog: false,
      dialog1:false,
      dialog2:false,

    }
  },
  methods: {
    userDetailShown() {
      this.userDetail = !this.userDetail;
    },

    add() {
      var flag1 = true
      if (this.AddUser.userid != "" && this.AddUser.password != "" && this.AddUser.role != "") {
        for (var i = 0; i < this.desserts.length; i++) {
          if (this.AddUser.userid == this.desserts[i].UserId) {
            this.dialog1 = true
            flag1 = false
          }
        }
        if (flag1 == true) {
          console.log("AddUser")
          EventServices.AddUser(this.AddUser)
            .then((response) => {
              console.log(response)
              var user=response.data.msg
              if(user==''){
                this.dialog2=true
              }
              else{
                alert("not defind")
              }

            })
            .catch((error) => {
              console.log(error)
              console.log("hjhk")
            })
        }

      }
      else {
        this.dialog = true
      }
       this.AddUser.userid="";
       this.AddUser.password="";
       this.AddUser.role="";
    },
    back() {
      this.$router.push("/dashboard")
    }
  },
  components: {
    dash
  },
  created() {
    console.log("")
    EventServices.GetUser()
      .then((response) => {
        console.log(response)
        this.desserts = response.data.user
        // for (var i = 0; i < desserts.length; i++) {
        //   this.item.push(desserts[i].medicinename)
        // }
        // this.desserts=response.data
      })
      .catch((error) => {
        console.log(error)
        console.log("hjhk1")
      })
  }
}
</script>