package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitTitTopicRouter(Router *gin.RouterGroup) {
	TitTopicRouter := Router.Group("topic").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		TitTopicRouter.POST("createTitTopic", v1.CreateTitTopic)   // 新建TitTopic
		TitTopicRouter.DELETE("deleteTitTopic", v1.DeleteTitTopic) // 删除TitTopic
		TitTopicRouter.PUT("updateTitTopic", v1.UpdateTitTopic)    // 更新TitTopic
		TitTopicRouter.GET("findTitTopic", v1.FindTitTopic)        // 根据ID获取TitTopic
		TitTopicRouter.GET("getTitTopicList", v1.GetTitTopicList)  // 获取TitTopic列表
	}
}

func InitTitTopicRouter2(Router *gin.RouterGroup) {
	TitTopicRouter := Router.Group("topicOption").Use(middleware.JWTTitAuth())
	{
		TitTopicRouter.GET("/:businessType", v1.QueryByBusinessType) // 根据 businessType 获取 题目及各个题目下的选项
	}
}
