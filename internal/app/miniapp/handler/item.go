package handler

import (
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-smart/v2/internal/service"
	"github.com/quarkcloudio/quark-smart/v2/pkg/utils"
)

// 结构体
type Item struct{}

// 商品列表
func (p *Item) Index(ctx *quark.Context) error {
	return ctx.JSONOk("Hello, world!")
}

// 商品详情
func (p *Item) Detail(ctx *quark.Context) error {
	return ctx.JSONOk("Hello, world!")
}

// 商品分类
func (p *Item) Category(ctx *quark.Context) error {
	itemService := service.NewItemService()
	itemCategories := itemService.GetCategoriesByPid(0)
	for index, itemCategory := range itemCategories {
		itemCategory.Title = "全部商品"
		itemCategory.CoverId = utils.GetImagePath(itemCategory.CoverId)
		itemCategories[index].Children = append(itemCategories[index].Children, itemCategory)
		children := itemService.GetCategoriesByPid(itemCategory.Id)
		for _, child := range children {
			child.CoverId = utils.GetImagePath(child.CoverId)
			itemCategories[index].Children = append(itemCategories[index].Children, child)
		}
	}
	return ctx.JSONOk("ok", itemCategories)
}
