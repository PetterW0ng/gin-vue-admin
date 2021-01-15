// 自动生成模板SysDict
package model

import (
	"github.com/jinzhu/gorm"
)

type SysDict struct {
	gorm.Model
	ParentId      int    `json:"parentId" form:"parentId" `
	PropertyName  string `json:"text" form:"propertyName" `
	PropertyValue int    `json:"propertyValue" form:"propertyValue" `
	SeqNumber     int    `json:"seqNumber" form:"seqNumber" `
	CodeName      string `json:"codeName" form:"codeName" `
	Code          string `json:"code" form:"code" `
}

// DictType, 来自于 SysDict 里面的 Code
type DictType string

func (dict DictType) Value() string {
	return string(dict)
}

const (
	GENDER          DictType = "gender"
	SCHOOL          DictType = "school"
	STUDY_MAJOR     DictType = "study_major"
	EDUCATION       DictType = "education"
	SCHOO_SYSTEM    DictType = "school_system"
	WORKING_TATE    DictType = "working_tate"
	SERVICE_YPE     DictType = "service_ype"
	INCOME          DictType = "income"
	BENEFITS        DictType = "benefits"
	CHILD_TYPE      DictType = "child_type"
	CHILD_AGE       DictType = "child_age"
	TRAINING_NUMBER DictType = "training_number"
	TRAINING_FEE    DictType = "training_fee"
	PAYMENT_WAY     DictType = "payment_way"
	SOURCE          DictType = "source"
	ENTRY_POINT     DictType = "entry_point"
	IS_EVALUATE     DictType = "is_evaluate"
	WORK_INTENTION  DictType = "work_intention"
	CUSTOMER_TAG    DictType = "customer_tag"
	USER_TYPE       DictType = "user_type"
)
