<template>
  <div>
    <div v-for="song in songs">
      <iframe width="95%" height="100" :src='"https://w.soundcloud.com/player/?url=https%3A//api.soundcloud.com/tracks/" + song.sound_cloud_id + ""'></iframe>
    </div>
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