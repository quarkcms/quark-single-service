package resource

import (
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/app/admin/searches"
	"github.com/quarkcloudio/quark-go/v3/template/admin/resource"
	"github.com/quarkcloudio/quark-smart/v2/internal/model"
)

type Bill struct {
	resource.Template
}

// 初始化
func (p *Bill) Init(ctx *quark.Context) interface{} {

	// 标题
	p.Title = "资金流水"

	// 模型
	p.Model = &model.UserBill{}

	// 分页
	p.PageSize = 10

	return p
}

func (p *Bill) Fields(ctx *quark.Context) []interface{} {
	field := &resource.Field{}

	return []interface{}{
		field.Hidden("id", "ID"),

		field.Text("bill_no", "交易单号"),

		field.Text("order_no", "关联订单"),

		field.Text("created_at", "交易时间"),

		field.Text("created_at", "交易金额"),

		field.Text("created_at", "交易用户"),

		field.Text("created_at", "支付方式"),

		field.Text("created_at", "备注"),
	}
}

// 搜索
func (p *Bill) Searches(ctx *quark.Context) []interface{} {
	return []interface{}{
		searches.Input("title", "标题"),
		searches.Status(),
		searches.DatetimeRange("created_at", "创建时间"),
	}
}

// 行为
func (p *Bill) Actions(ctx *quark.Context) []interface{} {
	return []interface{}{}
}
