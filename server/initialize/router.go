package initialize

import (
	_ "gin-vue-admin/docs"
	"gin-vue-admin/global"
	"gin-vue-admin/middleware"
	"gin-vue-admin/router"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// 初始化总路由

func Routers() *gin.Engine {
	var Router = gin.Default()
	// Router.Use(middleware.LoadTls())  // 打开就能玩https了
	global.GVA_LOG.Debug("use middleware logger")
	// 跨域
	Router.Use(middleware.Cors())
	global.GVA_LOG.Debug("use middleware cors")
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.GVA_LOG.Debug("register swagger handler")
	// 后台管理系统对应的路由
	ApiGroup := Router.Group("")
	{
		router.InitUserRouter(ApiGroup)                  // 注册用户路由
		router.InitBaseRouter(ApiGroup)                  // 注册基础功能路由 不做鉴权
		router.InitMenuRouter(ApiGroup)                  // 注册menu路由
		router.InitAuthorityRouter(ApiGroup)             // 注册角色路由
		router.InitApiRouter(ApiGroup)                   // 注册功能api路由
		router.InitFileUploadAndDownloadRouter(ApiGroup) // 文件上传下载功能路由
		router.InitWorkflowRouter(ApiGroup)              // 工作流相关路由
		router.InitCasbinRouter(ApiGroup)                // 权限相关路由
		router.InitJwtRouter(ApiGroup)                   // jwt相关路由
		router.InitSystemRouter(ApiGroup)                // system相关路由
		router.InitCustomerRouter(ApiGroup)              // 客户路由
		router.InitAutoCodeRouter(ApiGroup)              // 创建自动化代码
		router.InitCertHolderRouter(ApiGroup)            // 证书相关
		router.InitSysDictRouter(ApiGroup)               // 系统字典
		router.InitTitUserRouter(ApiGroup)               // 人才盘点用户
		router.InitTitUserBaseinfoRouter(ApiGroup)       // 人才盘点用户基础信息
		router.InitTitUserTopicAnswerRouter(ApiGroup)    // 人才盘点用户答题信息
		router.InitTitTrainingInfoRouter(ApiGroup)       // 人才盘点 用户基础信息 -- 培训信息
		router.InitTitTopicRouter(ApiGroup)              // 人才盘点题库
		router.InitTitTopicOptionRouter(ApiGroup)        // 人才盘点题库选项
	}

	// 人才盘点对应的路由
	TitGroup := Router.Group("/tit")
	{
		router.InitTitUserBaseRouter(TitGroup)         // 注册登录
		router.InitTitSysDictRouter(TitGroup)          // 查询字典数据
		router.InitAreaRouter(TitGroup)                // 省市区路由
		router.InitTitUserBasicinfoRouter(TitGroup)    // 增、改、查 人才基本信息
		router.InitTitTopicRouter2(TitGroup)           // 问题列表、选项相关
		router.InitTitUserTopicAnswerRouter2(TitGroup) // 人才用户答题相关
	}

	global.GVA_LOG.Info("router register success")
	return Router
}
