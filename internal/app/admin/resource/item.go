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
		SetBody(p.Field1(ctx))
	tabPanes = append(tabPanes, pane1)

	// 规格库存
	pane2 := (&tabs.TabPane{}).
		Init().
		SetTitle("规格库存").
		SetBody(p.Field2(ctx))
	tabPanes = append(tabPanes, pane2)

	// 商品详情
	pane3 := (&tabs.TabPane{}).
		Init().
		SetTitle("商品详情").
		SetBody(p.Field3(ctx))
	tabPanes = append(tabPanes, pane3)

	// 营销设置
	pane4 := (&tabs.TabPane{}).
		Init().
		SetTitle("营销设置").
		SetBody(p.Field4(ctx))
	tabPanes = append(tabPanes, pane4)

	// 其他设置
	pane5 := (&tabs.TabPane{}).
		Init().
		SetTitle("其他设置").
		SetBody(p.Field5(ctx))
	tabPanes = append(tabPanes, pane5)

	return tabPanes
}

// 基础字段
func (p *Item) Field1(ctx *quark.Context) []interface{} {
	field := &resource.Field{}

	treeData, _ := service.NewCategoryService().GetList("ITEM")

	return []interface{}{
		field.Hidden("id", "ID"),

		field.Text("name", "商品名称").
			SetRules([]rule.Rule{
				rule.Required("商品名称必须填写"),
			}),

		field.ImagePicker("slider_image", "商品轮播图").
			SetMode("multiple").
			SetLimitNum(10).
			SetRules([]rule.Rule{
				rule.Required("请上传商品轮播图"),
			}).
			SetHelp("建议尺寸：800*800，默认首张图为主图，支持拖拽排序，最多上传10张").
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

// 规格库存字段
func (p *Item) Field2(ctx *quark.Context) []interface{} {
	field := &resource.Field{}

	return []interface{}{

		field.Radio("spec_type", "规格类型").
			SetOptions([]radio.Option{
				field.RadioOption("单规格", 0),
				field.RadioOption("多规格", 1),
			}).
			SetDefault(0).
			SetWhen(0, func() interface{} {
				return []interface{}{
					field.ImagePicker("image", "图片").
						OnlyOnForms(),

					field.Number("price", "售价").
						SetPrecision(2).
						SetAddonAfter("元").
						SetDefault(0.00).
						OnlyOnForms(),

					field.Number("cost", "成本价").
						SetPrecision(2).
						SetAddonAfter("元").
						SetDefault(0.00).
						OnlyOnForms(),

					field.Number("ot_price", "划线价").
						SetPrecision(2).
						SetAddonAfter("元").
						SetDefault(0.00).
						OnlyOnForms(),

					field.Number("stock", "库存").
						SetAddonAfter("件").
						SetDefault(0).
						OnlyOnForms(),
				}
			}).
			SetWhen(1, func() interface{} {
				return []interface{}{
					field.Sku("attributes", "datasource", "商品规格", "商品属性").
						SetFields([]interface{}{
							field.ImagePicker("image", "图片").
								SetColumnWidth(140),

							field.Number("price", "售价").
								SetColumnWidth(120),

							field.Number("cost", "成本价").
								SetColumnWidth(120),

							field.Number("ot_price", "划线价").
								SetColumnWidth(120),

							field.Number("stock", "库存").
								SetColumnWidth(120),
						}).
						OnlyOnForms(),
				}
			}).
			OnlyOnForms(),
	}
}

// 商品详情
func (p *Item) Field3(ctx *quark.Context) []interface{} {
	field := &resource.Field{}

	return []interface{}{

		field.Editor("content", "商品详情").
			OnlyOnForms(),
	}
}

// 营销设置
func (p *Item) Field4(ctx *quark.Context) []interface{} {
	field := &resource.Field{}

	return []interface{}{

		field.Number("views", "浏览量").
			SetDefault(0).
			OnlyOnForms(),

		field.Number("ficti", "已售数量").
			SetAddonAfter("件").
			SetDefault(0).
			OnlyOnForms(),

		field.Number("sort", "排序").
			SetDefault(0).
			OnlyOnForms(),
	}
}

// 其他设置
func (p *Item) Field5(ctx *quark.Context) []interface{} {
	field := &resource.Field{}

	return []interface{}{

		field.Text("keyword", "关键词").
			OnlyOnForms(),

		field.TextArea("info", "简介").
			SetRules([]rule.Rule{
				rule.Max(200, "描述不能超过200个字符"),
			}).
			OnlyOnForms(),
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
		actions.FormStep(),
		actions.FormSubmit(),
		actions.FormExtraBack(),
	}
}
