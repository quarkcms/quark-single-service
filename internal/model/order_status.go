package model

import (
	"github.com/quarkcloudio/quark-go/v3/utils/datetime"
)

// OrderStatus 定义订单操作记录表的结构体
type OrderStatus struct {
	OrderId       int               `json:"order_id" gorm:"not null;comment:订单id;index:order_id"`                                                                                                                                                                                                   // 订单id
	ChangeType    string            `json:"change_type" gorm:"not null;comment:操作类型（create_order:订单生成,pay_success:用户付款成功;delivery_goods:已发货 快递公司：圆通速递 快递单号：YT46466545445555;take_delivery:已收货;check_order_over:用户评价;apply_refund:用户申请退款，原因：收货地址填错了;refund_price:退款给用户：124.63元;）;index:change_type"` // 操作类型
	ChangeMessage string            `json:"change_message" gorm:"not null;comment:操作备注"`                                                                                                                                                                                                            // 操作备注
	CreatedAt     datetime.Datetime `json:"created_at"`                                                                                                                                                                                                                                             // 创建时间
	UpdatedAt     datetime.Datetime `json:"updated_at"`                                                                                                                                                                                                                                             // 更新时间
}
