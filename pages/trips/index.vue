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
      <v-tabs-slider color="primary"></v-tabs-slider>
      <v-tab>My Trips</v-tab>
      <v-tab>Shared With Me</v-tab>
    </v-tabs>
    <v-row>
      <h1 class="trip-header">Current Trips</h1>
    </v-row>
    <v-divider class="thicc"></v-divider>
    <div v-if="tripsPlanned">
      <v-row v-for="trip in trips" :key="trip.id">
        <Trip
          @click="directToTrip(trip.id)"
          :image="trip.image"
          :location="trip.location"
          :date="trip.start_date"
        ></Trip>
      </v-row>
    </div>
    <div v-else>
      <v-row class="pl-3 pt-4">
        <h3>You have no trips planned, plan a trip today!</h3>
      </v-row>
    </div>

    <v-row class="mt-4">
      <h1 class="trip-header">Past Trips</h1>
    </v-row>
    <v-divider class="thicc"></v-divider>
    <div v-if="pastTripsFinished">
      <v-row v-for="trip in past_trips" :key="trip.id">
        <Trip
          :image="trip.image"
          :location="trip.location"
          :date="trip.start_date"
        ></Trip>
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
      sharedTrips: false,
      trips: [],
      past_trips: [],
    }
  },
  methods: {
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
  margin-bottom: 0px;
  padding-left: 8px;
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
}
</style>
