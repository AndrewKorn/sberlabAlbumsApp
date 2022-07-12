import {createRouter, createWebHashHistory} from "vue-router";
import ConcretePlaylist from "./ConcretePlaylist.vue";
import PlaylistTable from "./components/PlaylistTable.vue";
import App from "./App.vue";
import Home from "./Home.vue";


export default createRouter( {
    history: createWebHashHistory(),
    routes: [
        {path: "/", component: Home},
        {path: "/:name", component: ConcretePlaylist},
    ]
})