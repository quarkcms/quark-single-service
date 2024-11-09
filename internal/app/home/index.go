package home

import (
	"github.com/quarkcloudio/quark-go/v3"
)

// 结构体
type Index struct{}

// 首页
func (p *Index) Index(ctx *quark.Context) error {
	return ctx.Render(200, "index.html", map[string]interface{}{
		"content": "Hello, world!",
	})
}
