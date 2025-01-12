package handler

import (
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-smart/v2/internal/service"
	"github.com/quarkcloudio/quark-smart/v2/pkg/utils"
)

// 结构体
type Index struct{}

// 首页
func (p *Index) Index(ctx *quark.Context) error {
	return ctx.JSONOk("Hello, world!")
}

// 轮播图
func (p *Index) Banner(ctx *quark.Context) error {
	banners := service.NewBannerService().GetList()
	for index, banner := range banners {
		// 处理图片 url
		banners[index].CoverId = utils.GetImagePath(banner.CoverId)
	}
	return ctx.JSONOk("ok", banners)
}