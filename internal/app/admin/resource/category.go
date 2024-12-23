package resource

import (
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/app/admin/actions"
	"github.com/quarkcloudio/quark-go/v3/app/admin/searches"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/form/rule"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/tabs"
	"github.com/quarkcloudio/quark-go/v3/template/admin/resource"
	"github.com/quarkcloudio/quark-smart/v2/internal/model"
	"github.com/quarkcloudio/quark-smart/v2/internal/service"
	"gorm.io/gorm"
)

type Category struct {
	resource.Template
}

// 初始化
func (p *Category) Init(ctx *quark.Context) interface{} {

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
func (p *Category) Query(ctx *quark.Context, query *gorm.DB) *gorm.DB {
	return query.Where("type = ?", "ARTICLE")
}

func (p *Category) Fields(ctx *quark.Context) []interface{} {
	var tabPanes []interface{}

	// 基础字段
	basePane := (&tabs.TabPane{}).
		Init().
		SetTitle("基础").
		SetBody(p.BaseFields(ctx))
	tabPanes = append(tabPanes, basePane)

	// 扩展字段
	extendPane := (&tabs.TabPane{}).
		Init().
		SetTitle("扩展").
		SetBody(p.ExtendFields(ctx))
	tabPanes = append(tabPanes, extendPane)

	return tabPanes
}

// 基础字段
func (p *Category) BaseFields(ctx *quark.Context) []interface{} {
	field := &resource.Field{}

	// 分类列表
	categories, _ := service.NewCategoryService().GetListWithRoot("ARTICLE")

	return []interface{}{
		field.Hidden("id", "ID"),

		field.Hidden("pid", "父节点"),

		field.Text("title", "标题").
			SetRules([]rule.Rule{
				rule.Required("标题必须填写"),
			}),

		field.Text("name", "缩略名").
			SetRules([]rule.Rule{
				rule.Required("缩略名必须填写"),
			}),

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

// 扩展字段
func (p *Category) ExtendFields(ctx *quark.Context) []interface{} {
	field := &resource.Field{}

	return []interface{}{
		field.Image("cover_id", "封面图").
			SetMode("single").
			OnlyOnForms(),

		field.Text("index_tpl", "频道模板").
			OnlyOnForms(),

		field.Text("lists_tpl", "列表模板").
			OnlyOnForms(),

		field.Text("detail_tpl", "详情模板").
			OnlyOnForms(),

		field.Number("page_num", "分页数量").
			SetEditable(true).
			SetDefault(10),

		field.Switch("status", "状态").
			SetEditable(true).
			SetTrueValue("正常").
			SetFalseValue("禁用").
			SetDefault(true),
	}
}

// 搜索
func (p *Category) Searches(ctx *quark.Context) []interface{} {
	return []interface{}{
		searches.Input("title", "标题"),
		searches.Status(),
		searches.DatetimeRange("created_at", "创建时间"),
	}
}

// 行为
func (p *Category) Actions(ctx *quark.Context) []interface{} {
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
