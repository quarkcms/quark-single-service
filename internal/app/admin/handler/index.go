package handler

import (
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/message"
)

// 结构体
type Index struct{}

// 首页
func (p *Index) Index(ctx *quark.Context) error {
	return ctx.JSON(200, message.Success("Hello, world!"))
}
