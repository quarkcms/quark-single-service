package service

import (
	"github.com/quarkcloudio/quark-go/v3/dal/db"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/form/fields/treeselect"
	"github.com/quarkcloudio/quark-smart/v2/internal/model"
)

type NavigationService struct{}

func NewNavigationService() *NavigationService {
	return &NavigationService{}
}

// 获取TreeSelect组件数据
func (p *NavigationService) TreeSelect(root bool) (list []treeselect.TreeData, Error error) {
	if root {
		list = append(list, treeselect.TreeData{
			Title: "根节点",
			Value: 0,
		})
	}
	list = append(list, p.FindTreeSelectNode(0)...)
	return list, nil
}

// 递归获取TreeSelect组件数据
func (p *NavigationService) FindTreeSelectNode(pid int) (list []treeselect.TreeData) {
	navigations := []model.Navigation{}
	db.Client.
		Where("pid = ?", pid).
		Order("sort asc,id asc").
		Select("title", "id", "pid").
		Find(&navigations)
	if len(navigations) == 0 {
		return list
	}
	for _, v := range navigations {
		item := treeselect.TreeData{
			Value: v.Id,
			Title: v.Title,
		}
		children := p.FindTreeSelectNode(v.Id)
		if len(children) > 0 {
			item.Children = children
		}
		list = append(list, item)
	}
	return list
}
