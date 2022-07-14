<template>
  <div>
<!--    <div class="row row-cols-1 row-cols-md-3" style="margin: 0 auto; width: 80%">
      <div class="col mb-4" v-for="album in albums" v-on:click="$emit('openPlaylist', album.album_name)">
        <div class="card h-100 bg-light text-center p-3">
          <div class="card-body">
            <h5 class="card-title" style="text-align: center">{{album.album_name}}</h5>
            <p style="text-align: left; margin-top: 5px; margin-bottom: 0px"><strong>First 3 songs:</strong></p>
            <p class="card-text" v-for="song in album.songs.slice(0, 3)" style="text-align: left; margin-top: 5px; margin-bottom: 5px;">
              {{song.artist}} - {{song.song_name}} ({{song.duration}})
            </p>
          </div>
        </div>
      </div>
    </div>-->


    <table class="table">
      <thead class="thead-light">
      <tr>
        <th scope="col"><strong>№</strong></th>
        <th scope="col"><strong>Playlist Name</strong></th>
        <th scope="col"><strong>Songs count</strong></th>
        <th scope="col"></th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="(album, index) in albums">
        <th scope="row" v-on:click="$emit('openPlaylist', album.album_name)">{{index + 1}}</th >
        <td v-on:click="$emit('openPlaylist', album.album_name)">{{album.album_name}}</td>
        <td v-on:click="$emit('openPlaylist', album.album_name)">{{album.songs.length}}</td>
        <td>
          <button class="btn btn-outline-danger" v-on:click="$emit('deletePlaylist', album.album_name)">X</button>
        </td>
      </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  name: "PlaylistTable",
  props: ["albums"],
  components: {}
}
</script>

<style scoped>

  th {
    text-align: center;
  }
  td {
    text-align: center;
  }
</style>