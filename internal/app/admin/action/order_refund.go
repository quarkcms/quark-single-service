package action

import (
	"strconv"

	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/message"
	"github.com/quarkcloudio/quark-go/v3/template/admin/resource"
	"github.com/quarkcloudio/quark-go/v3/template/admin/resource/actions"
	"github.com/quarkcloudio/quark-smart/v2/internal/service"
	"gorm.io/gorm"
)

type OrderRefundAction struct {
	actions.ModalForm
}

// 订单退款
func OrderRefund() *OrderRefundAction {
	return &OrderRefundAction{}
}

// 初始化
func (p *OrderRefundAction) Init(ctx *quark.Context) interface{} {

	// 设置按钮文字
	p.Name = "立即退款"

	// 类型
	p.Type = "link"

	// 设置按钮大小,large | middle | small | default
	p.Size = "small"

	// 关闭时销毁 Modal 里的元素
	p.DestroyOnClose = true

	//  执行成功后刷新的组件
	p.Reload = "table"

	// 在表格行内展示
	p.SetOnlyOnIndexTableRow(true)

	// 行为接口接收的参数
	p.SetApiParams([]string{
		"id",
	})

	return p
}

// 字段
func (p *OrderRefundAction) Fields(ctx *quark.Context) []interface{} {
	field := &resource.Field{}
	return []interface{}{
		field.Hidden("id", "ID"),
		field.Display("支付金额${pay_price}"),
		field.Display("可退款金额${pay_price}"),
		field.Number("refund_price", "退款金额").SetPlaceholder("请输入退款金额"),
	}
}

// 表单数据（异步获取）
func (p *OrderRefundAction) Data(ctx *quark.Context) map[string]interface{} {
	id, _ := strconv.Atoi(ctx.Query("id").(string))
	order, _ := service.NewOrderService().GetOrderById(id)
	return map[string]interface{}{
		"id":        id,
		"pay_price": order.PayPrice,
	}
}

// 执行行为句柄
func (p *OrderRefundAction) Handle(ctx *quark.Context, query *gorm.DB) error {
	var refundReq struct {
		Id          int     `json:"id"`
		RefundPrice float64 `json:"refund_price"`
	}
	if err := ctx.Bind(&refundReq); err != nil {
		return ctx.JSON(200, message.Error(err.Error()))
	}
	if err := service.NewOrderService().Refund(refundReq.Id, refundReq.RefundPrice); err != nil {
		return ctx.JSON(200, message.Error("操作失败"))
	}
	return ctx.JSON(200, message.Success("操作成功"))
}
