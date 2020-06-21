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

// @Tags TitTrainingInfo
// @Summary 创建TitTrainingInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TitTrainingInfo true "创建TitTrainingInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /trainingInfo/createTitTrainingInfo [post]
func CreateTitTrainingInfo(c *gin.Context) {
	var trainingInfo model.TitTrainingInfo
	_ = c.ShouldBindJSON(&trainingInfo)
	err := service.CreateTitTrainingInfo(trainingInfo)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags TitTrainingInfo
// @Summary 删除TitTrainingInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TitTrainingInfo true "删除TitTrainingInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /trainingInfo/deleteTitTrainingInfo [delete]
func DeleteTitTrainingInfo(c *gin.Context) {
	var trainingInfo model.TitTrainingInfo
	_ = c.ShouldBindJSON(&trainingInfo)
	err := service.DeleteTitTrainingInfo(trainingInfo)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags TitTrainingInfo
// @Summary 更新TitTrainingInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TitTrainingInfo true "更新TitTrainingInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /trainingInfo/updateTitTrainingInfo [put]
func UpdateTitTrainingInfo(c *gin.Context) {
	var trainingInfo model.TitTrainingInfo
	_ = c.ShouldBindJSON(&trainingInfo)
	err := service.UpdateTitTrainingInfo(&trainingInfo)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags TitTrainingInfo
// @Summary 用id查询TitTrainingInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TitTrainingInfo true "用id查询TitTrainingInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /trainingInfo/findTitTrainingInfo [get]
func FindTitTrainingInfo(c *gin.Context) {
	var trainingInfo model.TitTrainingInfo
	_ = c.ShouldBindQuery(&trainingInfo)
	err, retrainingInfo := service.GetTitTrainingInfo(trainingInfo.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"retrainingInfo": retrainingInfo}, c)
	}
}

// @Tags TitTrainingInfo
// @Summary 分页获取TitTrainingInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取TitTrainingInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /trainingInfo/getTitTrainingInfoList [get]
func GetTitTrainingInfoList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetTitTrainingInfoInfoList(pageInfo)
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
