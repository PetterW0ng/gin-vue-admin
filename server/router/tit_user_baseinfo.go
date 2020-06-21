package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitTitUserBaseinfoRouter(Router *gin.RouterGroup) {
	TitUserBaseinfoRouter := Router.Group("userBaseinfo").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		TitUserBaseinfoRouter.POST("createTitUserBaseinfo", v1.CreateTitUserBaseinfo)   // 新建TitUserBaseinfo
		TitUserBaseinfoRouter.DELETE("deleteTitUserBaseinfo", v1.DeleteTitUserBaseinfo) // 删除TitUserBaseinfo
		TitUserBaseinfoRouter.PUT("updateTitUserBaseinfo", v1.UpdateTitUserBaseinfo)    // 更新TitUserBaseinfo
		TitUserBaseinfoRouter.GET("findTitUserBaseinfo", v1.FindTitUserBaseinfo)        // 根据ID获取TitUserBaseinfo
		TitUserBaseinfoRouter.GET("getTitUserBaseinfoList", v1.GetTitUserBaseinfoList)  // 获取TitUserBaseinfo列表
	}
}

func InitTitUserBasicinfoRouter(group *gin.RouterGroup) {
	TitUserBaseinfoRouter := group.Group("userBaseinfo").Use(middleware.JWTTitAuth())
	TitUserBaseinfoRouter.POST("create", v1.AddTitUserBaseinfo)   // 新建TitUserBaseinfo
	TitUserBaseinfoRouter.PUT("modify", v1.ModifyTitUserBaseinfo) // 更新TitUserBaseinfo
	TitUserBaseinfoRouter.GET("", v1.QueryTitUserBaseinfo)        // 根据ID获取 TitUserBaseinfo
}
