package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitSysCusScoreRouter(Router *gin.RouterGroup) {
	SysCusScoreRouter := Router.Group("sCusScore").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		SysCusScoreRouter.POST("createSysCusScore", v1.CreateSysCusScore)   // 新建SysCusScore
		SysCusScoreRouter.DELETE("deleteSysCusScore", v1.DeleteSysCusScore) // 删除SysCusScore
		SysCusScoreRouter.PUT("updateSysCusScore", v1.UpdateSysCusScore)    // 更新SysCusScore
		SysCusScoreRouter.GET("findSysCusScore", v1.FindSysCusScore)        // 根据ID获取SysCusScore
		SysCusScoreRouter.GET("getSysCusScoreList", v1.GetSysCusScoreList)  // 获取SysCusScore列表
	}
}
