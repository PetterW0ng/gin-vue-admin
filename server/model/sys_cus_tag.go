// 自动生成模板SysCusTag
package model

import (
	"github.com/jinzhu/gorm"
)

type SysCusTag struct {
	gorm.Model
	SysCustomerId int    `json:"cusID" form:"cusID" `
	TagCode       string `json:"tagCode" form:"tagCode" `
	TagValue      int    `json:"tagValue" form:"tagValue" `
}
