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

// @Tags SysCusScore
// @Summary 创建SysCusScore
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysCusScore true "创建SysCusScore"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sCusScore/createSysCusScore [post]
func CreateSysCusScore(c *gin.Context) {
	var sCusScore model.SysCusScore
	_ = c.ShouldBindJSON(&sCusScore)
	err := service.CreateSysCusScore(sCusScore)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags SysCusScore
// @Summary 删除SysCusScore
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysCusScore true "删除SysCusScore"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sCusScore/deleteSysCusScore [delete]
func DeleteSysCusScore(c *gin.Context) {
	var sCusScore model.SysCusScore
	_ = c.ShouldBindJSON(&sCusScore)
	err := service.DeleteSysCusScore(sCusScore)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SysCusScore
// @Summary 更新SysCusScore
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysCusScore true "更新SysCusScore"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sCusScore/updateSysCusScore [put]
func UpdateSysCusScore(c *gin.Context) {
	var sCusScore model.SysCusScore
	_ = c.ShouldBindJSON(&sCusScore)
	err := service.UpdateSysCusScore(&sCusScore)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags SysCusScore
// @Summary 用id查询SysCusScore
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysCusScore true "用id查询SysCusScore"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sCusScore/findSysCusScore [get]
func FindSysCusScore(c *gin.Context) {
	var sCusScore model.SysCusScore
	_ = c.ShouldBindQuery(&sCusScore)
	err, resCusScore := service.GetSysCusScore(sCusScore.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"resCusScore": resCusScore}, c)
	}
}

// @Tags SysCusScore
// @Summary 分页获取SysCusScore列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取SysCusScore列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sCusScore/getSysCusScoreList [get]
func GetSysCusScoreList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetSysCusScoreInfoList(pageInfo)
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
