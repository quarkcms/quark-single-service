package response

// 商品分类
type ItemCategoryResp struct {
	Id       int                `json:"id"`
	Pid      int                `json:"pid"`
	Title    string             `json:"title"`
	CoverId  string             `json:"cover_id,omitempty"`
	Children []ItemCategoryResp `json:"children,omitempty" gorm:"-"`
}

// 商品列表
type ItemIndexResp struct {
	Id         int     `json:"id"`          // 商品id
	Name       string  `json:"name"`        // 商品名称
	Image      string  `json:"image"`       // 商品图片
	Price      float64 `json:"price"`       // 商品价格
	FictiSales int     `json:"ficti_sales"` // 商品虚拟销量
}
