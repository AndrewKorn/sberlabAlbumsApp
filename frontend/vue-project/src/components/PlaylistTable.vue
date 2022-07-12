<template>
  <div>
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