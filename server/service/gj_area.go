package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
)

func QueryAllProvince() (err error, areas []model.GjArea) {
	err = global.GVA_DB.Where("layer = ?", 0).Find(&areas).Error
	return
}

func FindCityOrArea(pid string) (err error, areas []model.GjArea) {
	err = global.GVA_DB.Where("pid = ?", pid).Find(&areas).Error
	return
}

func GetAreaById(areaId string) (area model.GjArea) {
	global.GVA_DB.Where("id = ?", areaId).First(&area)
	return
}

func GetAreaAll() (area []model.GjArea) {
	global.GVA_DB.Find(&area)
	return
}
