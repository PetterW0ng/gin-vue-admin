package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
)

func QueryAllProvince(c *gin.Context) {
	err, provinces := service.QueryAllProvince()
	if err == nil {
		response.OkWithData(gin.H{"provinces": provinces}, c)
	} else {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	}
}

func FindByPid(c *gin.Context) {
	pid := c.Param("pid")
	err, children := service.FindCityOrArea(pid)
	if err == nil {
		response.OkWithData(gin.H{"children": children}, c)
	} else {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	}
}
