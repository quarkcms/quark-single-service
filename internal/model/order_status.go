package model

import (
	"github.com/quarkcloudio/quark-go/v3/utils/datetime"
)

// OrderStatus 定义订单操作记录表的结构体
type OrderStatus struct {
	OID           uint              `json:"oid" gorm:"not null;comment:'订单id';index:oid"`                         // 订单id
	ChangeType    string            `json:"change_type" gorm:"not null;comment:'操作类型';index:change_type"`         // 操作类型
	ChangeMessage string            `json:"change_message" gorm:"not null;comment:'操作备注'"`                        // 操作备注
	CreateTime    datetime.Datetime `json:"create_time" gorm:"not null;default:CURRENT_TIMESTAMP;comment:'操作时间'"` // 操作时间
}

// Seeder
func (m *OrderStatus) Seeder() {

}
