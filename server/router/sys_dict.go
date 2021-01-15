package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitSysDictRouter(Router *gin.RouterGroup) {
	SysDictRouter := Router.Group("dict").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		SysDictRouter.POST("createSysDict", v1.CreateSysDict)   // 新建SysDict
		SysDictRouter.DELETE("deleteSysDict", v1.DeleteSysDict) // 删除SysDict
		SysDictRouter.PUT("updateSysDict", v1.UpdateSysDict)    // 更新SysDict
		SysDictRouter.GET("findSysDict", v1.FindSysDict)        // 根据ID获取SysDict
		SysDictRouter.GET("getSysDictList", v1.GetSysDictList)  // 获取SysDict列表
	}
}

func InitTitSysDictRouter(Router *gin.RouterGroup) {
	SysDictRouter := Router.Group("dict")
	//SysDictRouter := Router.Group("dict").Use(middleware.JWTTitAuth())
	SysDictRouter.GET("/:code", v1.FindByCode)       // 根据code 获取SysDict 列表
	SysDictRouter.POST("baseForm", v1.QueryTitUsers) // 获取用户基本信息表单页的所有选项
}
