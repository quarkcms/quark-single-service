package handler

import (
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-smart/v2/internal/service"
	"github.com/quarkcloudio/quark-smart/v2/pkg/utils"
)

// 轮播图
type Banner struct{}

// 列表
func (p *Banner) Index(ctx *quark.Context) error {
	banners := service.NewBannerService().GetList()

	for index, banner := range banners {
		// 处理图片 url
		banners[index].CoverId = utils.GetImagePath(banner.CoverId)
	}

	return ctx.JSONOk("ok", banners)
}
