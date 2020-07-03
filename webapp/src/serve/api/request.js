import axios from 'axios'; // 引入axios
import {USER_TOKEN, LOGIN_OUT} from "../../store/mutation-type";
import {Toast} from 'vant';
import store from '../../store/store'

const service = axios.create({
    // baseURL: "/api/tit",
    baseURL: "http://titserve.pkucarenjk.com/tit",
    timeout: 6000
})
Toast.allowMultiple()
let acitveAxios = 0
let toast
const showLoading = () => {
    acitveAxios++
    if (acitveAxios > 0) {
        toast = Toast.loading({
            duration: 1000, // 持续展示 toast
            forbidClick: true,
            message: 'loading...',
        });
    }
}

const closeLoading = () => {
    acitveAxios--
    if (acitveAxios <= 0) {
        toast && toast.clear()
    }
}
//http request 拦截器
service.interceptors.request.use(
    config => {
        showLoading()
        const token = store.getters[USER_TOKEN]
        config.data = JSON.stringify(config.data);
        config.headers = {
            'Content-Type': 'application/json',
            'x-token': token
        }
        return config;
    },
    error => {
        closeLoading()
        Toast.fail(error);
        return Promise.reject(error);
    }
);


//http response 拦截器
service.interceptors.response.use(
    response => {
        closeLoading()
        if (response.data.code == 0 || response.headers.success === "true") {
            return response.data
        } else {
            Toast.fail(response.data.msg || decodeURI(response.headers.msg));
            if (response.data.data && response.data.data.reload) {
                store.commit(LOGIN_OUT)
            }
            return Promise.reject(response.data.msg)
        }
    },
    error => {
        closeLoading()
        Toast.fail(error);
        return Promise.reject(error)
    }
)

export default service

