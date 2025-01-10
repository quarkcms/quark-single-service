package model

import (
	"github.com/quarkcloudio/quark-go/v3/utils/datetime"
)

// BillRecord 账单记录表的结构体
type BillRecord struct {
	Id            int               `json:"id" gorm:"primaryKey;autoIncrement;comment:账单id"`                 // 账单记录id
	Title         string            `json:"title" gorm:"not null;default:'';comment:账单标题，例如：日账单、周账单、月账单"`    // 账单标题
	Day           string            `json:"day" gorm:"not null;default:'';comment:日期"`                       // 日期
	Type          int8              `json:"type" gorm:"not null;default:1;comment:账单类型(1:日账单,2:周账单,3:月账单)"`  // 账单类型
	EntryPrice    float64           `json:"entry_price" gorm:"unsigned;not null;default:0.00;comment:收入金额"`  // 收入金额
	ExpPrice      float64           `json:"exp_price" gorm:"unsigned;not null;default:0.00;comment:支出金额"`    // 支出金额
	IncomePrice   float64           `json:"income_price" gorm:"unsigned;not null;default:0.00;comment:入账金额"` // 入账金额
	StartDatetime datetime.Datetime `json:"start_datetime" gorm:"type:datetime(0);comment:账单周期开始时间"`         // 账单周期开始时间
	EndDatetime   datetime.Datetime `json:"end_datetime" gorm:"type:datetime(0);comment:账单周期结束时间"`           // 账单周期结束时间
	Status        int8              `json:"status" gorm:"not null;default:1;comment:状态"`                     // 状态
	CreatedAt     datetime.Datetime `json:"created_at" gorm:"type:datetime(0)"`
	UpdatedAt     datetime.Datetime `json:"updated_at" gorm:"type:datetime(0)"` // 记录更新时间
}
