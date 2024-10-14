package resource

import (
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/component/form/rule"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/service/actions"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/service/searches"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/template/resource"
	"github.com/quarkcloudio/quark-go/v3/pkg/builder"
	"github.com/quarkcloudio/quark-smart/v2/internal/model"
)

type Banner struct {
	resource.Template
}

// 初始化
func (p *Banner) Init(ctx *builder.Context) interface{} {

	// 标题
	p.Title = "广告"

	// 模型
	p.Model = &model.Banner{}

	// 分页
	p.PerPage = 10

	return p
}

func (p *Banner) Fields(ctx *builder.Context) []interface{} {
	field := &resource.Field{}

	// 获取分类
	categorys, _ := (&model.BannerCategory{}).Options()

	return []interface{}{
		field.ID("id", "ID"),

		field.Text("title", "标题").
			SetRules([]*rule.Rule{
				rule.Required(true, "标题必须填写"),
			}),

		field.Image("cover_id", "图片").
			OnlyOnForms(),

		field.Select("category_id", "广告位").
			SetOptions(categorys).
			SetRules([]*rule.Rule{
				rule.Required(true, "请选择广告位"),
			}).
			OnlyOnForms(),

		field.Number("sort", "排序").
			SetEditable(true).
			SetDefault(0),

		field.Text("url", "链接"),

		field.Switch("status", "状态").
			SetTrueValue("正常").
			SetFalseValue("禁用").
			SetEditable(true).
			SetDefault(true),
	}
}

// 搜索
func (p *Banner) Searches(ctx *builder.Context) []interface{} {
	options, _ := (&model.BannerCategory{}).Options()

	return []interface{}{
		searches.Input("title", "标题"),
		searches.Select("category_id", "广告位", options),
		searches.Status(),
		searches.DatetimeRange("created_at", "创建时间"),
	}
}

// 行为
func (p *Banner) Actions(ctx *builder.Context) []interface{} {
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
