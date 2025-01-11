package dto

type ItemDTO struct {
	Id          int            `json:"id"`           // 商品id
	MerchantId  int            `json:"merchant_id"`  // 商户Id(0为总后台管理员创建,不为0的时候是商户后台创建)
	Image       string         `json:"image"`        // 商品图片
	SliderImage string         `json:"slider_image"` // 轮播图
	Name        string         `json:"name"`         // 商品名称
	Keyword     string         `json:"keyword"`      // 关键字
	Description string         `json:"description"`  // 商品简介
	CategoryIds string         `json:"category_ids"` // 分类ids
	Price       float64        `json:"price"`        // 商品价格
	OtPrice     float64        `json:"ot_price"`     // 市场价
	Sort        int16          `json:"sort"`         // 排序
	Sales       int            `json:"sales"`        // 销量
	Stock       int            `json:"stock"`        // 库存
	Status      uint8          `json:"status"`       // 状态(0:未上架,1:上架)
	Cost        float64        `json:"cost"`         // 成本价
	FictiSales  int            `json:"ficti_sales"`  // 虚拟销量
	Views       int            `json:"views"`        // 浏览量
	FictiViews  int            `json:"ficti_views"`  // 虚拟浏览量
	SpecType    uint8          `json:"spec_type"`    // 规格:0单,1多
	Attrs       []AttrDTO      `json:"attrs"`        // 商品属性
	AttrValues  []AttrValueDTO `json:"attr_values"`  // 商品属性值
}

// 管理后台解析用
type ItemAttrItemDTO struct {
	Name string `json:"name"` // 属性值
}

// 管理后台解析用
type ItemAttrDTO struct {
	Name  string            `json:"name"` // 属性名
	Items []ItemAttrItemDTO `json:"itemAttrItem"`
}

// 商品属性值表
type AttrValueDTO struct {
	Id        int         `json:"id"`         // 主键
	ItemId    int         `json:"item_id"`    // 商品ID
	Suk       string      `json:"suk"`        // 商品属性索引值
	Stock     int         `json:"stock"`      // 属性对应的库存
	Sales     int         `json:"sales"`      // 销量，默认为0
	Price     float64     `json:"price"`      // 属性金额
	Image     interface{} `json:"image"`      // 图片
	ImageJson interface{} `json:"-"`          // 图片，Json字符串
	Cost      float64     `json:"cost"`       // 成本价，默认为0.00
	OtPrice   float64     `json:"ot_price"`   // 原价，默认为0.00
	AttrValue interface{} `json:"attr_value"` // attr_values 创建更新时的属性对应
	IsDefault bool        `json:"is_default"` // 是否默认
	Status    bool        `json:"status"`     // 是否上架
}

// 商品属性
type AttrDTO struct {
	Id         int         `json:"id"`          // 主键
	ItemId     int         `json:"item_id"`     // 商品ID
	AttrName   string      `json:"attr_name"`   // 属性名
	AttrValues string      `json:"attr_values"` // 属性值,字符串形式
	AttrItems  interface{} `json:"attr_items"`  // 属性值,json形式
}
