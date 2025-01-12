package service

import (
	"github.com/quarkcloudio/quark-go/v3/dal/db"
	"github.com/quarkcloudio/quark-go/v3/utils/datetime"
	"github.com/quarkcloudio/quark-smart/v2/internal/dto/response"
	"github.com/quarkcloudio/quark-smart/v2/internal/model"
)

type BannerService struct{}

func NewBannerService() *BannerService {
	return &BannerService{}
}

// 获取轮播列表
func (p *BannerService) GetList() []response.BannerListResp {
	banners := make([]response.BannerListResp, 0)
	db.Client.Model(model.Banner{}).
		Where("status = ?", 1).
		Where("deadline IS NULL OR deadline > ?", datetime.Now()).
		Order("sort, id").
		Find(&banners)
	return banners
}
