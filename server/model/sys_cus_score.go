// 自动生成模板SysCusScore
package model

import (
	"github.com/jinzhu/gorm"
)

type SysCusScore struct {
	gorm.Model
	Eid       string `json:"eid" form:"eid" `
	CourseNum int    `json:"courseNum" form:"courseNum" `
	Duration  int    `json:"duration" form:"duration" `
	StudyRank string `json:"studyRank"`
	SequenceA int    `json:"sequenceA" form:"sequenceA" `
	SequenceB int    `json:"sequenceB" form:"sequenceB" `
	SequenceC int    `json:"sequenceC" form:"sequenceC" `
	SequenceD int    `json:"sequenceD" form:"sequenceD" `
	Var1      string `json:"var1" form:"var1" `
}
