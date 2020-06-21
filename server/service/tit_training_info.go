package service

import (
	"bytes"
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateTitTrainingInfo
// @description   create a TitTrainingInfo
// @param     trainingInfo               model.TitTrainingInfo
// @auth      weiqin
// @return    err             error

func CreateTitTrainingInfo(trainingInfo model.TitTrainingInfo) (err error) {
	err = global.GVA_DB.Create(&trainingInfo).Error
	return err
}

// @title    DeleteTitTrainingInfo
// @description   delete a TitTrainingInfo
// @auth      weiqin
// @param     trainingInfo               model.TitTrainingInfo
// @return                    error

func DeleteTitTrainingInfo(trainingInfo model.TitTrainingInfo) (err error) {
	err = global.GVA_DB.Delete(trainingInfo).Error
	return err
}

// @title    UpdateTitTrainingInfo
// @description   update a TitTrainingInfo
// @param     trainingInfo          *model.TitTrainingInfo
// @auth      weiqin
// @return                    error

func UpdateTitTrainingInfo(trainingInfo *model.TitTrainingInfo) (err error) {
	err = global.GVA_DB.Save(trainingInfo).Error
	return err
}

// @title    GetTitTrainingInfo
// @description   get the info of a TitTrainingInfo
// @auth     weiqin
// @param     id              uint
// @return                    error
// @return    TitTrainingInfo        TitTrainingInfo

func GetTitTrainingInfo(id uint) (err error, trainingInfo model.TitTrainingInfo) {
	err = global.GVA_DB.Where("id = ?", id).First(&trainingInfo).Error
	return
}

// @title    GetTitTrainingInfoInfoList
// @description   get TitTrainingInfo list by pagination, 分页获取用户列表
// @auth      weiqin
// @param     info            PageInfo
// @return                    error

func GetTitTrainingInfoInfoList(info request.PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB
	var trainingInfos []model.TitTrainingInfo
	err = db.Find(&trainingInfos).Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&trainingInfos).Error
	return err, trainingInfos, total
}

func BatchAddTrainingInfo(trainingInfos []model.TitTrainingInfo) (err error) {
	var buffer bytes.Buffer
	batchInsert := "INSERT INTO `tit_training_infos` (`tit_user_baseinfo_id`,`training_course`,`begin_time`,`end_time`,`payment_way`) values "
	if _, err = buffer.WriteString(batchInsert); err != nil {
		return
	}
	db := global.GVA_DB
	for i, e := range trainingInfos {
		if i == len(trainingInfos)-1 {
			buffer.WriteString(fmt.Sprintf("('%d','%s','%s','%s', '%d');", e.TitUserBaseinfoId, e.TrainingCourse, e.BeginTime.Format("2006-01-02 15:04:05"), e.EndTime.Format("2006-01-02 15:04:05"), e.PaymentWay))
		} else {
			buffer.WriteString(fmt.Sprintf("('%d','%s','%s','%s', '%d'),", e.TitUserBaseinfoId, e.TrainingCourse, e.BeginTime.Format("2006-01-02 15:04:05"), e.EndTime.Format("2006-01-02 15:04:05"), e.PaymentWay))
		}
	}
	err = db.Exec(buffer.String()).Error
	return
}

func BatchModifyTrainingInfo(trainingInfos []model.TitTrainingInfo) (err error) {

	// 先删除之前关联的 trainingInfo
	db := global.GVA_DB
	db.Where("tit_user_baseinfo_id = ? ", trainingInfos[0].TitUserBaseinfoId).Delete(model.TitTrainingInfo{})

	var buffer bytes.Buffer
	batchInsert := "INSERT INTO `tit_training_infos` (`tit_user_baseinfo_id`,`training_course`,`begin_time`,`end_time`,`payment_way`) values "
	if _, err = buffer.WriteString(batchInsert); err != nil {
		return
	}

	for i, e := range trainingInfos {
		if i == len(trainingInfos)-1 {
			buffer.WriteString(fmt.Sprintf("('%d','%s','%s','%s', '%d');", e.TitUserBaseinfoId, e.TrainingCourse, e.BeginTime.Format("2006-01-02 15:04:05"), e.EndTime.Format("2006-01-02 15:04:05"), e.PaymentWay))
		} else {
			buffer.WriteString(fmt.Sprintf("('%d','%s','%s','%s', '%d'),", e.TitUserBaseinfoId, e.TrainingCourse, e.BeginTime.Format("2006-01-02 15:04:05"), e.EndTime.Format("2006-01-02 15:04:05"), e.PaymentWay))
		}
	}
	err = db.Exec(buffer.String()).Error
	return
}

func QueryTrainingInfoByBaseId(titUserBaseinfoId int) (trainingInfos []model.TitTrainingInfo) {
	db := global.GVA_DB
	db.Where("tit_user_baseinfo_id = ?", titUserBaseinfoId).Find(&trainingInfos)
	return
}
