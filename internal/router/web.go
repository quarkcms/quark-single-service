package router

import (
	"github.com/quarkcloudio/quark-go/v3/pkg/builder"
	"github.com/quarkcloudio/quark-smart/v2/internal/handler"
)

// 注册Web路由
func WebRegister(b *builder.Engine) {
	b.GET("/", (&handler.Home{}).Index)
}
