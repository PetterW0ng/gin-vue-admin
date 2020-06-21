// 自动生成模板SysDict
package model

import (
	"github.com/jinzhu/gorm"
)

type SysDict struct {
	gorm.Model
	ParentId      int    `json:"parentId" form:"parentId" `
	PropertyName  string `json:"propertyName" form:"propertyName" `
	PropertyValue string `json:"propertyValue" form:"propertyValue" `
	SeqNumber     int    `json:"seqNumber" form:"seqNumber" `
	CodeName      string `json:"codeName" form:"codeName" `
	Code          string `json:"code" form:"code" `
}

// DictType, 来自于 SysDict 里面的 Code
type DictType string
