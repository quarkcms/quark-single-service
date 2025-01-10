package model

import "github.com/quarkcloudio/quark-go/v3/utils/datetime"

// OrderDetail 定义订单购物详情表的结构体
type OrderDetail struct {
	Id          int               `json:"id" gorm:"primaryKey;autoIncrement;comment:主键"`                // 主键
	OrderId     int               `json:"order_id" gorm:"not null;comment:订单id;index:order_id(32)"`     // 订单id
	ItemId      int               `json:"item_id" gorm:"not null;default:0;comment:商品ID;index:item_id"` // 商品ID
	OrderNo     string            `json:"order_no" gorm:"size:32;not null;comment:订单号"`                 // 订单号
	Name        string            `json:"name" gorm:"not null;comment:商品名称"`                            // 商品名称
	Content     string            `json:"content" gorm:"type:text;default:null;comment:商品详情"`           // 商品详情
	AttrValueId int               `json:"attr_value_id" gorm:"unsigned;default:null;comment:规格属性值id"`   // 规格属性值id
	Image       string            `json:"image" gorm:"not null;comment:商品图片"`                           // 商品图片
	SKU         string            `json:"sku" gorm:"not null;comment:商品sku"`                            // 商品sku
	Price       float64           `json:"price" gorm:"unsigned;not null;comment:商品价格"`                  // 商品价格
	PayNum      int               `json:"pay_num" gorm:"unsigned;not null;default:0;comment:购买数量"`      // 购买数量
	CreatedAt   datetime.Datetime `json:"created_at"`
	UpdatedAt   datetime.Datetime `json:"updated_at"` // 记录更新时间
}
