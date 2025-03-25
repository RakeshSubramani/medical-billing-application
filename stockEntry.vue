<template>
  <div>
    <dash />
    <v-icon @click="back()">mdi-arrow-left</v-icon>
    <pre class="blue--text text-h5 font-weight-bold d-flex justify-center">stock Entry</pre>
    <div>

    </div>
    <div class="d-flex justify-center mt-16">
      <v-main>
        <div class="text-right">


          <v-btn color="primary" dark @click.stop="dialog = true">
            ADD
          </v-btn>
          <template>
      <div class="text-center">
        <v-dialog v-model="dialog" width="500">

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

          <v-dialog v-model="dialog" width="600">
            <v-app-bar color="blue" dense dark>
              <v-icon class="black--text" @click="dialog = false">
                mdi-minus


              </v-icon>

              <v-toolbar-title class="ml-5" @click="addStock()">Add stock</v-toolbar-title>
              <v-spacer></v-spacer>


            </v-app-bar>
            
            <div>
              <v-card class="">

                <v-card-text class="mt-">

                  <v-text-field v-model="addmedicine.medicinename" :counter="10" label="Medicine" hide-details required
                    outlined></v-text-field>

                  <v-text-field v-model="addmedicine.brand" :counter="10" label="Brand" hide-details required outlined
                    class="mt-4"></v-text-field>
                </v-card-text>

                <v-card-actions>
                  <v-spacer></v-spacer>
                  <v-btn color="green darken-1" text @click="add()" >
                    ADD
                  </v-btn>

                </v-card-actions>
              </v-card>
            </div>
          </v-dialog>

        </div>
      </v-main>
    </div>
    <div class="d-flex justify-center">
      <v-form class="ml-16 mt-16">
        <v-container>
          <v-row>
            <v-col cols="12" md="3"> <v-select v-model="Updatemedicine.medicinename" :items="item"
                item-text="medicineName" label="medicine Name" outlined></v-select>
            </v-col>
            <v-col cols="12" md="3">
              <v-text-field v-model="item1" label="Brand" disabled></v-text-field>
            </v-col>

            <v-col cols="12" md="3">
              <v-text-field type="Number" v-model.number="Updatemedicine.quantity" :counter="10" label="Quantity"
                hide-details required outlined></v-text-field>
            </v-col>

            <v-col cols="12" md="3">
              <v-text-field type="Number" v-model.number="Updatemedicine.unitprice" :counter="10" label="unitPrice"
                hide-details required outlined></v-text-field>
            </v-col>
          </v-row>
        </v-container>
      </v-form>
    </div>
    <div>
      <v-bottom-navigation background-color="indigo" shift color="primary" class="mt-16">
        <v-btn class="text-h6 font-weight-bold" @click="update()">
          update
        </v-btn>

      </v-bottom-navigation>
    </div>
    <template>
  <div class="text-center">
    

    <v-snackbar
      v-model="snackbar"

      color="primary"
    >
      the add field is empty

      <template v-slot:action="{ attrs }">
        <v-btn
          color=""
          text
          v-bind="attrs"
          @click="snackbar = false"
        >
          Close
        </v-btn>
      </template>
    </v-snackbar>
  </div>
</template>
<template>
  <div class="text-center">
    

    <v-snackbar
      v-model="snackbar1"
      color="primary"
    >
      user already exists

      <template v-slot:action="{ attrs }">
        <v-btn
          color=""
          text
          v-bind="attrs"
          @click="snackbar1 = false"
        >
          Close
        </v-btn>
      </template>
    </v-snackbar>
  </div>
</template>

<template>
  <div class="text-center">
    

    <v-snackbar
      v-model="snackbar2"
    >
      give correct required field

      <template v-slot:action="{ attrs }">
        <v-btn
          color="red"
          text
          v-bind="attrs"
          @click="snackbar2 = false"
        >
          Close
        </v-btn>
      </template>
    </v-snackbar>
  </div>
</template>
<template>
      <div class="text-center">
        <v-dialog v-model="dialog2" width="500">

          <v-card>
            <v-card-title class="text-h5 blue lighten-2">
              Add Stock
            </v-card-title>

            <v-card-text class="text-h5  green--text mt-5">
               add stock successfully </v-card-text>

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

    <template>
      <div class="text-center">
        <v-dialog v-model="dialog3" width="500">

          <v-card>
            <v-card-title class="text-h5 blue lighten-2">
              Update Stock
            </v-card-title>

            <v-card-text class="text-h5  green--text mt-5">
               Update stock successfully </v-card-text>

            <v-divider></v-divider>

            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn color="primary" text @click="dialog3 = false">
                ok
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-dialog>
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
      addmedicine: {
        medicinename: '',
        brand: '',
      },
      Updatemedicine: {
        medicinename: '',
        quantity: 0,
        unitprice: 0,
      },
      item: [],
      item1: "",
      userDetail: false,
      dialog: false,
     
      show: true,
      desserts: [],
      snackbar:false,
      snackbar1:false,
      snackbar2:false,
      dialog2:false,
      dialog3:false,
      
    }
  },
  components: {
    dash
  },
  methods: {
    userDetailShown() {
      this.userDetail = !this.userDetail;
    },
    add() {
      if (this.addmedicine.medicinename != "" && this.addmedicine.brand != "") {
        var flag1 = true;
        // var flag2=true;
        for (var i = 0; i < this.desserts.length; i++) {
          if (this.desserts[i].medicinename != this.addmedicine.medicinename) {
            console.log(this.desserts[i].medicinename)
            flag1 = false
          }
          else {
            flag1 = true
          }
        }
        if (flag1 == false) {
          console.log("Addmedicine")
          EventServices.AddMedicine(this.addmedicine)
            .then((response) => {
              console.log("dfsdf",response.data)
              if(response.data.status==""){
                this.dialog2=true
              }

            })
            .catch((error) => {
              console.log(error)
              console.log("hjhk")
            })
        }else{
         this.snackbar1=true
        }
        this.addmedicine.medicinename="";
        this.addmedicine.brand="";


      }
      else {
        this.snackbar=true
      }

      console.log("Getmedicine")
      EventServices.GetMedicine()
        .then((response) => {
          console.log(response)
          this.desserts = response.data.medicine
          for (var i = 0; i < this.desserts.length; i++) {
            this.item.push(this.desserts[i].medicinename)
          }
          // this.desserts=response.data
        })
        .catch((error) => {
          console.log(error)
          console.log("hjhk1")
        })



        this.dialog=false
    },
    addStock() {
      this.show = true
    },
    back() {
      this.$router.push("/dashboard")
    },
    update() {
       if(this.Updatemedicine.medicinename!=""&& this.Updatemedicine.quantity>0 && this.Updatemedicine.unitprice>0 && this.Updatemedicine.quantity % 1==0 && this.Updatemedicine.unitprice % 1==0){
        console.log("Updatemedicine")
      EventServices.UpdateMedicine(this.Updatemedicine)
        .then((response) => {
          console.log(response)
          if(response.data.status=="S"){
              this.dialog3=true
          }
          // this.desserts = response.data
        })
        .catch((error) => {
          console.log(error)
          console.log("hjhk")
        })
       }
     else{
      this.snackbar2=true
     }
     this.Updatemedicine.medicinename="";
     this.Updatemedicine.quantity=0;
     this.Updatemedicine.unitprice=0;
     this.item1='';
   
    },
  },
  created() {
    console.log("Getmedicine")
    EventServices.GetMedicine()
      .then((response) => {
        console.log(response)
        this.desserts = response.data.medicine
        for (var i = 0; i < this.desserts.length; i++) {
          this.item.push(this.desserts[i].medicinename)
        }
        // this.desserts=response.data
      })
      .catch((error) => {
        console.log(error)
        console.log("hjhk1")
      })
  },
  watch: {
    'Updatemedicine.medicinename': function () {
      EventServices.GetMedicine()
        .then((response) => {
          console.log(response);
          this.desserts = response.data.medicine;
          for (var j = 0; j < this.desserts.length; j++) {
            if (this.Updatemedicine.medicinename == this.desserts[j].medicinename) {
              this.item1 =this.desserts[j].brand;
              console.log(j);
              break;
            }
          }
        })
        .catch((error) => {
          console.log(error);
          console.log("Error fetching medicine data"); // Improve error handling
        });
    }
  }


}
</script>