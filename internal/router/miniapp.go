package router

import (
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-smart/v2/internal/app/miniapp/handler"
)

// 注册MiniApp路由
func MiniAppRegister(b *quark.Engine) {
	g := b.Group("/api/miniapp")
	g.GET("/", (&handler.Index{}).Index)
}