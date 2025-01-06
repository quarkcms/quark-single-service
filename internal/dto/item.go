package dto

type ItemDTO struct {
	Id          int     `json:"id" gorm:"primaryKey;autoIncrement;comment:商品id"`
	MerId       int     `json:"mer_id" gorm:"not null;default:0;comment:商户Id(0为总后台管理员创建,不为0的时候是商户后台创建)"`
	Image       string  `json:"image" gorm:"not null;size:256;comment:商品图片"`
	SliderImage string  `json:"slider_image" gorm:"not null;size:2000;comment:轮播图"`
	Name        string  `json:"name" gorm:"not null;size:128;comment:商品名称"`
	Keyword     string  `json:"keyword" gorm:"not null;size:256;comment:关键字"`
	Description string  `json:"description" gorm:"not null;size:256;comment:商品简介"`
	CategoryIds string  `json:"category_ids" gorm:"not null;size:64;comment:分类ids"`
	Price       float64 `json:"price" gorm:"not null;type:decimal(8,2);default:0.00;comment:商品价格"`
	OtPrice     float64 `json:"ot_price" gorm:"not null;type:decimal(8,2);default:0.00;comment:市场价"`
	Sort        int16   `json:"sort" gorm:"not null;default:0;comment:排序"`
	Sales       int     `json:"sales" gorm:"not null;default:0;comment:销量"`
	Stock       int     `json:"stock" gorm:"not null;default:0;comment:库存"`
	Status      uint8   `json:"status" gorm:"not null;default:1;comment:状态(0:未上架,1:上架)"`
	Cost        float64 `json:"cost" gorm:"not null;type:decimal(8,2);default:0.00;comment:成本价"`
	Ficti       int     `json:"ficti" gorm:"default:100;comment:虚拟销量"`
	Views       int     `json:"views" gorm:"default:0;comment:浏览量"`
	SpecType    uint8   `json:"spec_type" gorm:"not null;default:0;comment:规格:0单,1多"`
}

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
