<template>
  <div >
    <div class="header">
      <div class="logo">
        <h1 class="text-white" style="margin-top: 5px; margin-left: 10px; font-size: 60px">Be part'y</h1>
        <img src="jigsaw.png" width="80" height="80">
      </div>

      <button v-if="this.flg === 0" class="text-white" style="outline: none;background-color: black;margin-left: auto; margin-right: 5em;" data-toggle="modal" data-target="#createGroup"><h4>Create group</h4></button>
      <button v-if="this.flg === 0" class="text-white" style="outline: none;background-color: black;margin-left: 0; margin-right: 5em;" data-toggle="modal" data-target="#codeInput"><h4>Log in</h4></button>
      <button v-if="group_name !== ''" class="text-white" style="outline: none; background-color: black; margin-left: auto; margin-right: 5em"><h4>{{this.group_name.toUpperCase()}}</h4></button>

      <div class="modal fade text-black-50" id="createGroup" data-backdrop="static" tabindex="-1" role="dialog" aria-labelledby="staticBackdropLabel" aria-hidden="true">
        <div class="modal-dialog" role="document">
          <div class="modal-content">
            <div class="modal-header">
              <h5 class="modal-title" id="staticBackdropLabel">Enter your group name</h5>
              <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                <span aria-hidden="true">&times;</span>
              </button>
            </div>
            <div class="modal-body">
              <form>
                <div class="mb-3">
                  <label for="Group-name" class="col-form-label">Group name:</label>
                  <input type="text" class="form-control" id="Group-name" v-model="group_name">
                </div>
              </form>
            </div>
            <div class="modal-footer">
              <button type="button" class="btn btn-primary" data-dismiss="modal" data-toggle="modal" data-target="#codeInfo">Ok</button>
            </div>
          </div>
        </div>
      </div>

      <div class="modal fade text-black-50" id="codeInfo" data-backdrop="static" tabindex="-1" role="dialog" aria-labelledby="staticBackdropLabel" aria-hidden="true">
        <div class="modal-dialog" role="document">
          <div class="modal-content">
            <div class="modal-header">
              <h5 class="modal-title" id="staticBackdropLabel">Your group code</h5>
              <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                <span aria-hidden="true">&times;</span>
              </button>
            </div>
            <div class="modal-body">
              Your group code is #AS18QZA. Share it with your friends to make playlists together
            </div>
            <div class="modal-footer">
              <button type="button" class="btn btn-primary" data-dismiss="modal" v-on:click="changeFlg">Ok</button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="modal fade text-black-50" id="codeInput" data-backdrop="static" tabindex="-1" role="dialog" aria-labelledby="staticBackdropLabel" aria-hidden="true">
      <div class="modal-dialog" role="document">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title" id="staticBackdropLabel">Enter group code</h5>
            <button type="button" class="close" data-dismiss="modal" aria-label="Close">
              <span aria-hidden="true">&times;</span>
            </button>
          </div>
          <div class="modal-body">
            <form>
              <div class="mb-3">
                <label for="Group-code" class="col-form-label">Group code:</label>
                <input type="text" class="form-control" id="Group-code" v-model="group_code">
              </div>
            </form>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-primary" data-dismiss="modal">Log in</button>
          </div>
        </div>
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
    data() {
      return {
        group_name: "",
        group_code: "",
        flg: 0
      }
    },

   methods: {
      changeFlg() {
        this.flg = !this.flg
      }
   }
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
    display: flex;
  }

  .logo {
    display: flex;
  }

  img {
    margin-left: 30px; margin-top: 15px;
  }

</style>
