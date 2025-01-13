package request

// 分页
type PageReq struct {
	Page     int `query:"page" default:"1"`
	PageSize int `query:"page_size" default:"10"`
}
