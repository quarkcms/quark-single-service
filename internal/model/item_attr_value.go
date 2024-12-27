package model

// ItemAttrValue 定义商品属性值表的结构体
type ItemAttrValue struct {
	Id        uint    `json:"id" gorm:"primaryKey;autoIncrement;comment:主键"`                     // 主键
	ItemId    uint    `json:"item_id" gorm:"not null;comment:商品ID;index:store_id"`               // 商品ID
	Suk       string  `json:"suk" gorm:"not null;comment:商品属性索引值 (attr_value|attr_value[|....]"` // 商品属性索引值
	Stock     uint    `json:"stock" gorm:"not null;comment:属性对应的库存"`                             // 属性对应的库存
	Sales     uint    `json:"sales" gorm:"not null;default:0;comment:销量"`                        // 销量，默认为0
	Price     float64 `json:"price" gorm:"not null;comment:属性金额"`                                // 属性金额
	Image     string  `json:"image" gorm:"comment:图片"`                                           // 图片
	Cost      float64 `json:"cost" gorm:"not null;default:0.00;comment:成本价"`                     // 成本价，默认为0.00
	OtPrice   float64 `json:"ot_price" gorm:"not null;default:0.00;comment:原价"`                  // 原价，默认为0.00
	AttrValue string  `json:"attr_value" gorm:"type:text;comment:attr_values 创建更新时的属性对应"`        // attr_values 创建更新时的属性对应
}
