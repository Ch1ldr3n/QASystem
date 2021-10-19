import { createApp } from 'vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import App from './App.vue'

const app = createApp(App)
import axios from "axios"
app.config.globalProperties.$http = axios;
axios.defaults.baseURL = 'qanda-bauhinia.app.secoder.net/'
app.use(ElementPlus)
app.mount('#app')
