package response

type TitTopicRelated struct {
	ID              uint   `json:"id"`
	TitTopicId      int    `json:"titTopicId" form:"titTopicId" `
	RecommendType   int    `json:"recommendType" `
	RecommendObject string `json:"recommendObject"`
	ObjectLink      string `json:"objectLink"`
	Remark          string `json:"remark"`
}
