package model

import (
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/component/form/fields/selectfield"
	appmodel "github.com/quarkcloudio/quark-go/v3/pkg/app/admin/model"
	"github.com/quarkcloudio/quark-go/v3/pkg/dal/db"
	"github.com/quarkcloudio/quark-go/v3/pkg/utils/datetime"
	"gorm.io/gorm"
)

// 分类模型
type BannerCategory struct {
	Id        int               `json:"id" gorm:"autoIncrement"`
	Title     string            `json:"title" gorm:"size:200;not null"`
	Name      string            `json:"name" gorm:"size:100;default:null"`
	Width     int               `json:"width" gorm:"size:11;default:0;"`
	Height    int               `json:"height" gorm:"size:11;default:0;"`
	Status    int               `json:"status" gorm:"size:1;not null;default:1"`
	CreatedAt datetime.Datetime `json:"created_at"`
	UpdatedAt datetime.Datetime `json:"updated_at"`
	DeletedAt gorm.DeletedAt    `json:"deleted_at"`
}

// Seeder
func (m *BannerCategory) Seeder() {

	// 如果菜单已存在，不执行Seeder操作
	if (&appmodel.Menu{}).IsExist(106) {
		return
	}

	// 创建菜单
	menuSeeders := []*appmodel.Menu{
		{Id: 106, Name: "广告管理", GuardName: "admin", Icon: "icon-banner", Type: 1, Pid: 0, Sort: 0, Path: "/banner", Show: 1, IsEngine: 0, IsLink: 0, Status: 1},
		{Id: 107, Name: "广告位列表", GuardName: "admin", Icon: "", Type: 2, Pid: 106, Sort: 0, Path: "/api/admin/bannerCategory/index", Show: 1, IsEngine: 1, IsLink: 0, Status: 1},
	}
	db.Client.Create(&menuSeeders)

	// 创建默认内容
	seeders := []BannerCategory{
		{Title: "首页广告位", Name: "indexPage", Status: 1},
	}
	db.Client.Create(&seeders)
}

// 获取列表
func (model *BannerCategory) Options() (options []selectfield.Option, Error error) {
	getList := []BannerCategory{}
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
