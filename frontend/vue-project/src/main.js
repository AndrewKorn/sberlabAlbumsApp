import { createApp } from 'vue'
import App from './App.vue'
import router from "./router"

const app = createApp(App)
export const API_BASE_URL = "_API_BASE_URL"
app.use(router)
app.mount('#app')
