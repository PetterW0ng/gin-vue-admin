// 自动生成模板TitUser
package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type TitUser struct {
	gorm.Model
	Username                    string    `json:"username" form:"username" `
	Gender                      int       `json:"gender" form:"gender" `
	Birthday                    time.Time `json:"birthday" form:"birthday" `
	Telphone                    string    `json:"telphone" form:"telphone" `
	OpenId                      string    `json:"openId" form:"openId" `
	TitUserBaseinfoId           int       `json:"baseinfoId" form:"titUserBaseinfoId" `
	JobInfoBatchNum             int       `json:"jobInfoNum" form:"jobInfoBatchNum" `
	PerfessionKnowledgeBatchNum int       `json:"perfessionNum" form:"perfessionKnowledgeBatchNum" `
	IndustryPerspectiveBatchNum int       `json:"industryNum" form:"industryPerspectiveBatchNum" `
	IpAddress                   string    `json:"ipAddress" form:"ipAddress"`
}
