package handler

import (
	"github.com/quarkcloudio/quark-go/v3"
)

// 结构体
type Item struct{}

// 商品列表
func (p *Item) Index(ctx *quark.Context) error {
	return ctx.JSONOk("Hello, world!")
}

// 商品详情
func (p *Item) Detail(ctx *quark.Context) error {
	return ctx.JSONOk("Hello, world!")
}
