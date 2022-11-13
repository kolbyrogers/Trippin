<template>
  <div>
    <div id="sky">
      <div id="cloud1" class="cloud"></div>
      <div id="cloud2" class="cloud"></div>
      <div id="cloud3" class="cloud"></div>
      <div id="cloud4" class="cloud"></div>
      <div id="cloud5" class="cloud"></div>
      <div id="cloud6" class="cloud"></div>
      <div id="cloud7" class="cloud"></div>
      <div id="aeroplane"></div>
    </div>
    <h1>Trippin'</h1>
    <p>Vacations are better shared.</p>
    <button @click="signInWithGoogle()" type="button" class="login-with-google-btn">
      Sign in with Google
    </button>
  </div>

</template>

<script>
export default {
  name: 'IndexPage',
  methods: {
    signInWithGoogle: async function () {
      const provider = new $nuxt.$fireModule.auth.GoogleAuthProvider()
      this.$fire.auth.signInWithPopup(provider)
        .catch(function (err) {
          console.error(err)
        }).then((res) => {
          const newUser = {
            uid: res.user.uid,
            email: res.user.email,
            displayName: res.user.displayName,
          }
          this.createUser(newUser).then(() => {
            this.$router.push('/trips')
          })
        })
    },
    createUser: async function (user) {
      const data = {
        'id': user.uid,
        'displayName': user.displayName,
        'email': user.email,
      }
      this.$axios.post('http://localhost:8080/api/users', data, {
        headers: {
          'Content-Type': 'application/json'
        }
      })
        .then(function (response) {
        })
        .catch(function (error) {
          console.error(error);
        });
    },
  }
}
</script>
