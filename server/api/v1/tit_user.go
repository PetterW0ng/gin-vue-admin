package v1

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/global/response"
	"gin-vue-admin/middleware"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"gin-vue-admin/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

// @Tags TitUser
// @Summary 创建TitUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TitUser true "创建TitUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/createTitUser [post]
func CreateTitUser(c *gin.Context) {
	var user model.TitUser
	_ = c.ShouldBindJSON(&user)
	err := service.CreateTitUser(user)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags TitUser
// @Summary 删除TitUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TitUser true "删除TitUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /user/deleteTitUser [delete]
func DeleteTitUser(c *gin.Context) {
	var user model.TitUser
	_ = c.ShouldBindJSON(&user)
	err := service.DeleteTitUser(user)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags TitUser
// @Summary 更新TitUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TitUser true "更新TitUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /user/updateTitUser [put]
func UpdateTitUser(c *gin.Context) {
	var user model.TitUser
	_ = c.ShouldBindJSON(&user)
	err := service.UpdateTitUser(&user)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags TitUser
// @Summary 用id查询TitUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TitUser true "用id查询TitUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /user/findTitUser [get]
func FindTitUser(c *gin.Context) {
	var user model.TitUser
	_ = c.ShouldBindQuery(&user)
	err, reuser := service.GetTitUser(user.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"reuser": reuser}, c)
	}
}

// @Tags TitUser
// @Summary 分页获取TitUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取TitUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/getTitUserList [get]
func GetTitUserList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetTitUserInfoList(pageInfo)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(resp.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, c)
	}
}

// TitUserRegister, 注册人员信息至人才表
func TitUserRegister(c *gin.Context) {
	var rr request.TitUserRegister
	c.ShouldBindJSON(&rr)
	error := utils.Verify(rr, utils.Rules{
		"UserName":         {utils.NotEmpty()},
		"Gender":           {utils.NotEmpty()},
		"VerificationCode": {utils.NotEmpty()},
		"Birthday":         {utils.NotEmpty()},
		"Telphone":         {utils.NotEmpty()},
	})
	if error != nil {
		response.FailWithMessage(error.Error(), c)
		return
	}
	// 验证 验证码是否正确
	cmd := global.GVA_REDIS.Get(rr.Telphone)
	if cmd.Err() == nil && cmd.Val() == rr.VerificationCode {
		user := model.TitUser{Username: rr.Username, Telphone: rr.Telphone, Gender: rr.Gender, Birthday: rr.Birthday, IpAddress: c.ClientIP()}
		if err := service.CreateTitUser(user); err != nil {
			global.GVA_LOG.Error("用户注册失败了", err)
			response.FailWithMessage("手机号已存在", c)
		} else {
			if err, u := service.FindTitUserByPhone(rr.Telphone); err == nil {
				generateTitUserToke(c, u)
			} else {
				response.Fail(c)
			}
		}
	} else {
		response.FailWithMessage("验证码不正确", c)
	}
}

// TitUserLogin, 人才在手机端通过手机号登录
func TitUserLogin(c *gin.Context) {
	var lr request.TitUserLogin
	_ = c.ShouldBindJSON(&lr)
	error := utils.Verify(lr, utils.Rules{"VerificationCode": {utils.NotEmpty()}, "Telphone": {utils.NotEmpty()}})
	if error != nil {
		response.FailWithMessage(error.Error(), c)
		return
	}
	// 验证 手机号/验证码 是否正确
	// 验证 验证码是否正确
	cmd := global.GVA_REDIS.Get(lr.Telphone)
	if cmd.Err() == nil && cmd.Val() == lr.VerificationCode {
		if err, u := service.FindTitUserByPhone(lr.Telphone); err == nil {
			generateTitUserToke(c, u)
		} else {
			response.Fail(c)
		}
	} else {
		response.FailWithMessage("验证码不正确", c)
	}
}

// GetVerificationCode, 登录或注册时获取验证码逻辑
func GetVerificationCode(c *gin.Context) {
	// 如果是注册 需要首先对手机号判断是否注册验证
	var param struct {
		Phone   string `json:"phone"`
		IsLogin bool   `json:"isLogin"`
	}
	if e := c.BindJSON(&param); e != nil {
		global.GVA_LOG.Error("绑定参数失败了", e)
		response.FailWithMessage("参数错误", c)
		return
	}

	if param.IsLogin {
		// 校验是否已经注册
		if err, _ := service.FindTitUserByPhone(param.Phone); err != nil {
			response.FailWithMessage("请先注册", c)
			return
		}
	} else {
		if err, _ := service.FindTitUserByPhone(param.Phone); err == nil {
			response.FailWithMessage("改手机号已经注册，请直接登录！", c)
			return
		}
	}
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	error := utils.SendShotMessage(param.Phone, code)
	if error != nil {
		response.FailWithMessage("验证码发送失败了", c)
	} else {
		global.GVA_REDIS.Set(param.Phone, code, time.Minute*15)
		global.GVA_LOG.Info("%s 的验证码下发成功 code= %s", param.Phone, code)
		response.OkWithMessage("验证码发送成功", c)
	}
}

// 登录以后签发jwt
func generateTitUserToke(c *gin.Context, user model.TitUser) {
	j := &middleware.JWT{
		SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey), // 唯一签名
	}
	clams := request.TitUserClaims{
		ID:       user.ID,
		Telphone: user.Telphone,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,       // 签名生效时间
			ExpiresAt: time.Now().Unix() + 60*60*24*7, // 过期时间 一周
			Issuer:    "wq",                           // 签名的发行者
		},
	}
	token, err := j.CreateTitUserToken(clams)
	if err != nil {
		response.FailWithMessage("获取token失败", c)
		return
	}
	// 设置 用户 token 至 redis
	if err := global.GVA_REDIS.Set(fmt.Sprintf("tit:token:%s", user.Telphone), token, time.Hour*24*7).Err(); err != nil {
		response.FailWithMessage("设置登录状态失败", c)
		return
	} else {
		response.OkWithData(resp.TitLoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: clams.StandardClaims.ExpiresAt * 1000,
		}, c)
	}
}

func TitUser(c *gin.Context) {
	claims, _ := c.Get("claims")
	currentUser := claims.(*request.TitUserClaims)
	err, tu := service.GetTitUser(currentUser.ID)
	if err != nil {
		global.GVA_LOG.Error("查询用户信息出错", err)
		response.FailWithMessage("用户不存在", c)
	} else {
		response.OkWithData(tu, c)
	}
}
