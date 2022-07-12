<template>
  <div class="app">
    <h1>Hello, there you can make your own playlist with your friends!</h1>
    <hr>
    <router-view/>
  </div>
</template>

<script>
  export default {
    name: "app",
    components: {

    }
  }
</script>


<style>
  .app {
    font-family: 'Avenir', Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: #2c3e50;
    margin-top: 60px;
  }
</style>
