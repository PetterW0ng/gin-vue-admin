package utils

import (
	"testing"

	_ "gin-vue-admin/config"
)

func TestSendSMS(t *testing.T) {
	SendShotMessage("13718320428", "123456")
}
