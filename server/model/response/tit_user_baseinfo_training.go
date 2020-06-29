package response

import (
	"time"
)

// TitUserBaseinfoTraining, 数据已经转化过了，可以直接在页面展示
type TitUserBaseinfoTraining struct {
	// 教学背景
	School           string `json:"school" form:"school" `
	MajorsStudied    string `json:"majorsStudied" form:"majorsStudied" `
	HighestEducation string `json:"highestEducation" form:"highestEducation" `
	SchoolSystem     string `json:"schoolSystem" form:"schoolSystem" `
	// 工作单位
	WorkingState string `json:"workingState" form:"workingState" `
	Company      string `json:"company" form:"company" `
	Area         string `json:"area" form:"area" `
	JobTitle     string `json:"jobTitle" form:"jobTitle" `
	ServiceType  string `json:"serviceType" form:"serviceType" `
	Income       string `json:"income" form:"income" `
	Benefits     string `json:"benefits" form:"benefits" `
	ChildType    string `json:"childType" form:"childType" `
	ChildAge     string `json:"childAge" form:"childAge" `
	// 培训经历
	TrainingNumber string         `json:"trainingNumber" form:"trainingNumber" `
	TrainingFee    string         `json:"trainingFee" form:"trainingFee" `
	TrainingInfos  []TrainingInfo `json:"trainingInfos" form:"trainingInfos" `
}

type TrainingInfo struct {
	TrainingCourse string    `json:"trainingCourse" form:"trainingCourse" `
	BeginTime      time.Time `json:"beginTime" form:"beginTime" `
	EndTime        time.Time `json:"endTime" form:"endTime" `
	PaymentWay     string    `json:"paymentWay" form:"paymentWay" `
}

// TitUserBase， 该类型为表里面的原始数据 ，拿来修改数据使用
type TitUserBase struct {
	ID               uint              `json:"id"`
	TitUserId        uint              `json:"titUserId" form:"titUserId" `
	School           int               `json:"school" form:"school" `
	MajorsStudied    int               `json:"majorsStudied" form:"majorsStudied" `
	HighestEducation int               `json:"highestEducation" form:"highestEducation" `
	SchoolSystem     int               `json:"schoolSystem" form:"schoolSystem" `
	WorkingState     int               `json:"workingState" form:"workingState" `
	Company          string            `json:"company" form:"company" `
	Area             string            `json:"area" form:"area" `
	JobTitle         string            `json:"jobTitle" form:"jobTitle" `
	ServiceType      int               `json:"serviceType" form:"serviceType" `
	Income           int               `json:"income" form:"income" `
	Benefits         string            `json:"benefits" form:"benefits" `
	ChildType        string            `json:"childType" form:"childType" `
	ChildAge         int               `json:"childAge" form:"childAge" `
	TrainingNumber   string            `json:"trainingNumber" form:"trainingNumber" `
	TrainingFee      int               `json:"trainingFee" form:"trainingFee" `
	TrainingInfo     []TitTrainingInfo `json:"trainingInfos" form:"trainingInfos" `
}

type TitTrainingInfo struct {
	TitUserBaseinfoId int       `json:"titUserBaseinfoId" form:"titUserBaseinfoId" `
	TrainingCourse    string    `json:"trainingCourse" form:"trainingCourse" `
	BeginTime         time.Time `json:"beginTime" form:"beginTime" `
	EndTime           time.Time `json:"endTime" form:"endTime" `
	PaymentWay        int       `json:"paymentWay" form:"paymentWay" `
}
