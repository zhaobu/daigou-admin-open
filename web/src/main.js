import { auth } from '@/directive/auth'
import '@/plugins'
// 引入封装的router
import router from '@/router/index'
import { store } from '@/store/index'
import 'quill/dist/quill.bubble.css'
import 'quill/dist/quill.core.css'
import 'quill/dist/quill.snow.css'
import Vue from 'vue'
// time line css
import '../node_modules/timeline-vuejs/dist/timeline-vuejs.css'
import App from './App.vue'

Vue.config.productionTip = false

// 按钮权限指令
auth(Vue)

new Vue({
  render: (h) => h(App),
  router,
  store
}).$mount('#app')
