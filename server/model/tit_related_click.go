package model

import "github.com/jinzhu/gorm"

// TitRelatedClick , 用来记录用户对 得分页上的课程推荐的点击情况
// 对后面的每个连接的点击事件统计 做准备
// 1、查询出最近一周的时间、次数、对象：select date_format(c.created_at, '%Y-%m-%d'),c.recommend_object , count(1) from tit_related_clicks c where c.deleted_at is null and c.created_at >= date_sub(CURDATE(), INTERVAL 6 DAY) group by date_format(c.created_at, '%Y-%m-%d'),c.recommend_object;
type TitRelatedClick struct {
	gorm.Model
	UserPhone       string `json:"userPhone"`
	TitUserId       uint   `json:"titUserId"`
	TitRelatedId    int    `json:"titRelatedId"`
	RecommendURL    string `json:"recommendURL"`
	RecommendObject string `json:"recommendObject"`
}
