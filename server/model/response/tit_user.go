package response

import "gin-vue-admin/model"

type TitLoginResponse struct {
	User      model.TitUser `json:"user"`
	Token     string        `json:"token"`
	ExpiresAt int64         `json:"expiresAt"`
}
