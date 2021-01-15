// 自动生成模板SysCustomer
package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type SysCustomer struct {
	gorm.Model
	RegisterTime time.Time   `json:"registerTime" form:"registerTime"  gorm:"column:register_time"`
	EID          string      `json:"eID" form:"eID"  gorm:"column:eid"`
	Name         string      `json:"name" form:"name" `
	Nickname     string      `json:"nickname" form:"nickname"  gorm:"column:nickname"`
	Gender       int         `json:"gender" form:"gender"  gorm:"column:gender"`
	Birthday     time.Time   `json:"birthday" form:"birthday" `
	Phone        string      `json:"phone" form:"phone" `
	Email        string      `json:"email" form:"email" `
	WeixinNum    string      `json:"weixinNum" form:"weixinNum" `
	Source       int         `json:"source" form:"source" `
	EntryPoint   int         `json:"entryPoint" form:"entryPoint" `
	IsEvaluate   bool        `json:"isEvaluate" form:"isEvaluate" `
	CourseType   int         `json:"courseType" form:"courseType"`
	SysCusTags   []SysCusTag `json:"tags" form:"tags" `
	UserType     int         `json:"userType" form:"userType"`
}
