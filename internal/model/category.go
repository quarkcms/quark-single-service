package model

import (
	"github.com/quarkcloudio/quark-go/v3/dal/db"
	appmodel "github.com/quarkcloudio/quark-go/v3/model"
	"github.com/quarkcloudio/quark-go/v3/service"
	"github.com/quarkcloudio/quark-go/v3/utils/datetime"
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
	if service.NewMenuService().IsExist(102) {
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
