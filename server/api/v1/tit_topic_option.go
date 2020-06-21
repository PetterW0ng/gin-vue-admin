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

// @Tags TitTopicOption
// @Summary 创建TitTopicOption
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TitTopicOption true "创建TitTopicOption"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /topicOption/createTitTopicOption [post]
func CreateTitTopicOption(c *gin.Context) {
	var topicOption model.TitTopicOption
	_ = c.ShouldBindJSON(&topicOption)
	err := service.CreateTitTopicOption(topicOption)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags TitTopicOption
// @Summary 删除TitTopicOption
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TitTopicOption true "删除TitTopicOption"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /topicOption/deleteTitTopicOption [delete]
func DeleteTitTopicOption(c *gin.Context) {
	var topicOption model.TitTopicOption
	_ = c.ShouldBindJSON(&topicOption)
	err := service.DeleteTitTopicOption(topicOption)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags TitTopicOption
// @Summary 更新TitTopicOption
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TitTopicOption true "更新TitTopicOption"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /topicOption/updateTitTopicOption [put]
func UpdateTitTopicOption(c *gin.Context) {
	var topicOption model.TitTopicOption
	_ = c.ShouldBindJSON(&topicOption)
	err := service.UpdateTitTopicOption(&topicOption)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags TitTopicOption
// @Summary 用id查询TitTopicOption
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TitTopicOption true "用id查询TitTopicOption"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /topicOption/findTitTopicOption [get]
func FindTitTopicOption(c *gin.Context) {
	var topicOption model.TitTopicOption
	_ = c.ShouldBindQuery(&topicOption)
	err, retopicOption := service.GetTitTopicOption(topicOption.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"retopicOption": retopicOption}, c)
	}
}

// @Tags TitTopicOption
// @Summary 分页获取TitTopicOption列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取TitTopicOption列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /topicOption/getTitTopicOptionList [get]
func GetTitTopicOptionList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetTitTopicOptionInfoList(pageInfo)
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
