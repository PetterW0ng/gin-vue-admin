import {
    USER_SEX, USER_TOKEN
} from "./mutation-type";

// 引入本地存储
import {
    getLocalStore,
    setLocalStore,
    removeLocalStore
} from '../config/global'

export default {
    // 性别
    [USER_SEX](state) {
        if (state.userInfo.sex == '1') {
            return '女'
        } else if (state.userInfo.sex == '2') {
            return '男';
        } else {
            return '未填写'
        }
    },
    [USER_TOKEN](state) {
        let token = state.token;
        if (!token) {
            token = getLocalStore('token')
            state.token = token
        }
        return token
    }

}
