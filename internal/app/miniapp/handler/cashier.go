package handler

import (
	"github.com/quarkcloudio/quark-go/v3"
)

// 结构体
type Cashier struct{}

// 收银台
func (p *Cashier) Index(ctx *quark.Context) error {
	return ctx.JSONOk("Hello, world!")
}

// 调用支付
func (p *Cashier) Pay(ctx *quark.Context) error {
	return ctx.JSONOk("Hello, world!")
}
