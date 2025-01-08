package resource

import (
	"strconv"

	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/app/admin/searches"
	"github.com/quarkcloudio/quark-go/v3/template/admin/resource"
	"github.com/quarkcloudio/quark-smart/v2/internal/app/admin/action"
	"github.com/quarkcloudio/quark-smart/v2/internal/model"
	"gorm.io/gorm"
)

type BillRecord struct {
	resource.Template
}

// 初始化
func (p *BillRecord) Init(ctx *quark.Context) interface{} {

	// 标题
	p.Title = "账单记录"

	// 模型
	p.Model = &model.BillRecord{}

	// 分页
	p.PageSize = 10

	return p
}

// 查询
func (p *BillRecord) Query(ctx *quark.Context, query *gorm.DB) *gorm.DB {
	activeKey := ctx.QueryParam("activeKey")
	switch activeKey {
	case "day":
		query.Where("type", 1)
	case "week":
		query.Where("type", 2)
	case "month":
		query.Where("type", 3)
	default:
		query.Where("type", 1)
	}
	return query
}

// 菜单
func (p *BillRecord) Menus(ctx *quark.Context) interface{} {
	return map[string]interface{}{
		"type": "tab",
		"items": []map[string]string{
			{
				"key":   "day",
				"label": "日账单",
			},
			{
				"key":   "week",
				"label": "周账单",
			},
			{
				"key":   "month",
				"label": "月账单",
			},
		},
	}
}

func (p *BillRecord) Fields(ctx *quark.Context) []interface{} {
	field := &resource.Field{}

	return []interface{}{
		field.ID("id", "ID").
			SetColumnWidth(60),

		field.Text("title", "标题").
			SetColumnWidth(100).
			SetEllipsis(true),

		field.Text("day", "日期").
			SetColumnWidth(100),

		field.Number("entry_price", "收入金额", func(row map[string]interface{}) interface{} {
			return "￥" + strconv.FormatFloat(row["entry_price"].(float64), 'f', 2, 64)
		}).
			SetColumnWidth(100),

		field.Number("exp_price", "支出金额", func(row map[string]interface{}) interface{} {
			if row["exp_price"].(float64) == 0 {
				return "￥" + strconv.FormatFloat(row["exp_price"].(float64), 'f', 2, 64)
			}
			return "￥-" + strconv.FormatFloat(row["exp_price"].(float64), 'f', 2, 64)
		}).
			SetColumnWidth(100),

		field.Number("income_price", "入账金额", func(row map[string]interface{}) interface{} {
			return "￥" + strconv.FormatFloat(row["income_price"].(float64), 'f', 2, 64)
		}).
			SetColumnWidth(100),
	}
}

// 搜索
func (p *BillRecord) Searches(ctx *quark.Context) []interface{} {
	return []interface{}{
		searches.DatetimeRange("created_at", "创建时间"),
	}
}

// 行为
func (p *BillRecord) Actions(ctx *quark.Context) []interface{} {
	return []interface{}{
		action.BillDetail("账单详情"),
	}
}
