package handler

import (
	"github.com/quarkcloudio/quark-go/v3"
)

// 结构体
type User struct{}

// 用户中心
func (p *User) Index(ctx *quark.Context) error {
	return ctx.JSONOk("Hello, world!")
}

// 更新用户信息
func (p *User) Save(ctx *quark.Context) error {
	return ctx.JSONOk("Hello, world!")
}
