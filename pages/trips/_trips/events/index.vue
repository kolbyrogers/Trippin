<template>
  <v-container>
    <v-row>
      <v-col cols="2">
        <v-btn class="white--text" color="primary" @click="goBack">Back</v-btn>
      </v-col>
      <v-col cols="10"></v-col>
    </v-row>
    <v-row justify="center">
      <h1 class="trip-heading">New Event</h1>
    </v-row>
    <v-row>
      <v-col cols="6">
        <img src="/images/new_event.svg" width="75%" height="75%" />
      </v-col>
      <v-col cols="6">
        <v-row :class="this.$vuetify.breakpoint.sm ? 'pr-4' : ''">
          <v-text-field v-model="name" label="Name" solo></v-text-field>
        </v-row>
        <v-row :class="this.$vuetify.breakpoint.sm ? 'pr-4' : ''">
          <v-text-field v-model="location" label="Location" solo></v-text-field>
        </v-row>
        <v-row>
          <!-- Import Google Map Component -->
        </v-row>
        <v-row :class="this.$vuetify.breakpoint.sm ? 'pr-4' : ''">
          <v-dialog
            ref="dialog1"
            v-model="modal_start"
            :return-value.sync="date_start"
            persistent
            width="290px"
          >
            <template v-slot:activator="{ on, attrs }">
              <v-text-field
                solo
                dense
                v-model="date_start"
                label="Start Date"
                readonly
                v-bind="attrs"
                v-on="on"
              >
              </v-text-field>
            </template>
            <v-date-picker v-model="date_start" scrollable>
              <v-spacer></v-spacer>
              <v-btn text color="primary" @click="modal_start = false">
                Cancel
              </v-btn>
              <v-btn
                text
                color="primary"
                @click="$refs.dialog1.save(date_start)"
              >
                OK
              </v-btn>
            </v-date-picker>
          </v-dialog>
        </v-row>
        <v-row :class="this.$vuetify.breakpoint.sm ? 'pr-4' : ''">
          <v-dialog
            ref="dialog"
            v-model="modal"
            :return-value.sync="time"
            persistent
            width="290px"
          >
            <template v-slot:activator="{ on, attrs }">
              <v-text-field
                v-model="time"
                label="Time"
                solo
                readonly
                v-bind="attrs"
                v-on="on"
              ></v-text-field>
            </template>
            <v-time-picker v-if="modal" v-model="time" full-width>
              <v-spacer></v-spacer>
              <v-btn text color="primary" @click="modal = false">
                Cancel
              </v-btn>
              <v-btn text color="primary" @click="$refs.dialog.save(time)">
                OK
              </v-btn>
            </v-time-picker>
          </v-dialog>
        </v-row>
      </v-col>
    </v-row>
    <v-row justify="center">
      <v-btn color="primary" class="white--text" @click="createEvent"
        >Create</v-btn
      >
    </v-row>
  </v-container>
</template>

<script>
export default {
  data() {
    return {
      location: '',
      name: '',
      modal_start: false,
      date_start: new Date(Date.now() - new Date().getTimezoneOffset() * 60000)
        .toISOString()
        .substr(0, 10),
      time: null,
      modal: false,
    }
  },
  methods: {
    goBack() {
      this.$router.go(-1)
    },
    createEvent() {
      const data = {
        name: this.name,
        location: this.location,
        date: this.date_start,
        time: this.time,
        tripId: this.$route.params.trips,
      }
      console.log('Sending trip:', data)
      this.$axios
        .post('http://localhost:8080/api/events', data, {
          headers: {
            'Content-Type': 'application/json',
          },
        })
        .then(function (response) {
          console.log(response)
        })
        .catch(function (error) {
          console.log(error)
        })
      this.name = ''
      this.location = ''
      this.time = ''
      this.$router.go(-1)
    },
    print() {
      console.log('clicked')
    },
  },
}
</script>

<style>
.trip-heading {
  font-size: 2.5rem;
  font-weight: 700;
  color: #3c3c3c;
}
</style>
