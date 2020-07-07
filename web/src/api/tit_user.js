import service from '@/utils/request'

// @Tags TitUser
// @Summary 创建TitUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TitUser true "创建TitUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/createTitUser [post]
export const createTitUser = (data) => {
    return service({
        url: "/user/createTitUser",
        method: 'post',
        data
    })
}


// @Tags TitUser
// @Summary 删除TitUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TitUser true "删除TitUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /user/deleteTitUser [delete]
export const deleteTitUser = (data) => {
    return service({
        url: "/user/deleteTitUser",
        method: 'delete',
        data
    })
}

// @Tags TitUser
// @Summary 更新TitUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TitUser true "更新TitUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /user/updateTitUser [put]
export const updateTitUser = (data) => {
    return service({
        url: "/user/updateTitUser",
        method: 'put',
        data
    })
}


// @Tags TitUser
// @Summary 用id查询TitUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TitUser true "用id查询TitUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /user/findTitUser [get]
export const findTitUser = (params) => {
    return service({
        url: "/user/findTitUser",
        method: 'get',
        params
    })
}


// @Tags TitUser
// @Summary 分页获取TitUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取TitUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/getTitUserList [get]
export const getTitUserList = (params) => {
    return service({
        url: "/user/getTitUserList",
        method: 'get',
        params
    })
}


export const findBaseInfo = (params) => {
    return service({
        url: "/user/getTitUserBaseInfo",
        method: 'get',
        params
    })
}
