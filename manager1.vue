<template>
    <div>
        <div class="d-flex justify-center font-weight-bold text-h4 black--text mt-10">
            Todays Sale:
        </div>
        <div class="d-flex justify-center font-weight-bold text-h4 pink--text mt-2">
            {{ todaySale }}
        </div>

        <div>
            <currentinventry class="mt-0 "></currentinventry>
        </div>
    </div>
</template>
<script>
import currentinventry from "@/components/currentinventry.vue"
import EventServices from "@/Service/EventServices";
export default {
    name: "",
    data() {
        return {
            sale: 0,
            todaySale: 0,
            billdetail: this.$store.state.billDetails,
        }
    },
    created() {
        console.log("Dash3")
        EventServices.ManagerDash()
            .then((response) => {
                console.log(response)
                this.todaySale = response.data.managertodaysale
                this.todaySale=parseFloat(this.todaySale.toFixed(2))
            })
            .catch((error) => {
                console.log(error)
                console.log("hjhk")
            })
    },
    components: {
        currentinventry
    }

}
</script>