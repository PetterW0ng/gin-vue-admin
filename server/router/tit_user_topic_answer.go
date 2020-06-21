package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitTitUserTopicAnswerRouter(Router *gin.RouterGroup) {
	TitUserTopicAnswerRouter := Router.Group("userTopicAnswer").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		TitUserTopicAnswerRouter.POST("createTitUserTopicAnswer", v1.CreateTitUserTopicAnswer)   // 新建TitUserTopicAnswer
		TitUserTopicAnswerRouter.DELETE("deleteTitUserTopicAnswer", v1.DeleteTitUserTopicAnswer) // 删除TitUserTopicAnswer
		TitUserTopicAnswerRouter.PUT("updateTitUserTopicAnswer", v1.UpdateTitUserTopicAnswer)    // 更新TitUserTopicAnswer
		TitUserTopicAnswerRouter.GET("findTitUserTopicAnswer", v1.FindTitUserTopicAnswer)        // 根据ID获取TitUserTopicAnswer
		TitUserTopicAnswerRouter.GET("getTitUserTopicAnswerList", v1.GetTitUserTopicAnswerList)  // 获取TitUserTopicAnswer列表
	}
}

func InitTitUserTopicAnswerRouter2(Router *gin.RouterGroup) {
	TitUserTopicAnswerRouter := Router.Group("userTopicAnswer").Use(middleware.JWTTitAuth())
	{
		TitUserTopicAnswerRouter.POST("add", v1.AddTitUserTopicAnswer)    // 新增用户的答题记录
		TitUserTopicAnswerRouter.GET("", v1.QueryTitUserTopicAnswer)      // 查看用户 businessType 下的记录
		TitUserTopicAnswerRouter.GET("analysis", v1.QueryTitUserAnalysis) // 根据积分题 对用户进行分析
	}
}
