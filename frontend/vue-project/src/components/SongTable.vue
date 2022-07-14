<template>
  <div>

    <div v-for="song in songs">
      <iframe width="95%" height="100" :src='"https://w.soundcloud.com/player/?url=https%3A//api.soundcloud.com/tracks/" + song.song_id + ""'></iframe><div style="font-size: 10px; color: #cccccc;line-break: anywhere;word-break: normal;overflow: hidden;white-space: nowrap;text-overflow: ellipsis; font-family: Interstate,Lucida Grande,Lucida Sans Unicode,Lucida Sans,Garuda,Verdana,Tahoma,sans-serif;font-weight: 100;"></div>
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