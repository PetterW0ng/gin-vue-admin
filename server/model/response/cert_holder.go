package response

type Holder struct {
	UserName        string `json:"userName" bson:"userName"`
	SerialNum       string `json:"serialNum" bson:"serialNum"`
	CertificateName string `json:"certificateName" bson:"certificateName"`
	IssueTime       string `json:"issueTime" bson:"issueTime"`
	IssuingUnit     string `json:"issuingUnit" bson:"issuingUnit"`
	CertificateImg  string `json:"certificateImg" bson:"certificateImg"`
	CustomerTag     int    `json:"customerTag" bson:"customerTag"`
}
