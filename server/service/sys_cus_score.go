package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateSysCusScore
// @description   create a SysCusScore
// @param     sCusScore               model.SysCusScore
// @auth      weiqin
// @return    err             error

func CreateSysCusScore(sCusScore model.SysCusScore) (err error) {
	err = global.GVA_DB.Create(&sCusScore).Error
	return err
}

// @title    DeleteSysCusScore
// @description   delete a SysCusScore
// @auth      weiqin
// @param     sCusScore               model.SysCusScore
// @return                    error

func DeleteSysCusScore(sCusScore model.SysCusScore) (err error) {
	err = global.GVA_DB.Delete(sCusScore).Error
	return err
}

// @title    UpdateSysCusScore
// @description   update a SysCusScore
// @param     sCusScore          *model.SysCusScore
// @auth      weiqin
// @return                    error

func UpdateSysCusScore(sCusScore *model.SysCusScore) (err error) {
	err = global.GVA_DB.Save(sCusScore).Error
	return err
}

// @title    GetSysCusScore
// @description   get the info of a SysCusScore
// @auth     weiqin
// @param     id              uint
// @return                    error
// @return    SysCusScore        SysCusScore

func GetSysCusScore(id uint) (err error, sCusScore model.SysCusScore) {
	err = global.GVA_DB.Where("id = ?", id).First(&sCusScore).Error
	return
}

func QuerySysCusScore(eid string) (err error, sCusScore model.SysCusScore) {
	err = global.GVA_DB.Where("eid = ?", eid).First(&sCusScore).Error
	return
}

// @title    GetSysCusScoreInfoList
// @description   get SysCusScore list by pagination, 分页获取用户列表
// @auth      weiqin
// @param     info            PageInfo
// @return                    error

func GetSysCusScoreInfoList(info request.PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB
	var sCusScores []model.SysCusScore
	err = db.Find(&sCusScores).Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&sCusScores).Error
	return err, sCusScores, total
}
