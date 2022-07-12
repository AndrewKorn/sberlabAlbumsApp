<template>
  <div>

    <div v-for="song in songs">
      <iframe width="95%" height="100" :src='"https://w.soundcloud.com/player/?url=https%3A//api.soundcloud.com/tracks/" + song.song_id + ""'></iframe><div style="font-size: 10px; color: #cccccc;line-break: anywhere;word-break: normal;overflow: hidden;white-space: nowrap;text-overflow: ellipsis; font-family: Interstate,Lucida Grande,Lucida Sans Unicode,Lucida Sans,Garuda,Verdana,Tahoma,sans-serif;font-weight: 100;"></div>
    </div>

<!--    <table class="table table-hover">
      <thead class="thead-dark">
      <tr>
        <th scope="col">№</th>
        <th scope="col">Singer</th>
        <th scope="col">Song title</th>
        <th scope="col">Duration</th>
        <th scope="col"></th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="(song, index) in songs">
        <th scope="row" v-on:click="$emit('openSong', song)">{{index + 1}}</th >
        <td v-on:click="$emit('openSong', song)">{{song.artist}}</td>
        <td v-on:click="$emit('openSong', song)">{{song.song_name}}</td>
        <td v-on:click="$emit('openSong', song)">{{song.duration}}</td>
        <td>
          <button class="btn btn-outline-danger" v-on:click="$emit('deleteSong', song)">X</button>
        </td>
      </tr>
      </tbody>
    </table>-->
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
  iframe {
    margin-top: 10px;
    margin-bottom: 10px;
  }
</style>