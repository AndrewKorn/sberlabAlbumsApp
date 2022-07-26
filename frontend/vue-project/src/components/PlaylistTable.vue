<template>
  <div>

    <div :class='"scrolling-wrapper row row-cols-1 row-cols-md-3"'>
      <div class="col mb-4" v-for="(album, index) in albums">
        <div class="card h-100 bg-light text-center p-3">
          <div class="modal-header">
            <h4 class="modal-title" v-on:click="$emit('openPlaylist', album.album_name)">{{album.album_name}}</h4>
            <button type="button" class="close" data-dismiss="modal" aria-label="Close" v-on:click="$emit('deletePlaylist', album.album_name)">
              <span aria-hidden="true">&times;</span>
            </button>
          </div>
          <div v-on:click="$emit('openPlaylist', album.album_name)">
          <img :src='"" + this.pic[index % this.pic.length] + ""' height="200" class="card-img-top"><!--
          <div class="card-body">
            <p style="text-align: left; margin-top: 5px; margin-bottom: 0"><strong>Preview:</strong></p>
            <p class="card-text" v-for="song in album.songs" style="text-align: left; margin-top: 5px; margin-bottom: 5px;" v-if="album.songs.length > 0">
              {{song.artist}} - {{song.song_name}}
            </p>
            <p class="card-text" v-else>There is no songs yet</p>
          </div>-->
          </div>
        </div>
      </div>
    </div>

  </div>
</template>

<script>
export default {
  name: "PlaylistTable",
  props: ["albums"],
  components: {},

  data() {
    return {
      pic: ["pic1.jpg", "pic2.jpg", "pic3.jpg", "pic4.jpg", "pic5.jpg"]
    }
  },

  mounted() {
    setTimeout(() => {
      for (let i = 0; i < this.albums.length; ++i) {
        console.log(this.albums[i].songs.length)
      }
    }, 1000)
  }
}
</script>

<style scoped>
.scrolling-wrapper {
  display: flex;
  flex-wrap: nowrap;
  overflow-x: auto;
}
.card {
  flex: 0 0 auto;
}

</style>