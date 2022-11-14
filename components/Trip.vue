<template>
    <v-card width="100%" class="mb-5 text-left">
        <v-img class="trip-img" @click="directToTrip(trip.id)" :src="trip.imageURL" height="300px"
            gradient="to top left, rgba(100,115,201,.33), rgba(25,32,72,.7)">
            <h2 class="pt-4 pl-4 ml-4 white--text">{{ trip.location }}</h2>
            <h4 class="pl-4 ml-4 white--text">{{ trip.startDate.substring(0, 4) }}</h4>
        </v-img>
    </v-card>
</template>

<script>
export default {
    props: ['trip'],
    data() {
        return {
        }
    },
    methods: {
        directToTrip(id) {
            this.$router.push('/trips/' + id)
        },
    }
}
</script>

<style>
.trip-img:hover {
    cursor: pointer;
}
</style>