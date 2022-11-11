<template>
    <v-card width="100%" class="mb-5 text-left">
        <v-img class="event-img" :src="event.image" height="300px"
            gradient="to top left, rgba(100,115,201,.33), rgba(25,32,72,.7)">
            <h2 class="pa-4 white--text">{{ event.name }}</h2>
            <h3 class="pa-4 white--text">{{ event.location }}</h3>
        </v-img>
    </v-card>
</template>

<script>
export default {
    props: ['event'],
    data() {
        return {
        }
    },
}
</script>

<style>
.event-img:hover {
    cursor: pointer;
}
</style>