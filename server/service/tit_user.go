package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
)

// @title    CreateTitUser
// @description   create a TitUser
// @param     user               model.TitUser
// @auth                     （2020/04/05  20:22）
// @return    err             error
func CreateTitUser(user model.TitUser) (err error) {
	err = global.GVA_DB.Create(&user).Error
	return err
}

// @title    DeleteTitUser
// @description   delete a TitUser
// @auth                     （2020/04/05  20:22）
// @param     user               model.TitUser
// @return                    error

func DeleteTitUser(user model.TitUser) (err error) {
	err = global.GVA_DB.Delete(user).Error
	return err
}

// @title    UpdateTitUser
// @description   update a TitUser
// @param     user          *model.TitUser
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateTitUser(user *model.TitUser) (err error) {
	err = global.GVA_DB.Save(user).Error
	return err
}

// @title    GetTitUser
// @description   get the info of a TitUser
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    TitUser        TitUser

func GetTitUser(id uint) (err error, user model.TitUser) {
	err = global.GVA_DB.Where("id = ?", id).First(&user).Error
	return
}

// @title    GetTitUserInfoList
// @description   get TitUser list by pagination, 分页获取用户列表
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetTitUserInfoList(info request.PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB
	var users []response.TitUserInfoForBackend
	var titUser []model.TitUser
	err = db.Find(&titUser).Count(&total).Error
	//err = db.Order("created_at desc").Limit(limit).Offset(offset).Find(&users).Error
	err = db.Order("u.created_at desc").Limit(limit).Offset(offset).Table("tit_users u").Select("u.id, u.created_at, u.username, u.gender, u.birthday, u.telphone, uta.topic_option_ids change_job_option").Joins("left join tit_user_topic_answers uta on u.id = uta.tit_user_id and u.job_info_batch_num = uta.batch_num and uta.tit_topic_id = 12").Where("u.deleted_at is null").Scan(&users).Error
	return err, users, total
}

func FindTitUserByPhone(telphone string) (err error, u model.TitUser) {
	db := global.GVA_DB
	err = db.First(&u, " telphone = ?", telphone).Error
	return
}

func ModifyTitUser(user *model.TitUser) (err error) {
	db := global.GVA_DB
	err = db.Model(&user).Update(user).Error
	return
}
