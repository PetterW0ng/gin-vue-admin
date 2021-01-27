package response

type CustomerStudy struct {
	XeOrderArray   []XeOrder `json:"orders"`
	CertArray      []Cert    `json:"certs"`
	CourseNum      int       `json:"courseNum"`
	Duration       int       `json:"duration"`
	StudyRank      string    `json:"studyRank"`
	SequenceA      int       `json:"sequenceA" form:"sequenceA" `
	SequenceB      int       `json:"sequenceB" form:"sequenceB" `
	SequenceC      int       `json:"sequenceC" form:"sequenceC" `
	SequenceD      int       `json:"sequenceD" form:"sequenceD" `
	CertNumRank    string    `json:"certNumRank"`
	PersistenceDay int       `json:"persistenceDay"`
}

type XeOrder struct {
	CourseName string `json:"courseName"`
	PayTime    string `json:"payTime"`
	Price      int    `json:"price"`
}

type Cert struct {
	CertName  string `json:"certName"`
	IssueTime string `json:"issueTime"`
}
