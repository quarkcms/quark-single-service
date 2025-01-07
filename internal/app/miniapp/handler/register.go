package handler

import (
	"github.com/quarkcloudio/quark-go/v3"
)

// 结构体
type Register struct{}

// 用户注册
func (p *Register) Index(ctx *quark.Context) error {
	return ctx.JSONOk("Hello, world!")
}
