// 自动生成模板TitTrainingInfo
package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type TitTrainingInfo struct {
	gorm.Model
	TitUserBaseinfoId int       `json:"titUserBaseinfoId" form:"titUserBaseinfoId" `
	TrainingCourse    string    `json:"trainingCourse" form:"trainingCourse" `
	BeginTime         time.Time `json:"beginTime" form:"beginTime" `
	EndTime           time.Time `json:"endTime" form:"endTime" `
	PaymentWay        int       `json:"paymentWay" form:"paymentWay" `
}
