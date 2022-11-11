<template>
  <v-container>
    <v-row>
      <v-col cols="2">
        <v-btn color="primary" class="white--text" @click="back"> Back </v-btn>
      </v-col>
      <v-col offset="6" cols="2">
        <v-btn color="primary" class="white--text" @click="directToShare">
          Share
        </v-btn>
      </v-col>
      <v-col cols="2">
        <v-btn color="primary" class="white--text" @click="directToNewEvent">
          New Event
        </v-btn>
      </v-col>
    </v-row>
    <h1>Trip Location</h1>
    <v-row justify="center" v-for="event in events" :key="event.id">
      <Event :event="event"></Event>
    </v-row>
  </v-container>
</template>

<script>
export default {
  async asyncData({ params }) {
    const id = params.id
    return { id }
  },
  data() {
    return {
      events: [
        {
          id: 12,
          name: 'Snorkeling',
          location: 'Coral Reef',
          image: '/images/snorkel.jpg',
        },
        {
          id: 21,
          name: 'Hiking',
          location: 'Mount Everest',
          image: '/images/snorkel.jpg',
        },
      ],
    }
  },
  methods: {
    directToNewEvent() {},
    directToShare() {
      this.$router.push('/trips/' + this.id + '/share')
    },
    back() {
      this.$router.go(-1)
    },
  },
}
</script>
