package response

import "time"

type TitUserBaseinfoTraining struct {
	// 教学背景
	Schoole          string `json:"schoole" form:"schoole" `
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
