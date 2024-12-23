package resource

import (
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/app/admin/actions"
	"github.com/quarkcloudio/quark-go/v3/app/admin/searches"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/form/rule"
	"github.com/quarkcloudio/quark-go/v3/template/admin/resource"
	"github.com/quarkcloudio/quark-smart/v2/internal/model"
	"github.com/quarkcloudio/quark-smart/v2/internal/service"
)

type Navigation struct {
	resource.Template
}

// 初始化
func (p *Navigation) Init(ctx *quark.Context) interface{} {

	// 标题
	p.Title = "导航"

	// 模型
	p.Model = &model.Navigation{}

	// 默认排序
	p.IndexQueryOrder = "sort asc"

	// 树形表格
	p.TableListToTree = true

	// 分页
	p.PageSize = false

	return p
}

func (p *Navigation) Fields(ctx *quark.Context) []interface{} {
	field := &resource.Field{}

	// 获取分类
	categorys, _ := service.NewNavigationService().TreeSelect(true)

	return []interface{}{
		field.Hidden("id", "ID"),

		field.Hidden("pid", "父节点"),

		field.Text("title", "标题").
			SetRules([]rule.Rule{
				rule.Required("标题必须填写"),
			}),

		field.TreeSelect("pid", "父节点").
			SetTreeData(categorys).
			SetDefault(0).
			OnlyOnForms(),

		field.Image("cover_id", "封面图").
			OnlyOnForms(),

		field.Number("sort", "排序").
			SetEditable(true).
			SetDefault(0),

		field.Text("url", "链接"),

		field.Switch("status", "状态").
			SetTrueValue("正常").
			SetFalseValue("禁用").
			SetDefault(true).
			OnlyOnForms(),
	}
}

// 搜索
func (p *Navigation) Searches(ctx *quark.Context) []interface{} {

	return []interface{}{
		searches.Input("title", "标题"),
		searches.Status(),
		searches.DatetimeRange("created_at", "创建时间"),
	}
}

// 行为
func (p *Navigation) Actions(ctx *quark.Context) []interface{} {

	return []interface{}{
		actions.CreateLink(),
		actions.BatchDelete(),
		actions.BatchDisable(),
		actions.BatchEnable(),
		actions.EditLink(),
		actions.Delete(),
		actions.FormSubmit(),
		actions.FormReset(),
		actions.FormBack(),
		actions.FormExtraBack(),
	}
}
