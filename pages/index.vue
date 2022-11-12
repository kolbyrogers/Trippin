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
          console.log(err)
        }).then((user) => {
          const newUser = {
            uid: user.user.uid,
            email: user.user.email,
            displayName: user.user.displayName,
          }
          this.createUser(newUser)
          $nuxt.$router.push('/trips')
        })
    },
    createUser: async function (user) {
      const data = {
        'id': user.uid,
        'displayName': user.displayName,
        'email': user.email,
      }
      console.log("Sending user:", JSON.stringify(data));
      this.$axios.post('http://localhost:8080/api/users', data, {
        headers: {
          'Content-Type': 'application/json'
        }
      })
        .then(function (response) {
          console.log(response);
        })
        .catch(function (error) {
          console.log(error);
        });
    },
  }
}
</script>
