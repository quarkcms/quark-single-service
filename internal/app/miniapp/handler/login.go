package handler

import (
	"github.com/dchest/captcha"
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-smart/v2/internal/dto/request"
	"github.com/quarkcloudio/quark-smart/v2/internal/service"
)

// 结构体
type Login struct{}

// 用户名、密码登录
func (p *Login) Index(ctx *quark.Context) error {
	loginReq := &request.LoginReq{}
	if err := ctx.Bind(loginReq); err != nil {
		return ctx.JSONError(err.Error())
	}
	if loginReq.Captcha.Id == "" || loginReq.Captcha.Value == "" {
		return ctx.JSONError("验证码不能为空")
	}

	verifyResult := captcha.VerifyString(loginReq.Captcha.Id, loginReq.Captcha.Value)
	if !verifyResult {
		return ctx.JSONError("验证码错误")
	}
	captcha.Reload(loginReq.Captcha.Id)

	if loginReq.Username == "" || loginReq.Password == "" {
		return ctx.JSONError("用户名或密码不能为空")
	}
	token, err := service.NewAuthService(ctx).Login(loginReq.Username, loginReq.Password)
	if err != nil {
		return ctx.JSONError(err.Error())
	}

	return ctx.JSONOk("获取成功", map[string]interface{}{
		"token": token,
	})
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
