import {createApp} from 'vue'
import "./assets/css/tailwind.css";
import '@mdi/font/css/materialdesignicons.css'
import 'vuetify/styles'
import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'
import router from './router/router.js'
import App from './App.vue'
import './style.css';
const vuetify = createVuetify({
    components,
    directives
})
const app = createApp(App)
app.use(vuetify)
app.use(router)
app.mount('#app')