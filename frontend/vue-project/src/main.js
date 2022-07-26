import { createApp } from 'vue'
import App from './App.vue'
import router from "./router"

const app = createApp(App)

export const API_BASE_URL = "_API_BASE_URL"
export const API_PARSER_URL = "_API_PARSER_URL"
/*export const API_BASE_URL = "http://localhost:1337/"
export const API_PARSER_URL = "http://localhost:7331/"*/

app.use(router)
app.mount('#app')
