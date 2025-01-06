package resource

import (
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/app/admin/searches"
	"github.com/quarkcloudio/quark-go/v3/template/admin/resource"
	"github.com/quarkcloudio/quark-smart/v2/internal/model"
)

type Order struct {
	resource.Template
}

// 初始化
func (p *Order) Init(ctx *quark.Context) interface{} {

	// 标题
	p.Title = "订单"

	// 模型
	p.Model = &model.Order{}

	// 分页
	p.PageSize = 10

	return p
}

func (p *Order) Fields(ctx *quark.Context) []interface{} {
	field := &resource.Field{}

	return []interface{}{
		field.Hidden("id", "ID"),

		field.Text("order_no", "订单号"),

		field.Text("name", "商品信息"),

		field.Text("user_info", "用户信息"),

		field.Text("pay_price", "支付金额"),

		field.Text("pay_type", "支付方式"),

		field.Text("pay_time", "支付时间"),

		field.Text("status", "订单状态"),
	}
}

// 搜索
func (p *Order) Searches(ctx *quark.Context) []interface{} {
	return []interface{}{
		searches.Input("order_no", "订单号"),
		searches.DatetimeRange("pay_time", "支付时间"),
	}
}

// 行为
func (p *Order) Actions(ctx *quark.Context) []interface{} {
	return []interface{}{}
}
