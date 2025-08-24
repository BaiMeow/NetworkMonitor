import { createApp } from 'vue'
import App from './App.vue'
import './style.css'
import 'element-plus/theme-chalk/dark/css-vars.css'
import router from './router'

const app = createApp(App)
app.use(router)
app.mount('#app')
