// 自动生成模板TitUserBaseinfo
package model

import (
	"github.com/jinzhu/gorm"
)

type TitUserBaseinfo struct {
	gorm.Model
	TitUserId        uint   `json:"titUserId" form:"titUserId" `
	Schoole          int    `json:"schoole" form:"schoole" `
	MajorsStudied    int    `json:"majorsStudied" form:"majorsStudied" `
	HighestEducation int    `json:"highestEducation" form:"highestEducation" `
	SchoolSystem     int    `json:"schoolSystem" form:"schoolSystem" `
	WorkingState     int    `json:"workingState" form:"workingState" `
	Company          string `json:"company" form:"company" `
	Area             string `json:"area" form:"area" `
	JobTitle         string `json:"jobTitle" form:"jobTitle" `
	ServiceType      int    `json:"serviceType" form:"serviceType" `
	Income           int    `json:"income" form:"income" `
	Benefits         string `json:"benefits" form:"benefits" `
	ChildType        string `json:"childType" form:"childType" `
	ChildAge         int    `json:"childAge" form:"childAge" `
	TrainingNumber   string `json:"trainingNumber" form:"trainingNumber" `
	TrainingFee      int    `json:"trainingFee" form:"trainingFee" `
}
