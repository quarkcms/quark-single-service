package handler

import (
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-smart/v2/config"
)

// 结构体
type Login struct{}

// 用户登录
func (p *Login) Index(ctx *quark.Context) error {
	return ctx.JSONOk("Hello, world!")
}

// 用户模拟登录
func (p *Login) Mock(ctx *quark.Context) error {
	if !(config.App.Env == "develop" || config.App.Env == "dev" || config.App.Env == "development") {
		return ctx.JSONOk("It must be a development environment!")
	}
	return ctx.JSONOk("11")
}
