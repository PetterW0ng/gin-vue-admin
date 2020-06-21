package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateTitUserBaseinfo
// @description   create a TitUserBaseinfo
// @param     userBaseinfo               model.TitUserBaseinfo
// @auth      weiqin
// @return    err             error

func CreateTitUserBaseinfo(userBaseinfo *model.TitUserBaseinfo) (err error) {
	err = global.GVA_DB.Save(&userBaseinfo).Error
	return err
}

// @title    DeleteTitUserBaseinfo
// @description   delete a TitUserBaseinfo
// @auth      weiqin
// @param     userBaseinfo               model.TitUserBaseinfo
// @return                    error

func DeleteTitUserBaseinfo(userBaseinfo model.TitUserBaseinfo) (err error) {
	err = global.GVA_DB.Delete(userBaseinfo).Error
	return err
}

// @title    UpdateTitUserBaseinfo
// @description   update a TitUserBaseinfo
// @param     userBaseinfo          *model.TitUserBaseinfo
// @auth      weiqin
// @return                    error

func UpdateTitUserBaseinfo(userBaseinfo *model.TitUserBaseinfo) (err error) {
	err = global.GVA_DB.Save(userBaseinfo).Error
	return err
}

// @title    GetTitUserBaseinfo
// @description   get the info of a TitUserBaseinfo
// @auth     weiqin
// @param     id              uint
// @return                    error
// @return    TitUserBaseinfo        TitUserBaseinfo

func GetTitUserBaseinfo(id uint) (err error, userBaseinfo model.TitUserBaseinfo) {
	err = global.GVA_DB.Where("id = ?", id).First(&userBaseinfo).Error
	return
}

// @title    GetTitUserBaseinfoInfoList
// @description   get TitUserBaseinfo list by pagination, 分页获取用户列表
// @auth      weiqin
// @param     info            PageInfo
// @return                    error

func GetTitUserBaseinfoInfoList(info request.PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB
	var userBaseinfos []model.TitUserBaseinfo
	err = db.Find(&userBaseinfos).Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&userBaseinfos).Error
	return err, userBaseinfos, total
}
