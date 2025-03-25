<template>
  <div>
    <dash/>
    <v-icon @click="back()">mdi-arrow-left</v-icon>
   <pre class="blue--text text-h5 font-weight-bold d-flex justify-center mt-7">login History</pre>
    <v-container class="mt-10">  
    <v-data-table
    :headers="headers"
    :items="desserts"
    class="elevation-19 font-weight-bold"
  ></v-data-table>
</v-container> 
</div>  
</template>
<script>
import dash from '../components/dash.vue';
import EventServices from "@/Service/EventServices";

  export default {
    data () {
      return {
        headers:    [
          {
            text: 'User_id',
            align: 'start',
            sortable: false,
            value: 'user_id',
          },
          { text: 'logindate',value:"logindate"},
          { text: 'logintime',value:"logintime"},
          { text: 'logoutdate',value:"logoutdate"},
          { text: 'logouttime',value:"logouttime"},

          
        ],
        desserts:[]

      }
    },
    methods:{
      back(){
      this.$router.push("/dashboard")
    }
    },
    components:{
      dash
    },
    created(){
      console.log("LoginHistory")
      EventServices.LoginHistory()
      .then((response)=>{
        console.log(response)
        this.desserts=response.data.loginhistory_arr
          })
      .catch((error)=>{
        console.log(error)
        console.log("hjhk")
      })
      
    },
   

    }
  
</script>