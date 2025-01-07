package handler

import (
	"github.com/quarkcloudio/quark-go/v3"
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
	return ctx.JSONOk("Hello, world!")
}

// 取消订单
func (p *Order) Cancel(ctx *quark.Context) error {
	return ctx.JSONOk("Hello, world!")
}
