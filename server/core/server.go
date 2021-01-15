package core

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/initialize"
	"net/http"
	"time"
)

func RunWindowsServer() {
	if global.GVA_CONFIG.System.UseMultipoint {
		// 初始化redis服务
		initialize.Redis()
	}
	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")
	// 设置 静态文件访问路径
	Router.StaticFile("/MP_verify_h0cLD7HLadxHZUGK.txt", "./resource/MP_verify_h0cLD7HLadxHZUGK.txt")
	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := &http.Server{
		Addr:           address,
		Handler:        Router,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20, // header 大小为1M
	}
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.GVA_LOG.Debug("server run success on ", address)
	global.GVA_LOG.Error(s.ListenAndServe())
}
