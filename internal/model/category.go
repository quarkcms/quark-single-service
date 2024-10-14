package model

import (
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/component/form/fields/treeselect"
	appmodel "github.com/quarkcloudio/quark-go/v3/pkg/app/admin/model"
	"github.com/quarkcloudio/quark-go/v3/pkg/dal/db"
	"github.com/quarkcloudio/quark-go/v3/pkg/utils/datetime"
	"gorm.io/gorm"
)

// 分类模型
type Category struct {
	Id          int               `json:"id" gorm:"autoIncrement"`
	Pid         int               `json:"pid"`
	Title       string            `json:"title" gorm:"size:200;not null"`
	Sort        int               `json:"sort" gorm:"size:11;default:0;"`
	CoverId     string            `json:"cover_id" gorm:"size:500;default:null"`
	Name        string            `json:"name" gorm:"size:100;default:null"`
	Description string            `json:"description" gorm:"size:500;default:null"`
	Count       int               `json:"count" gorm:"size:11;default:10;"`
	IndexTpl    string            `json:"index_tpl" gorm:"size:100;"`
	ListTpl     string            `json:"list_tpl" gorm:"size:100;"`
	DetailTpl   string            `json:"detail_tpl" gorm:"size:100;"`
	PageNum     int               `json:"page_num" gorm:"size:11;default:10;"`
	Type        string            `json:"type" gorm:"size:200;not null;default:ARTICLE"`
	Status      int               `json:"status" gorm:"size:1;not null;default:1"`
	CreatedAt   datetime.Datetime `json:"created_at"`
	UpdatedAt   datetime.Datetime `json:"updated_at"`
	DeletedAt   gorm.DeletedAt    `json:"deleted_at"`
}

// Seeder
func (m *Category) Seeder() {

	// 如果菜单已存在，不执行Seeder操作
	if (&appmodel.Menu{}).IsExist(102) {
		return
	}

	// 创建菜单
	menuSeeders := []*appmodel.Menu{
		{Id: 102, Name: "分类列表", GuardName: "admin", Icon: "", Type: 2, Pid: 101, Sort: 0, Path: "/api/admin/category/index", Show: 1, IsEngine: 1, IsLink: 0, Status: 1},
	}
	db.Client.Create(&menuSeeders)

	// 创建默认内容
	seeders := []Category{
		{Title: "默认分类", Name: "default", Type: "ARTICLE", Status: 1},
	}
	db.Client.Create(&seeders)
}

// 获取TreeSelect组件数据
func (model *Category) TreeSelect(root bool) (list []*treeselect.TreeData, Error error) {

	// 是否有根节点
	if root {
		list = append(list, &treeselect.TreeData{
			Title: "根节点",
			Value: 0,
		})
	}

	list = append(list, model.FindTreeSelectNode(0)...)

	return list, nil
}

// 递归获取SelectTree组件数据
func (model *Category) FindTreeSelectNode(pid int) (list []*treeselect.TreeData) {
	categories := []Category{}
	db.Client.
		Where("pid = ?", pid).
		Order("sort asc,id asc").
		Select("title", "id", "pid").
		Find(&categories)

	if len(categories) == 0 {
		return list
	}

	for _, v := range categories {
		item := &treeselect.TreeData{
			Value: v.Id,
			Title: v.Title,
		}

		children := model.FindTreeSelectNode(v.Id)
		if len(children) > 0 {
			item.Children = children
		}

		list = append(list, item)
	}

	return list
}
