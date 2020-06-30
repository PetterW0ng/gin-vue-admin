package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateTitTopic
// @description   create a TitTopic
// @param     topic               model.TitTopic
// @auth      weiqin
// @return    err             error

func CreateTitTopic(topic model.TitTopic) (err error) {
	err = global.GVA_DB.Create(&topic).Error
	return err
}

// @title    DeleteTitTopic
// @description   delete a TitTopic
// @auth      weiqin
// @param     topic               model.TitTopic
// @return                    error

func DeleteTitTopic(topic model.TitTopic) (err error) {
	err = global.GVA_DB.Delete(topic).Error
	return err
}

// @title    UpdateTitTopic
// @description   update a TitTopic
// @param     topic          *model.TitTopic
// @auth      weiqin
// @return                    error

func UpdateTitTopic(topic *model.TitTopic) (err error) {
	err = global.GVA_DB.Save(topic).Error
	return err
}

// @title    GetTitTopic
// @description   get the info of a TitTopic
// @auth     weiqin
// @param     id              uint
// @return                    error
// @return    TitTopic        TitTopic

func GetTitTopic(id uint) (err error, topic model.TitTopic) {
	err = global.GVA_DB.Where("id = ?", id).First(&topic).Error
	return
}

// @title    GetTitTopicInfoList
// @description   get TitTopic list by pagination, 分页获取用户列表
// @auth      weiqin
// @param     info            PageInfo
// @return                    error

func GetTitTopicInfoList(info request.PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB
	var topics []model.TitTopic
	err = db.Find(&topics).Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&topics).Error
	return err, topics, total
}

func QueryByBusinessType(businessType int) (err error, topics []model.TitTopic) {
	db := global.GVA_DB
	err = db.Order("order").Where(" business_type = ?", businessType).Preload("TitTopicOptions").Find(&topics).Error
	return err, topics
}

func QueryTopicByBusinessType(businessType int) (err error, topics []model.TitTopic) {
	db := global.GVA_DB
	err = db.Order("order").Where(" business_type = ?", businessType).Preload("TitTopicRelatedList").Find(&topics).Error
	return err, topics
}
