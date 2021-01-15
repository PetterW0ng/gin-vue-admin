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
	"net/http"
	"qiniupkg.com/x/log.v7"
	"strconv"
	"strings"
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

func GetTitUserBaseInfo(c *gin.Context) {
	titUserId, _ := strconv.Atoi(c.Query("ID"))
	if err, tu := service.GetTitUser(uint(titUserId)); err == nil {
		if err, tub := service.GetTitUserBaseinfo(uint(tu.TitUserBaseinfoId)); err == nil {
			dictGroup := service.GetGroupedDict()
			var area string
			if len(tub.Area) > 0 {
				area = service.GetAreaById(tub.Area).DisplayName
			} else {
				area = ""
			}
			var benefits, childType, trainingTimes = strings.Split(tub.Benefits, ","), strings.Split(tub.ChildType, ","), strings.Split(tub.TrainingNumber, ",")
			var benefitsStr, childTypeStr, trainingTimesStr string
			for inx, v := range benefits {
				i, _ := strconv.Atoi(v)
				if inx == len(benefits)-1 {
					benefitsStr += dictGroup["benefits"][i]
				} else {
					benefitsStr += dictGroup["benefits"][i] + "、"
				}
			}
			for inx, v := range trainingTimes {
				i, _ := strconv.Atoi(v)
				if inx == len(benefits)-1 {
					trainingTimesStr += dictGroup["training_number"][i]
				} else {
					trainingTimesStr += dictGroup["training_number"][i] + "、"
				}

			}
			for inx, v := range childType {
				i, _ := strconv.Atoi(v)
				if inx == len(benefits)-1 {
					childTypeStr += dictGroup["child_type"][i]
				} else {
					childTypeStr += dictGroup["child_type"][i] + "、"
				}
			}

			tubtR := resp.TitUserBaseinfoTraining{
				School:           tub.School, // int 转 string
				MajorsStudied:    dictGroup["study_major"][tub.MajorsStudied],
				HighestEducation: dictGroup["highest_education"][tub.HighestEducation],
				SchoolSystem:     dictGroup["school_system"][tub.SchoolSystem],
				WorkingState:     dictGroup["working_state"][tub.WorkingState],
				Company:          tub.Company,
				Area:             area,
				JobTitle:         tub.JobTitle,

				ServiceType: dictGroup["service_type"][tub.ServiceType],
				Income:      dictGroup["income"][tub.Income],
				Benefits:    benefitsStr,
				ChildType:   childTypeStr,
				ChildAge:    dictGroup["child_age"][tub.ChildAge],

				TrainingNumber: trainingTimesStr,
				TrainingFee:    dictGroup["training_fee"][tub.TrainingFee],
			}
			// 根据 TitUserBaseinfoId 查询 training 信息
			trainingInfos := service.QueryTrainingInfoByBaseId(tu.TitUserBaseinfoId)
			var trainingInfosReturn []resp.TrainingInfo
			for _, item := range trainingInfos {
				trainingInfosReturn = append(trainingInfosReturn, resp.TrainingInfo{
					PaymentWay:     dictGroup["payment_way"][item.PaymentWay],
					TrainingCourse: item.TrainingCourse,
					BeginTime:      item.BeginTime,
					EndTime:        item.EndTime,
				})
			}
			tubtR.TrainingInfos = trainingInfosReturn
			response.OkWithData(gin.H{"userBaseinfo": tubtR}, c)
			return
		} else {
			response.FailWithMessage("该用户没有填写基本信息", c)
		}
	} else {
		response.FailWithMessage(fmt.Sprintf("用户不存在或已删除，%v", err), c)
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
				GenerateTitUserToke(c, u)
			} else {
				response.Fail(c)
			}
		}
	} else {
		response.FailWithMessage("验证码不正确", c)
	}
}

// 微信服务回调的接口 获取openId 和 token
func DoWxCallback(c *gin.Context) {
	// 获取 code , state
	code := c.Query("code")
	state := c.Query("state")
	var redirectUrl string
	var weixinServer service.WeixinServer
	if server, err := service.WxServerNew(code); err == nil {
		weixinServer = server
		openId := weixinServer.GetOpenId()
		wxToken := weixinServer.GetToken()
		log.Debug("openId")
		// 判断是否已经绑定
		if err, wxUser := service.GetSysWxUserByOpenId(openId); err != nil {
			// 微信用户不存在， 录入微信用户信息
			if wxUserInfo, err := weixinServer.GetUserInfo(); err == nil {
				service.CreateSysWxUser(wxUserInfo)
			} else {
				log.Error("微信接口回调出错了", err)
			}
			redirectUrl = fmt.Sprintf(global.GVA_CONFIG.WeiXin.RedirectUrl, openId, state, wxToken, 0)
			c.Redirect(http.StatusMovedPermanently, redirectUrl)
			return
		} else {
			if err, u := service.FindTitUserByPhone(wxUser.Phone); err == nil {
				if token, err := createUserToken(u); err == nil {
					redirectUrl = fmt.Sprintf(global.GVA_CONFIG.WeiXin.RedirectUrl, openId, 0, 0, token)
					c.Redirect(http.StatusMovedPermanently, redirectUrl)
					return
				} else {
					log.Error("微信接口回调出错了", err)
				}
			} else {
				log.Warn("根据手机号查找时 titUser 出错了", wxUser.Phone, err)
			}
			redirectUrl = fmt.Sprintf(global.GVA_CONFIG.WeiXin.RedirectUrl, openId, state, wxToken, 0)
			c.Redirect(http.StatusMovedPermanently, redirectUrl)
		}
	} else {
		log.Error("微信接口回调出错了", err)
		response.FailWithMessage("接口出错了, 请稍后再试！", c)
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
			GenerateTitUserToke(c, u)
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

// GetVerificationCode, 登录或注册时获取验证码逻辑
func GetVerificationCode2(c *gin.Context) {
	// 如果是注册 需要首先对手机号判断是否注册验证
	var param struct {
		Phone string `json:"phone"`
	}
	if e := c.BindJSON(&param); e != nil {
		global.GVA_LOG.Error("绑定参数失败了", e)
		response.FailWithMessage("参数错误", c)
		return
	}

	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	error := utils.SendShotMessage(param.Phone, code)
	if error != nil {
		response.FailWithMessage("验证码发送失败了", c)
	} else {
		global.GVA_REDIS.Set("code:"+param.Phone, code, time.Minute*15)
		global.GVA_LOG.Info("%s 的验证码下发成功 code= %s", param.Phone, code)
		response.OkWithMessage("验证码发送成功", c)
	}
}

func createUserToken(user model.TitUser) (token string, err error) {
	j := &middleware.JWT{
		SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey), // 唯一签名
	}
	clams := request.TitUserClaims{
		ID:       user.ID,
		Telphone: user.Telphone,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,        // 签名生效时间
			ExpiresAt: time.Now().Unix() + 60*60*24*15, // 过期时间 两周
			Issuer:    "wq",                            // 签名的发行者
		},
	}
	token, err = j.CreateTitUserToken(clams)
	return
}

// 登录以后签发jwt
func GenerateTitUserToke(c *gin.Context, user model.TitUser) {
	token, err := createUserToken(user)
	if err != nil {
		response.FailWithMessage("获取token失败", c)
		return
	}
	// 设置用户 token 至 redis ，可以把用户拉入黑名单
	if err := global.GVA_REDIS.Set(fmt.Sprintf("tit:token:%s", user.Telphone), token, time.Hour*24*15).Err(); err != nil {
		response.FailWithMessage("设置登录状态失败", c)
		return
	} else {
		response.OkWithData(resp.TitLoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: time.Now().Unix() + 60*60*24*15*1000,
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

func GetStudyReport(c *gin.Context) {
	openId := c.Query("openId")
	_, wxUser := service.GetSysWxUserByOpenId(openId)
	generateStudyReport(c, wxUser.Phone)
}

func generateStudyReport(c *gin.Context, phone string) {
	customerStudy := resp.CustomerStudy{}
	err, customer := service.GetSysCustomerByPhone(phone)
	if err == nil {
		// 查找xiaoe订单, 筛选
		if err, xetOrderArray := service.QueryUserOrder(customer.EID); err == nil {
			xetOrder := make([]resp.XeOrder, len(xetOrderArray))
			for i, item := range xetOrderArray {
				xetOrder[i] = resp.XeOrder{item.PurchaseName, item.PayTime, item.Price}
			}
			customerStudy.XeOrderArray = xetOrder
		}
	}
	if err, sCusScore := service.QuerySysCusScore(customer.EID); err == nil {
		customerStudy.Duration = sCusScore.Duration / 60
		customerStudy.SequenceA = sCusScore.SequenceA
		customerStudy.SequenceB = sCusScore.SequenceB
		customerStudy.SequenceC = sCusScore.SequenceC
		customerStudy.SequenceD = sCusScore.SequenceD
		customerStudy.CourseNum = sCusScore.CourseNum
		customerStudy.StudyRank = sCusScore.StudyRank
	}
	if err, holderArray := service.GetCertByPhone(customer.Phone); err == nil {
		certArray := make([]resp.Cert, len(holderArray))
		for i, item := range holderArray {
			certArray[i] = resp.Cert{item.CertificateName, item.IssueTime}
		}
		customerStudy.CertArray = certArray
	}
	// todo 证书排名
	customerStudy.CertNumRank = "12%"
	response.OkWithData(customerStudy, c)
}

func MyStudyReport(c *gin.Context) {
	claims, _ := c.Get("claims")
	currentUser := claims.(*request.TitUserClaims)
	generateStudyReport(c, currentUser.Telphone)
}
