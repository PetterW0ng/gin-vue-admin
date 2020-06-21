import service from '@/utils/request'

// @Tags SysDict
// @Summary 创建SysDict
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysDict true "创建SysDict"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dict/createSysDict [post]
export const createSysDict = (data) => {
    return service({
        url: "/dict/createSysDict",
        method: 'post',
        data
    })
}


// @Tags SysDict
// @Summary 删除SysDict
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysDict true "删除SysDict"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dict/deleteSysDict [delete]
export const deleteSysDict = (data) => {
    return service({
        url: "/dict/deleteSysDict",
        method: 'delete',
        data
    })
}

// @Tags SysDict
// @Summary 更新SysDict
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysDict true "更新SysDict"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dict/updateSysDict [put]
export const updateSysDict = (data) => {
    return service({
        url: "/dict/updateSysDict",
        method: 'put',
        data
    })
}


// @Tags SysDict
// @Summary 用id查询SysDict
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysDict true "用id查询SysDict"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dict/findSysDict [get]
export const findSysDict = (params) => {
    return service({
        url: "/dict/findSysDict",
        method: 'get',
        params
    })
}


// @Tags SysDict
// @Summary 分页获取SysDict列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取SysDict列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dict/getSysDictList [get]
export const getSysDictList = (params) => {
    return service({
        url: "/dict/getSysDictList",
        method: 'get',
        params
    })
}
