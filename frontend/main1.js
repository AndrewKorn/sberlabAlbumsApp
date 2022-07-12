const app = new Vue({
    el: '#app',
    data: {
        message: 'Hello, there you can make your own playlist with your friends!',
        albums:[],
        albumName:"",
        vis:"visibility: hidden"
    },
    methods: {
        loadPage: function () {
            this.$http.get("http://localhost:8080/albums").then(response => {
                response.data.forEach(element => {
                    this.albums.push(element)
                })
            }, () => {
                console.log("error")
            });
        },

        openPlaylist: function (name) {
            this.$http.get("http://localhost:8080/albums/" + name + "").then(response => {
                console.log(response.data)
                this.vis = "visibility: visible"
            }, () => {
                console.log("err")
            });
        },

        createPlaylist: function () {
            if (this.albumName !== "") {
                this.$http.post("http://localhost:8080/albums", {album_name: this.albumName}).then(response => {
                    console.log(this.albumName)
                }, () => {
                    console.log("err")
                });
            }
        },

        deletePlaylist: function (name) {
            this.$http.delete("http://localhost:8080/albums/" + name + "").then(response => {

                const index = this.albums.indexOf(name);
                if (index !== -1) {
                    this.albums.splice(index, 1)
                }
                window.location.reload()
            }, () => {
                console.log("err")
            })

        }
    },
    beforeMount() {
        this.loadPage()
    }
})