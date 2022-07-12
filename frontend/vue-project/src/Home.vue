<template>
  <div>
    <h3><strong>Current playlists:</strong></h3>

      <div class="pltab">
        <PlaylistTable
            v-bind:albums="albums"
            @deletePlaylist="deletePlaylist"
            @openPlaylist="openPlaylist"
        />
      </div>

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
    fetch("http://localhost:1337/albums")
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
      fetch("http://localhost:1337/albums/" + name + "", {method: "DELETE"})
          .then(response => {
            this.albums = this.albums.filter(album => album.album_name !== name)
          })
    },

    createPlaylist(album) {
      console.log(album.album_name)
      fetch("http://localhost:1337/albums", {
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
    padding-left: 2%;
    padding-bottom: 10px;
  }

  .pltab {
    margin-left: 2%;
    margin-right: 2%;
  }
</style>