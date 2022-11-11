<template>
    <v-card @click="directToEvent(event.id)" width="50%" class="my-6" style="border: 1px solid #424242;">
        <div class="d-flex flex-no-wrap justify-space-between text-left">
            <div>
                <v-card-title>
                    <h2 class="pb-6">{{ event.name }}</h2>
                </v-card-title>
                <v-card-subtitle>
                    <h2 class="font-italic">{{ event.location }}</h2>
                </v-card-subtitle>
            </div>
            <v-avatar justify="right" size="200" rounded="0">
                <v-img :src="event.image"></v-img>
            </v-avatar>
        </div>
    </v-card>
</template>

<script>
export default {
    props: ['event'],
    data() {
        return {
        }
    },
    methods: {
        directToEvent(id) {
            console.log("clicked event", id);
            this.$router.push('/trips/' + this.$route.params.trips + '/events/' + id)
        },
    }
}
</script>

<style>
.event-img:hover {
    cursor: pointer;
}
</style>