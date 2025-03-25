<template>
  <div>
    <dash></dash>
    <v-icon @click="back()">mdi-arrow-left</v-icon>
    <pre class="blue--text text-h4 font-weight-bold d-flex justify-center mt-16 ">Stock view</pre>
    <v-container class="mt-16">
    <v-data-table 
    :headers="headers"
    :items="desserts"
    class="elevation-19 font-weight-bold" 
    color="blue"
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
            text: 'Medicinename',
            align: 'start',
            sortable: false,
            value: 'medicinename',
          },
          { text: 'Brand',value:"brand"},
          { text: 'Quantity',value:"quantity"},
          { text: 'Unitprice',value:"unitprice"},
          
        ],
        desserts:[]
      }
    },
    components:{
      dash
    },
    methods:{
      back(){
      this.$router.push("/dashboard")
    }
    },
    created(){
      console.log("stockview")
      EventServices.StockView()
      .then((response)=>{
        console.log(response)
        this.desserts=response.data.stockviewArr
          })
      .catch((error)=>{
        console.log(error)
        console.log("hjhk")
      })
      
    },
  }
</script>