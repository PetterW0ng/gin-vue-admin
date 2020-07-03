package router

import (
	v1 "gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitAreaRouter(Router *gin.RouterGroup) {
	SysDictRouter := Router.Group("area").Use(middleware.JWTTitAuth())
	SysDictRouter.GET("/:pid", v1.FindByPid)               // 根据 pid 获取省的市或市的区
	SysDictRouter.POST("allProvince", v1.QueryAllProvince) // 获取所有省份
	SysDictRouter.POST("tree", v1.AreaList)                // 获取所有省份
	SysDictRouter.POST("", v1.Area)                        // 获取所有省份
}
