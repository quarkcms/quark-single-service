package service

import (
	"github.com/quarkcloudio/quark-go/v3/dal/db"
	"github.com/quarkcloudio/quark-smart/v2/internal/model"
)

type CategoryService struct{}

func NewCategoryService() *CategoryService {
	return &CategoryService{}
}

// 获取菜单列表
func (p *CategoryService) GetList() (categories []model.Category, Error error) {
	list := []model.Category{}
	err := db.Client.
		Where("status = ?", 1).
		Order("sort asc,id asc").
		Select("title", "id", "pid").
		Find(&list).Error
	return list, err
}

// 获取菜单列表携带根节点
func (p *CategoryService) GetListWithRoot() (categories []model.Category, Error error) {
	list, err := p.GetList()
	if err != nil {
		return list, err
	}
	list = append(list, model.Category{Id: 0, Pid: -1, Title: "根节点"})
	return list, err
}
