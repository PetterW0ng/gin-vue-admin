package service

import (
	"encoding/json"
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"io/ioutil"
	"net/http"
	"qiniupkg.com/x/log.v7"
)

const (
	ACCESS_TOKEN_URL = "/sns/oauth2/access_token?appid=%v&secret=%v&code=%v&grant_type=authorization_code"
	GET_USERINFO     = "/sns/userinfo?access_token=%v&openid=%v&lang=zh_CN"
	WX_TOKEN         = "weixin:token"
)

type WeixinServer interface {
	GetUserInfo() (wxUserInfo model.SysWxUser, err error)
	GetOpenId() string
	GetToken() string
}

type WxService struct {
	openId string
	token  string
	scope  string
}

func (wxSer *WxService) GetOpenId() string {
	return wxSer.openId
}

func (wxSer *WxService) GetToken() string {
	return wxSer.token
}

func WxServerNew(code string) (wxSer *WxService, err error) {

	var path = fmt.Sprintf(ACCESS_TOKEN_URL, global.GVA_CONFIG.WeiXin.Appkey, global.GVA_CONFIG.WeiXin.SecretKey, code)
	req, err := http.NewRequest("GET", global.GVA_CONFIG.WeiXin.Url+path, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	//body := []byte("{\"access_token\":\"41_YAwEPj5ToMNNyVD5duN917XgKrx_8aoOMFeQj-qg3rbWpJ_zfOGaTJgK0YlmDSKM4P5g8YKfNvpS53gw0WHhHA\",\"expires_in\":7200,\"refresh_token\":\"41_zdUUz0WUh9B6pBzjAOgbvRbMTtj0Ky-ffGs8tsdGZC6-3wUgoTSraT9zsnIkwmGEwYhePClNuG2Gd9HrMFFehw\",\"openid\":\"oezDYvpOC--3PEOUB2yIzRIEe6-8\",\"scope\":\"snsapi_userinfo\"}")
	log.Info("向微信发起了获取 token 的请求， response= ", string(body))
	var respMap map[string]interface{}
	if err := json.Unmarshal(body, &respMap); err != nil {
		return nil, err
	}
	if v, ok := respMap["access_token"]; ok {
		return &WxService{respMap["openid"].(string), v.(string), respMap["scope"].(string)}, nil
	} else {
		return nil, err
	}
}

func (wxSer *WxService) GetUserInfo() (wxUserInfo model.SysWxUser, err error) {
	var path = fmt.Sprintf(GET_USERINFO, wxSer.token, wxSer.openId)
	req, err := http.NewRequest("GET", global.GVA_CONFIG.WeiXin.Url+path, nil)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	log.Info("向微信发起了获取 userInfo 的请求， response= ", string(body))
	err = json.Unmarshal(body, &wxUserInfo)
	return
}
