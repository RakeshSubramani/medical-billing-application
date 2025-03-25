<template>

  <div class="home " id="img">
    <v-app-bar app color="#01579B
" dark>
      <p class="">MEDAPP</p>
      <v-spacer></v-spacer>
    </v-app-bar>
    <template>
      <div class="text-center">
        <v-dialog v-model="dialog" width="500">

          <v-card>
            <v-card-title class="text-h5 blue lighten-2">
              Login Validation
            </v-card-title>

            <v-card-text class="text-h5  lighten-2">
              invalid user </v-card-text>

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


    <div class="d-flex justify-center mt-16 ">
      <v-hover close-delay="533">
        <v-card height="450" width="400" color="#4DD0E1


" class="d-flex justify-center">
<pre class=" text-h4 font-weight-medium blue--text mt-3" style="position: fixed;">LOGIN</pre>

          <v-card height="350" width="300"  class="mt-16" >
            <!-- <v-container> -->

              <v-form ref="form" style="position: fixed;" class="mt-16">
                <v-text-field v-model="validation.user_id" label="UserId" required class="ml-10 mt-1"
                  color="black"></v-text-field>

                <v-text-field v-model="validation.password" label="Password" type="password" required class="ml-10 mt-6"
                  color="black"></v-text-field>
                <!-- <v-container> -->
                  <div class="ml-10"><v-btn color="#3949AB
" class="ml-16 mt-10" @click="login()" v-show="loginbtn">
                    lOGIN
                  </v-btn></div>
                <!-- </v-container> -->
              </v-form>
            <!-- </v-container> -->
          </v-card>
        </v-card>
      </v-hover>
    </div>
  </div>
</template>

<script>
import EventServices from "@/Service/EventServices"
export default {
  name: "Home",
  data() {
    return {
      validation: {
        user_id: "",
        password: "",
      },
      job: "",
      loginbtn: false,
      loginuser: this.$store.state.login,
      loginhistory: {},
      dialog: false,
    }
  },
  watch: {
    'validation.password': function () {
      if ((this.validation.password != "") && (this.validation.user_id != "")) {
        this.loginbtn = true;
      }
    }

  },
  methods: {
    login() {
      console.log("LoginValidation")
      EventServices.LoginValidation(this.validation)
        .then((response) => {
          console.log("res", response)
          console.log(response.data.role)
          this.job = response.data.role
          this.$store.commit("UPDATE_ROLE", response.data.role)
          if (this.job == "Biller") {

            this.$router.push("/dashboard");
          }
          else if (this.job == "Manager") {

            this.$router.push("/dashboard");
          }
          else if (this.job == "SystemAdmin") {
            this.$router.push("/dashboard");

          }
          else if (this.job == "Inventry") {
            this.$router.push("/dashboard");

          }
          else {
            this.dialog = true
          }
        })
        .catch((error) => {
          console.log(error)
          console.log("hjhk")
        })
    },
  }
};
</script>