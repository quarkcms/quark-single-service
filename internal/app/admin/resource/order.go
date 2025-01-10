package resource

import (
	"fmt"

	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/app/admin/searches"
	"github.com/quarkcloudio/quark-go/v3/template/admin/resource"
	"github.com/quarkcloudio/quark-smart/v2/internal/model"
	"github.com/quarkcloudio/quark-smart/v2/internal/service"
	"github.com/quarkcloudio/quark-smart/v2/pkg/utils"
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

func (p *Order) Fields(ctx *quark.Context) []interface{} {
	field := &resource.Field{}

	return []interface{}{
		field.Hidden("id", "ID"),

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

		field.Text("user_info", "用户信息"),

		field.Text("pay_price", "支付金额"),

		field.Text("pay_type", "支付方式"),

		field.Text("pay_time", "支付时间"),

		field.Text("status", "订单状态"),
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
	return []interface{}{}
}
