package resource

import (
	"fmt"

	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/app/admin/actions"
	"github.com/quarkcloudio/quark-go/v3/app/admin/searches"
	"github.com/quarkcloudio/quark-go/v3/template/admin/resource"
	"github.com/quarkcloudio/quark-smart/v2/internal/app/admin/action"
	"github.com/quarkcloudio/quark-smart/v2/internal/model"
	"github.com/quarkcloudio/quark-smart/v2/internal/service"
	"github.com/quarkcloudio/quark-smart/v2/pkg/utils"
	"gorm.io/gorm"
)

type Order struct {
	resource.Template
}

// 初始化
func (p *Order) Init(ctx *quark.Context) interface{} {

	// 标题
	p.Title = "订单"

	// 模型
	p.Model = &model.Order{}

	// 分页
	p.PageSize = 10

	return p
}

// 查询类型
func (p *Order) Query(ctx *quark.Context, query *gorm.DB) *gorm.DB {
	activeKey := ctx.QueryParam("activeKey")
	switch activeKey {
	case "all":
		// 全部
	case "pendingPayment":
		// 待支付
		query.Where("paid", 0)
	case "pendingShipment":
		// 待发货（预留）
		query.Where("paid", 1).Where("status = ?", 0).Where("shipping_type = ?", 1)
	case "pendingVerify":
		// 待核销
		query.Where("paid", 1).Where("status = ?", 0).Where("shipping_type = ?", 2)
	case "pendingReceipt":
		// 待收货（预留）
		query.Where("paid", 1).Where("status = ?", 1).Where("shipping_type = ?", 1)
	case "pendingReview":
		// 待评价（预留）
		query.Where("paid", 1).Where("status = ?", 2)
	case "completed":
		// 已完成
		query.Where("paid", 1).Where("status = ?", 3)
	case "refunded":
		// 已退款
		query.Where("paid", 1).Where("status = ?", -2)
	case "deleted":
		// 已删除
		query.Unscoped().Where("deleted_at IS NOT NULL")
	}
	return query
}

// 菜单
func (p *Order) Menus(ctx *quark.Context) interface{} {
	orderService := service.NewOrderService()

	return map[string]interface{}{
		"type": "tab",
		"items": []map[string]string{
			{
				"key":   "all",
				"label": "全部",
			},
			{
				"key":   "pendingPayment",
				"label": fmt.Sprintf("待支付(%d)", orderService.GetNumByStatus("pendingPayment")),
			},
			{
				"key":   "pendingVerify",
				"label": fmt.Sprintf("待核销(%d)", orderService.GetNumByStatus("pendingVerify")),
			},
			{
				"key":   "completed",
				"label": "已完成",
			},
			{
				"key":   "refunded",
				"label": "已退款",
			},
			{
				"key":   "deleted",
				"label": "已删除",
			},
		},
	}
}

func (p *Order) Fields(ctx *quark.Context) []interface{} {
	field := &resource.Field{}

	return []interface{}{
		field.Hidden("id", "ID"),

		field.Hidden("paid", "PAID"),

		field.Text("order_no", "订单号").SetColumnWidth(200),

		field.Text("name", "商品信息", func(row map[string]interface{}) interface{} {
			result := ""
			orderDetails, err := service.NewOrderService().GetOrderDetailsByOrderId(row["id"])
			if err != nil {
				return result
			}
			for k, orderDetail := range orderDetails {
				name := orderDetail.Name
				image := utils.GetImagePath(orderDetail.Image)
				style := ""
				if k != 0 {
					style = "margin-top:5px"
				}
				price := orderDetail.Price * float64(orderDetail.PayNum)
				title := fmt.Sprintf("商品名称：%s\r\n规格名称：%s\r\n支付价格：¥%.2f\r\n购买数量：%d", name, orderDetail.SKU, price, orderDetail.PayNum)
				result = result + fmt.Sprintf("<div title='%s' style='%s'><img src='%s' height=40 width=40 /> %s</div>", title, style, image, name)
			}
			return result
		}).SetColumnWidth(250),

		field.Text("user_info", "用户信息", func(row map[string]interface{}) interface{} {
			userInfo, err := service.NewUserService().GetInfoById(row["uid"])
			if err != nil {
				return nil
			}
			return fmt.Sprintf("用户ID：%d</br>用户账号：%s</br>用户昵称：%s", userInfo.Id, userInfo.Username, userInfo.Nickname)
		}),

		field.Text("pay_price", "实际支付", func(row map[string]interface{}) interface{} {
			if row["paid"].(uint8) == 0 {
				return nil
			}
			return row["pay_price"]
		}),

		field.Text("pay_type", "支付方式"),

		field.Text("pay_time", "支付时间"),

		// 0:待发货:,1:待收货,2:已收货,待评价,3:已完成
		field.Radio("status", "订单状态", func(row map[string]interface{}) interface{} {
			if row["paid"].(uint8) == 0 {
				return "未付款"
			}
			result := ""
			switch row["status"] {
			case -2:
				result = "退款成功"
			case -1:
				result = "申请退款"
			case 0:
				result = "待发货"
			case 1:
				result = "待收货"
			case 2:
				result = "已收货,待评价"
			case 3:
				result = "已完成"
			}
			return result
		}),
	}
}

// 搜索
func (p *Order) Searches(ctx *quark.Context) []interface{} {
	return []interface{}{
		searches.Input("order_no", "订单号"),
		searches.DatetimeRange("pay_time", "支付时间"),
	}
}

// 行为
func (p *Order) Actions(ctx *quark.Context) []interface{} {
	return []interface{}{
		action.OrderBatchDelete(),
		action.OrderBatchVerify(),
		actions.More("更多", []interface{}{
			action.OrderDelete(),
			action.OrderRefund(),
			action.OrderVerify(),
			action.OrderDetail(),
		}),
	}
}
