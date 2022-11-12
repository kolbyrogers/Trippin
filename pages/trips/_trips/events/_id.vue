<template>
    <div>
        <h1>Event page</h1>
        <input type="file" @change="onFileChange" />
        <div class="ma-12 pa-12" v-for="image in images" :key="image.id">
            <v-img :src="image"></v-img>
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            images: [],
        }
    },
    methods: {
        onFileChange(e) {
            const file = e.target.files[0]
            this.uploadFile(file)
        },
        onSubmit(e) {
            e.preventDefault()
        },
        uploadFile(file) {
            const that = this;
            that.images = []
            const storageRef = this.$fire.storage.ref()
            const fileRef = storageRef.child(this.$route.params.id + '/' + file.name)
            fileRef.put(file).then(() => {
                console.log('Uploaded a file!')
                const imageListRef = storageRef.child(this.$route.params.id + '/')
                console.log(imageListRef);
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
            })
        },
    },
}
</script>