import Vue from 'vue'
import Moment from "moment";
// 金钱过滤 ¥xx.xx
Vue.filter('moneyFormat', (value) => {
    return '¥' + Number(value).toFixed(2);
})

Vue.filter('dateFormat', (value, format) => {
    return Moment(value).format(format ? format : 'YYYY-MM-DD HH:mm:ss')
})
