package model

// ItemAttr 定义商品属性表的结构体
type ItemAttr struct {
	Id         uint   `json:"id" gorm:"primaryKey;autoIncrement;comment:'主键'"`                // 主键
	ItemId     uint   `json:"item_id" gorm:"not null;default:0;comment:'商品ID';index:item_id"` // 商品ID
	AttrName   string `json:"attr_name" gorm:"not null;comment:'属性名'"`                        // 属性名
	AttrValues string `json:"attr_values" gorm:"not null;comment:'属性值'"`                      // 属性值
}
