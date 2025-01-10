package resource

import (
	"fmt"

	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/app/admin/actions"
	"github.com/quarkcloudio/quark-go/v3/app/admin/searches"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/form/fields/radio"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/form/rule"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/tabs"
	"github.com/quarkcloudio/quark-go/v3/template/admin/resource"
	"github.com/quarkcloudio/quark-smart/v2/internal/model"
	"github.com/quarkcloudio/quark-smart/v2/internal/service"
	"gorm.io/gorm"
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

// 查询类型
func (p *Item) Query(ctx *quark.Context, query *gorm.DB) *gorm.DB {
	activeKey := ctx.QueryParam("activeKey")
	switch activeKey {
	case "onSale":
		query.Where("status", 1)
	case "offSale":
		query.Where("status", 0)
	}
	return query
}

// 菜单
func (p *Item) Menus(ctx *quark.Context) interface{} {
	totalNum := service.NewItemService().GetNumByStatus(nil)
	onSalelNum := service.NewItemService().GetNumByStatus(1)
	offSaleNum := service.NewItemService().GetNumByStatus(0)
	return map[string]interface{}{
		"type": "tab",
		"items": []map[string]string{
			{
				"key":   "all",
				"label": fmt.Sprintf("全部商品(%d)", totalNum),
			},
			{
				"key":   "onSale",
				"label": fmt.Sprintf("出售中的商品(%d)", onSalelNum),
			},
			{
				"key":   "offSale",
				"label": fmt.Sprintf("仓库中的商品(%d)", offSaleNum),
			},
		},
	}
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
		field.ID("id", "ID"),

		field.Hidden("attrs", "Attrs"),

		field.Text("name", "商品名称").
			SetColumnWidth(160).
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
					field.ImagePicker("image", "图片"),

					field.Number("price", "售价").
						SetPrecision(2).
						SetAddonAfter("元").
						SetDefault(0.00),

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
						SetDefault(0),
				}
			}).
			SetWhen(1, func() interface{} {
				return []interface{}{
					field.Sku("attrs", "attr_values", "商品规格", "商品属性").
						SetFields([]interface{}{
							field.ImagePicker("image", "图片").
								SetColumnWidth(140),

							field.Number("price", "售价").
								SetColumnWidth(120).
								SetDefault(0),

							field.Number("cost", "成本价").
								SetColumnWidth(120).
								SetDefault(0),

							field.Number("ot_price", "划线价").
								SetColumnWidth(120).
								SetDefault(0),

							field.Number("stock", "库存").
								SetColumnWidth(120).
								SetDefault(0),
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

		field.Number("ficti_views", "浏览量").
			SetDefault(0).
			OnlyOnForms(),

		field.Number("ficti_sales", "已售数量").
			SetAddonAfter("件").
			SetDefault(0),

		field.Number("sort", "排序").
			SetEditable(true).
			SetDefault(0),

		field.Switch("status", "状态").
			SetTrueValue("上架").
			SetFalseValue("下架").
			SetEditable(true).
			SetDefault(true).
			OnlyOnIndex(),
	}
}

// 其他设置
func (p *Item) Field5(ctx *quark.Context) []interface{} {
	field := &resource.Field{}

	return []interface{}{

		field.Text("keyword", "关键词").
			OnlyOnForms(),

		field.TextArea("description", "简介").
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
		searches.DatetimeRange("created_at", "上架时间"),
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

// 保存数据后回调
func (p *Item) AfterSaved(ctx *quark.Context, id int, data map[string]interface{}, result *gorm.DB) (err error) {
	// 更新商品分类表
	service.NewCategoryService().StoreItemCategory(id, data["category_ids"])
	// 更新商品属性
	service.NewItemService().StoreOrUpdateItemAttrs(id, data["attrs"])
	// 更新商品属性值
	service.NewItemService().StoreOrUpdateItemAttrValues(id, data["attr_values"])
	return err
}
