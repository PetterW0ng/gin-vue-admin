package request

import "time"

type TitUserRegister struct {
	Username         string    `json:"username" form:"username" `
	Gender           int       `json:"gender" form:"gender" `
	Birthday         time.Time `json:"birthday" form:"birthday" `
	Telphone         string    `json:"telphone" form:"telphone" `
	VerificationCode string    `json:"verificationCode" form:"verificationCode" `
}

type TitUserLogin struct {
	Telphone         string `json:"telphone" form:"telphone" `
	VerificationCode string `json:"verificationCode" form:"verificationCode" `
}
