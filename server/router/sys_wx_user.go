package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitSysWxUserRouter(Router *gin.RouterGroup) {
	SysWxUserRouter := Router.Group("wxUser").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		SysWxUserRouter.POST("createSysWxUser", v1.CreateSysWxUser)   // 新建SysWxUser
		SysWxUserRouter.DELETE("deleteSysWxUser", v1.DeleteSysWxUser) // 删除SysWxUser
		SysWxUserRouter.PUT("updateSysWxUser", v1.UpdateSysWxUser)    // 更新SysWxUser
		SysWxUserRouter.GET("findSysWxUser", v1.FindSysWxUser)        // 根据ID获取SysWxUser
		SysWxUserRouter.GET("getSysWxUserList", v1.GetSysWxUserList)  // 获取SysWxUser列表
	}
}
