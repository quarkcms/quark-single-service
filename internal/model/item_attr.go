package model

import (
	"github.com/quarkcloudio/quark-go/v3/utils/datetime"
	"gorm.io/gorm"
)

// ItemAttr 定义商品属性表的结构体
type ItemAttr struct {
	Id         int               `json:"id" gorm:"primaryKey;autoIncrement;comment:主键"`                // 主键
	ItemId     int               `json:"item_id" gorm:"not null;default:0;comment:商品ID;index:item_id"` // 商品ID
	AttrName   string            `json:"attr_name" gorm:"not null;comment:属性名"`                        // 属性名
	AttrValues string            `json:"attr_values" gorm:"not null;comment:属性值,字符串形式"`                // 属性值,字符串形式
	AttrItems  string            `json:"attr_items" gorm:"not null;comment:属性值,json形式"`                // 属性值,json形式
	CreatedAt  datetime.Datetime `json:"created_at"`
	UpdatedAt  datetime.Datetime `json:"updated_at"`
	DeletedAt  gorm.DeletedAt    `json:"deleted_at"`
}
