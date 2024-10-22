import {createApp, reactive} from 'vue'
import App from './App.vue'
import router from './router'
import axios from './services/axios.js';
import utils from './services/utils.js';
import LoadingSpinner from './components/LoadingSpinner.vue'

import ErrorMsg from './components/ErrorMsg.vue'
import Post from './components/Post.vue'
import ProfilesList from './components/ProfilesList.vue'

import './assets/main.css'
import './assets/home.css'
import './assets/login_style.css'
import './assets/profile_style.css'

import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { fas } from '@fortawesome/free-solid-svg-icons'
import { far } from '@fortawesome/free-regular-svg-icons'


const app = createApp(App)
app.config.globalProperties.$axios = axios;
app.config.globalProperties.$utils = utils;

//--- COMPONENT ---//
app.component("ErrorMsg", ErrorMsg);
app.component("Post", Post);
app.component("ProfilesList", ProfilesList);
app.component("LoadingSpinner", LoadingSpinner);
//---------------//


library.add(fas);
library.add(far);

app.component('font-awesome-icon', FontAwesomeIcon)

app.use(router)
app.mount('#app')
