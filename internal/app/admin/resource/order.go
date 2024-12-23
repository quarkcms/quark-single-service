package resource

import (
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/app/admin/searches"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/form/rule"
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

		field.Text("title", "标题").
			SetRules([]rule.Rule{
				rule.Required("标题必须填写"),
			}),

		field.Text("name", "缩略名").
			SetRules([]rule.Rule{
				rule.Required("缩略名必须填写"),
			}),

		field.Image("cover_id", "封面图").
			SetMode("single").
			OnlyOnForms(),

		field.TextArea("description", "描述").
			OnlyOnForms(),

		field.Number("sort", "排序").
			SetEditable(true),

		field.Switch("status", "状态").
			SetTrueValue("正常").
			SetFalseValue("禁用").
			SetDefault(true).
			OnlyOnForms(),
	}
}

// 搜索
func (p *Order) Searches(ctx *quark.Context) []interface{} {
	return []interface{}{
		searches.Input("title", "标题"),
		searches.Status(),
		searches.DatetimeRange("created_at", "创建时间"),
	}
}

// 行为
func (p *Order) Actions(ctx *quark.Context) []interface{} {
	return []interface{}{}
}
