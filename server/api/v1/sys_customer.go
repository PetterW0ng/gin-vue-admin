package v1

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"gin-vue-admin/utils"
	"github.com/gin-gonic/gin"
)

// @Tags SysCustomer
// @Summary 创建SysCustomer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysCustomer true "创建SysCustomer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/createSysCustomer [post]
func CreateSysCustomer(c *gin.Context) {
	var customer model.SysCustomer
	_ = c.ShouldBindJSON(&customer)
	err := service.CreateSysCustomer(customer)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags SysCustomer
// @Summary 删除SysCustomer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysCustomer true "删除SysCustomer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /customer/deleteSysCustomer [delete]
func DeleteSysCustomer(c *gin.Context) {
	var customer model.SysCustomer
	_ = c.ShouldBindJSON(&customer)
	err := service.DeleteSysCustomer(customer)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SysCustomer
// @Summary 更新SysCustomer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysCustomer true "更新SysCustomer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /customer/updateSysCustomer [put]
func UpdateSysCustomer(c *gin.Context) {
	var customer model.SysCustomer
	_ = c.ShouldBindJSON(&customer)
	err := service.UpdateSysCustomer(&customer)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags SysCustomer
// @Summary 用id查询SysCustomer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysCustomer true "用id查询SysCustomer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /customer/findSysCustomer [get]
func FindSysCustomer(c *gin.Context) {
	var customer model.SysCustomer
	_ = c.ShouldBindQuery(&customer)
	err, recustomer := service.GetSysCustomer(customer.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"recustomer": recustomer}, c)
	}
}

func GetUserXetOrder(c *gin.Context) {
	eid := c.Param("eid")
	if err, list := service.QueryUserOrder(eid); err != nil {
		response.FailWithMessage("获取订单列表失败了，联系管理员", c)
	} else {
		response.OkWithData(list, c)
	}
}

type RegisterCustomerStruct struct {
	Telphone         string `json:"telphone"`
	VerificationCode string `json:"verificationCode"`
	CourseType       int    `json:"courseType"`
}

func RegisterCustomer(c *gin.Context) {
	var rr RegisterCustomerStruct
	_ = c.ShouldBindJSON(&rr)
	CustomerVerify := utils.Rules{
		"Telphone":         {utils.NotEmpty()},
		"VerificationCode": {utils.NotEmpty()},
		"CourseType":       {utils.NotEmpty()},
	}
	customerVerifyErr := utils.Verify(rr, CustomerVerify)
	if customerVerifyErr != nil {
		response.FailWithMessage(customerVerifyErr.Error(), c)
		return
	}
	// 验证 验证码是否正确
	cmd := global.GVA_REDIS.Get(rr.Telphone)
	if cmd.Err() == nil && cmd.Val() == rr.VerificationCode {
		customer := model.SysCustomer{Phone: rr.Telphone, CourseType: rr.CourseType, IsEvaluate: false}
		if err := service.CreateSysCustomer(customer); err != nil {
			global.GVA_LOG.Error("用户注册失败了", err)
			response.FailWithMessage("手机号已存在", c)
		} else {
			response.OkWithMessage("注册成功", c)
			/*if err, u := service.FindTitUserByPhone(rr.Telphone); err == nil {
				generateTitUserToke(c, u)
			} else {
				response.Fail(c)
			}*/
		}
	} else {
		response.FailWithMessage("验证码不正确", c)
	}
}

// @Tags SysCustomer
// @Summary 分页获取SysCustomer列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取SysCustomer列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/getSysCustomerList [get]
func GetSysCustomerList(c *gin.Context) {
	var sysCustomerQuery request.PaginatedSysCustomer
	_ = c.ShouldBindQuery(&sysCustomerQuery)
	err, list, total := service.GetSysCustomerInfoList(sysCustomerQuery)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		groupedDict := service.GetGroupedDict()
		var (
			genderMap, sourceMap, entryPoinMap, evaluateMap = groupedDict[model.GENDER.Value()], groupedDict[model.SOURCE.Value()], groupedDict[model.ENTRY_POINT.Value()], groupedDict[model.IS_EVALUATE.Value()]
		)
		// 由 dbModel 转换成
		var respSysCustomerList = make([]resp.SysCustomer, len(list))
		for i, customerInfo := range list {
			var respSysCustomer = resp.SysCustomer{}
			respSysCustomer.SetValueFromDBModel(customerInfo)
			respSysCustomer.Source = sourceMap[customerInfo.Source]
			respSysCustomer.EntryPoint = evaluateMap[customerInfo.EntryPoint]
			respSysCustomer.Gender = genderMap[customerInfo.Gender]
			if customerInfo.IsEvaluate {
				respSysCustomer.EntryPoint = entryPoinMap[1]
			} else {
				respSysCustomer.EntryPoint = entryPoinMap[2]
			}
			var respTags = make([]resp.SysCusTag, len(customerInfo.SysCusTags))
			for j, tag := range customerInfo.SysCusTags {
				respTags[j] = resp.SysCusTag{TagCode: tag.TagCode, TagValue: tag.TagValue, TagName: groupedDict[tag.TagCode][tag.TagValue]}
			}
			respSysCustomer.Tags = respTags
			respSysCustomerList[i] = respSysCustomer
		}

		response.OkWithData(resp.PageResult{
			List:     respSysCustomerList,
			Total:    total,
			Page:     sysCustomerQuery.Page,
			PageSize: sysCustomerQuery.PageSize,
		}, c)
	}
}
