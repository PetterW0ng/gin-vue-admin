package response

type WXConfigData struct {
	NonceStr  string   `json:"nonceStr" mapstructure:"nonceStr"`
	Timestamp string   `json:"timestamp" mapstructure:"timestamp"`
	AppId     string   `json:"appId" mapstructure:"appId"`
	Signature string   `json:"signature" mapstructure:"signature"`
	JsApiList []string `json:"jsApiList" mapstructure:"jsApiList"`
	ShareURL  string   `json:"shareUrl" mapstucture:"shareUrl"`
}
