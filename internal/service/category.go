package service

import (
	"github.com/quarkcloudio/quark-go/v3/dal/db"
	"github.com/quarkcloudio/quark-smart/v2/internal/model"
)

type CategoryService struct{}

func NewCategoryService() *CategoryService {
	return &CategoryService{}
}

// 获取分类列表
func (p *CategoryService) GetList(categoryType string) (categories []model.Category, err error) {
	list := []model.Category{}
	err = db.Client.
		Where("status = ?", 1).
		Where("type = ?", categoryType).
		Order("sort asc, id asc").
		Select("title", "id", "pid").
		Find(&list).Error
	return list, err
}

// 获取分类列表携带根节点
func (p *CategoryService) GetListWithRoot(categoryType string) (categories []model.Category, err error) {
	list, err := p.GetList(categoryType)
	if err != nil {
		return list, err
	}
	list = append(list, model.Category{Id: 0, Pid: -1, Title: "根节点"})
	return list, err
}

// 删除商品类别
func (p *CategoryService) DeleteItemCategoryByItemId(itemId interface{}) {
	db.Client.
		Where("item_id = ?", itemId).
		Delete(&model.ItemCategory{})
}

// 创建商品类别
func (p *CategoryService) StoreItemCategory(itemId int, categoryIds interface{}) {
	p.DeleteItemCategoryByItemId(itemId)
	if categoryIds == nil {
		return
	}
	for _, item := range categoryIds.([]interface{}) {
		db.Client.Create(&model.ItemCategory{
			ItemId:     itemId,
			CategoryId: int(item.(float64)),
		})
	}
}
