import service from '@/utils/request'

// @Tags SysApi
// @Summary 获取权限客户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "获取权限客户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cert/getCertHolderList [post]
export const getCertHolderList = (data) => {
    return service({
        url: "/cert/getCertHolderList",
        method: 'post',
        data
    })
}

export const getCertHolder = (data) => {
    return service({
        url: "/cert/certHolder",
        method: 'get',
        data
    })
}

export const deleteCertHolder = (data) => {
    return service({
        url: "/cert/certHolder",
        method: 'delete',
        data
    })
}

export const updateCertHolder = (data) => {
    return service({
        url: "/cert/certHolder",
        method: 'put',
        data
    })
}