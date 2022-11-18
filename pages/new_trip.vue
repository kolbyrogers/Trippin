<template>
  <v-container>
    <v-row>
      <v-col cols="2">
        <v-btn class="white--text" color="primary" @click="goBack">Back</v-btn>
      </v-col>
      <v-col cols="10"></v-col>
    </v-row>
    <v-row justify="center">
      <h1 class="trip-heading">New Trip</h1>
    </v-row>
    <v-row>
      <v-col cols="6">
        <v-row :class="this.$vuetify.breakpoint.sm ? 'pl-3' : ''">
          <v-text-field v-model="location" label="Location" solo></v-text-field>
        </v-row>
        <v-row>
          <!-- TODO: Import Google Map Component -->
        </v-row>
        <v-row :class="this.$vuetify.breakpoint.sm ? 'pl-3' : ''">
          <v-dialog ref="dialog1" v-model="modal_start" :return-value.sync="date_start" persistent width="290px">
            <template v-slot:activator="{ on, attrs }">
              <v-text-field outlined dense v-model="date_start" label="Start Date" readonly v-bind="attrs" v-on="on">
              </v-text-field>
            </template>
            <v-date-picker v-model="date_start" scrollable :min="current_day">
              <v-spacer></v-spacer>
              <v-btn text color="primary" @click="modal_start = false">
                Cancel
              </v-btn>
              <v-btn text color="primary" @click="$refs.dialog1.save(date_start)">
                OK
              </v-btn>
            </v-date-picker>
          </v-dialog>
        </v-row>
        <v-row :class="this.$vuetify.breakpoint.sm ? 'pl-3' : ''">
          <v-dialog ref="dialog2" v-model="modal_end" :return-value.sync="date_end" persistent width="290px">
            <template v-slot:activator="{ on, attrs }">
              <v-text-field outlined dense v-model="date_end" label="End Date" readonly v-bind="attrs" v-on="on">
              </v-text-field>
            </template>
            <v-date-picker v-model="date_end" scrollable :min="date_start">
              <v-spacer></v-spacer>
              <v-btn text color="primary" @click="modal_end = false">
                Cancel
              </v-btn>
              <v-btn text color="primary" @click="$refs.dialog2.save(date_end)">
                OK
              </v-btn>
            </v-date-picker>
          </v-dialog>
        </v-row>
      </v-col>
      <v-col cols="6">
        <img src="/images/new_trip.svg" width="75%" height="75%" />
      </v-col>
    </v-row>
    <v-row justify="center">
      <v-btn color="primary" class="white--text" @click="createTrip" :disabled="
        !location ||
        !(
          (date_start == current_day && date_end == current_tomorrow) ||
          date_end != current_tomorrow
        )
      ">Create</v-btn>
    </v-row>
  </v-container>
</template>

<script>
export default {
  data() {
    return {
      location: '',
      current_day: new Date(Date.now() - new Date().getTimezoneOffset() * 60000)
        .toISOString()
        .substr(0, 10),
      modal_start: false,
      modal_end: false,
      date_start: new Date(Date.now() - new Date().getTimezoneOffset() * 60000)
        .toISOString()
        .substr(0, 10),
      date_end: new Date(
        Date.now() - new Date().getTimezoneOffset() * 60000 + 86400000
      )
        .toISOString()
        .substr(0, 10),
      current_tomorrow: new Date(
        Date.now() - new Date().getTimezoneOffset() * 60000 + 86400000
      )
        .toISOString()
        .substr(0, 10),
    }
  },
  methods: {
    goBack() {
      this.$router.go(-1)
    },
    createTrip: async function () {
      let that = this
      let response = await this.$axios.get(
        'https://api.unsplash.com/search/photos/?client_id=' +
        this.$config.unsplashID +
        '&query=' +
        this.location +
        '&orientation=landscape'
      )
      let responseURL = response.data.results[0].urls.regular

      const data = {
        id: '',
        location: this.location,
        startDate: this.date_start,
        endDate: this.date_end,
        editors: [this.$store.state.user.uid],
        viewers: [],
        imageURL: responseURL,
      }
      this.$axios
        .post('http://localhost:8080/api/trips', data, {
          headers: {
            'Content-Type': 'application/json',
          },
        })
        .then(function (response) {
          that.$router.push('/trips')
        })
        .catch(function (error) {
          console.error(error)
        })
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
