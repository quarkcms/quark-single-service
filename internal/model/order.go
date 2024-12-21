package model

import (
	"github.com/quarkcloudio/quark-go/v3/utils/datetime"
)

// Order 定义订单表的结构体
type Order struct {
	ID                     uint              `json:"id" gorm:"primaryKey;autoIncrement;comment:'订单ID'"`                             // 订单ID
	OrderID                string            `json:"order_id" gorm:"not null;uniqueIndex;comment:'订单号'"`                            // 订单号
	UID                    uint              `json:"uid" gorm:"not null;comment:'用户id';index:uid"`                                  // 用户id
	Realname               string            `json:"realname" gorm:"not null;comment:'用户姓名'"`                                       // 用户姓名
	UserPhone              string            `json:"user_phone" gorm:"not null;comment:'用户电话'"`                                     // 用户电话
	UserAddress            string            `json:"user_address" gorm:"not null;comment:'详细地址'"`                                   // 详细地址
	TotalNum               uint              `json:"total_num" gorm:"not null;default:0;comment:'订单商品总数'"`                          // 订单商品总数
	TotalPrice             float64           `json:"total_price" gorm:"not null;default:0.00;comment:'订单总价'"`                       // 订单总价
	PayPrice               float64           `json:"pay_price" gorm:"not null;default:0.00;comment:'实际支付金额'"`                       // 实际支付金额
	Paid                   uint8             `json:"paid" gorm:"not null;default:0;comment:'支付状态'"`                                 // 支付状态
	PayTime                datetime.Datetime `json:"pay_time" gorm:"comment:'支付时间'"`                                                // 支付时间
	PayType                string            `json:"pay_type" gorm:"not null;comment:'支付方式'"`                                       // 支付方式
	Status                 uint8             `json:"status" gorm:"not null;default:0;comment:'订单状态（0：待发货；1：待收货；2：已收货，待评价；3：已完成；）'"` // 订单状态
	RefundStatus           uint8             `json:"refund_status" gorm:"not null;default:0;comment:'0 未退款 1 申请中 2 已退款 3 退款中'"`     // 退款状态
	RefundReasonWapImg     string            `json:"refund_reason_wap_img" gorm:"comment:'退款图片'"`                                   // 退款图片
	RefundReasonWapExplain string            `json:"refund_reason_wap_explain" gorm:"comment:'退款用户说明'"`                             // 退款用户说明
	RefundReasonWap        string            `json:"refund_reason_wap" gorm:"comment:'前台退款原因'"`                                     // 前台退款原因
	RefundReason           string            `json:"refund_reason" gorm:"comment:'不退款的理由'"`                                         // 不退款的理由
	RefundReasonTime       datetime.Datetime `json:"refund_reason_time" gorm:"comment:'退款时间'"`                                      // 退款时间
	RefundPrice            float64           `json:"refund_price" gorm:"not null;default:0.00;comment:'退款金额'"`                      // 退款金额
	Remark                 string            `json:"remark" gorm:"comment:'管理员备注'"`                                                 // 管理员备注
	MerID                  uint              `json:"mer_id" gorm:"not null;default:0;comment:'商户ID'"`                               // 商户ID
	IsMerCheck             uint8             `json:"is_mer_check" gorm:"not null;default:0"`                                        // 是否商户审核
	Cost                   float64           `json:"cost" gorm:"not null;comment:'成本价'"`                                            // 成本价
	VerifyCode             string            `json:"verify_code" gorm:"not null;default:'';comment:'核销码'"`                          // 核销码
	ShippingType           uint8             `json:"shipping_type" gorm:"not null;default:1;comment:'配送方式 1=快递 ，2=门店自提'"`           // 配送方式
	ClerkID                uint              `json:"clerk_id" gorm:"not null;default:0;comment:'店员id/核销员id'"`                       // 店员id/核销员id
	OutTradeNo             string            `json:"out_trade_no" gorm:"comment:'商户系统内部的订单号,32个字符内、可包含字母'"`                         // 商户系统内部的订单号
	CreatedAt              datetime.Datetime `json:"created_at"`
	UpdatedAt              datetime.Datetime `json:"updated_at"` // 记录更新时间
}

// Seeder
func (m *Order) Seeder() {

}
