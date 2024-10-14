package handler

import (
	"github.com/quarkcloudio/quark-go/v3/pkg/builder"
)

// 结构体
type Home struct{}

// 首页
func (p *Home) Index(ctx *builder.Context) error {

	return ctx.Render(200, "index.html", map[string]interface{}{
		"content": "Hello, world!",
	})
}
