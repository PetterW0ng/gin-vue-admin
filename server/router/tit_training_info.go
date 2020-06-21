package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitTitTrainingInfoRouter(Router *gin.RouterGroup) {
	TitTrainingInfoRouter := Router.Group("trainingInfo").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		TitTrainingInfoRouter.POST("createTitTrainingInfo", v1.CreateTitTrainingInfo)   // 新建TitTrainingInfo
		TitTrainingInfoRouter.DELETE("deleteTitTrainingInfo", v1.DeleteTitTrainingInfo) // 删除TitTrainingInfo
		TitTrainingInfoRouter.PUT("updateTitTrainingInfo", v1.UpdateTitTrainingInfo)    // 更新TitTrainingInfo
		TitTrainingInfoRouter.GET("findTitTrainingInfo", v1.FindTitTrainingInfo)        // 根据ID获取TitTrainingInfo
		TitTrainingInfoRouter.GET("getTitTrainingInfoList", v1.GetTitTrainingInfoList)  // 获取TitTrainingInfo列表
	}
}
