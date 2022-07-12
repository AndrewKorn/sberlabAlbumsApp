<template>
  <div>
      <h3>Current playlists:</h3>

      <PlaylistTable
          v-bind:albums="albums"
          @deletePlaylist="deletePlaylist"
          @openPlaylist="openPlaylist"
      />

      <AddPlaylistForm
        @createPlaylist="createPlaylist"
      />
  </div>
</template>

<script>
import PlaylistTable from "./components/PlaylistTable.vue";
import AddPlaylistForm from "./components/AddPlaylistForm.vue";
import router from "./router";
export default {
  name: "Home",
  data() {
    return {
      album_name: "",
      albums: []
    }
  },

  mounted() {
    fetch("http://localhost:8080/albums")
        .then(response => response.json())
        .then(json => {
          this.albums = json
        })
  },

  components: {
    AddPlaylistForm,
    PlaylistTable
  },

  methods: {
    deletePlaylist(name) {
      console.log(name)
      fetch("http://localhost:8080/albums/" + name + "", {method: "DELETE"})
          .then(response => {
            this.albums = this.albums.filter(album => album.album_name !== name)
          })
    },

    createPlaylist(album) {
      console.log(album.album_name)
      fetch("http://localhost:8080/albums", {
        method: "POST",
        body: JSON.stringify({
          album_name: album.album_name
        })})
          .then(response => {
            this.albums.push(album)
          })
    },

    openPlaylist(album_name) {
      router.push('/' + album_name)
    }
  }
}
</script>

<style scoped>
h3 {
  text-align: left;
  padding-top: 50px;
  padding-left: 10px;
}
</style>