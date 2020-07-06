package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
)

func CreateTitRelatedClick(userClick model.TitRelatedClick) (err error) {
	err = global.GVA_DB.Create(&userClick).Error
	return err
}
