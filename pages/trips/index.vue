<template>
  <v-container>
    <v-row>
      <v-col offset="10" cols="2">
        <v-btn color="primary" class="white--text" @click="directToNewTrip">
          New
        </v-btn>
      </v-col>
    </v-row>
    <v-tabs centered class="pb-4">
      <v-tab>My Trips</v-tab>
      <v-tab-item>
        <v-row>
          <h1 class="trip-header mx-4">Planned Trips</h1>
        </v-row>
        <v-divider class="thicc"></v-divider>
        <div v-if="tripsPlanned">
          <v-row v-for="trip in trips" :key="trip.id">
            <Trip :trip="trip"></Trip>
          </v-row>
        </div>
        <div v-else>
          <v-row class="pa-4">
            <h3>You have no trips planned.</h3>
          </v-row>
        </div>
        <v-row class="mt-4">
          <h1 class="trip-header mx-4">Past Trips</h1>
        </v-row>
        <v-divider class="thicc"></v-divider>
        <div v-if="pastTripsFinished">
          <v-row fluid v-for="trip in past_trips" :key="trip.id">
            <Trip :trip="trip"></Trip>
          </v-row>
        </div>
        <div v-else>
          <v-row class="pa-4">
            <h3>You have no past trips.</h3>
          </v-row>
        </div>
      </v-tab-item>
      <v-tab>Shared With Me</v-tab>
      <v-tab-item>
        <v-row>
          <h1 class="trip-header mx-4">Shared Trips</h1>
        </v-row>
        <v-divider class="thicc" />
        <div v-if="tripsShared">
          <v-row v-for="trip in shared_trips" :key="trip.id">
            <Trip :trip="trip"></Trip>
          </v-row>
        </div>
        <div v-else>
          <v-row class="pa-4">
            <h3>You have no shared trips.</h3>
          </v-row>
        </div>
      </v-tab-item>
    </v-tabs>
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
      shared_trips: [],
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
          that.trips = response.data;
          that.filterTrips();
        })
        .catch(function (error) {
          console.error(error);
        });
    },
    filterTrips() {
      this.shared_trips = this.trips.filter(trip => trip.viewers != null && trip.viewers.includes(this.$store.state.user.uid));
      this.past_trips = this.trips.filter(trip => trip.endDate < new Date().toISOString().slice(0, 10) && !this.shared_trips.includes(trip));
      this.trips = this.trips.filter(trip => trip.endDate >= new Date().toISOString().slice(0, 10) && !this.shared_trips.includes(trip));
    },
    changeTripPerspective() {
      this.sharedTrips = !this.sharedTrips
    },
    directToNewTrip() {
      this.$router.push('/new_trip')
    },
    directToTrip(id) {
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
    tripsShared() {
      return this.shared_trips.length > 0 ? true : false
    },
    pastSharedTripsFinished() {
      return this.past_shared_trips.length > 0 ? true : false
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
