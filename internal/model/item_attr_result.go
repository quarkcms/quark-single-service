package model

// ItemAttrResult 定义商品属性详情表的结构体
type ItemAttrResult struct {
	ID         uint   `json:"id" gorm:"primaryKey;autoIncrement;comment:'主键'"`       // 主键
	ItemID     uint   `json:"item_id" gorm:"not null;comment:'商品ID';index:item_id"`  // 商品ID
	Result     string `json:"result" gorm:"type:longtext;not null;comment:'商品属性参数'"` // 商品属性参数
	ChangeTime uint   `json:"change_time" gorm:"not null;comment:'上次修改时间'"`          // 上次修改时间
}

// Seeder
func (m *ItemAttrResult) Seeder() {

}
