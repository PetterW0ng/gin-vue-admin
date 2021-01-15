package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateSysWxUser
// @description   create a SysWxUser
// @param     wxUser               model.SysWxUser
// @auth      weiqin
// @return    err             error

func CreateSysWxUser(wxUser model.SysWxUser) (err error) {
	err = global.GVA_DB.Create(&wxUser).Error
	return err
}

// @title    DeleteSysWxUser
// @description   delete a SysWxUser
// @auth      weiqin
// @param     wxUser               model.SysWxUser
// @return                    error

func DeleteSysWxUser(wxUser model.SysWxUser) (err error) {
	err = global.GVA_DB.Delete(wxUser).Error
	return err
}

// @title    UpdateSysWxUser
// @description   update a SysWxUser
// @param     wxUser          *model.SysWxUser
// @auth      weiqin
// @return                    error

func UpdateSysWxUser(wxUser *model.SysWxUser) (err error) {
	err = global.GVA_DB.Save(wxUser).Error
	return err
}

// @title    GetSysWxUser
// @description   get the info of a SysWxUser
// @auth     weiqin
// @param     id              uint
// @return                    error
// @return    SysWxUser        SysWxUser

func GetSysWxUser(id uint) (err error, wxUser model.SysWxUser) {
	err = global.GVA_DB.Where("id = ?", id).First(&wxUser).Error
	return
}

// 跟据 openId 查找
func GetSysWxUserByOpenId(openId string) (err error, wxUser model.SysWxUser) {
	err = global.GVA_DB.Where("open_id = ?", openId).First(&wxUser).Error
	return
}

// @title    GetSysWxUserInfoList
// @description   get SysWxUser list by pagination, 分页获取用户列表
// @auth      weiqin
// @param     info            PageInfo
// @return                    error

func GetSysWxUserInfoList(info request.PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB
	var wxUsers []model.SysWxUser
	err = db.Find(&wxUsers).Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&wxUsers).Error
	return err, wxUsers, total
}
