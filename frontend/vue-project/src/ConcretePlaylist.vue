<template>
  <div class="app">

    <div class="top">
      <button class="btn btn-outline-dark" v-on:click="backBtn"><strong>◀</strong></button>
      <h3><strong>{{$route.params.album_name}}:</strong></h3>
    </div>

    <SongTable
      v-if="songs.length > 0"
      v-bind:songs="songs"
      @deleteSong="deleteSong"
    />

    <div v-else>
      <h2 style="text-align: center">There is no songs yet</h2>
    </div>

    <AddSongForm
      @createSong="createSong"
    />
  </div>
</template>

<script>

import SongTable from "./components/SongTable.vue";
import AddSongForm from "./components/AddSongForm.vue";
import router from "./router";
import {API_BASE_URL, API_PARSER_URL} from "./main";
export default {
  name: "ConcretePlaylist",
  components: {AddSongForm, SongTable},
  data() {
    return {
      album_name: "",
      songs: []
    }
  },


  mounted() {
    this.album_name = this.$route.params.album_name
    console.log(this.album_name)
    fetch(API_BASE_URL + "" + this.$route.params.group_name + "/albums/" + this.album_name + "")
        .then(response => response.json())
        .then(json => {
          this.songs = json
        })
  },

  methods: {
    deleteSong(song) {
      fetch(API_BASE_URL + "" + this.$route.params.group_name + "/albums/" + this.album_name + "/" + song.id + "", {method: "DELETE"})
          .then(response => {
            console.log(response)
            this.songs = this.songs.filter(s => s.song_name !== song.song_name)
          })
    },

    createSong(song) {
      fetch(API_PARSER_URL + "" + song.artist + "/" + song.song_name + "")
          .then(response => response.json())
          .then(json => {
            song.sound_cloud_id = json
          })
          .then(response => {
                fetch(API_BASE_URL + "" + this.$route.params.group_name + "/albums/" + this.album_name + "", {
                  method: "POST",
                  body: JSON.stringify({
                    sound_cloud_id: song.sound_cloud_id,
                    artist: song.artist,
                    song_name: song.song_name,
                    duration: song.duration
                  })
                })
                    .then(response => {
                          this.songs.push(song)
                        }
                    )
              }
          )
      /*const options = {
        method: 'GET',
        headers: {
          'X-RapidAPI-Key': '24c67493b3mshab853c823f5dadfp1c2d93jsn0996aee90e2c',
          'X-RapidAPI-Host': 'soundcloud-scraper.p.rapidapi.com'
        }
      };

      fetch('https://soundcloud-scraper.p.rapidapi.com/v1/search/tracks?term=' + song.artist + '%20' + song.song_name + '', options)
          .then(response => response.json())
          .then(response => {
            console.log(response)
            song.sound_cloud_id = response.tracks.items[0].id.toString()
            song.duration = response.tracks.items[0].durationText
            console.log(song.song_id)
          })
          .then(response => {
                fetch(API_BASE_URL + "" + this.$route.params.group_name + "/albums/" + this.album_name + "", {
                  method: "POST",
                  body: JSON.stringify({
                    sound_cloud_id: song.sound_cloud_id,
                    artist: song.artist,
                    song_name: song.song_name,
                    duration: song.duration
                  })
                })
                    .then(response => {
                          this.songs.push(song)
                        }
                    )
              }
          )*/
    },

    backBtn() {
      router.push('/' + this.$route.params.group_name + "/albums")
    }
  }
}
</script>

<style scoped>
  .app {
    font-family: 'Avenir', Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: #2c3e50;
  }

  h3 {
    text-align: left;
    padding-left: 1%;
  }

  .top {
    display: flex;
  }

  button {
    text-align: center;
    height: 35px;
    margin-left: 2.5%;
  }
</style>