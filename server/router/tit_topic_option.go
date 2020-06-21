package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitTitTopicOptionRouter(Router *gin.RouterGroup) {
	TitTopicOptionRouter := Router.Group("topicOption").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		TitTopicOptionRouter.POST("createTitTopicOption", v1.CreateTitTopicOption)   // 新建TitTopicOption
		TitTopicOptionRouter.DELETE("deleteTitTopicOption", v1.DeleteTitTopicOption) // 删除TitTopicOption
		TitTopicOptionRouter.PUT("updateTitTopicOption", v1.UpdateTitTopicOption)    // 更新TitTopicOption
		TitTopicOptionRouter.GET("findTitTopicOption", v1.FindTitTopicOption)        // 根据ID获取TitTopicOption
		TitTopicOptionRouter.GET("getTitTopicOptionList", v1.GetTitTopicOptionList)  // 获取TitTopicOption列表
	}
}
