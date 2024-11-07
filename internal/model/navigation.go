package model

import (
	"github.com/quarkcloudio/quark-go/v3/dal/db"
	appmodel "github.com/quarkcloudio/quark-go/v3/model"
	"github.com/quarkcloudio/quark-go/v3/service"
	"github.com/quarkcloudio/quark-go/v3/utils/datetime"
	"gorm.io/gorm"
)

// 导航
type Navigation struct {
	Id        int               `json:"id" gorm:"autoIncrement"`
	Pid       int               `json:"pid"`
	Title     string            `json:"title" gorm:"size:200;not null"`
	CoverId   string            `json:"cover_id" gorm:"size:500;default:null"`
	Sort      int               `json:"sort" gorm:"size:11;default:0;"`
	UrlType   int               `json:"url_type" gorm:"size:1;not null;default:1"`
	Url       string            `json:"url" gorm:"size:200;not null"`
	Status    int               `json:"status" gorm:"size:1;not null;default:1"`
	CreatedAt datetime.Datetime `json:"created_at"`
	UpdatedAt datetime.Datetime `json:"updated_at"`
	DeletedAt gorm.DeletedAt    `json:"deleted_at"`
}

// Seeder
func (m *Navigation) Seeder() {

	// 如果菜单已存在，不执行Seeder操作
	if service.NewMenuService().IsExist(109) {
		return
	}

	// 创建菜单
	menuSeeders := []*appmodel.Menu{
		{Id: 109, Name: "导航管理", GuardName: "admin", Icon: "", Type: 2, Pid: 7, Sort: 0, Path: "/api/admin/navigation/index", Show: 1, IsEngine: 1, IsLink: 0, Status: 1},
	}
	db.Client.Create(&menuSeeders)

	// 创建默认内容
	seeders := []Navigation{
		{Title: "默认导航", Status: 1},
	}
	db.Client.Create(&seeders)
}
