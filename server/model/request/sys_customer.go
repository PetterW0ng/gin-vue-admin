// 自动生成模板SysCustomer
package request

import (
	"time"
)

type PaginatedSysCustomer struct {
	RegisterTimeBegin time.Time `json:"registerTimeBegin" form:"registerTimeBegin" `
	RegisterTimeEnd   time.Time `json:"registerTimeEnd" form:"registerTimeEnd" `
	Gender            int       `json:"gender" form:"gender" `
	Birthday          time.Time `json:"birthday" form:"birthday" `
	Source            int       `json:"source" form:"source" `
	EntryPoint        int       `json:"entryPoint" form:"entryPoint" `
	IsEvaluate        int       `json:"isEvaluate" form:"isEvaluate" `
	Query             string    `json:"queryStr" form:"queryStr"`
	PageInfo
}
