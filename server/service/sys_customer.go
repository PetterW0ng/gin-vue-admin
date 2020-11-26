package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateSysCustomer
// @description   create a SysCustomer
// @param     customer               model.SysCustomer
// @auth      weiqin
// @return    err             error

func CreateSysCustomer(customer model.SysCustomer) (err error) {
	err = global.GVA_DB.Create(&customer).Error
	return err
}

// @title    DeleteSysCustomer
// @description   delete a SysCustomer
// @auth      weiqin
// @param     customer               model.SysCustomer
// @return                    error

func DeleteSysCustomer(customer model.SysCustomer) (err error) {
	err = global.GVA_DB.Delete(customer).Error
	return err
}

// @title    UpdateSysCustomer
// @description   update a SysCustomer
// @param     customer          *model.SysCustomer
// @auth      weiqin
// @return                    error

func UpdateSysCustomer(customer *model.SysCustomer) (err error) {
	err = global.GVA_DB.Save(customer).Error
	return err
}

// @title    GetSysCustomer
// @description   get the info of a SysCustomer
// @auth     weiqin
// @param     id              uint
// @return                    error
// @return    SysCustomer        SysCustomer

func GetSysCustomer(id uint) (err error, customer model.SysCustomer) {
	err = global.GVA_DB.Where("id = ?", id).First(&customer).Error
	return
}

// @title    GetSysCustomerInfoList
// @description   get SysCustomer list by pagination, 分页获取客户列表
// @auth      weiqin
// @param     info            PaginatedSysCustomer
// @return                    error

func GetSysCustomerInfoList(info request.PaginatedSysCustomer) (err error, list []model.SysCustomer, total int) {
	limit := info.PageSize
	if limit == 0 {
		limit = 10
	}
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB
	db.LogMode(true)
	var customers []model.SysCustomer
	if info.Gender != 0 {
		db.Where("gender = ?", info.Gender)
	}
	if info.EntryPoint != 0 {
		db.Where("entry_point = ?", info.EntryPoint)
	}
	if info.Source != 0 {
		db.Where("source = ?", info.Source)
	}
	if !info.RegisterTimeBegin.IsZero() {
		db.Where("register_time >= ?", info.RegisterTimeBegin)
	}
	if !info.RegisterTimeEnd.IsZero() {
		db.Where("register_time <= ?", info.RegisterTimeEnd)
	}
	if !info.Birthday.IsZero() {
		db.Where("birthday = ?", info.Birthday)
	}
	if info.Query != "" {
		db.Where("(name = ? or nickname = ? or phone = ?)", info.Query, info.Query, info.Query)
	}
	db.Where("is_evaluate = ?", info.IsEvaluate)

	err = db.Model(&customers).Count(&total).Error
	// Related 和 Preload 的使用
	err = db.Limit(limit).Offset(offset).Preload("SysCusTags").Find(&customers).Error
	return err, customers, total
}
