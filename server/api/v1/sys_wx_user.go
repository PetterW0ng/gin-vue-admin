package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
)

// @Tags SysWxUser
// @Summary 创建SysWxUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysWxUser true "创建SysWxUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wxUser/createSysWxUser [post]
func CreateSysWxUser(c *gin.Context) {
	var wxUser model.SysWxUser
	_ = c.ShouldBindJSON(&wxUser)
	err := service.CreateSysWxUser(wxUser)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags SysWxUser
// @Summary 删除SysWxUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysWxUser true "删除SysWxUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wxUser/deleteSysWxUser [delete]
func DeleteSysWxUser(c *gin.Context) {
	var wxUser model.SysWxUser
	_ = c.ShouldBindJSON(&wxUser)
	err := service.DeleteSysWxUser(wxUser)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SysWxUser
// @Summary 更新SysWxUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysWxUser true "更新SysWxUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wxUser/updateSysWxUser [put]
func UpdateSysWxUser(c *gin.Context) {
	var wxUser model.SysWxUser
	_ = c.ShouldBindJSON(&wxUser)
	err := service.UpdateSysWxUser(&wxUser)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags SysWxUser
// @Summary 用id查询SysWxUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysWxUser true "用id查询SysWxUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wxUser/findSysWxUser [get]
func FindSysWxUser(c *gin.Context) {
	var wxUser model.SysWxUser
	_ = c.ShouldBindQuery(&wxUser)
	err, rewxUser := service.GetSysWxUser(wxUser.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"rewxUser": rewxUser}, c)
	}
}

// @Tags SysWxUser
// @Summary 分页获取SysWxUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取SysWxUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wxUser/getSysWxUserList [get]
func GetSysWxUserList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetSysWxUserInfoList(pageInfo)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(resp.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, c)
	}
}
