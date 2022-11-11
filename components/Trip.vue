<template>
    <v-card width="100%" class="mb-5 text-left">
        <v-img @click="directToTrip(trip.id)" :src="trip.image" height="300px"
            gradient="to top left, rgba(100,115,201,.33), rgba(25,32,72,.7)">
            <h2 class="pa-4 white--text">{{ trip.location }}</h2>
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
            console.log("clicked trip", id);
            this.$router.push('/trips/' + id)
        },
    }
}
</script>