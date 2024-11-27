package resource

import (
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/app/admin/actions"
	"github.com/quarkcloudio/quark-go/v3/app/admin/searches"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/form/rule"
	"github.com/quarkcloudio/quark-go/v3/template/admin/resource"
	"github.com/quarkcloudio/quark-go/v3/utils/lister"
	"github.com/quarkcloudio/quark-smart/v2/internal/model"
	"github.com/quarkcloudio/quark-smart/v2/internal/service"
	"gorm.io/gorm"
)

type Page struct {
	resource.Template
}

// 初始化
func (p *Page) Init(ctx *quark.Context) interface{} {

	// 标题
	p.Title = "单页"

	// 模型
	p.Model = &model.Post{}

	// 默认排序
	p.IndexQueryOrder = "id asc"

	// 分页
	p.PageSize = false

	return p
}

// 只查询单页类型
func (p *Page) Query(ctx *quark.Context, query *gorm.DB) *gorm.DB {
	return query.Where("type", "PAGE")
}

// 字段
func (p *Page) Fields(ctx *quark.Context) []interface{} {
	field := &resource.Field{}

	// 获取分类
	pages, _ := service.NewPostService().TreeSelect(true)

	return []interface{}{
		field.Hidden("id", "ID"),

		field.Hidden("pid", "父节点"),

		field.Hidden("adminid", "AdminID"),

		field.Hidden("type", "类型").
			SetDefault("PAGE"),

		field.Text("title", "标题").
			SetRules([]rule.Rule{
				rule.Required("标题必须填写"),
			}),
		field.Text("name", "缩略名").
			OnlyOnForms(),

		field.TextArea("description", "描述").
			SetRules([]rule.Rule{
				rule.Max(200, "描述不能超过200个字符"),
			}).
			OnlyOnForms(),

		field.TreeSelect("pid", "根节点").
			SetTreeData(pages).
			OnlyOnForms(),

		field.Editor("content", "内容").OnlyOnForms(),

		field.Datetime("created_at", "创建时间").OnlyOnIndex(),

		field.Switch("status", "状态").
			SetTrueValue("正常").
			SetFalseValue("禁用").
			SetEditable(true).
			SetDefault(true),
	}
}

// 搜索
func (p *Page) Searches(ctx *quark.Context) []interface{} {

	return []interface{}{
		searches.Input("title", "标题"),
		searches.Status(),
		searches.DatetimeRange("created_at", "创建时间"),
	}
}

// 行为
func (p *Page) Actions(ctx *quark.Context) []interface{} {

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

// 列表页面显示前回调
func (p *Page) BeforeIndexShowing(ctx *quark.Context, list []map[string]interface{}) []interface{} {
	data := ctx.AllQuerys()
	if search, ok := data["search"].(map[string]interface{}); ok && search != nil {
		result := []interface{}{}
		for _, v := range list {
			result = append(result, v)
		}

		return result
	}

	// 转换成树形表格
	tree, _ := lister.ListToTree(list, "id", "pid", "children", 0)

	return tree
}
