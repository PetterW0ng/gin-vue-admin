// 自动生成模板TitUserTopicAnswer
package model

import (
	"github.com/jinzhu/gorm"
)

type TitUserTopicAnswer struct {
	gorm.Model
	TitUserId           int    `json:"titUserId" form:"titUserId" `
	TitTopicId          int    `json:"titTopicId" form:"titTopicId" `
	TopicTitle          string `json:"topicTitle" form:"topicTitle" `
	TopicOptionSelected string `json:"topicOptionSelected" form:"topicOptionSelected" `
	Answer              string `json:"answer" form:"answer" `
	BatchNum            int    `json:"batchNum" form:"batchNum"`
	BusinessType        int    `json:"businessType" form:"businessType"`
	Score               int    `json:"score" form:"score" `
	Order               int    `json:"order" form:"order" `
}
