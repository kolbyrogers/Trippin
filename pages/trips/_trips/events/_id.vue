<template>
  <v-container>
    <v-row>
      <v-col cols="2">
        <v-btn color="primary" class="white--text" @click="back"> Back </v-btn>
      </v-col>
      <v-col offset="8" cols="2">
        <v-btn color="primary" class="white--text" @click="chooseFiles">
          Add Photos
        </v-btn>
      </v-col>
    </v-row>
    <h1>Event page</h1>
    <input id="fileUpload" type="file" @change="onFileChange" hidden />
    <div class="ma-12 pa-12" v-for="image in images" :key="image.id">
      <v-img :src="image"></v-img>
    </div>
  </v-container>
</template>

<script>
export default {
  data() {
    return {
      images: [],
    }
  },
  mounted() {
    this.getImages();
  },
  methods: {
    chooseFiles() {
      document.getElementById("fileUpload").click()
    },
    back() {
      this.$router.go(-1)
    },
    onFileChange(e) {
      const file = e.target.files[0]
      this.uploadFile(file)
    },
    onSubmit(e) {
      e.preventDefault()
    },
    getImages() {
      const that = this;
      const storageRef = this.$fire.storage.ref()
      const imageListRef = storageRef.child(this.$route.params.id + '/')
      imageListRef.listAll().then(function (res) {
        console.log("rnning listAll");
        res.items.forEach(function (itemRef) {
          console.log(itemRef.getDownloadURL());
          itemRef.getDownloadURL().then(function (res) {
            console.log(res);
            that.images.push(res)
          })
        });
      }).catch(function (error) {
        // Uh-oh, an error occurred!
      });
    },
    uploadFile(file) {
      const that = this;
      that.images = []
      const storageRef = this.$fire.storage.ref()
      const fileRef = storageRef.child(this.$route.params.id + '/' + file.name)
      fileRef.put(file).then(() => {
        console.log('Uploaded a file!')
        const imageListRef = storageRef.child(this.$route.params.id + '/')
        this.getImages()
      })
    },
  },
}
</script>