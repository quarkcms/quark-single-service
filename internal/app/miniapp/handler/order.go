package handler

import (
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-smart/v2/internal/dto/request"
	"github.com/quarkcloudio/quark-smart/v2/internal/dto/response"
	"github.com/quarkcloudio/quark-smart/v2/internal/service"
)

// 结构体
type Order struct{}

// 订单列表
func (p *Order) Index(ctx *quark.Context) error {
	return ctx.JSONOk("Hello, world!")
}

// 订单详情
func (p *Order) Detail(ctx *quark.Context) error {
	return ctx.JSONOk("Hello, world!")
}

// 提交订单
func (p *Order) Submit(ctx *quark.Context) error {
	submitOrderReq := request.SubmitOrderReq{}
	ctx.Bind(&submitOrderReq)

	// 获取用户id
	uid, err := service.NewAuthService(ctx).GetUid()
	if err != nil {
		return ctx.JSONError(err.Error())
	}

	orderNo, err := service.NewOrderService().Submit(uid, submitOrderReq)
	if err != nil {
		return ctx.JSONError(err.Error())
	}
	return ctx.JSONOk("ok", response.SubmitOrderResp{
		OrderNo: orderNo,
	})
}

// 取消订单
func (p *Order) Cancel(ctx *quark.Context) error {
	return ctx.JSONOk("Hello, world!")
}
