<template>
  <v-container>
    <v-row>
      <v-col cols="2">
        <v-btn @click="goBack" color="primary" class="white--text">
          Back
        </v-btn>
      </v-col>
    </v-row>
    <v-row justify="center">
      <h1>Share</h1>
    </v-row>
    <v-row justify="center" width="50%">
      <v-col cols="4"></v-col>
      <v-col cols="4">
        <v-text-field solo placeholder="Enter email address" v-model="email">
        </v-text-field>
      </v-col>
      <v-col cols="4"></v-col>
    </v-row>
    <v-row justify="center" class="pt-2">
      <v-btn color="primary" class="color--text" @click="shareTrip" :disabled="validEmail">Share</v-btn>
    </v-row>
    <v-row justify="center">
      <img src="/images/share.svg" />
    </v-row>
  </v-container>
</template>

<script>
export default {
  data() {
    return {
      email: '',
      can_edit: false,
    }
  },
  methods: {
    shareTrip: async function () {
      let that = this;
      const data = {
        email: this.email,
        editor: false,
      }
      this.$axios.put('http://localhost:8080/api/trips/' + this.$route.params.trips, data, {
        headers: {
          'Content-Type': 'application/json',
        }
      })
        .then(function (response) {
          if (response.status == 200) {
            that.goBack();
          }
        })
        .catch(function (error) {
          that.$toast.show('Error sharing trip. Check the email and try again', {
            // override the global option
            type: 'error',
            duration: 3000,
          })
        });
    },
    goBack() {
      this.$router.go(-1)
    },
  },
  computed: {
    validEmail() {
      return !String(this.email)
        .toLowerCase()
        .match(
          /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
        )
    },
  },
}
</script>

<style>

</style>
