package response

// 列表
type BannerListResp struct {
	Id         int    `json:"id"`
	CategoryId int    `json:"categoryId"`
	Title      string `json:"title"`
	UrlType    int    `json:"urlType"`
	Url        string `json:"url"`
	CoverId    string `json:"coverId"`
}
