package model

// ItemCategory 定义商品分类辅助表的结构体
type ItemCategory struct {
	Id         int `json:"id" gorm:"primaryKey;autoIncrement"`                 // 主键ID
	ItemId     int `json:"item_id" gorm:"not null;default:0;comment:商品id"`     // 商品ID
	CategoryId int `json:"category_id" gorm:"not null;default:0;comment:分类id"` // 分类ID
}
