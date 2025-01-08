package resource

import (
	"strconv"

	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/app/admin/searches"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/form/fields/selectfield"
	"github.com/quarkcloudio/quark-go/v3/template/admin/resource"
	"github.com/quarkcloudio/quark-go/v3/utils/datetime"
	"github.com/quarkcloudio/quark-smart/v2/internal/app/admin/action"
	"gorm.io/gorm"
)

// BillResource 定义后台用户账单资源的结构体
type BillResource struct {
	Id        int               `json:"id"`         // 用户账单id
	LinkId    int               `json:"link_id"`    // 订单id（orders.id）
	BillNo    string            `json:"bill_no"`    // 交易单号
	OrderNo   string            `json:"order_no"`   // 订单号
	CreatedAt datetime.Datetime `json:"created_at"` // 交易时间
	Number    float64           `json:"number"`     // 交易金额
	Realname  string            `json:"realname"`   // 交易用户`
	PayType   string            `json:"pay_type"`   // 支付方式
	Mark      string            `json:"mark"`       // 备注
	Pm        int               `json:"pm"`         // 0=支出,1=获得
}

func (BillResource) TableName() string {
	return "bills"
}

type Bill struct {
	resource.Template
}

// 初始化
func (p *Bill) Init(ctx *quark.Context) interface{} {

	// 标题
	p.Title = "资金流水"

	// 模型
	p.Model = BillResource{}

	// 分页
	p.PageSize = 10

	return p
}

// 查询
func (p *Bill) Query(ctx *quark.Context, query *gorm.DB) *gorm.DB {

	return query.Select(
		"bills.id",
		"bills.link_id",
		"bills.bill_no",
		"bills.pm",
		"bills.number",
		"bills.mark",
		"bills.created_at",
		"orders.order_no",
		"orders.realname",
		"orders.pay_type",
	).
		Joins("JOIN orders ON orders.id = bills.link_id")
}

func (p *Bill) Fields(ctx *quark.Context) []interface{} {
	field := &resource.Field{}

	return []interface{}{
		field.Hidden("id", "ID"),

		field.Text("bill_no", "交易单号").
			SetColumnWidth(180).
			SetEllipsis(true),

		field.Text("order_no", "关联订单").
			SetColumnWidth(180).
			SetEllipsis(true),

		field.Datetime("created_at", "交易时间").
			SetColumnWidth(160),

		field.Text("number", "交易金额", func(row map[string]interface{}) interface{} {
			if row["pm"] == 1 {
				return "+" + strconv.FormatFloat(row["number"].(float64), 'f', 2, 64)
			}

			return "-" + strconv.FormatFloat(row["number"].(float64), 'f', 2, 64)
		}).
			SetColumnWidth(100),

		field.Text("realname", "交易用户").
			SetColumnWidth(100),

		field.Select("pay_type", "支付方式").
			SetOptions([]selectfield.Option{
				{
					Label: "微信支付",
					Value: "WECHAT_PAY",
				},
				{
					Label: "支付宝支付",
					Value: "ALI_PAY",
				},
				{
					Label: "线下支付",
					Value: "OFFLINE_PAY",
				},
			}).
			SetColumnWidth(100),

		field.TextArea("mark", "备注").
			SetEllipsis(true),
	}
}

// 搜索
func (p *Bill) Searches(ctx *quark.Context) []interface{} {
	return []interface{}{
		searches.DatetimeRange("bills.created_at", "订单时间"),
		searches.Input("orders.order_no", "关联订单"),
		searches.Input("orders.realname", "交易用户"),
	}
}

// 行为
func (p *Bill) Actions(ctx *quark.Context) []interface{} {
	return []interface{}{
		action.BillMark("备注"),
	}
}
