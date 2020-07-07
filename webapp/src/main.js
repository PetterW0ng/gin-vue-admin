import Vue from 'vue'
import App from './App.vue'
import router from './router/router'
import store from './store/store'
import FastClick from 'fastclick'
import VTop from './components/backToTop/ToTop.vue'
import 'element-ui/lib/theme-chalk/index.css';
import './icons' // icon
// 解决移动端点击延迟200ms的问题
if ('addEventListener' in document) {
    document.addEventListener('DOMContentLoaded', function () {
        FastClick.attach(document.body);
    }, false);
}

// 注册为全局组件
Vue.component('v-top', VTop);

// 瀑布流
import waterfall from 'vue-waterfall2'

Vue.use(waterfall)

import {Select, Option, FormItem, Form} from 'element-ui'

Vue.use(Select)
Vue.use(Option)
Vue.use(FormItem)
Vue.use(Form)


Vue.config.productionTip = false

// 全局引入UI库 vant
import '@/plugins/vant'
// 全局引入rem
import '@/config/rem'
// 全局引入 Reset-css
import 'reset-css'
//  全局引入 过滤器
import '@/config/filter'


new Vue({
    router,
    store,
    render: h => h(App)
}).$mount('#app')
