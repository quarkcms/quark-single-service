package model

import "github.com/quarkcloudio/quark-go/v3/utils/datetime"

// OrderDetail 定义订单购物详情表的结构体
type OrderDetail struct {
	Id          uint              `json:"id" gorm:"primaryKey;autoIncrement;comment:'主键'"`                   // 主键
	OrderId     uint              `json:"order_id" gorm:"not null;comment:'订单id';uniqueIndex:oid"`           // 订单id
	ItemId      uint              `json:"item_id" gorm:"not null;default:0;comment:'商品ID';index:product_id"` // 商品ID
	Info        string            `json:"info" gorm:"type:text;not null;comment:'购买东西的详细信息'"`                // 购买东西的详细信息
	Unique      string            `json:"unique" gorm:"not null;comment:'唯一id';uniqueIndex:oid"`             // 唯一id
	OrderNo     string            `json:"order_no" gorm:"not null;comment:'订单号'"`                            // 订单号
	Name        string            `json:"name" gorm:"not null;comment:'商品名称'"`                               // 商品名称
	AttrValueId *uint             `json:"attr_value_id" gorm:"unsigned;default:null;comment:'规格属性值id'"`      // 规格属性值id
	Image       string            `json:"image" gorm:"not null;comment:'商品图片'"`                              // 商品图片
	SKU         string            `json:"sku" gorm:"not null;comment:'商品sku'"`                               // 商品sku
	Price       float64           `json:"price" gorm:"unsigned;not null;comment:'商品价格'"`                     // 商品价格
	PayNum      uint              `json:"pay_num" gorm:"unsigned;not null;default:0;comment:'购买数量'"`         // 购买数量
	CreatedAt   datetime.Datetime `json:"created_at"`
	UpdatedAt   datetime.Datetime `json:"updated_at"` // 记录更新时间
}
