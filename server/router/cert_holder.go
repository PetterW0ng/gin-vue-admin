package router

import (
	v1 "gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitCertHolderRouter(Router *gin.RouterGroup) {
	CertHolderRouter := Router.Group("cert").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	CertHolderRouter.POST("getCertHolderList", v1.GetCertHolderList) // 分页获取获得证书用户的列表
}
