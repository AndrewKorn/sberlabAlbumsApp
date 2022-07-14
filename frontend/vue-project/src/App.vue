<template>
  <div >
    <div class="header">
      <div class="qwe">
        <h1 style="margin-top: 5px; margin-left: 10px; font-size: 60px">Be part'y</h1>
        <img src="jigsaw.png" width="80" height="80">
      </div>
    </div>
    <h1 class="hello">Hello, there you can make your own playlist with your friends!</h1>
    <hr/>
    <router-view/>
  </div>
</template>

<script>
 export default {
    name: "app",
    components: {},
  }
</script>


<style scoped>
  .hello{
    font-family: 'Avenir', Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: #2c3e50;
    margin-top: 60px;
  }

  .header {
    color: #cccccc;
    text-align: left;
    height: 100px;
    margin-bottom: 10px;
    background-color: black;
  }

  .qwe {
    display: flex;
  }

  img {
    margin-left: 30px; margin-top: 15px;
  }
</style>
