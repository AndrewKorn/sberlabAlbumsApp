<template>
  <div>
    <h3><strong>Current playlists:</strong></h3>

    <div class="playlistTab" v-if="albums.length > 0">
      <PlaylistTable
          v-bind:albums="albums"
          @deletePlaylist="deletePlaylist"
          @openPlaylist="openPlaylist"
      />
    </div>

    <div v-else>
      <h2 style="text-align: center">There is no albums yet</h2>
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
import {API_BASE_URL} from "./main";
export default {
  name: "Home",
  data() {
    return {
      album_name: "",
      albums: []
    }
  },

  mounted() {
    fetch(API_BASE_URL + "" + this.$route.params.group_name + "/albums")
        .then(response => response.json())
        .then(json => {
          this.albums = json
          for (let i = 0; i < this.albums.length; ++i) {
            fetch(API_BASE_URL + "" + this.$route.params.group_name + "/albums/" + this.albums[i].album_name + "")
                .then(response => response.json())
                .then(json => {
                  this.albums[i].songs = json
                })
        }})
  },

  components: {
    AddPlaylistForm,
    PlaylistTable
  },

  methods: {
    deletePlaylist(name) {
      fetch(API_BASE_URL + "" + this.$route.params.group_name + "/albums/" + name + "" , {method: "DELETE"})
          .then(response => {
            this.albums = this.albums.filter(album => album.album_name !== name)
          })
    },

    createPlaylist(album) {

      fetch(API_BASE_URL + "" + this.$route.params.group_name + "/albums", {
        method: "POST",
        body: JSON.stringify({
          album_name: album.album_name
        })})
          .then(response => {
            this.albums.push(album)
          })
    },

    openPlaylist(album_name) {
      router.push('/' + this.$route.params.group_name + "/albums/" + album_name + "")
    }
  }
}
</script>

<style scoped>
  h3 {
    text-align: left;
    padding-left: 2%;
    padding-bottom: 10px;
  }

  .playlistTab {
    margin-left: 2%;
    margin-right: 2%;
  }
</style>