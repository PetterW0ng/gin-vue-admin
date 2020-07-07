// 引入mutation-type
import {
    USER_INFO,
    INIT_USER_INFO,
    CHANGE_USER_NICK_NAME,
    USER_INFO_BRITHDAY,
    USER_INFO_SEX,
    LOGIN_OUT,
    USER_TOKEN
} from './mutation-type'
import Vue from 'vue'

// 引入本地存储
import {
    getLocalStore,
    setLocalStore,
    removeLocalStore
} from '../config/global'

export default {
    // 注意:外界传值的参数一定要和定义的参数一致 例如 goodsID  isCheckedAll
    // 1.保存用户信息到本地
    [USER_INFO](state, {
        userInfo
    }) {
        // 1.1 把外界传来的userInfo保存到state中的userInfo
        state.userInfo = userInfo;
        // 1.2 保存到本地缓存中
        setLocalStore('userInfo', state.userInfo);
    },
    [USER_TOKEN](state, {
        token
    }) {
        // 1.1 把外界传来的userInfo保存到state中的userInfo
        state.token = token;
        // 1.2 保存到本地缓存中
        setLocalStore('token', state.token);
    },
    //  2.初始化获取用户信息
    [INIT_USER_INFO](state) {
        // 2.1 先存本地用户数据
        let initUserInfo = getLocalStore('userInfo');
        if (initUserInfo) {
            state.userInfo = JSON.parse(initUserInfo);
        }
    },
    // 3.修改昵称
    [CHANGE_USER_NICK_NAME](state, {
        nickName
    }) {
        // 3.1 从state中取出userInfo
        let userInfo = state.userInfo;
        // 3.2 遍历userInfo的key取出User_name,替换Value值
        Object.keys(userInfo).forEach((info, index) => {
            if (info == 'user_name') {
                userInfo['user_name'] = nickName;
            }
        });
        // 3.3 同步state数据
        state.userInfo = {
            ...userInfo
        };
        // 3.4 将数据更新到本地
        setLocalStore('userInfo', state.userInfo);
    },

    // 4.用户生日
    [USER_INFO_BRITHDAY](state, {
        brithday
    }) {
        // 4.1 取出state中的用户信息
        let userInfo = state.userInfo;
        // 4.2 遍历userInfo的value值
        Object.values(userInfo).forEach((info, index) => {
            // 4.3 判断是否有brithday
            if (info.brithday) {
                info.brithday = brithday;
            } else {
                Vue.set(userInfo, 'brithday', brithday);
            }
        });
        // 4.4 同步state数据
        state.userInfo = {
            ...userInfo
        };
        // 4.5 将数据更新到本地
        setLocalStore('userInfo', state.userInfo);
    },
    // 5.用户性别
    [USER_INFO_SEX](state, {
        sex
    }) {
        // 5.1 取出用户信息
        let userInfo = state.userInfo;
        Object.values(userInfo).forEach((info, index) => {
            // 5.2 判断是否有sex
            if (info.sex) {
                info.sex = sex;
            } else {
                Vue.set(userInfo, 'sex', sex);
            }
        });
        // 5.3 同步state数据
        state.userInfo = {
            ...userInfo
        };
        // 5.4 将数据更新到本地
        setLocalStore('userInfo', state.userInfo);
    },

    // 6. 退出登录
    [LOGIN_OUT](state) {
        state.userInfo = {};
        state.token = ''
        removeLocalStore('userInfo');
        removeLocalStore('token');
    }
}
