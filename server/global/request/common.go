package request

// Paging common input parameter structure
type PageInfo struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"pageSize" form:"pageSize"`
}

// Find by id structure
type GetById struct {
	Id float64 `form:"id" form:"id"`
}

type IdsReq struct {
	Ids []int `form:"ids" form:"ids"`
}
