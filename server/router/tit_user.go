package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitTitUserRouter(Router *gin.RouterGroup) {
	TitUserRouter := Router.Group("user").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		TitUserRouter.POST("createTitUser", v1.CreateTitUser)   // 新建TitUser
		TitUserRouter.DELETE("deleteTitUser", v1.DeleteTitUser) // 删除TitUser
		TitUserRouter.PUT("updateTitUser", v1.UpdateTitUser)    // 更新TitUser
		TitUserRouter.GET("findTitUser", v1.FindTitUser)        // 根据ID获取TitUser
		TitUserRouter.GET("getTitUserList", v1.GetTitUserList)  // 获取TitUser列表
	}
}

func InitTitUserBaseRouter(Router *gin.RouterGroup) {
	TitUserRouter := Router.Group("user")
	{
		TitUserRouter.POST("register", v1.TitUserRegister)             // tit 用户注册
		TitUserRouter.POST("login", v1.TitUserLogin)                   // tit 用户登录
		TitUserRouter.POST("verificationCode", v1.GetVerificationCode) // 获取验证码
	}
}
