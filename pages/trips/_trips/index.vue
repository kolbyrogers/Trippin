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
    <h1>{{ trip.location }}</h1>
    <p>{{ trip.startDate }} - {{ trip.endDate }}</p>
    <div v-if="!eventsPlanned">
      <h2 class="pt-6">No Events Planned</h2>
    </div>
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
      events: [],
      trip: {},
    }
  },
  methods: {
    getEvents: async function () {
      const that = this;
      this.$axios.get('http://localhost:8080/api/events/trips/' + this.$route.params.trips, {
        headers: {
          'Content-Type': 'application/json',
        }
      })
        .then(function (response) {
          that.events = response.data;
        })
        .catch(function (error) {
          console.error(error);
        });
    },
    getTrip: async function () {
      const that = this;
      this.$axios.get('http://localhost:8080/api/trips/' + this.$route.params.trips, {
        headers: {
          'Content-Type': 'application/json',
        }
      })
        .then(function (response) {
          that.trip = response.data;
        })
        .catch(function (error) {
          console.error(error);
        });
    },
    directToNewEvent() {
      this.$router.push('/trips/' + this.$route.params.trips + '/events')
    },
    directToShare() {
      this.$router.push('/trips/' + this.$route.params.trips + '/share')
    },
    back() {
      this.$router.go(-1)
    },
  },
  computed: {
    eventsPlanned() {
      return this.events.length > 0 ? true : false
    },
  },
  async beforeMount() {
    await this.getTrip(this.$route.params.trips);
    await this.getEvents()
  },
}
</script>
