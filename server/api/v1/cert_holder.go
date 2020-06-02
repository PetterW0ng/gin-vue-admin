package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
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
