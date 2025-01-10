package request

// 分页
type PageReq struct {
	Page     int `query:"page" default:"1"`
	PageSize int `query:"pageSize" default:"10"`
}
