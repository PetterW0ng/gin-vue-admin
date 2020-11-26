package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitSysCustomerRouter(Router *gin.RouterGroup) {
	SysCustomerRouter := Router.Group("customer").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		SysCustomerRouter.POST("createSysCustomer", v1.CreateSysCustomer)   // 新建SysCustomer
		SysCustomerRouter.DELETE("deleteSysCustomer", v1.DeleteSysCustomer) // 删除SysCustomer
		SysCustomerRouter.PUT("updateSysCustomer", v1.UpdateSysCustomer)    // 更新SysCustomer
		SysCustomerRouter.GET("findSysCustomer", v1.FindSysCustomer)        // 根据ID获取SysCustomer
		SysCustomerRouter.POST("getSysCustomerList", v1.GetSysCustomerList) // 获取SysCustomer列表
		SysCustomerRouter.GET("order/:eid", v1.GetUserXetOrder)             // 获取用户订单
	}
}
