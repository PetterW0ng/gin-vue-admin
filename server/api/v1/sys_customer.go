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
	err := service.CreateSysCustomer(&customer)
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
	EntryPoint       int    `json:"entryPoint"`
	Source           int    `json:"source"`
	OpenId           string `json:"openId"`
	State            string `json:"state"`
	WxToken          string `json:"wxToken"`
	UserType         int    `json:"userType"`
}

func RegisterCustomer(c *gin.Context) {
	var rr RegisterCustomerStruct
	_ = c.ShouldBindJSON(&rr)
	CustomerVerify := utils.Rules{
		"Telphone":         {utils.NotEmpty()},
		"VerificationCode": {utils.NotEmpty()},
		"CourseType":       {utils.NotEmpty()},
		"OpenId":           {utils.NotEmpty()},
		"UserType":         {utils.NotEmpty()},
	}
	customerVerifyErr := utils.Verify(rr, CustomerVerify)
	if customerVerifyErr != nil {
		response.FailWithMessage(customerVerifyErr.Error(), c)
		return
	}
	// 验证 验证码是否正确
	cmd := global.GVA_REDIS.Get("code:" + rr.Telphone)
	if cmd.Err() == nil && cmd.Val() == rr.VerificationCode {
		customer := model.SysCustomer{Phone: rr.Telphone, CourseType: rr.CourseType, EntryPoint: rr.EntryPoint, Source: rr.Source, IsEvaluate: false, UserType: rr.UserType}
		// 查看是否注册了 sysCustomer
		if err, _ := service.GetSysCustomerByPhone(rr.Telphone); err != nil {
			service.CreateSysCustomer(&customer)
			// 向小鹅通注册用户数据
			go service.RegisterToXiaoet(&customer)
		}
		// 查看是否注册了 tit
		if err, _ := service.FindTitUserByPhone(rr.Telphone); err != nil {
			service.CreateTitUser(model.TitUser{Telphone: rr.Telphone, OpenId: rr.OpenId})
		}
		// 绑定微信
		if err, wxUser := service.GetSysWxUserByOpenId(rr.OpenId); err == nil {
			if wxUser.Phone != rr.Telphone {
				wxUser.Phone = rr.Telphone
				service.UpdateSysWxUser(&wxUser)
			}
		}
		// 自动登录
		_, u := service.FindTitUserByPhone(rr.Telphone)
		GenerateTitUserToke(c, u)
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
	_ = c.ShouldBindJSON(&sysCustomerQuery)
	err, list, total := service.GetSysCustomerInfoList(sysCustomerQuery)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		groupedDict := service.GetGroupedDict()
		var (
			genderMap, sourceMap, entryPoinMap, evaluateMap, userTypeMap = groupedDict[model.GENDER.Value()], groupedDict[model.SOURCE.Value()], groupedDict[model.ENTRY_POINT.Value()], groupedDict[model.IS_EVALUATE.Value()], groupedDict[model.USER_TYPE.Value()]
		)
		// 由 dbModel 转换成
		var respSysCustomerList = make([]resp.SysCustomer, len(list))
		for i, customerInfo := range list {
			var respSysCustomer = resp.SysCustomer{}
			respSysCustomer.SetValueFromDBModel(customerInfo)
			respSysCustomer.Source = sourceMap[customerInfo.Source]
			respSysCustomer.EntryPoint = entryPoinMap[customerInfo.EntryPoint]
			respSysCustomer.Gender = genderMap[customerInfo.Gender]
			respSysCustomer.UserType = userTypeMap[customerInfo.UserType]

			if customerInfo.IsEvaluate {
				respSysCustomer.IsEvaluate = evaluateMap[1]
			} else {
				respSysCustomer.IsEvaluate = evaluateMap[2]
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
