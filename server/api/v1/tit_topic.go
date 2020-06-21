package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

// @Tags TitTopic
// @Summary 创建TitTopic
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TitTopic true "创建TitTopic"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /topic/createTitTopic [post]
func CreateTitTopic(c *gin.Context) {
	var topic model.TitTopic
	_ = c.ShouldBindJSON(&topic)
	err := service.CreateTitTopic(topic)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags TitTopic
// @Summary 删除TitTopic
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TitTopic true "删除TitTopic"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /topic/deleteTitTopic [delete]
func DeleteTitTopic(c *gin.Context) {
	var topic model.TitTopic
	_ = c.ShouldBindJSON(&topic)
	err := service.DeleteTitTopic(topic)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags TitTopic
// @Summary 更新TitTopic
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TitTopic true "更新TitTopic"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /topic/updateTitTopic [put]
func UpdateTitTopic(c *gin.Context) {
	var topic model.TitTopic
	_ = c.ShouldBindJSON(&topic)
	err := service.UpdateTitTopic(&topic)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags TitTopic
// @Summary 用id查询TitTopic
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TitTopic true "用id查询TitTopic"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /topic/findTitTopic [get]
func FindTitTopic(c *gin.Context) {
	var topic model.TitTopic
	_ = c.ShouldBindQuery(&topic)
	err, retopic := service.GetTitTopic(topic.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"retopic": retopic}, c)
	}
}

// @Tags TitTopic
// @Summary 分页获取TitTopic列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取TitTopic列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /topic/getTitTopicList [get]
func GetTitTopicList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetTitTopicInfoList(pageInfo)
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

func QueryByBusinessType(c *gin.Context) {
	businessType := c.Param("businessType")
	bt, _ := strconv.Atoi(businessType)
	err, topicOptions := service.QueryByBusinessType(bt)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
		return
	} else {
		response.OkWithData(gin.H{"topicOptions": topicOptions}, c)
	}
}
