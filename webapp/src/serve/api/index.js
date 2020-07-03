import service from './request.js'

/***** 登录界面接口 *********/
// 1.获取手机验证码(GET)
export const getPhoneCaptcha = (data) => {
    return service({
        url: "/user/verificationCode",
        method: 'post',
        data
    })
}

// 2.注册
export const register = (data) => {
    return service({
        url: "/user/register",
        method: 'post',
        data
    })
}

// 3.登录
export const login = (phone, captcha) => {
    return service({
        url: "/user/login",
        method: 'post',
        data: {telphone: phone, verificationCode: captcha}
    })
}

/***** 主页的接口 ************/
// 题目及选项 相关页
export const queryTopicList = (businessType) => {
    return service({
        url: "/topicOption/" + businessType,
        method: 'GET'
    })
}

export const addUserAnswer = (data) => {
    return service({
        url: "/userTopicAnswer/add",
        method: "POST",
        data
    })
}
// 用户答案详情页
export const loadUserTopicAnswers = (businessType, batchNum) => {
    return service({
        url: "/userTopicAnswer?businessType=" + businessType + "&batchNum=" + batchNum,
        method: "GET"
    })
}


// 用户得分页
export const userScore = () => {
    return service({
        url: "/userTopicAnswer/analysis",
        method: "GET"
    })
}

// 课程的连接点击
export const toStore = (relatedId) => {
    return service({
        url: "/userTopicAnswer/toStore?relatedId=" + relatedId + "&random=" + new Date(),
        method: "GET"
    })
}

// home 页
export const loadUserInfo = () => {
    return service({
        url: "/user",
        method: 'GET'
    })
}


// 基本信息页
export const selectOptions = () => {
    return service({
        url: "/dict/baseForm",
        method: 'POST'
    })
}
export const areaList = () => {
    return service({
        url: "/area/tree",
        method: 'POST'
    })
}
export const getArea = (id) => {
    return service({
        url: "/area?id=" + id,
        method: 'POST'
    })
}
export const saveBaseinfo = (data) => {
    return service({
        url: "/userBaseinfo/create",
        method: "POST",
        data
    })
}

export const userBaseinfo = () => {
    return service({
        url: "/userBaseinfo",
        method: "GET"
    })
}

export const initUserBaseinfo = () => {
    return service({
        url: "/userBaseinfo",
        method: "POST"
    })
}




