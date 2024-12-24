package resource

import (
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/app/admin/actions"
	"github.com/quarkcloudio/quark-go/v3/app/admin/searches"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/form/fields/radio"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/form/rule"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/tabs"
	"github.com/quarkcloudio/quark-go/v3/template/admin/resource"
	"github.com/quarkcloudio/quark-smart/v2/internal/model"
	"github.com/quarkcloudio/quark-smart/v2/internal/service"
)

type Item struct {
	resource.Template
}

// 初始化
func (p *Item) Init(ctx *quark.Context) interface{} {

	// 标题
	p.Title = "商品"

	// 模型
	p.Model = &model.Item{}

	// 分页
	p.PageSize = 10

	return p
}

func (p *Item) Fields(ctx *quark.Context) []interface{} {
	var tabPanes []interface{}

	// 基础信息
	pane1 := (&tabs.TabPane{}).
		Init().
		SetTitle("基础信息").
		SetBody(p.BaseFields(ctx))
	tabPanes = append(tabPanes, pane1)

	// 规格库存
	pane2 := (&tabs.TabPane{}).
		Init().
		SetTitle("规格库存").
		SetBody(p.ExtendFields(ctx))
	tabPanes = append(tabPanes, pane2)

	// 商品详情
	pane3 := (&tabs.TabPane{}).
		Init().
		SetTitle("商品详情").
		SetBody(p.ExtendFields(ctx))
	tabPanes = append(tabPanes, pane3)

	// 营销设置
	pane4 := (&tabs.TabPane{}).
		Init().
		SetTitle("营销设置").
		SetBody(p.ExtendFields(ctx))
	tabPanes = append(tabPanes, pane4)

	// 其他设置
	pane5 := (&tabs.TabPane{}).
		Init().
		SetTitle("其他设置").
		SetBody(p.ExtendFields(ctx))
	tabPanes = append(tabPanes, pane5)

	return tabPanes
}

// 基础字段
func (p *Item) BaseFields(ctx *quark.Context) []interface{} {
	field := &resource.Field{}

	treeData, _ := service.NewCategoryService().GetList("ITEM")

	return []interface{}{
		field.Hidden("id", "ID"),

		field.Text("name", "商品名称").
			SetRules([]rule.Rule{
				rule.Required("商品名称必须填写"),
			}),

		field.Image("slider_image", "商品轮播图").
			SetMode("multiple").
			SetLimitNum(10).
			SetRules([]rule.Rule{
				rule.Required("请上传商品轮播图"),
			}).
			SetHelp("建议尺寸：800*800，默认首张图为主图，最多上传10张").
			OnlyOnForms(),

		field.TreeSelect("category_ids", "商品分类").
			SetTreeData(treeData, "pid", "title", "id").
			SetMultiple(true).
			SetRules([]rule.Rule{
				rule.Required("请选择商品分类"),
			}).
			OnlyOnForms(),

		field.Radio("status", "商品状态").
			SetOptions([]radio.Option{
				field.RadioOption("上架", 1),
				field.RadioOption("下架", 0),
			}).
			SetDefault(1).
			OnlyOnForms(),
	}
}

// 扩展字段
func (p *Item) ExtendFields(ctx *quark.Context) []interface{} {
	field := &resource.Field{}

	return []interface{}{

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
func (p *Item) Searches(ctx *quark.Context) []interface{} {
	return []interface{}{
		searches.Input("name", "商品名称"),
		searches.Status(),
		searches.DatetimeRange("created_at", "创建时间"),
	}
}

// 行为
func (p *Item) Actions(ctx *quark.Context) []interface{} {
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
