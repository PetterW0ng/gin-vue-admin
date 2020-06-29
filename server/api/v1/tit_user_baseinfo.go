package v1

import (
	"encoding/json"
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
	"School":           {utils.NotEmpty()},
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

	trainingNumber, _ := json.Marshal(ubt.TrainingNumber)
	benefits, _ := json.Marshal(ubt.Benefits)
	childType, _ := json.Marshal(ubt.ChildType)
	userBaseinfo := model.TitUserBaseinfo{TitUserId: currentUser.ID, School: ubt.School, MajorsStudied: ubt.MajorsStudied, HighestEducation: ubt.HighestEducation,
		SchoolSystem: ubt.SchoolSystem, WorkingState: ubt.WorkingState, TrainingNumber: string(trainingNumber[1 : len(trainingNumber)-1]), TrainingFee: ubt.TrainingFee, Company: ubt.Company,
		Area: ubt.Area, JobTitle: ubt.JobTitle, ServiceType: ubt.ServiceType, Income: ubt.Income, Benefits: string(benefits[1 : len(benefits)-1]), ChildType: string(childType[1 : len(childType)-1]), ChildAge: ubt.ChildAge}

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

func GetTitUserBaseinfo(c *gin.Context) {
	claims, _ := c.Get("claims")
	currentUser := claims.(*request.TitUserClaims)
	if err, tu := service.GetTitUser(currentUser.ID); err == nil {
		if err, t := service.GetTitUserBaseinfo(uint(tu.TitUserBaseinfoId)); err == nil {
			titUserBase := resp.TitUserBase{ID: t.ID, TitUserId: t.TitUserId, School: t.School, MajorsStudied: t.MajorsStudied, HighestEducation: t.HighestEducation, SchoolSystem: t.SchoolSystem, WorkingState: t.WorkingState, Company: t.Company,
				Area: t.Area, JobTitle: t.JobTitle, ServiceType: t.ServiceType, Income: t.Income, Benefits: t.Benefits, ChildAge: t.ChildAge, ChildType: t.ChildType, TrainingNumber: t.TrainingNumber, TrainingFee: t.TrainingFee}
			trainingInfos := service.QueryTrainingInfoByBaseId(tu.TitUserBaseinfoId)
			var trainingInfosReturn []resp.TitTrainingInfo
			for _, item := range trainingInfos {
				trainingInfosReturn = append(trainingInfosReturn, resp.TitTrainingInfo{
					TitUserBaseinfoId: item.TitUserBaseinfoId,
					PaymentWay:        item.PaymentWay,
					TrainingCourse:    item.TrainingCourse,
					BeginTime:         item.BeginTime,
					EndTime:           item.EndTime,
				})
			}
			titUserBase.TrainingInfo = trainingInfosReturn
			response.OkWithData(gin.H{"userBaseinfo": titUserBase}, c)
			return
		}
	}
	response.FailWithMessage("获取基本信息失败了", c)
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
			for inx, v := range benefits {
				i, _ := strconv.Atoi(v)
				if inx == len(benefits)-1 {
					benefitsStr += dictGroup["benefits"][i]
				} else {
					benefitsStr += dictGroup["benefits"][i] + "、"
				}
			}
			for inx, v := range trainingTimes {
				i, _ := strconv.Atoi(v)
				if inx == len(benefits)-1 {
					trainingTimesStr += dictGroup["training_number"][i]
				} else {
					trainingTimesStr += dictGroup["training_number"][i] + "、"
				}

			}
			for inx, v := range childType {
				i, _ := strconv.Atoi(v)
				if inx == len(benefits)-1 {
					childTypeStr += dictGroup["child_type"][i]
				} else {
					childTypeStr += dictGroup["child_type"][i] + "、"
				}
			}

			tubtR := resp.TitUserBaseinfoTraining{
				School:           dictGroup["school"][tub.School], // int 转 string
				MajorsStudied:    dictGroup["study_major"][tub.MajorsStudied],
				HighestEducation: dictGroup["highest_education"][tub.HighestEducation],
				SchoolSystem:     dictGroup["school_system"][tub.SchoolSystem],
				WorkingState:     dictGroup["working_state"][tub.WorkingState],
				Company:          tub.Company,
				Area:             area,
				JobTitle:         tub.JobTitle,

				ServiceType: dictGroup["service_type"][tub.ServiceType],
				Income:      dictGroup["income"][tub.Income],
				Benefits:    benefitsStr,
				ChildType:   childTypeStr,
				ChildAge:    dictGroup["child_age"][tub.ChildAge],

				TrainingNumber: trainingTimesStr,
				TrainingFee:    dictGroup["training_fee"][tub.TrainingFee],
			}
			// 根据 TitUserBaseinfoId 查询 training 信息
			trainingInfos := service.QueryTrainingInfoByBaseId(tu.TitUserBaseinfoId)
			var trainingInfosReturn []resp.TrainingInfo
			for _, item := range trainingInfos {
				trainingInfosReturn = append(trainingInfosReturn, resp.TrainingInfo{
					PaymentWay:     dictGroup["payment_way"][item.PaymentWay],
					TrainingCourse: item.TrainingCourse,
					BeginTime:      item.BeginTime,
					EndTime:        item.EndTime,
				})
			}
			tubtR.TrainingInfos = trainingInfosReturn
			response.OkWithData(gin.H{"userBaseinfo": tubtR}, c)
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
	userBaseinfo := model.TitUserBaseinfo{Model: gorm.Model{ID: uint(tu.TitUserBaseinfoId)}, TitUserId: currentUser.ID, School: ubt.School, MajorsStudied: ubt.MajorsStudied, HighestEducation: ubt.HighestEducation,
		SchoolSystem: ubt.SchoolSystem, WorkingState: ubt.WorkingState, TrainingNumber: fmt.Sprint(ubt.TrainingNumber), TrainingFee: ubt.TrainingFee, Company: ubt.Company,
		Area: ubt.Area, JobTitle: ubt.JobTitle, ServiceType: ubt.ServiceType, Income: ubt.Income, Benefits: fmt.Sprint(ubt.Benefits), ChildType: fmt.Sprint(ubt.ChildType), ChildAge: ubt.ChildAge}

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
