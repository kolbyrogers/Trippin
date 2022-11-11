<template>
  <v-container>
    <v-row>
      <h1>Share</h1>
    </v-row>
    <v-row>
      <v-text-field> </v-text-field>
    </v-row>
    <v-row>
      <v-radio label="Has Edit Access" v-model="can_edit"></v-radio>
    </v-row>
    <v-row>
      <v-btn color="primary" class="color--text" @click="shareTrip"
        >Share</v-btn
      >
    </v-row>
    <v-row>
      <img src="/images/share.svg" />
    </v-row>
  </v-container>
</template>

<script>
export default {
  data() {
    return {
      can_edit: false,
    }
  },
  methods: {
    shareTrip() {
      console.log('Share Trip')
      this.router.go(-1)
    },
  },
}
</script>

<style></style>
