package router

import (
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-smart/v2/internal/app/home"
)

// 注册Web路由
func WebRegister(b *quark.Engine) {
	b.GET("/", (&home.Index{}).Index)
}
