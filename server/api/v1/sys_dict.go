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

const (
	GENDER, SCHOOLE, STUDY_MAJOR, EDUCATION, SCHOO_SYSTEM, WORKING_TATE, SERVICE_YPE, INCOME, BENEFITS, CHILD_TYPE, CHILD_AGE, TRAINING_NUMBER, TRAINING_FEE model.DictType = "gender", "schoole", "study_major", "education", "schoo_system", "working_tate", "service_ype", "income", "benefits", "child_type", "", "training_number", "training_fee"
)

// @Tags SysDict
// @Summary 创建SysDict
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysDict true "创建SysDict"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dict/createSysDict [post]
func CreateSysDict(c *gin.Context) {
	var dict model.SysDict
	_ = c.ShouldBindJSON(&dict)
	err := service.CreateSysDict(dict)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags SysDict
// @Summary 删除SysDict
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysDict true "删除SysDict"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dict/deleteSysDict [delete]
func DeleteSysDict(c *gin.Context) {
	var dict model.SysDict
	_ = c.ShouldBindJSON(&dict)
	err := service.DeleteSysDict(dict)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SysDict
// @Summary 更新SysDict
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysDict true "更新SysDict"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dict/updateSysDict [put]
func UpdateSysDict(c *gin.Context) {
	var dict model.SysDict
	_ = c.ShouldBindJSON(&dict)
	err := service.UpdateSysDict(&dict)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags SysDict
// @Summary 用id查询SysDict
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysDict true "用id查询SysDict"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dict/findSysDict [get]
func FindSysDict(c *gin.Context) {
	var dict model.SysDict
	_ = c.ShouldBindQuery(&dict)
	err, redict := service.GetSysDict(dict.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"redict": redict}, c)
	}
}

// @Tags SysDict
// @Summary 分页获取SysDict列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取SysDict列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dict/getSysDictList [get]
func GetSysDictList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetSysDictInfoList(pageInfo)
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

func FindByCode(c *gin.Context) {

	code := c.Param("code")
	if len(code) == 0 {
		response.FailWithMessage("参数错误", c)
		return
	}
	err, SysDictList := service.FindByCode(code)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{code: SysDictList}, c)
	}
}

func QueryTitUsers(c *gin.Context) {
	// 查询所有的 字典数据
	var dictGroup = make(map[string][]model.SysDict, 16)
	if err, dictList := service.GetAllDict(); err == nil {
		for _, value := range dictList {
			v, ok := dictGroup[value.Code]
			if ok {
				v = append(v, value)
			} else {
				v = []model.SysDict{}
				v = append(v, value)
			}
			dictGroup[value.Code] = v
		}
		response.OkWithData(gin.H{"dictGroup": dictGroup}, c)
	} else {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	}

}
