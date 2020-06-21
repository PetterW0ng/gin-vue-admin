package model

import "github.com/jinzhu/gorm"

type TitTopicRelated struct {
	gorm.Model
	TitTopicId      int    `json:"titTopicId" form:"titTopicId" `
	RecommendType   int    `json:"recommendType" `
	RecommendObject string `json:"recommendObject"`
	ObjectLink      string `json:"objectLink"`
	Remark          string `json:"remark"`
	Order           int    `json:"order" form:"order" `
}
