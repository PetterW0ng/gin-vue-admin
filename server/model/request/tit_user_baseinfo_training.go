package request

import "time"

type TitUserBaseinfoTraining struct {
	School           string         `json:"school" form:"school" `
	MajorsStudied    int            `json:"majorsStudied" form:"majorsStudied" `
	HighestEducation int            `json:"highestEducation" form:"highestEducation" `
	SchoolSystem     int            `json:"schoolSystem" form:"schoolSystem" `
	WorkingState     int            `json:"workingState" form:"workingState" `
	Company          string         `json:"company" form:"company" `
	Area             string         `json:"area" form:"area" `
	JobTitle         string         `json:"jobTitle" form:"jobTitle" `
	ServiceType      int            `json:"serviceType" form:"serviceType" `
	Income           int            `json:"income" form:"income" `
	Benefits         []int          `json:"benefits" form:"benefits" `
	ChildType        []int          `json:"childType" form:"childType" `
	ChildAge         int            `json:"childAge" form:"childAge" `
	TrainingNumber   []int          `json:"trainingNumber" form:"trainingNumber" `
	TrainingFee      int            `json:"trainingFee" form:"trainingFee" `
	TrainingInfos    []TrainingInfo `json:"trainingInfos" form:"trainingInfos" `
}

type TrainingInfo struct {
	TrainingCourse string    `json:"trainingCourse" form:"trainingCourse" `
	BeginTime      time.Time `json:"beginTime" form:"beginTime" `
	EndTime        time.Time `json:"endTime" form:"endTime" `
	PaymentWay     int       `json:"paymentWay" form:"paymentWay" `
}
