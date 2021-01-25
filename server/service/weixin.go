package service

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/response"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

const (
	WEB_ACCESS_TOKEN_URL = "/sns/oauth2/access_token?appid=%v&secret=%v&code=%v&grant_type=authorization_code"
	GET_USERINFO_URL     = "/sns/userinfo?access_token=%v&openid=%v&lang=zh_CN"
	WX_TOKEN             = "weixin:token"
	JSAPI_TICKET         = "weixin:jsticket"
	WX_TOKEN_EXPIRE      = time.Second * 7200 // second
	ACCESS_TOKEN_URL     = "/cgi-bin/token?grant_type=client_credential&appid=%v&secret=%v"
	JSAPI_TICKET_URL     = "/cgi-bin/ticket/getticket?access_token=%v&type=jsapi"
	LETTERRUNES          = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
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

	var path = fmt.Sprintf(WEB_ACCESS_TOKEN_URL, global.GVA_CONFIG.WeiXin.Appkey, global.GVA_CONFIG.WeiXin.SecretKey, code)
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
	global.GVA_LOG.Info("向微信发起了获取 token 的请求， response= ", string(body))
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
	var path = fmt.Sprintf(GET_USERINFO_URL, wxSer.token, wxSer.openId)
	req, err := http.NewRequest("GET", global.GVA_CONFIG.WeiXin.Url+path, nil)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	global.GVA_LOG.Info("向微信发起了获取 userInfo 的请求， response= ", string(body))
	err = json.Unmarshal(body, &wxUserInfo)
	return
}

// 获取 Access_token
func getWxAccessToken() (token string, err error) {
	if cmd := global.GVA_REDIS.Get(WX_TOKEN); cmd.Err() == nil && len(cmd.Val()) > 0 {
		return cmd.Val(), nil
	} else {
		return queryAndSetWXToken()
	}
}

var wxTokenMutex, jsapiMutex sync.Mutex

func getWxJsapiTicket() (jsapiTicket string, err error) {
	if cmd := global.GVA_REDIS.Get(JSAPI_TICKET); cmd.Err() == nil && len(cmd.Val()) > 0 {
		return cmd.Val(), nil
	} else {
		jsapiMutex.Lock()
		defer jsapiMutex.Unlock()
		if token, err := getWxAccessToken(); err == nil {
			var path = fmt.Sprintf(JSAPI_TICKET_URL, token)
			req, err := http.NewRequest("GET", global.GVA_CONFIG.WeiXin.Url+path, nil)
			if err != nil {
				return "", err
			}
			req.Header.Set("Content-Type", "application/json")
			client := &http.Client{}
			resp, err := client.Do(req)
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)
			global.GVA_LOG.Info("向微信发起了获取jsapiTicket 的请求， response= ", string(body))
			var respMap map[string]interface{}
			if err := json.Unmarshal(body, &respMap); err != nil {
				return "", err
			}
			if v, ok := respMap["ticket"]; ok {
				global.GVA_REDIS.Set(JSAPI_TICKET, v, WX_TOKEN_EXPIRE)
				return v.(string), nil
			} else {
				return "", fmt.Errorf("向微信发起了获取jsapiTicket 请求失败了，原因：s%", respMap["errmsg"])
			}
		} else {
			return "", err
		}
	}
}

func GetSignatureConfig(openId, url string) (config response.WXConfigData, err error) {
	nonceStr := RandStringRunes(16)
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	if ticket, err := getWxJsapiTicket(); err != nil {
		global.GVA_LOG.Error("获取微信 ticket 出错了", err)
		return config, err
	} else {
		//keys := []string{"nonceStr", "url", "timestamp"}
		longstr := "jsapi_ticket=" + ticket + "&noncestr=" + nonceStr + "&timestamp=" + timestamp + "&url=" + url
		global.GVA_LOG.Info(longstr)
		h := sha1.New()
		h.Write([]byte(longstr))
		signature := fmt.Sprintf("%x", h.Sum(nil))
		shareUrl := "https://open.weixin.qq.com/connect/oauth2/authorize?appid=wx031c58989e81ab49&redirect_uri=https%3a%2f%2ftit.pkucarenjk.com%2f%23%2fSummarizing&response_type=code&scope=snsapi_userinfo&state=STATE#wechat_redirect"
		return response.WXConfigData{NonceStr: nonceStr, Timestamp: timestamp, AppId: global.GVA_CONFIG.WeiXin.Appkey, Signature: signature, ShareURL: shareUrl, JsApiList: []string{"updateTimelineShareData", "onMenuShareAppMessage"}}, nil
	}

}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = rune(LETTERRUNES[rand.Intn(len(LETTERRUNES))])
	}
	return string(b)
}

func queryAndSetWXToken() (token string, err error) {
	wxTokenMutex.Lock()
	defer wxTokenMutex.Unlock()
	if cmd := global.GVA_REDIS.Get(WX_TOKEN); cmd.Err() == nil && len(cmd.Val()) > 0 {
		return cmd.Val(), nil
	} else {
		var path = fmt.Sprintf(ACCESS_TOKEN_URL, global.GVA_CONFIG.WeiXin.Appkey, global.GVA_CONFIG.WeiXin.SecretKey)
		req, err := http.NewRequest("GET", global.GVA_CONFIG.WeiXin.Url+path, nil)
		if err != nil {
			return "", err
		}
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		global.GVA_LOG.Info("向微信发起了获取token 的请求， response= ", string(body))
		var respMap map[string]interface{}
		if err := json.Unmarshal(body, &respMap); err != nil {
			return "", err
		}
		if v, ok := respMap["access_token"]; ok {
			global.GVA_REDIS.Set(WX_TOKEN, v, WX_TOKEN_EXPIRE)
			return v.(string), nil
		} else {
			return "", fmt.Errorf("向微信发起了获取token 请求失败了，原因：s%", respMap["errmsg"])
		}
	}
}
