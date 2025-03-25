<template>
    <div>
        <dash/>
        
       
            <v-card class="d-flex justify-center mt-16">    <div style="position: fixed;">
<systemAdminAndInventry v-if="role=='SystemAdmin'" ></systemAdminAndInventry>
<systemAdminAndInventry v-if="role=='Inventry'" ></systemAdminAndInventry>
<manager1 v-if="role=='Manager'" > </manager1>
<biller v-if="role=='Biller'" > </biller>
</div>
</v-card>
<!-- </v-img> -->

    </div>
</template>
<script>
import dash from "../components/dash.vue";
import systemAdminAndInventry from "../components/systemAdminAndInventry.vue";
import manager1 from "../components/manager1.vue";
import biller from "../components/biller.vue";
import EventServices from "@/Service/EventServices";
export default{
    name:"",
    data(){
        return{
           role:'' 
        }
    },
    components:{
        dash,systemAdminAndInventry,manager1,biller
    },
    created(){
        console.log("Dash4")
      EventServices.Dash()
      .then((response)=>{
        console.log(response)
        this.role=response.data.role
          })
      .catch((error)=>{
        console.log(error)
        console.log("hjhk")
      })
    }
}
</script>