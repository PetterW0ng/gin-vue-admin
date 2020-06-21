// 自动生成模板TitTopic
package model

import (
	"github.com/jinzhu/gorm"
)

type TitTopic struct {
	gorm.Model
	Title           string           `json:"title" form:"title" `
	TopicType       int              `json:"topicType" form:"topicType" `
	BusinessType    int              `json:"businessType" form:"businessType" `
	IsRequired      int              `json:"isRequired" form:"isRequired" `
	IsScored        int              `json:"isScored" form:"isScored" `
	Order           int              `json:"order" form:"order" `
	SurveyLatitude  string           `json:"surveyLatitude" form:"surveyLatitude"`
	Score           string           `json:"score" form:"score"`
	TitTopicOptions []TitTopicOption `json:"options" form:"options"`
}
