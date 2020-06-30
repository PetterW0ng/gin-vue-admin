package v1

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
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

// ToStore 重定向至 小鹅通商店
func ToStore(c *gin.Context) {
	claims, _ := c.Get("claims")
	currentUser := claims.(*request.TitUserClaims)
	ID, _ := strconv.Atoi(c.Query("relatedId"))
	var url string
	err, related := service.GetTitTopicRelated(ID)
	if err == nil {
		url = related.ObjectLink
		global.GVA_LOG.Info(currentUser.Telphone, "访问了连接:", url)
		//c.Redirect(http.StatusMovedPermanently, url)
		response.OkWithData(gin.H{"url": url}, c)
	} else {
		response.FailWithMessage(fmt.Sprintf("访问商店失败了，%v", err), c)
	}
}

func QueryTitUserAnalysis(c *gin.Context) {
	const PERFESSION_BUSINESS_TYPE, WORK_AGE_TOPIC, SCORED = 2, 13, 1
	claims, _ := c.Get("claims")
	currentUser := claims.(*request.TitUserClaims)

	if err, tu := service.GetTitUser(currentUser.ID); err == nil {
		err1, topicAnswers := service.QueryTitUserTopicAnswer(currentUser.ID, PERFESSION_BUSINESS_TYPE, tu.PerfessionKnowledgeBatchNum)
		err2, topics := service.QueryTopicByBusinessType(PERFESSION_BUSINESS_TYPE)
		var surveyDimensions, gtStandard, ltStandard []string
		var standardScores, userScores []int
		var topicRelatedList []resp.TitTopicRelated
		if err1 == nil && err2 == nil {
			// 找出 用户答题时选择的从业年龄
			var workAge string
			var ix int
			for _, item := range topicAnswers {
				if item.TitTopicId == WORK_AGE_TOPIC {
					workAge = item.TopicOptionIds
					break
				}
			}
			switch workAge {
			case "64":
				ix = 0
			case "65":
				ix = 1
			case "66":
				ix = 2
			default:
				ix = -1
			}
			global.GVA_LOG.Info("用户的工作年限答题时选择的是：", workAge)
			for _, item := range topics {
				if item.IsScored != SCORED {
					continue
				}
				surveyDimensions = append(surveyDimensions, item.SurveyLatitude)
				standardScore, _ := strconv.Atoi(strings.Split(item.Score, ",")[ix])
				standardScores = append(standardScores, standardScore)
				for _, ans := range topicAnswers {
					if uint(ans.TitTopicId) == item.ID {
						userScores = append(userScores, ans.Score)
						if ans.Score >= standardScore {
							gtStandard = append(gtStandard, item.SurveyLatitude)
						} else {
							ltStandard = append(ltStandard, item.SurveyLatitude)
							// 设置推荐的内容 课程、书籍、工具
							for _, relatedItem := range item.TitTopicRelatedList {
								topicRelatedList = append(topicRelatedList, resp.TitTopicRelated{ID: relatedItem.ID, TitTopicId: relatedItem.TitTopicId, RecommendType: relatedItem.RecommendType, RecommendObject: relatedItem.RecommendObject, ObjectLink: "", Remark: relatedItem.Remark})
							}
						}
					}
				}
			}
			response.OkWithData(gin.H{"surveyDimensions": surveyDimensions, "standardScores": standardScores, "userScores": userScores, "gtStandard": gtStandard, "ltStandard": ltStandard, "topicRelatedList": topicRelatedList}, c)
		} else {
			response.FailWithMessage(fmt.Sprintf("查询用户得分情况失败了，%v", err), c)
		}
	}
}
