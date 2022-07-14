<template>
  <div class="app">
    <h3>{{$route.params.name}}:</h3>

    <SongTable
      v-bind:songs="songs"
      @openSong="openSong"
      @deleteSong="deleteSong"
    />

    <AddSongForm
      @createSong="createSong"
    />
  </div>
</template>

<script>

import SongTable from "./components/SongTable.vue";
import AddSongForm from "./components/AddSongForm.vue";
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
    this.album_name = this.$route.params.name
    console.log(this.album_name)
    fetch("http://localhost:1337/albums/" + this.album_name + "")
        .then(response => response.json())
        .then(json => {
          this.songs = json.songs
        })
  },

  methods: {
    openSong(song) {
      const options = {
        method: 'GET',
        headers: {
          'X-RapidAPI-Key': '24c67493b3mshab853c823f5dadfp1c2d93jsn0996aee90e2c',
          'X-RapidAPI-Host': 'genius.p.rapidapi.com'
        }
      };

      fetch('https://genius.p.rapidapi.com/search?q=' + song.artist + '%20' + song.song_name + '', options)
          .then(response => response.json())
          .then(response =>  {
            //console.log(response.response.hits[0].result)
            window.location = response.response.hits[0].result.url
          })
          .catch(err => console.error(err));
    },

    deleteSong(song) {
      console.log("http://localhost:1337/albums/" + this.album_name + "/" + song.id + "")
      fetch("http://localhost:1337/albums/" + this.album_name + "/" + song.id + "", {method: "DELETE"})
          .then(response => {
            console.log(response)
            this.songs = this.songs.filter(s => s.song_name !== song.song_name)
          })
    },

    createSong(song) {
      const options = {
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
            song.song_id = response.tracks.items[0].id.toString()
            song.duration = response.tracks.items[0].durationText
            console.log(song.song_id)
          })
          .then(response => {
                fetch("http://localhost:1337/albums/" + this.album_name + "", {
                  method: "POST",
                  body: JSON.stringify({
                    song_id: song.song_id,
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
    margin-top: 60px;
  }

  h3 {
    text-align: left;
    padding-top: 50px;
    padding-left: 10px;
  }
</style>