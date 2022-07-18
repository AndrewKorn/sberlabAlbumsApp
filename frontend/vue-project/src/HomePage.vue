<template>
  <div>
    <h2>To start creating or editing playlists, you need to make your own new group or log in to an existing one</h2>
    <div class="text-center">
      <div class="spinner-grow text-dark" style="width: 80px; height: 80px;" role="status">
        <span class="sr-only">Loading...</span>
      </div>
    </div>

  </div>
</template>

<script>
export default {
  name: "HomePage"
}
</script>

<style scoped>
  h2 {
    font-family: 'Avenir', Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: #2c3e50;
    margin-top: 100px;
  }
</style>