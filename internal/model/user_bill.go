package model

import "github.com/quarkcloudio/quark-go/v3/utils/datetime"

// UserBill 定义用户账单表的结构体
type UserBill struct {
	Id        uint              `json:"id" gorm:"primaryKey;autoIncrement;comment:'用户账单id'"`               // 用户账单id
	Uid       uint              `json:"uid" gorm:"not null;default:0;comment:'用户uid';index:openid"`        // 用户uid
	LinkId    string            `json:"link_id" gorm:"not null;default:'0';comment:'关联id'"`                // 关联id
	PM        uint8             `json:"pm" gorm:"not null;default:0;comment:'0 = 支出 1 = 获得'"`              // 0 = 支出 1 = 获得
	Title     string            `json:"title" gorm:"not null;default:'';comment:'账单标题'"`                   // 账单标题
	Category  string            `json:"category" gorm:"not null;default:'';comment:'明细种类'"`                // 明细种类
	Type      string            `json:"type" gorm:"not null;default:'';comment:'明细类型'"`                    // 明细类型
	Number    float64           `json:"number" gorm:"unsigned;not null;default:0.00;comment:'明细数字'"`       // 明细数字
	Balance   float64           `json:"balance" gorm:"unsigned;not null;default:0.00;comment:'剩余'"`        // 剩余
	Mark      string            `json:"mark" gorm:"not null;default:'';comment:'备注'"`                      // 备注
	Status    int8              `json:"status" gorm:"not null;default:1;comment:'0 = 带确定 1 = 有效 -1 = 无效'"` // 0 = 带确定 1 = 有效 -1 = 无效
	CreatedAt datetime.Datetime `json:"created_at"`
	UpdatedAt datetime.Datetime `json:"updated_at"` // 记录更新时间
}
