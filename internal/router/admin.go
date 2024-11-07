package router

import (
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-smart/v2/internal/app/admin/handler"
)

// 注册Admin路由
func AdminRegister(b *quark.Engine) {
	g := b.Group("/api/admin")
	g.GET("/index/index", (&handler.Index{}).Index)
}
