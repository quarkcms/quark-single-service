package model

import (
	"github.com/quarkcloudio/quark-go/v3/utils/datetime"
)

// OrderStatus 定义订单操作记录表的结构体
type OrderStatus struct {
	OrderId       int               `json:"order_id" gorm:"not null;comment:订单id;index:order_id"`       // 订单id
	ChangeType    string            `json:"change_type" gorm:"not null;comment:操作类型;index:change_type"` // 操作类型
	ChangeMessage string            `json:"change_message" gorm:"not null;comment:操作备注"`                // 操作备注
	CreatedAt     datetime.Datetime `json:"created_at"`                                                 // 创建时间
	UpdatedAt     datetime.Datetime `json:"updated_at"`                                                 // 更新时间
}
