package response

// 商品分类
type ItemCategory struct {
	Id       int            `json:"id"`
	Pid      int            `json:"pid"`
	Title    string         `json:"title"`
	CoverId  string         `json:"coverId,omitempty"`
	Children []ItemCategory `json:"children,omitempty" gorm:"-"`
}
