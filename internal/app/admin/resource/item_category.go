package resource

import (
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/app/admin/actions"
	"github.com/quarkcloudio/quark-go/v3/app/admin/searches"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/form/rule"
	"github.com/quarkcloudio/quark-go/v3/template/admin/resource"
	"github.com/quarkcloudio/quark-smart/v2/internal/model"
	"github.com/quarkcloudio/quark-smart/v2/internal/service"
	"gorm.io/gorm"
)

type ItemCategory struct {
	resource.Template
}

// 初始化
func (p *ItemCategory) Init(ctx *quark.Context) interface{} {

	// 标题
	p.Title = "分类"

	// 模型
	p.Model = &model.Category{}

	// 默认排序
	p.IndexQueryOrder = "sort asc"

	// 树形表格
	p.TableListToTree = true

	// 分页
	p.PageSize = false

	return p
}

// 全局查询
func (p *ItemCategory) Query(ctx *quark.Context, query *gorm.DB) *gorm.DB {
	return query.Where("type = ?", "ITEM")
}

func (p *ItemCategory) Fields(ctx *quark.Context) []interface{} {
	field := &resource.Field{}

	// 分类列表
	categories, _ := service.NewCategoryService().GetListWithRoot("ITEM")

	return []interface{}{
		field.Hidden("id", "ID"),

		field.Hidden("pid", "父节点"),

		field.Hidden("type", "类型").SetDefault("ITEM"),

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

		field.TreeSelect("pid", "父节点").
			SetTreeData(categories, -1, "pid", "title", "id").
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
func (p *ItemCategory) Searches(ctx *quark.Context) []interface{} {
	return []interface{}{
		searches.Input("title", "标题"),
		searches.Status(),
		searches.DatetimeRange("created_at", "创建时间"),
	}
}

// 行为
func (p *ItemCategory) Actions(ctx *quark.Context) []interface{} {
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
