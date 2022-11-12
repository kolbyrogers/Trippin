<template>
  <v-container>
    <v-row>
      <v-col offset="10" cols="2">
        <v-btn color="primary" class="white--text" @click="directToNewTrip">
          New
        </v-btn>
      </v-col>
    </v-row>
    <v-tabs centered class="pb-4" v-model="activeView">
      <v-tabs-slider color="primary"></v-tabs-slider>
      <v-tab>My Trips</v-tab>
      <v-tab>Shared With Me</v-tab>
    </v-tabs>
    <v-row>
      <h1 v-if="activeView" class="trip-header">Shared Trips</h1>
      <h1 v-else class="trip-header">Current Trips</h1>
    </v-row>
    <v-divider class="thicc"></v-divider>
    <div v-if="tripsPlanned">
      <v-row v-for="trip in trips" :key="trip.id">
        <Trip :trip="trip"></Trip>
      </v-row>
    </div>
    <div v-else>
      <v-row class="pl-3 pt-4">
        <h3>You have no trips planned, plan a trip today!</h3>
      </v-row>
    </div>
    <v-row class="mt-4">
      <h1 v-if="activeView" class="trip-header">Previous Shared Trips</h1>
      <h1 v-else class="trip-header">Past Trips</h1>
    </v-row>
    <v-divider class="thicc"></v-divider>
    <div v-if="pastTripsFinished">
      <v-row fluid v-for="trip in past_trips" :key="trip.id">
        <Trip :trip="trip"></Trip>
      </v-row>
    </div>
    <div v-else>
      <v-row class="pl-3 pt-4">
        <h3>You have no past trips.</h3>
      </v-row>
    </div>
  </v-container>
</template>

<script>
export default {
  data() {
    return {
      activeView: 0,
      sharedTrips: false,
      trips: [],
      past_trips: [],
    }
  },
  methods: {
    getTrips: async function () {
      const that = this;
      this.$axios.get('http://localhost:8080/api/trips/users/' + this.$store.state.user.uid, {
        headers: {
          'Content-Type': 'application/json',
        }
      })
        .then(function (response) {
          console.log(response.data);
          that.trips = response.data;
        })
        .catch(function (error) {
          console.log(error);
        });
    },
    changeTripPerspective() {
      this.sharedTrips = !this.sharedTrips
    },
    directToNewTrip() {
      this.$router.push('/new_trip')
    },
    directToTrip(id) {
      console.log('clicked trip', id)
      this.$router.push('/trips/' + id)
    },
  },
  async mounted() {
    await this.getTrips()
  },
  computed: {
    tripsPlanned() {
      return this.trips.length > 0 ? true : false
    },
    pastTripsFinished() {
      return this.past_trips.length > 0 ? true : false
    },
  },
}
</script>

<style lang="scss">
.trip-header {
  padding-left: 8px;
  margin: 0, 0;
}

.header {
  font-weight: bold;
  text-decoration: underline;
  text-decoration-color: var(--primary);
}

.change-cursor {
  cursor: pointer;
}

.thicc {
  border-width: 2px;
  border-color: black;
  margin-bottom: 25px;
}
</style>
