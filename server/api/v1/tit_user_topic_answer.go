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

// @Tags TitUserTopicAnswer
// @Summary 创建TitUserTopicAnswer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TitUserTopicAnswer true "创建TitUserTopicAnswer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userTopicAnswer/createTitUserTopicAnswer [post]
func CreateTitUserTopicAnswer(c *gin.Context) {
	var userTopicAnswer model.TitUserTopicAnswer
	_ = c.ShouldBindJSON(&userTopicAnswer)
	err := service.CreateTitUserTopicAnswer(userTopicAnswer)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags TitUserTopicAnswer
// @Summary 删除TitUserTopicAnswer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TitUserTopicAnswer true "删除TitUserTopicAnswer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userTopicAnswer/deleteTitUserTopicAnswer [delete]
func DeleteTitUserTopicAnswer(c *gin.Context) {
	var userTopicAnswer model.TitUserTopicAnswer
	_ = c.ShouldBindJSON(&userTopicAnswer)
	err := service.DeleteTitUserTopicAnswer(userTopicAnswer)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags TitUserTopicAnswer
// @Summary 更新TitUserTopicAnswer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TitUserTopicAnswer true "更新TitUserTopicAnswer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userTopicAnswer/updateTitUserTopicAnswer [put]
func UpdateTitUserTopicAnswer(c *gin.Context) {
	var userTopicAnswer model.TitUserTopicAnswer
	_ = c.ShouldBindJSON(&userTopicAnswer)
	err := service.UpdateTitUserTopicAnswer(&userTopicAnswer)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags TitUserTopicAnswer
// @Summary 用id查询TitUserTopicAnswer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TitUserTopicAnswer true "用id查询TitUserTopicAnswer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userTopicAnswer/findTitUserTopicAnswer [get]
func FindTitUserTopicAnswer(c *gin.Context) {
	var userTopicAnswer model.TitUserTopicAnswer
	_ = c.ShouldBindQuery(&userTopicAnswer)
	err, reuserTopicAnswer := service.GetTitUserTopicAnswer(userTopicAnswer.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"reuserTopicAnswer": reuserTopicAnswer}, c)
	}
}

// @Tags TitUserTopicAnswer
// @Summary 分页获取TitUserTopicAnswer列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取TitUserTopicAnswer列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userTopicAnswer/getTitUserTopicAnswerList [get]
func GetTitUserTopicAnswerList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetTitUserTopicAnswerInfoList(pageInfo)
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

func AddTitUserTopicAnswer(c *gin.Context) {
	// 获取当前登录用户的 userId
	claims, _ := c.Get("claims")
	currentUser := claims.(*request.TitUserClaims)
	var userTopicAnswer request.UserTopicAnswer
	_ = c.BindJSON(&userTopicAnswer)
	err := service.SaveTitUserAnswer(userTopicAnswer, currentUser.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
		return
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func QueryTitUserTopicAnswer(c *gin.Context) {
	batchNum, _ := strconv.Atoi(c.Query("batchNum"))
	businessType, _ := strconv.Atoi(c.Query("businessType"))
	if batchNum < 1 || businessType < 1 {
		response.FailWithMessage("参数错误", c)
		return
	}
	// 获取当前登录用户的 userId
	claims, _ := c.Get("claims")
	currentUser := claims.(*request.TitUserClaims)
	err, topicAnswers := service.QueryTitUserTopicAnswer(currentUser.ID, businessType, batchNum)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
		return
	} else {
		response.OkWithData(gin.H{"topicAnswers": topicAnswers}, c)
	}
}

func QueryTitUserAnalysis(c *gin.Context) {
	var userId int
	err, surveyDimensions, standardScores, userScores, courseRecommends, bookRecommends := service.QueryTitUserAnalysis(userId)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"surveyDimensions": surveyDimensions, "standardScores": standardScores, "userScores": userScores, "courseRecommends": courseRecommends, "bookRecommends": bookRecommends}, c)
	}
}
