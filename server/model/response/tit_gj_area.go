package response

type AreaTree struct {
	Id       string     `json:"id"  `
	Name     string     `json:"text" `
	Pid      string     `json:"pid"`
	Children []AreaTree `json:"children"`
}
