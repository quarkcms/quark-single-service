package handler

import (
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-smart/v2/internal/service"
)

// 结构体
type Login struct{}

// 用户名、密码登录
func (p *Login) Index(ctx *quark.Context) error {
	return ctx.JSONOk("Hello, world!")
}

// 模拟登录
func (p *Login) Mock(ctx *quark.Context) error {
	token, err := service.NewAuthService(ctx).MockLogin()
	if err != nil {
		return ctx.JSONError(err.Error())
	}
	return ctx.JSONOk("获取成功", map[string]interface{}{
		"token": token,
	})
}
