package model

// GjArea, 是国家行政区域及区域编码
type GjArea struct {
	Id          string `json:"id"  `
	Name        string `json:"name" `
	DisplayName string `json:"displayName" `
	Path        string `json:"path" `
	Layer       int    `json:"layer" `
	Sort        int    `json:"sort" `
	Pid         string `json:"pid"`
	Memo        string `json:"memo"`
	RecStat     string `json:"recStat"`
	Creator     string `json:"creator"`
	Modifier    string `json:"modifier"`
}
