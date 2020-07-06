package model

import "github.com/jinzhu/gorm"

// TitRelatedClick , 用来记录用户对 得分页上的课程推荐的点击情况
type TitRelatedClick struct {
	gorm.Model
	UserPhone    string `json:"userPhone"`
	TitUserId    uint   `json:"titUserId"`
	TitRelatedId int    `json:"titRelatedId"`
	RecommendURL string `json:"recommendURL"`
}
