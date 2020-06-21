package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"gin-vue-admin/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
	"strings"
)

// @Tags TitUserBaseinfo
// @Summary 创建TitUserBaseinfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TitUserBaseinfo true "创建TitUserBaseinfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userBaseinfo/createTitUserBaseinfo [post]
func CreateTitUserBaseinfo(c *gin.Context) {
	var userBaseinfo model.TitUserBaseinfo
	_ = c.ShouldBindJSON(&userBaseinfo)
	err := service.CreateTitUserBaseinfo(&userBaseinfo)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags TitUserBaseinfo
// @Summary 删除TitUserBaseinfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TitUserBaseinfo true "删除TitUserBaseinfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userBaseinfo/deleteTitUserBaseinfo [delete]
func DeleteTitUserBaseinfo(c *gin.Context) {
	var userBaseinfo model.TitUserBaseinfo
	_ = c.ShouldBindJSON(&userBaseinfo)
	err := service.DeleteTitUserBaseinfo(userBaseinfo)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags TitUserBaseinfo
// @Summary 更新TitUserBaseinfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TitUserBaseinfo true "更新TitUserBaseinfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userBaseinfo/updateTitUserBaseinfo [put]
func UpdateTitUserBaseinfo(c *gin.Context) {
	var userBaseinfo model.TitUserBaseinfo
	_ = c.ShouldBindJSON(&userBaseinfo)
	err := service.UpdateTitUserBaseinfo(&userBaseinfo)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags TitUserBaseinfo
// @Summary 用id查询TitUserBaseinfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TitUserBaseinfo true "用id查询TitUserBaseinfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userBaseinfo/findTitUserBaseinfo [get]
func FindTitUserBaseinfo(c *gin.Context) {
	var userBaseinfo model.TitUserBaseinfo
	_ = c.ShouldBindQuery(&userBaseinfo)
	err, reuserBaseinfo := service.GetTitUserBaseinfo(userBaseinfo.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"reuserBaseinfo": reuserBaseinfo}, c)
	}
}

// @Tags TitUserBaseinfo
// @Summary 分页获取TitUserBaseinfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取TitUserBaseinfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userBaseinfo/getTitUserBaseinfoList [get]
func GetTitUserBaseinfoList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetTitUserBaseinfoInfoList(pageInfo)
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

var userBaseinfoRules utils.Rules = utils.Rules{
	"Schoole":          {utils.NotEmpty()},
	"MajorsStudied":    {utils.NotEmpty()},
	"HighestEducation": {utils.NotEmpty()},
	"SchoolSystem":     {utils.NotEmpty()},
	"WorkingState":     {utils.NotEmpty()},
	"TrainingNumber":   {utils.NotEmpty()},
	"TrainingFee":      {utils.NotEmpty()},
}

func AddTitUserBaseinfo(c *gin.Context) {
	claims, _ := c.Get("claims")
	currentUser := claims.(*request.TitUserClaims)
	var ubt request.TitUserBaseinfoTraining
	_ = c.ShouldBindJSON(&ubt)
	// 验证 字段是否齐全
	if verifyErr := utils.Verify(ubt, userBaseinfoRules); verifyErr != nil {
		response.FailWithMessage(verifyErr.Error(), c)
		return
	}
	userBaseinfo := model.TitUserBaseinfo{TitUserId: currentUser.ID, Schoole: ubt.Schoole, MajorsStudied: ubt.MajorsStudied, HighestEducation: ubt.HighestEducation,
		SchoolSystem: ubt.SchoolSystem, WorkingState: ubt.WorkingState, TrainingNumber: ubt.TrainingNumber, TrainingFee: ubt.TrainingFee, Company: ubt.Company,
		Area: ubt.Area, JobTitle: ubt.JobTitle, ServiceType: ubt.ServiceType, Income: ubt.Income, Benefits: ubt.Benefits, ChildType: ubt.ChildType, ChildAge: ubt.ChildAge}

	err := service.CreateTitUserBaseinfo(&userBaseinfo)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
		return
	} else {
		// 添加培训信息
		var titTrainingInfoList []model.TitTrainingInfo
		for _, value := range ubt.TrainingInfos {
			titTrainingInfoList = append(titTrainingInfoList, model.TitTrainingInfo{TitUserBaseinfoId: int(userBaseinfo.ID), TrainingCourse: value.TrainingCourse, BeginTime: value.BeginTime, EndTime: value.EndTime, PaymentWay: value.PaymentWay})
		}
		if err := service.BatchAddTrainingInfo(titTrainingInfoList); err != nil {
			response.FailWithMessage(fmt.Sprintf("添加培训信息失败，%v", err), c)
			return
		}
		// 更新 titUser 下关联的 basicId
		user := model.TitUser{Model: gorm.Model{ID: currentUser.ID}, TitUserBaseinfoId: int(userBaseinfo.ID)}
		if err := service.ModifyTitUser(&user); err != nil {
			response.FailWithMessage(fmt.Sprintf("修改关系时出错了，%v", err), c)
			return
		}
		response.OkWithMessage("创建成功", c)
	}
}

func QueryTitUserBaseinfo(c *gin.Context) {
	claims, _ := c.Get("claims")
	currentUser := claims.(*request.TitUserClaims)
	if err, tu := service.GetTitUser(currentUser.ID); err == nil {
		if err, tub := service.GetTitUserBaseinfo(uint(tu.TitUserBaseinfoId)); err == nil {
			dictGroup := service.GetGroupedDict()
			var area string
			if len(tub.Area) > 0 {
				area = service.GetAreaById(tub.Area).DisplayName
			} else {
				area = ""
			}
			var benefits, childType, trainingTimes = strings.Split(tub.Benefits, ","), strings.Split(tub.ChildType, ","), strings.Split(tub.TrainingNumber, ",")
			var benefitsStr, childTypeStr, trainingTimesStr string
			for _, v := range benefits {
				benefitsStr += dictGroup["benefits"][v] + "、"
			}
			for _, v := range trainingTimes {
				trainingTimesStr += dictGroup["training_number"][v] + "、"
			}
			for _, v := range childType {
				childTypeStr += dictGroup["child_type"][v] + "、"
			}

			tubtR := resp.TitUserBaseinfoTraining{
				Schoole:          dictGroup["schoole"][strconv.Itoa(tub.Schoole)], // int 转 string
				MajorsStudied:    dictGroup["study_major"][strconv.Itoa(tub.MajorsStudied)],
				HighestEducation: dictGroup["highest_education"][strconv.Itoa(tub.HighestEducation)],
				SchoolSystem:     dictGroup["school_system"][strconv.Itoa(tub.SchoolSystem)],
				WorkingState:     dictGroup["working_state"][strconv.Itoa(tub.WorkingState)],
				Company:          tub.Company,
				Area:             area,
				JobTitle:         tub.JobTitle,

				ServiceType: dictGroup["service_type"][strconv.Itoa(tub.ServiceType)],
				Income:      dictGroup["income"][strconv.Itoa(tub.Income)],
				Benefits:    benefitsStr,
				ChildType:   childTypeStr,
				ChildAge:    dictGroup["child_age"][strconv.Itoa(tub.ChildAge)],

				TrainingNumber: trainingTimesStr,
				TrainingFee:    dictGroup["training_fee"][strconv.Itoa(tub.TrainingFee)],
			}
			// 根据 TitUserBaseinfoId 查询 training 信息
			trainingInfos := service.QueryTrainingInfoByBaseId(tu.TitUserBaseinfoId)
			var trainingInfosReturn []resp.TrainingInfo
			for _, item := range trainingInfos {
				trainingInfosReturn = append(trainingInfosReturn, resp.TrainingInfo{
					PaymentWay:     dictGroup["payment_way"][strconv.Itoa(item.PaymentWay)],
					TrainingCourse: item.TrainingCourse,
					BeginTime:      item.BeginTime,
					EndTime:        item.EndTime,
				})
			}
			tubtR.TrainingInfos = trainingInfosReturn
			response.OkWithData(tubtR, c)
			return
		}
	} else {
		response.FailWithMessage(fmt.Sprintf("用户不存在或已删除，%v", err), c)
	}
}

func ModifyTitUserBaseinfo(c *gin.Context) {
	claims, _ := c.Get("claims")
	currentUser := claims.(*request.TitUserClaims)
	var ubt request.TitUserBaseinfoTraining
	_, tu := service.GetTitUser(currentUser.ID)
	_ = c.ShouldBindJSON(&ubt)
	// 验证 字段是否齐全
	if verifyErr := utils.Verify(ubt, userBaseinfoRules); verifyErr != nil {
		response.FailWithMessage(verifyErr.Error(), c)
		return
	}
	userBaseinfo := model.TitUserBaseinfo{Model: gorm.Model{ID: uint(tu.TitUserBaseinfoId)}, TitUserId: currentUser.ID, Schoole: ubt.Schoole, MajorsStudied: ubt.MajorsStudied, HighestEducation: ubt.HighestEducation,
		SchoolSystem: ubt.SchoolSystem, WorkingState: ubt.WorkingState, TrainingNumber: ubt.TrainingNumber, TrainingFee: ubt.TrainingFee, Company: ubt.Company,
		Area: ubt.Area, JobTitle: ubt.JobTitle, ServiceType: ubt.ServiceType, Income: ubt.Income, Benefits: ubt.Benefits, ChildType: ubt.ChildType, ChildAge: ubt.ChildAge}

	err := service.UpdateTitUserBaseinfo(&userBaseinfo)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
		return
	} else {
		// 添加培训信息
		var titTrainingInfoList []model.TitTrainingInfo
		for _, value := range ubt.TrainingInfos {
			titTrainingInfoList = append(titTrainingInfoList, model.TitTrainingInfo{TitUserBaseinfoId: int(userBaseinfo.ID), TrainingCourse: value.TrainingCourse, BeginTime: value.BeginTime, EndTime: value.EndTime, PaymentWay: value.PaymentWay})
		}
		if err := service.BatchModifyTrainingInfo(titTrainingInfoList); err != nil {
			response.FailWithMessage(fmt.Sprintf("添加培训信息失败，%v", err), c)
			return
		}
		response.OkWithMessage("修改成功", c)
	}
}
