package dto

type ItemAttrItemDTO struct {
	Name string `json:"name"` // 属性值
}

type ItemAttrDTO struct {
	Name  string            `json:"name"` // 属性名
	Items []ItemAttrItemDTO `json:"itemAttrItem"`
}

// ItemAttrValueDTO 定义商品属性值表的结构体
type ItemAttrValueDTO struct {
	Id        int         `json:"id"`         // 主键
	ItemId    int         `json:"item_id"`    // 商品ID
	Suk       string      `json:"suk"`        // 商品属性索引值
	Stock     int         `json:"stock"`      // 属性对应的库存
	Sales     int         `json:"sales"`      // 销量，默认为0
	Price     float64     `json:"price"`      // 属性金额
	Image     interface{} `json:"image"`      // 图片
	Cost      float64     `json:"cost"`       // 成本价，默认为0.00
	OtPrice   float64     `json:"ot_price"`   // 原价，默认为0.00
	AttrValue interface{} `json:"attr_value"` // attr_values 创建更新时的属性对应
	IsDefault bool        `json:"is_default"` // 是否默认
	Status    bool        `json:"status"`     // 是否上架
}
