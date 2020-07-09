package response

import (
	"gin-vue-admin/model"
	"github.com/jinzhu/gorm"
	"time"
)

// TitLoginResponse, 用户在web 端登录时，返回的数据
type TitLoginResponse struct {
	User      model.TitUser `json:"user"`
	Token     string        `json:"token"`
	ExpiresAt int64         `json:"expiresAt"`
}

// TitUserInfoForBackend， 用户在 后台 查询时 返回
type TitUserInfoForBackend struct {
	gorm.Model
	Username        string    `json:"username" form:"username" `
	Gender          int       `json:"gender" form:"gender" `
	Birthday        time.Time `json:"birthday" form:"birthday" `
	Telphone        string    `json:"telphone" form:"telphone" `
	OpenId          string    `json:"openId" form:"openId" `
	ChangeJobOption int       `json:"changeJobOption" form:"changeJobOption" `
}
