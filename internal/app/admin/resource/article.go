package resource

import (
	"time"

	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/app/admin/actions"
	"github.com/quarkcloudio/quark-go/v3/app/admin/searches"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/form/fields/checkbox"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/form/fields/radio"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/form/rule"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/tabs"
	"github.com/quarkcloudio/quark-go/v3/template/admin/resource"
	"github.com/quarkcloudio/quark-smart/v2/internal/model"
	"github.com/quarkcloudio/quark-smart/v2/internal/service"
	"gorm.io/gorm"
)

type Article struct {
	resource.Template
}

// 初始化
func (p *Article) Init(ctx *quark.Context) interface{} {

	// 标题
	p.Title = "文章"

	// 模型
	p.Model = &model.Post{}

	// 分页
	p.PageSize = 10

	return p
}

// 只查询文章类型
func (p *Article) Query(ctx *quark.Context, query *gorm.DB) *gorm.DB {
	return query.Where("type", "ARTICLE")
}

func (p *Article) Fields(ctx *quark.Context) []interface{} {
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
func (p *Article) BaseFields(ctx *quark.Context) []interface{} {
	field := &resource.Field{}

	// 分类列表
	categories, _ := service.NewCategoryService().GetList("ARTICLE")

	return []interface{}{
		field.ID("id", "ID"),

		field.Hidden("adminid", "AdminID"),

		field.Hidden("cover_ids", "封面图"),

		field.Hidden("type", "类型").
			SetDefault("ARTICLE"),

		field.Text("title", "标题").
			SetRules([]rule.Rule{
				rule.Required("标题必须填写"),
			}),

		field.TextArea("description", "描述").
			SetRules([]rule.Rule{
				rule.Max(200, "描述不能超过200个字符"),
			}).
			OnlyOnForms(),

		field.Text("author", "作者"),

		field.Number("level", "排序").
			SetEditable(true),

		field.Text("source", "来源").
			OnlyOnForms(),

		field.Checkbox("position", "推荐位").
			SetOptions([]checkbox.Option{
				field.CheckboxOption("首页推荐", 1),
				field.CheckboxOption("频道推荐", 2),
				field.CheckboxOption("列表推荐", 3),
				field.CheckboxOption("详情推荐", 4),
			}),

		field.Radio("show_type", "展现形式").
			SetOptions([]radio.Option{
				field.RadioOption("无图", 1),
				field.RadioOption("单图", 2),
				field.RadioOption("多图", 3),
			}).
			SetWhen(2, func() interface{} {
				return []interface{}{
					field.Image("single_cover_ids", "封面图").
						SetMode("multiple").
						SetLimitNum(1).
						OnlyOnForms(),
				}
			}).
			SetWhen(3, func() interface{} {
				return []interface{}{
					field.Image("multiple_cover_ids", "封面图").
						SetMode("multiple").
						OnlyOnForms(),
				}
			}).
			SetDefault(1).
			OnlyOnForms(),

		field.TreeSelect("category_id", "分类目录").
			SetTreeData(categories, "pid", "title", "id").
			SetRules([]rule.Rule{
				rule.Required("请选择分类目录"),
			}).
			OnlyOnForms(),

		field.Editor("content", "内容").OnlyOnForms(),

		field.Datetime("created_at", "发布时间").
			SetDefault(time.Now().Format("2006-01-02 15:04:05")),

		field.Switch("status", "状态").
			SetTrueValue("正常").
			SetFalseValue("禁用").
			OnlyOnForms(),
	}
}

// 扩展字段
func (p *Article) ExtendFields(ctx *quark.Context) []interface{} {
	field := &resource.Field{}

	return []interface{}{
		field.Text("name", "缩略名").
			OnlyOnForms(),

		field.Number("level", "排序").
			OnlyOnForms(),

		field.Number("view", "浏览量").
			OnlyOnForms(),

		field.Number("comment", "评论量").
			OnlyOnForms(),

		field.Text("password", "访问密码").
			OnlyOnForms(),

		field.File("file_ids", "附件").
			OnlyOnForms(),

		field.Switch("comment_status", "允许评论").
			SetEditable(true).
			SetTrueValue("正常").
			SetFalseValue("禁用").
			SetDefault(true),

		field.Datetime("created_at", "发布时间").
			OnlyOnForms(),

		field.Switch("status", "状态").
			SetEditable(true).
			SetTrueValue("正常").
			SetFalseValue("禁用").
			SetDefault(true),
	}
}

// 搜索
func (p *Article) Searches(ctx *quark.Context) []interface{} {
	options, _ := service.NewCategoryService().GetList("ARTICLE")

	return []interface{}{
		searches.Input("title", "标题"),
		searches.TreeSelect("category_id", "分类目录").SetTreeData(options, "pid", "title", "id"),
		searches.Status(),
		searches.DatetimeRange("created_at", "创建时间"),
	}
}

// 行为
func (p *Article) Actions(ctx *quark.Context) []interface{} {
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

// 编辑页面显示前回调
func (p *Article) BeforeEditing(request *quark.Context, data map[string]interface{}) map[string]interface{} {
	if data["show_type"] == 2 {
		data["single_cover_ids"] = data["cover_ids"]
	}

	if data["show_type"] == 3 {
		data["multiple_cover_ids"] = data["cover_ids"]
	}

	return data
}

// 保存数据前回调
func (p *Article) BeforeSaving(ctx *quark.Context, submitData map[string]interface{}) (map[string]interface{}, error) {
	if int(submitData["show_type"].(float64)) == 2 {
		submitData["cover_ids"] = submitData["single_cover_ids"]
	}

	if int(submitData["show_type"].(float64)) == 3 {
		submitData["cover_ids"] = submitData["multiple_cover_ids"]
	}

	return submitData, nil
}
