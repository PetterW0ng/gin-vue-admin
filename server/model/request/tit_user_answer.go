package request

type UserTopicAnswer struct {
	BusinessType int      `json:"businessType" form:"businessType" `
	Answers      []Answer `json:"answers" form:"answers"`
}

type Answer struct {
	TitTopicId        int    `json:"topicId" form:"topicId"`
	TitTopicOptionIds []int  `json:"optionIds" from:"optionIds"`
	Others            string `json:"others" form:"others"`
}
