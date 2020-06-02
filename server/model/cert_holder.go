package model

type Holder struct {
	UserName        string `json:"userName" bson:"userName"`
	SerialNum       string `json:"serialNum" bson:"serialNum"`
	CertificateName string `json:"certificateName" bson:"certificateName"`
	IssueTime       string `json:"issueTime" bson:"issueTime"`
	IssuingUnit     string `json:"issuingUnit" bson:"issuingUnit"`
	IdCard          string `json:"idCard" bson:"idCard"`
	Phone           string `json:"phone" bson:"phone"`
	CreateTime      string `json:"createTime" bson:"createTime"`
	CertificateImg  string `json:"certificateImg" bson:"certificateImg"`
}
