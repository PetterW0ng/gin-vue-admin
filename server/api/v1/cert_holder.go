package v1

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	"qiniupkg.com/x/log.v7"
	"regexp"
	"strconv"
)

func GetCertHolderList(c *gin.Context) {
	var paQuery request.PaginatedQuery
	_ = c.ShouldBindJSON(&paQuery)
	err, list, total := service.GetCertHolderList(paQuery)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(resp.PageResult{
			List:     list,
			Total:    total,
			Page:     paQuery.Page,
			PageSize: paQuery.PageSize,
		}, c)
	}
}

func GetCertByPhone(c *gin.Context) {
	phone := c.Param("phone")
	reg := regexp.MustCompile("^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$")
	if !reg.MatchString(phone) {
		response.FailWithMessage("手机号不正确", c)
		return
	}
	if err, list := service.GetCertByPhone(phone); err == nil {
		for i, item := range list {
			item.FileName = global.GVA_CONFIG.System.CertURL + item.FileName + ".jpg"
			list[i] = item
		}
		response.OkWithData(list, c)
	} else {
		log.Error("证书查找失败", err)
		response.FailWithMessage(fmt.Sprintf("证书查找失败，%v", err), c)
	}
}

func ExportCertInfo(c *gin.Context) {
	xlsx := excelize.NewFile()
	paQuery := request.PaginatedQuery{}
	paQuery.Page = 1
	paQuery.PageSize = 1000000
	if err, list, _ := service.GetCertHolderList(paQuery); err == nil {
		for i, d := range list {
			if i == 0 {
				xlsx.SetSheetRow("Sheet1", "A1", &[]interface{}{
					"姓名", "证书编号", "手机", "发证时间",
					"身份证号", "创建日期", "证书名称", "发证单位",
				})
			}
			xlsx.SetSheetRow("Sheet1", "A"+strconv.Itoa(i+2), &[]interface{}{
				d.UserName, d.SerialNum, d.Phone, d.IssueTime,
				d.IdCard, d.CreateTime, d.CertificateName, d.IssuingUnit,
			})
		}
	} else {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
		return
	}
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename=证书列表.xlsx")
	c.Header("Content-Transfer-Encoding", "binary")

	_ = xlsx.Write(c.Writer)
}
