package service

import (
	"github.com/quarkcloudio/quark-go/v3/dal/db"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/form/fields/selectfield"
	"github.com/quarkcloudio/quark-smart/v2/internal/model"
)

type BannerCategoryService struct{}

func NewBannerCategoryService() *BannerCategoryService {
	return &BannerCategoryService{}
}

// 获取列表
func (p *BannerCategoryService) Options() (options []selectfield.Option, Error error) {
	getList := []model.BannerCategory{}
	err := db.Client.Find(&getList).Error
	if err != nil {
		return options, err
	}
	for _, v := range getList {
		option := selectfield.Option{
			Label: v.Title,
			Value: v.Id,
		}
		options = append(options, option)
	}
	return options, nil
}
