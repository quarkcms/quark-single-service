package handler

import (
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/component/message"
	"github.com/quarkcloudio/quark-go/v3/pkg/builder"
)

// 结构体
type Index struct{}

// 首页
func (p *Index) Index(ctx *builder.Context) error {
	return ctx.JSON(200, message.Success("Hello, world!"))
}
