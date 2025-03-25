<template>
    
<div>
    <div class="d-flex justify-center font-weight-medium text-h4 black--text mt-16">
            current inventry value
            
        </div>
        <div class="d-flex justify-center font-weight-medium text-h4 mt-7 pink--text">
            {{ sale }}
        </div>
</div>
</template>
<script>
import EventServices from "@/Service/EventServices";

export default{
    name:"",
    data(){
        return{
               sale:0,
        }
    },
mounted(){
    console.log("Dash3")
      EventServices.ManagerDash()
      .then((response)=>{
        console.log(response)
        this.sale=response.data.inventaryvalue
          })
      .catch((error)=>{
        console.log(error)
        console.log("hjhk")
      })
    }
}
</script>