package resource

import (
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/app/admin/searches"
	"github.com/quarkcloudio/quark-go/v3/template/admin/resource"
	"github.com/quarkcloudio/quark-smart/v2/internal/model"
)

type RefundOrder struct {
	resource.Template
}

// 初始化
func (p *RefundOrder) Init(ctx *quark.Context) interface{} {

	// 标题
	p.Title = "售后订单"

	// 模型
	p.Model = &model.Order{}

	// 分页
	p.PageSize = 10

	return p
}

func (p *RefundOrder) Fields(ctx *quark.Context) []interface{} {
	field := &resource.Field{}

	return []interface{}{
		field.Hidden("id", "ID"),

		field.Text("refund_no", "退款单号"),

		field.Text("order_no", "原订单号"),

		field.Text("item_info", "商品信息"),

		field.Text("user_info", "用户信息"),

		field.Text("total_pay", "支付金额"),

		field.Text("refund_at", "发起退款时间"),

		field.Text("status", "退款状态"),

		field.Text("order_status", "订单状态"),

		field.Text("refund_info", "退款信息"),
	}
}

// 搜索
func (p *RefundOrder) Searches(ctx *quark.Context) []interface{} {
	return []interface{}{
		searches.Input("title", "标题"),
		searches.Status(),
		searches.DatetimeRange("created_at", "创建时间"),
	}
}

// 行为
func (p *RefundOrder) Actions(ctx *quark.Context) []interface{} {
	return []interface{}{}
}
