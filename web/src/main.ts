import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'

/* 引入 ArcoVue UI组件: https://arco.design/vue/component/icon */
import ArcoVue from '@arco-design/web-vue'
// 额外引入图标库
import ArcoVueIcon from '@arco-design/web-vue/es/icon'
import '@arco-design/web-vue/dist/arco.css'

const app = createApp(App)

/* 引入 ArcoVue UI组件 */
app.use(ArcoVue)
app.use(ArcoVueIcon)

app.use(createPinia())
app.use(router)

app.mount('#app')
