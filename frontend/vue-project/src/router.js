import {createRouter, createWebHashHistory} from "vue-router";
import ConcretePlaylist from "./ConcretePlaylist.vue";
import HomePage from "./HomePage.vue";
import ConcreteGroup from "./ConcreteGroup.vue";


export default createRouter( {
    history: createWebHashHistory(),
    routes: [
        {path: "/", component: HomePage},
        {path: "/:group_name/albums/:album_name", component: ConcretePlaylist},
        {path: "/:group_name/albums", component: ConcreteGroup}
    ]
})