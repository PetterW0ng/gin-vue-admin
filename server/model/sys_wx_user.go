// 自动生成模板SysWxUser
package model

import (
	"github.com/jinzhu/gorm"
)

type SysWxUser struct {
	gorm.Model
	OpenId     string `json:"openid" form:"openId" sql:"unique_index:idx_openid"`
	Nickname   string `json:"nickname" form:"nickname" `
	Gender     int    `json:"sex" form:"sex" `
	Country    string `json:"country" form:"country" `
	Province   string `json:"province" form:"province" `
	City       string `json:"city" form:"city" `
	HeadImgurl string `json:"headimgurl" form:"headImgurl" `
	Unionid    string `json:"unionid" form:"unionid" `
	Phone      string `json:"phone" form:"phone" `
}
