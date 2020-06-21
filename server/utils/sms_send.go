package utils

import (
	"fmt"
	"gin-vue-admin/global"
	"github.com/gwpp/alidayu-go"
	"github.com/gwpp/alidayu-go/request"
)

func SendShotMessage(phone, code string) error {
	client := alidayu.NewTopClient(global.GVA_CONFIG.Dayu.Appkey, global.GVA_CONFIG.Dayu.SecretKey)
	req := request.NewAlibabaAliqinFcSmsNumSendRequest()
	req.Extend = "normal"
	req.SmsFreeSignName = global.GVA_CONFIG.Dayu.SignName
	req.RecNum = phone
	req.SmsTemplateCode = "SMS_4395149"
	req.SmsParam = fmt.Sprintf(`{"code":"%s", "product":"人才盘点工具"}`, code)
	response, err := client.Execute(req)
	global.GVA_LOG.Info(response)
	if err != nil {
		global.GVA_LOG.Error("验证码发送失败！", err)
		return err
	}
	return nil
}
