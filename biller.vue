<template>
    <v-main>
    <div class="">
        <!-- {{ sale1 }}{{ sale2 }} -->
        <p class="font-weight-medium text-h4 mt-7 black--text">
            Your Todays Sales:
        </p>
        <!-- <p class="font-weight-medium text-h4 mt-7  green--text d-flex justify-center" v-if="sale1==0">
            0<v-icon class=" green--text ml-2" large>mdi-arrow-down</v-icon>
        </p> -->
        <p class="font-weight-medium text-h4 mt-7  green--text d-flex justify-center" v-if="sale2 < sale1 ">
            {{ sale1 }}<v-icon class=" green--text ml-2" large>mdi-arrow-up</v-icon>
        </p>
        <p class="font-weight-medium text-h4 mt-7 red--text d-flex justify-center" v-if="sale2>sale1">
        
            {{ sale1 }}<v-icon class="red--text ml-2" large>mdi-arrow-down</v-icon>
            
        </p>

    </div>
</v-main>
</template>
<script>
import EventServices from "@/Service/EventServices";
export default {
    name: "",
    data() {
        return {
            sale1:0,
            sale2:0,
            // billdetail: this.$store.state.billDetails,
            // totalSale:0,
            // user:this.$store.state.loginHistory1,
        }
    },
    beforeCreate(){
    //    this.user=this.$store.state.loginHistory1
    },
    created() {
        console.log("billerSale")
      EventServices.BillerDash()
      .then((response)=>{
        console.log(response)
        this.sale1=response.data.billertodaysale
        this.sale1=parseFloat(this.sale1.toFixed(2))
        console.log(this.sale1)
        this.sale2=response.data.yesterdaysale
        this.sale2=parseFloat(this.sale2.toFixed(2))
        // this.totalSale=this.sale1-this.sale2
        // console.log(this.totalSale)
          })
      .catch((error)=>{
        console.log(error)
        console.log("hjhk")
      })
      
    }
}
</script>