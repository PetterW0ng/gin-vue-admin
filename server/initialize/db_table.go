package initialize

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
)

// 注册数据库表专用
func DBTables() {
	db := global.GVA_DB
	db.AutoMigrate(model.SysUser{},
		model.SysAuthority{},
		model.SysApi{},
		model.SysBaseMenu{},
		model.JwtBlacklist{},
		model.SysWorkflow{},
		model.SysWorkflowStepInfo{},
		model.ExaFileUploadAndDownload{},
		model.ExaFile{},
		model.ExaFileChunk{},
		model.ExaCustomer{},
		model.SysDict{},
		model.TitUser{},
		model.TitUserBaseinfo{},
		model.TitUserTopicAnswer{},
		model.TitTrainingInfo{},
		model.TitTopic{},
		model.TitTopicOption{},
		model.TitTopicRelated{},
		model.TitRelatedClick{},
	)
	global.GVA_LOG.Debug("register table success")
}
