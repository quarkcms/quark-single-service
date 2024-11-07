package model

import (
	"github.com/quarkcloudio/quark-go/v3/dal/db"
	appmodel "github.com/quarkcloudio/quark-go/v3/model"
	"github.com/quarkcloudio/quark-go/v3/service"
	"github.com/quarkcloudio/quark-go/v3/utils/datetime"
	"gorm.io/gorm"
)

// 文章模型
type Post struct {
	Id            int               `json:"id" gorm:"autoIncrement"`
	Adminid       int               `json:"adminid"`
	Uid           int               `json:"uid"`
	CategoryId    int               `json:"category_id"`
	Tags          string            `json:"tags" gorm:"size:200;default:null"`
	Title         string            `json:"title" gorm:"size:200;not null"`
	Name          string            `json:"name" gorm:"size:200;default:null"`
	Author        string            `json:"author" gorm:"size:200;default:null"`
	Source        string            `json:"source" gorm:"size:200;default:null"`
	Description   string            `json:"description" gorm:"size:200;default:null"`
	Password      string            `json:"password" gorm:"size:200;default:null"`
	CoverIds      string            `json:"cover_ids" gorm:"size:1000;default:null"`
	Pid           int               `json:"pid" gorm:"default:0"`
	Level         int               `json:"level" gorm:"size:11;default:0"`
	Type          string            `json:"type" gorm:"size:200;not null;default:ARTICLE"`
	ShowType      int               `json:"show_type" gorm:"size:4;default:0"`
	Position      string            `json:"position" gorm:"size:100;default:null"`
	Link          string            `json:"link" gorm:"size:100;default:null"`
	Content       string            `json:"content" gorm:"type:text;default:null"`
	Comment       int               `json:"comment" gorm:"default:0"`
	View          int               `json:"view" gorm:"default:0"`
	PageTpl       string            `json:"page_tpl" gorm:"size:100"`
	CommentStatus int               `json:"comment_status" gorm:"size:1;not null;default:0"`
	FileIds       string            `json:"file_ids" gorm:"size:1000;default:null"`
	Status        int               `json:"status" gorm:"size:1;not null;default:1"`
	CreatedAt     datetime.Datetime `json:"created_at"`
	UpdatedAt     datetime.Datetime `json:"updated_at"`
	DeletedAt     gorm.DeletedAt    `json:"deleted_at"`
}

// Seeder
func (m *Post) Seeder() {

	// 如果菜单已存在，不执行Seeder操作
	if service.NewMenuService().IsExist(101) {
		return
	}

	// 创建菜单
	menuSeeders := []*appmodel.Menu{
		{Id: 101, Name: "内容管理", GuardName: "admin", Icon: "icon-read", Type: 1, Pid: 0, Sort: 0, Path: "/post", Show: 1, IsEngine: 0, IsLink: 0, Status: 1},
		{Id: 103, Name: "文章列表", GuardName: "admin", Icon: "", Type: 2, Pid: 101, Sort: 0, Path: "/api/admin/article/index", Show: 1, IsEngine: 1, IsLink: 0, Status: 1},
		{Id: 104, Name: "单页管理", GuardName: "admin", Icon: "icon-page", Type: 1, Pid: 0, Sort: 0, Path: "/page", Show: 1, IsEngine: 0, IsLink: 0, Status: 1},
		{Id: 105, Name: "单页列表", GuardName: "admin", Icon: "", Type: 2, Pid: 104, Sort: 0, Path: "/api/admin/page/index", Show: 1, IsEngine: 1, IsLink: 0, Status: 1},
	}
	db.Client.Create(&menuSeeders)

	// 创建默认内容
	seeders := []Post{
		{Title: "关于我们", Name: "aboutus", Content: "关于我们", Status: 1, Type: "PAGE"},
	}
	db.Client.Create(&seeders)
}
