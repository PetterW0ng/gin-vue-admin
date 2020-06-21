// 自动生成模板TitTopicOption
package model

import (
	"github.com/jinzhu/gorm"
)

type TitTopicOption struct {
	gorm.Model
	TitTopicId int    `json:"titTopicId" form:"titTopicId" `
	Title      string `json:"title" form:"title" `
	Score      int    `json:"score" form:"score" `
	Order      int    `json:"order" form:"order" `
}
