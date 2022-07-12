<template>
  <div>
    <table class="table table-hover">
      <thead class="thead-dark">
      <tr>
        <th scope="col">№</th>
        <th scope="col">Singer</th>
        <th scope="col">Song title</th>
<!--        <th scope="col"></th>-->
      </tr>
      </thead>
      <tbody>
      <tr v-for="(song, index) in songs">
        <th scope="row" v-on:click="$emit('openSong', song)">{{index + 1}}</th >
        <td v-on:click="$emit('openSong', song)">{{song.artist}}</td>
        <td v-on:click="$emit('openSong', song)">{{song.song_name}}</td>
<!--        <td>
          <button class="btn btn-outline-danger" v-on:click="$emit('deleteSong', song)">X</button>
        </td>-->
      </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  name: "SongTable",
  props: ["songs"],
  components: {}
}
</script>

<style scoped>
ul {
  list-style: none;
  margin: 0;
  padding: 0;
}
</style>