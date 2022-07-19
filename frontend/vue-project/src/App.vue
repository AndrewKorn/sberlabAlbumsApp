<template>
  <div >
    <div class="header">
      <div class="logo">
        <h1 class="text-white" style="margin-top: 5px; margin-left: 10px; font-size: 60px">Be part'y</h1>
        <img src="jigsaw.png" width="80" height="80">
      </div>

      <button v-if="this.$route.path === '/'" class="text-white" style="outline: none; border: none; background-color: black;margin-left: auto; margin-right: 5em;" data-toggle="modal" data-target="#createGroup"><h4>Create group</h4></button>
      <button v-if="this.$route.path === '/'" class="text-white" style="outline: none; border: none; background-color: black;margin-left: 0; margin-right: 5em;" data-toggle="modal" data-target="#codeInput"><h4>Log in</h4></button>
      <div class="dropdown" v-if="this.$route.path !== '/'" style="margin-left: auto">
        <button class="text-white dropdown-toggle" style="outline: none; background-color: black; margin-left: auto; margin-right: 10em; margin-top: 25px" type="button" id="dropdownMenuButton" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
          <h4>{{this.$route.params.group_name}}</h4>
        </button>
        <div class="dropdown-menu" aria-labelledby="dropdownMenuButton">
          <a class="dropdown-item" href="#">Back to home page</a>
          <a class="dropdown-item" href="#" data-toggle="modal" data-target="#codeInput">Log in to another group</a>
          <a class="dropdown-item" href="#" data-toggle="modal" data-target="#codeInfo">Get group code</a>
        </div>
      </div>


      <div class="modal fade text-black-50" id="createGroup" data-backdrop="static" tabindex="-1" role="dialog" aria-labelledby="staticBackdropLabel" aria-hidden="true" v-on:focus="makeID">
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
              Your group code is {{this.group_code}}. Share it with your friends to make playlists together
            </div>
            <div class="modal-footer">
              <button v-if="this.$route.path === '/'" type="button" class="btn btn-primary" data-dismiss="modal" v-on:click="createGroup">Ok</button>
              <button v-if="this.$route.path !== '/'" type="button" class="btn btn-primary" data-dismiss="modal">Ok</button>
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
                  <input type="text" class="form-control" id="Group-code" v-model="input">
                </div>
              </form>
            </div>
            <div class="modal-footer">
              <button type="button" class="btn btn-primary" data-dismiss="modal" v-on:click="logIn">Log in</button>
            </div>
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
 import router from "./router";

 export default {
    name: "app",
    components: {},
    data() {
      return {
        input: "",
        group_name: "",
        group_code: "",
        groups: []
      }
    },

   mounted() {
     this.input = ""
     fetch("http://localhost:1337/")
         .then(response => response.json())
         .then(json => {
           console.log(json)
           this.groups = json
         })
   },

   methods: {
      makeID() {
        let result = '';
        const characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
        const charactersLength = characters.length;
        for (let i = 0; i < 12; i++ ) {
          result += characters.charAt(Math.floor(Math.random() *
              charactersLength));
        }
        this.group_code = result
      },

      createGroup() {
        fetch("http://localhost:1337/", {
          method: "POST",
          body: JSON.stringify({
            group_code: this.group_code,
            group_name: this.group_name
          })})
            .then(response => {
              const newGroup = {
                group_code: this.group_code,
                group_name: this.group_name,
                albums: []
              }

              this.groups.push(newGroup)
            })
        router.push("/" + this.group_name + "/albums")
      },

      goBack() {
        this.group_code = ""
        this.group_name = ""
        this.input = ""
        router.push("/")
      },

      logIn() {
        if (this.input.trim()) {
          for (let i = 0; i < this.groups.length; ++i) {
            if (this.groups[i].group_code === this.input) {
              this.input = ""
              router.push("/" + this.groups[i].group_name + "/albums")
            }
          }
        }
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
