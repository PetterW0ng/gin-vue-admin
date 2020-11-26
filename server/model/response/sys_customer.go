// 自动生成模板SysCustomer
package response

import (
	"gin-vue-admin/model"
	"github.com/jinzhu/gorm"
	"time"
)

type SysCustomer struct {
	gorm.Model
	RegisterTime time.Time   `json:"registerTime" form:"registerTime"  gorm:"column:register_time"`
	EID          string      `json:"eID" form:"eID"  gorm:"column:e_id"`
	Name         string      `json:"name" form:"name" `
	Nickname     string      `json:"nickname" form:"nickname"  gorm:"column:nickname"`
	Gender       string      `json:"gender" form:"gender"  gorm:"column:gender"`
	Birthday     time.Time   `json:"birthday" form:"birthday" `
	Phone        string      `json:"phone" form:"phone" `
	Email        string      `json:"email" form:"email" `
	WeixinNum    string      `json:"weixinNum" form:"weixinNum" `
	Source       string      `json:"source" form:"source" `
	EntryPoint   string      `json:"entryPoint" form:"entryPoint" `
	IsEvaluate   string      `json:"isEvaluate" form:"isEvaluate" `
	Tags         []SysCusTag `json:"tags" form:"tags"`
}

// 设置不需要转换的 properties from model.SysCustomer
func (customer *SysCustomer) SetValueFromDBModel(sysCustomer model.SysCustomer) {
	customer.RegisterTime = sysCustomer.RegisterTime
	customer.EID = sysCustomer.EID
	customer.Name = sysCustomer.Name
	customer.Nickname = sysCustomer.Nickname
	customer.Birthday = sysCustomer.Birthday
	customer.Phone = sysCustomer.Phone
	customer.Email = sysCustomer.Email
	customer.WeixinNum = sysCustomer.WeixinNum
	customer.ID = sysCustomer.ID
}

type SysCusTag struct {
	TagCode  string `json:"tagCode" form:"tagCode" `
	TagValue int    `json:"tagValue" form:"tagValue" `
	TagName  string `json:"tagName" form:"tagName"`
}
