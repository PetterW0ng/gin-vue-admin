import Vue from 'vue'
import Router from 'vue-router'

// 引入一级组件
import {USER_TOKEN} from "../store/mutation-type";
import store from '../store/store'

// 懒加载二级组件 Tarbar
const Home = () => import('../views/mine/Home.vue');
const Baseinfo = () => import('../views/mine/Children/Baseinfo.vue')
const TopicOptions = () => import('../views/mine/Children/TopicOptions.vue')
const UserTopicAnswer = () => import('../views/mine/Children/UserTopicAnswer.vue')
const UserBaseinfo = () => import('../views/mine/Children/UserBaseinfo.vue')
const UserScore = () => import('../views/mine/Children/UserScore.vue')
const EmptyScore = () => import('../views/mine/Children/Empty.vue')

// 解决多次点击重复路由报错
const originalPush = Router.prototype.push
Router.prototype.push = function push(location) {
    return originalPush.call(this, location).catch(err => err)
}
// 注册登录
const Login = () => import('../views/login/Login.vue');
Vue.use(Router)

const router = new Router({
    // 解决路由跳转页面没有置顶
    scrollBehavior(to, from, savedPosition) {
        if (savedPosition) {
            return savedPosition
        } else {
            return {
                x: 0,
                y: 0
            }
        }
    },
    //mode: 'history',
    // !!注意: 二级路由不需要加 '/'
    routes: [
        {
            path: '/',
            redirect: '/home',
            // 是否数据缓存
            meta: {
                keepAlive: true
            },
        }, {
            // 基本信息 新增
            path: '/baseinfo',
            name: 'baseinfo',
            component: Baseinfo,
            meta: {
                requireAuth: true,
                keepAlive: true
            }
        }, {
            // 用户基本信息 详情
            path: '/userBaseinfo',
            name: 'userBaseinfo',
            component: UserBaseinfo,
            meta: {
                requireAuth: true,
                keepAlive: true
            }
        }, {
            // 用户得分 及 分析情况
            path: '/userScore',
            name: 'userScore',
            component: UserScore,
            meta: {
                requireAuth: true,
                keepAlive: true
            }
        }, {
            // 题目及选项
            path: '/topicOptions/:topicType',
            name: 'topicOptions',
            component: TopicOptions,
            meta: {
                requireAuth: true,
                keepAlive: true
            }
        }, {
            // 题目没有答完
            path: '/emptyScore',
            name: 'emptyScore',
            component: EmptyScore,
            meta: {
                requireAuth: true,
                keepAlive: true
            }
        }, {
            // 用户答题情况
            path: '/userTopicAnswer/:topicType',
            name: 'userTopicAnswer',
            component: UserTopicAnswer,
            meta: {
                requireAuth: true,
                keepAlive: true
            }
        }, {
            // homePage
            path: '/home',
            name: 'home',
            component: Home,
            meta: {
                requireAuth: true
            },
            children: []
        }, {
            // 登录 或 注册
            path: '/Login',
            name: 'login',
            component: Login
        }
    ]
})

//路由守卫
router.beforeEach((to, from, next) => {
    if (to.meta.requireAuth) {
        if (store.getters[USER_TOKEN]) {
            next()
        } else {
            next({
                path: '/login'
            })
        }
    } else {
        next()
    }
})

export default router
