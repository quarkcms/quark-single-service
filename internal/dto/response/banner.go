package response

// 列表
type BannerListResp struct {
	Id         int    `json:"id"`
	CategoryId int    `json:"category_id"`
	Title      string `json:"title"`
	UrlType    int    `json:"url_type"`
	Url        string `json:"url"`
	CoverId    string `json:"cover_id"`
}
