package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateTitTopicOption
// @description   create a TitTopicOption
// @param     topicOption               model.TitTopicOption
// @auth      weiqin
// @return    err             error

func CreateTitTopicOption(topicOption model.TitTopicOption) (err error) {
	err = global.GVA_DB.Create(&topicOption).Error
	return err
}

// @title    DeleteTitTopicOption
// @description   delete a TitTopicOption
// @auth      weiqin
// @param     topicOption               model.TitTopicOption
// @return                    error

func DeleteTitTopicOption(topicOption model.TitTopicOption) (err error) {
	err = global.GVA_DB.Delete(topicOption).Error
	return err
}

// @title    UpdateTitTopicOption
// @description   update a TitTopicOption
// @param     topicOption          *model.TitTopicOption
// @auth      weiqin
// @return                    error

func UpdateTitTopicOption(topicOption *model.TitTopicOption) (err error) {
	err = global.GVA_DB.Save(topicOption).Error
	return err
}

// @title    GetTitTopicOption
// @description   get the info of a TitTopicOption
// @auth     weiqin
// @param     id              uint
// @return                    error
// @return    TitTopicOption        TitTopicOption

func GetTitTopicOption(id uint) (err error, topicOption model.TitTopicOption) {
	err = global.GVA_DB.Where("id = ?", id).First(&topicOption).Error
	return
}

// @title    GetTitTopicOptionInfoList
// @description   get TitTopicOption list by pagination, 分页获取用户列表
// @auth      weiqin
// @param     info            PageInfo
// @return                    error

func GetTitTopicOptionInfoList(info request.PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB
	var topicOptions []model.TitTopicOption
	err = db.Find(&topicOptions).Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&topicOptions).Error
	return err, topicOptions, total
}
