<template>
    <v-card width="100%" style="border: 1px solid #424242;" class="mb-5 text-left">
        <v-img :src="this.image" height="300px" gradient="to top left, rgba(100,115,201,.33), rgba(25,32,72,.7)">
            <h2 class="pa-4 white--text">{{ this.location }}</h2>
        </v-img>
    </v-card>
</template>

<script>
export default {
    props: ['location', 'date', 'image'],
    data() {
        return {
        }
    }
}
</script>