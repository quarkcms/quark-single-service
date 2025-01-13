package request

// 商品列表查询
type ItemIndexQueryReq struct {
	PageReq
	CategoryId      int    `query:"category_id"`                    // 商品分类id：categoryies表中type为ITEM的分类
	ItemNameKeyword string `query:"item_name_keyword"`              // 模糊搜索：支持商品名称和关键字
	OrderByColumn   string `query:"order_by_column" default:"sort"` // 排序字段：默认sort asc排序，支持：sort、price、sales
	IsAsc           bool   `query:"is_asc" default:"true"`          // 是否正序：默认true
}
