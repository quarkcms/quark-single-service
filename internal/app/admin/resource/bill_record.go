package resource

import (
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/app/admin/searches"
	"github.com/quarkcloudio/quark-go/v3/template/admin/resource"
	"github.com/quarkcloudio/quark-smart/v2/internal/model"
)

type BillRecord struct {
	resource.Template
}

// 初始化
func (p *BillRecord) Init(ctx *quark.Context) interface{} {

	// 标题
	p.Title = "账单记录"

	// 模型
	p.Model = &model.UserBill{}

	// 分页
	p.PageSize = 10

	return p
}

func (p *BillRecord) Fields(ctx *quark.Context) []interface{} {
	field := &resource.Field{}

	return []interface{}{
		field.ID("id", "ID"),

		field.Text("bill_no", "标题"),

		field.Text("order_no", "日期"),

		field.Text("created_at", "收入金额"),

		field.Text("created_at", "支出金额"),

		field.Text("created_at", "入账金额"),
	}
}

// 搜索
func (p *BillRecord) Searches(ctx *quark.Context) []interface{} {
	return []interface{}{
		searches.Input("title", "标题"),
		searches.Status(),
		searches.DatetimeRange("created_at", "创建时间"),
	}
}

// 行为
func (p *BillRecord) Actions(ctx *quark.Context) []interface{} {
	return []interface{}{}
}
