package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateSysDict
// @description   create a SysDict
// @param     dict               model.SysDict
// @auth      weiqin
// @return    err             error

func CreateSysDict(dict model.SysDict) (err error) {
	err = global.GVA_DB.Create(&dict).Error
	return err
}

// @title    DeleteSysDict
// @description   delete a SysDict
// @auth      weiqin
// @param     dict               model.SysDict
// @return                    error

func DeleteSysDict(dict model.SysDict) (err error) {
	err = global.GVA_DB.Delete(dict).Error
	return err
}

// @title    UpdateSysDict
// @description   update a SysDict
// @param     dict          *model.SysDict
// @auth      weiqin
// @return                    error

func UpdateSysDict(dict *model.SysDict) (err error) {
	err = global.GVA_DB.Save(dict).Error
	return err
}

// @title    GetSysDict
// @description   get the info of a SysDict
// @auth     weiqin
// @param     id              uint
// @return                    error
// @return    SysDict        SysDict

func GetSysDict(id uint) (err error, dict model.SysDict) {
	err = global.GVA_DB.Where("id = ?", id).First(&dict).Error
	return
}

// @title    GetSysDictInfoList
// @description   get SysDict list by pagination, 分页获取用户列表
// @auth      weiqin
// @param     info            PageInfo
// @return                    error

func GetSysDictInfoList(info request.PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB
	var dicts []model.SysDict
	err = db.Find(&dicts).Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&dicts).Error
	return err, dicts, total
}

// 根据 SysDictExample 查找列表
func FindByCode(code model.DictType) (err error, dicts []model.SysDict) {
	db := global.GVA_DB
	err = db.Select("parent_id, property_name, property_value, code, code_name").Order("seq_number").Where(" code = ? ", code.Value()).Find(&dicts).Error
	return err, dicts
}

func GetAllDict() (err error, dicts []model.SysDict) {
	db := global.GVA_DB
	err = db.Select("parent_id, property_name, property_value, code, code_name").Order("seq_number").Find(&dicts).Error
	return err, dicts
}

func GetGroupedDict() (dictGroup map[string]map[int]string) {
	dictGroup = make(map[string]map[int]string)
	if err, dictList := GetAllDict(); err == nil {
		for _, value := range dictList {
			v, ok := dictGroup[value.Code]
			if ok {
				v[value.PropertyValue] = value.PropertyName
			} else {
				v = map[int]string{}
				v[value.PropertyValue] = value.PropertyName
			}
			dictGroup[value.Code] = v
		}
	}
	return
}
