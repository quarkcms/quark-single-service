package handler

import (
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-smart/v2/config"
	"github.com/quarkcloudio/quark-smart/v2/internal/service"
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
		return ctx.JSONError("It must be a development environment!")
	}
	uid := ctx.Query("uid", 1)
	userService := service.NewUserService()

	// 获取用户信息
	userInfo, err := userService.GetInfoById(uid)
	if err != nil {
		return ctx.JSONError(err.Error())
	}

	// 获取登录token
	token, err := ctx.JwtToken(userService.GetUserClaims(userInfo))
	if err != nil {
		return ctx.JSONError(err.Error())
	}

	return ctx.JSONOk("获取成功", map[string]interface{}{
		"token": token,
	})
}
