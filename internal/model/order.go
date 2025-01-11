package model

import (
	"github.com/quarkcloudio/quark-go/v3/dal/db"
	appmodel "github.com/quarkcloudio/quark-go/v3/model"
	"github.com/quarkcloudio/quark-go/v3/service"
	"github.com/quarkcloudio/quark-go/v3/utils/datetime"
	"gorm.io/gorm"
)

// Order 定义订单表的结构体
type Order struct {
	Id                    int               `json:"id" gorm:"primaryKey;autoIncrement;comment:订单ID"`                                            // 订单ID
	OrderNo               string            `json:"order_no" gorm:"size:32;not null;uniqueIndex:order_no(32);comment:订单号"`                      // 订单号
	Uid                   int               `json:"uid" gorm:"not null;comment:用户id;index:uid"`                                                 // 用户id
	Realname              string            `json:"realname" gorm:"not null;comment:用户姓名"`                                                      // 用户姓名
	UserPhone             string            `json:"user_phone" gorm:"not null;comment:用户电话"`                                                    // 用户电话
	UserAddress           string            `json:"user_address" gorm:"not null;comment:详细地址"`                                                  // 详细地址
	TotalNum              int               `json:"total_num" gorm:"not null;default:0;comment:订单商品总数"`                                         // 订单商品总数
	TotalPrice            float64           `json:"total_price" gorm:"not null;default:0.00;comment:订单总价"`                                      // 订单总价
	PayPrice              float64           `json:"pay_price" gorm:"not null;default:0.00;comment:实际支付金额"`                                      // 实际支付金额
	Paid                  uint8             `json:"paid" gorm:"not null;default:0;comment:支付状态"`                                                // 支付状态
	PayTime               datetime.Datetime `json:"pay_time" gorm:"type:datetime(0);comment:支付时间"`                                              // 支付时间
	PayType               string            `json:"pay_type" gorm:"not null;comment:支付方式"`                                                      // 支付方式
	Status                uint8             `json:"status" gorm:"not null;default:0;comment:订单状态(-1:申请退款;-2:退款成功;0:待发货;1:待收货;2:已收货,待评价;3:已完成)"` // 订单状态
	RefundStatus          uint8             `json:"refund_status" gorm:"not null;default:0;comment:0:未退款,1:申请中,2:已退款"`                          // 退款状态
	RefundReasonImg       string            `json:"refund_reason_img" gorm:"comment:退款图片"`                                                      // 退款图片
	RefundReasonExplain   string            `json:"refund_reason_explain" gorm:"comment:退款用户说明"`                                                // 退款用户说明
	RefundReason          string            `json:"refund_reason" gorm:"comment:前台退款原因"`                                                        // 前台退款原因
	RefundRejectionReason string            `json:"refund_rejection_reason" gorm:"comment:不退款的理由"`                                              // 不退款的理由
	RefundReasonTime      datetime.Datetime `json:"refund_reason_time" gorm:"type:datetime(0);comment:申请退款时间"`                                  // 退款时间
	RefundPrice           float64           `json:"refund_price" gorm:"not null;default:0.00;comment:退款金额"`                                     // 退款金额
	Remark                string            `json:"remark" gorm:"comment:管理员备注"`                                                                // 管理员备注
	MerchantId            int               `json:"merchant_id" gorm:"not null;default:0;comment:预留字段:商户ID"`                                    // 预留字段:商户ID
	IsMerchantCheck       uint8             `json:"is_merchant_check" gorm:"not null;default:0;comment:是否由商户核销"`                                // 是否已核销
	Cost                  float64           `json:"cost" gorm:"not null;comment:成本价"`                                                           // 成本价
	VerifyCode            string            `json:"verify_code" gorm:"not null;default:'';comment:核销码"`                                         // 核销码
	ShippingType          uint8             `json:"shipping_type" gorm:"not null;default:1;comment:配送方式:1=快递,2=门店自提"`                           // 配送方式
	ClerkId               int               `json:"clerk_id" gorm:"not null;default:0;comment:店员id/核销员id"`                                      // 店员id/核销员id
	IsCancel              int               `json:"is_cancel" gorm:"not null;default:0;comment:是否取消订单(0:否,1:是)"`                                // 是否取消订单
	IsAllRefund           int               `json:"is_all_refund" gorm:"not null;default:0;comment:是否全部退款(0:否,1:是)"`                            // 是否全部退款
	IsSystemDel           int               `json:"is_system_del" gorm:"not null;default:0;comment:是否为管理员删除(0:否,1:是)"`                          // 是否为管理员删除
	OutTradeNo            string            `json:"out_trade_no" gorm:"comment:商户系统内部的订单号,32个字符内、可包含字母"`                                        // 商户系统内部的订单号
	CreatedAt             datetime.Datetime `json:"created_at" gorm:"type:datetime(0)"`
	UpdatedAt             datetime.Datetime `json:"updated_at" gorm:"type:datetime(0)"` // 记录更新时间
	DeletedAt             gorm.DeletedAt    `json:"deleted_at"`
}

// Seeder
func (m *Order) Seeder() {

	// 如果菜单已存在，不执行Seeder操作
	if service.NewMenuService().IsExist(93) {
		return
	}

	// 创建菜单
	menuSeeders := []*appmodel.Menu{
		{Id: 93, Name: "订单管理", GuardName: "admin", Icon: "icon-orderedlist", Type: 1, Pid: 0, Sort: 0, Path: "/order", Show: 1, IsEngine: 0, IsLink: 0, Status: 1},
		{Id: 94, Name: "订单列表", GuardName: "admin", Icon: "", Type: 2, Pid: 93, Sort: 0, Path: "/api/admin/order/index", Show: 1, IsEngine: 1, IsLink: 0, Status: 1},
		{Id: 95, Name: "售后订单", GuardName: "admin", Icon: "", Type: 2, Pid: 93, Sort: 0, Path: "/api/admin/refundOrder/index", Show: 1, IsEngine: 1, IsLink: 0, Status: 1},
		{Id: 96, Name: "核销记录", GuardName: "admin", Icon: "", Type: 2, Pid: 93, Sort: 0, Path: "/api/admin/verifyOrder/index", Show: 1, IsEngine: 1, IsLink: 0, Status: 1},
	}
	db.Client.Create(&menuSeeders)
}
