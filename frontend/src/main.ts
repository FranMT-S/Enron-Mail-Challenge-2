import { createApp } from 'vue'
import App from './App.vue'
import router from './routes'
import { pinia } from './plugins/pinia'


import './styles.css'

createApp(App)
.use(router)
.use(pinia)
.mount('#app')
