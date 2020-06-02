package request

// Paging common input parameter structure
type PageInfo struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"pageSize" form:"pageSize"`
}

type PaginatedQuery struct {
	PageInfo
	Query string `json:"userName" form:"userName"`
}

// Find by id structure
type GetById struct {
	Id float64 `json:"id" form:"id"`
}
