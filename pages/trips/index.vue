<template>
  <v-container>
    <v-row>
      <v-col offset="10" cols="2">
        <v-btn color="primary" class="white--text" @click="directToNewTrip">
          New
        </v-btn>
      </v-col>
    </v-row>
    <v-row justify="center">
      <div>
        <h3
          class="change-cursor"
          @click="changeTripPerspective"
          :class="{ header: !sharedTrips }"
        >
          My Trips
        </h3>
      </div>
      <div class="pl-4">
        <h3
          class="change-cursor"
          @click="changeTripPerspective"
          :class="{ header: sharedTrips }"
        >
          Shared With Me
        </h3>
      </div>
    </v-row>
    <v-row>
      <h1 class="trip-header">Current Trips</h1>
    </v-row>
    <v-divider class="thicc"></v-divider>
    <v-row class="mt-4">
      <h1 class="trip-header">Past Trips</h1>
    </v-row>
    <v-divider class="thicc"></v-divider>
  </v-container>
</template>

<script>
export default {
  data() {
    return {
      sharedTrips: false,
    }
  },
  methods: {
    changeTripPerspective() {
      this.sharedTrips = !this.sharedTrips
    },
    directToNewTrip() {
      this.$router.push('/new_trip')
    },
  },
}
</script>

<style lang="scss">
.trip-header {
  margin-bottom: 0px;
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
