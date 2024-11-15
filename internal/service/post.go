package service

import (
	"github.com/quarkcloudio/quark-go/v3/dal/db"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/form/fields/treeselect"
	"github.com/quarkcloudio/quark-smart/v2/internal/model"
)

type PostService struct{}

func NewPostService() *PostService {
	return &PostService{}
}

// 获取TreeSelect组件数据
func (p *PostService) TreeSelect(root bool) (list []treeselect.TreeData, Error error) {
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
func (p *PostService) FindTreeSelectNode(pid int) (list []treeselect.TreeData) {
	posts := []model.Post{}
	db.Client.
		Where("pid = ?", pid).
		Where("type", "PAGE").
		Order("id asc").
		Select("title", "id", "pid").
		Find(&posts)
	if len(posts) == 0 {
		return list
	}
	for _, v := range posts {
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
