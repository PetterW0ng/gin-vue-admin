package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gin-vue-admin/global"
	resp "gin-vue-admin/model/response"
	"io/ioutil"
	"net/http"
	"qiniupkg.com/x/log.v7"
	"sync"
	"time"
)

var (
	mutex              sync.Mutex
	xiaoe_token        = "xiaoe:token"
	xiaoe_token_expire = time.Second * 7200 // second
	uri_token          = "/token"
	uri_user_order     = "/xe.get.user.orders/1.0.0"
	orderStateMap      = map[int]string{0: "未支付", 1: "支付成功", 2: "支付失败", 6: "订单过期", 10: "退款成功"}
	paymentTypeMap     = map[int]string{2: "单笔", 3: "付费产品包", 4: "团购", 5: "单笔的购买赠送", 6: "产品包的购买赠送", 7: "问答提问", 8: "问答偷听", 11: "付费活动报名", 12: "打赏类型", 13: "拼团单个资源", 14: "拼团产品包", 15: "超级会员"}
	resourceTypeMap    = map[int]string{0: "无", 1: "图文", 2: "音频", 3: "视频", 4: "直播", 5: "活动报名", 6: "专栏/会员", 7: "社群", 8: "大专栏", 20: "电子书", 21: "实物商品", 23: "超级会员", 25: "训练营", 29: "面授课"}
	wxAppTypeMap       = map[int]string{0: "小程序", 1: "公众号", 10: "开放API"}
)

type tokenResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type userOrderRequest struct {
	AccessToken string `json:"access_token"`
	UserId      string `json:"user_id"`
}

type userOrderResponse struct {
	Code          int         `json:"code"`
	Msg           string      `json:"msg"`
	Data          interface{} `json:"data"`
	RequestParams interface{} `json:"request_params"`
}

func QueryUserOrder(eid string) (err error, list []resp.XetOrder) {
	toaken, err := getToken()
	if err != nil {
		log.Error("请求用户订单时，获取token出错了", err)
		return
	}
	userOrderRequest := userOrderRequest{toaken, eid}
	requestBody, _ := json.Marshal(userOrderRequest)
	req, err := http.NewRequest("POST", global.GVA_CONFIG.XiaoETong.URL+uri_user_order, bytes.NewBuffer(requestBody))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	httpResp, err := client.Do(req)
	defer httpResp.Body.Close()
	body, _ := ioutil.ReadAll(httpResp.Body)
	var userOrderResponse userOrderResponse
	if err = json.Unmarshal(body, &userOrderResponse); err == nil {
		if userOrderResponse.Code == 0 {
			data := userOrderResponse.Data.(map[string]interface{})
			orderOuterList := data["list"].([]interface{})
			for _, item := range orderOuterList {
				it := item.(resp.XetOrderOuter)
				xetOrder := resp.XetOrder{}
				xetOrder.SetValues(it)
				xetOrder.PaymentType = paymentTypeMap[it.PaymentType]
				xetOrder.OrderState = orderStateMap[it.OrderState]
				xetOrder.WxAppType = wxAppTypeMap[it.WxAppType]
				xetOrder.ResourceType = wxAppTypeMap[it.ResourceType]
				list = append(list, xetOrder)
			}
			return
		} else {
			return fmt.Errorf("向小鹅通发起了获取order 请求失败了，原因：s%", userOrderResponse.Msg), nil
		}
	} else {
		return err, nil
	}
}

func getToken() (token string, err error) {
	if cmd := global.GVA_REDIS.Get(xiaoe_token); cmd.Err() == nil && len(cmd.Val()) > 0 {
		return cmd.Val(), nil
	} else {
		return queryAndSetToken()
	}
}

func queryAndSetToken() (token string, err error) {
	mutex.Lock()
	defer mutex.Unlock()
	requestBody, _ := json.Marshal(global.GVA_CONFIG.XiaoETong)
	req, err := http.NewRequest("GET", global.GVA_CONFIG.XiaoETong.URL+uri_token, bytes.NewBuffer(requestBody))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	log.Info("向小鹅通发起了获取token 的请求， response= s%", resp.Body)
	body, _ := ioutil.ReadAll(resp.Body)
	var tokenResponse tokenResponse
	if err = json.Unmarshal(body, &tokenResponse); err == nil {
		if tokenResponse.Code == 0 {
			// 保存 token 至 redis
			data := tokenResponse.Data.(map[string]interface{})
			token := data["access_token"]
			cmd := global.GVA_REDIS.Set(xiaoe_token, token, xiaoe_token_expire)
			if cmd.Err() == nil {
				return token.(string), nil
			} else {
				log.Warn("向小鹅通发起了获取token 请求成功，缓存时失败了")
				return token.(string), cmd.Err()
			}
		} else {
			return "", fmt.Errorf("向小鹅通发起了获取token 请求失败了，原因：s%", tokenResponse.Msg)
		}
	} else {
		return "", err
	}
}
