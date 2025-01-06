package handler

import (
	"github.com/quarkcloudio/quark-go/v3"
)

// 结构体
type Index struct{}

// 首页
func (p *Index) Index(ctx *quark.Context) error {
	return ctx.JSONOk("Hello, world!")
}
