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
        <v-row>
          <v-text-field label="Location" solo></v-text-field>
        </v-row>
        <v-row>
          <!-- Import Google Map Component -->
        </v-row>
        <v-row>
          <v-dialog
            ref="dialog1"
            v-model="modal_start"
            :return-value.sync="date_start"
            persistent
            width="290px"
          >
            <template v-slot:activator="{ on, attrs }">
              <v-text-field
                outlined
                dense
                v-model="date_start"
                label="Start Date"
                readonly
                v-bind="attrs"
                v-on="on"
              ></v-text-field>
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
        <!-- TODO: Add a checker to make sure end date is AFTER start date. Format inputs -->
        <v-row>
          <v-dialog
            ref="dialog2"
            v-model="modal_end"
            :return-value.sync="date_end"
            persistent
            width="290px"
          >
            <template v-slot:activator="{ on, attrs }">
              <v-text-field
                outlined
                dense
                v-model="date_end"
                label="End Date"
                readonly
                v-bind="attrs"
                v-on="on"
              ></v-text-field>
            </template>
            <v-date-picker v-model="date_end" scrollable>
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
        <!-- TODO: Mobile formatting of svg -->
        <img src="/images/new_trip.svg" />
      </v-col>
    </v-row>
    <v-row justify="center">
      <v-btn color="primary" class="white--text" @click="createTrip"
        >Create</v-btn
      >
    </v-row>
  </v-container>
</template>

<script>
export default {
  data() {
    return {
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
    }
  },
  computed: {
    computedDateFormatted() {
      return this.date
    },
  },
  methods: {
    goBack() {
      this.$router.go(-1)
    },
    createTrip() {
      console.log('Create Trip')
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
