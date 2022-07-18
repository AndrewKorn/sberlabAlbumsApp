<template>
  <form class="form-inline" @submit.prevent="onSubmit">
    <div class="form-group mx-sm-3 mb">
      <input type="text" class="form-control" id="inputNewPlaylist" placeholder="New playlist name" v-model="album_name">
    </div>
    <button class="btn btn-secondary">Create new playlist</button>
  </form>
</template>

<script>
export default {
  name: "AddPlaylistForm",
  data() {
    return {
      album_name: ""
    }
  },

  methods: {
    onSubmit() {
      if (this.album_name.trim()) {
        const newAlbum = {
          album_name: this.album_name,
          songs: []
        }
        this.$emit("createPlaylist", newAlbum)
      }
      this.album_name = ''
    }
  }
}
</script>

<style scoped>
  form {
    width: 30%;
    margin: 0 auto;
  }
</style>