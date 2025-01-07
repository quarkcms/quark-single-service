package router

import (
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-smart/v2/internal/app/miniapp/handler"
	"github.com/quarkcloudio/quark-smart/v2/internal/middleware"
)

// 注册MiniApp路由
func MiniAppRegister(b *quark.Engine) {

	// 不需要认证路由组
	g := b.Group("/api/miniapp")
	g.GET("/index/index", (&handler.Index{}).Index)
	g.GET("/login/index", (&handler.Login{}).Index)
	g.GET("/login/mock", (&handler.Login{}).Mock)

	// 需要登录认证路由组
	ag := b.Group("/api/miniapp", middleware.MiniAppMiddleware)
	ag.GET("/user/index", (&handler.User{}).Index)
}
