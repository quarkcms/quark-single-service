package model

import (
	"github.com/quarkcloudio/quark-go/v3/dal/db"
	appmodel "github.com/quarkcloudio/quark-go/v3/model"
	"github.com/quarkcloudio/quark-go/v3/service"
	"github.com/quarkcloudio/quark-go/v3/utils/datetime"
	"gorm.io/gorm"
)

type Item struct {
	Id          uint              `json:"id" gorm:"primaryKey;autoIncrement;comment:商品id"`
	MerId       uint              `json:"mer_id" gorm:"not null;default:0;comment:商户Id(0为总后台管理员创建,不为0的时候是商户后台创建)"`
	Image       string            `json:"image" gorm:"not null;size:256;comment:商品图片"`
	SliderImage string            `json:"slider_image" gorm:"not null;size:2000;comment:轮播图"`
	Name        string            `json:"name" gorm:"not null;size:128;comment:商品名称"`
	Keyword     string            `json:"keyword" gorm:"not null;size:256;comment:关键字"`
	Description string            `json:"description" gorm:"not null;size:256;comment:商品简介"`
	Content     string            `json:"content" gorm:"type:text;default:null;comment:商品详情"`
	CategoryIds string            `json:"category_ids" gorm:"not null;size:64;comment:分类ids"`
	Price       float64           `json:"price" gorm:"not null;type:decimal(8,2);default:0.00;comment:商品价格"`
	OtPrice     float64           `json:"ot_price" gorm:"not null;type:decimal(8,2);default:0.00;comment:市场价"`
	Sort        int16             `json:"sort" gorm:"not null;default:0;comment:排序"`
	Sales       uint              `json:"sales" gorm:"not null;default:0;comment:销量"`
	Stock       uint              `json:"stock" gorm:"not null;default:0;comment:库存"`
	Status      uint8             `json:"status" gorm:"not null;default:1;comment:状态(0:未上架,1:上架)"`
	Cost        float64           `json:"cost" gorm:"not null;type:decimal(8,2);default:0.00;comment:成本价"`
	Ficti       int               `json:"ficti" gorm:"default:100;comment:虚拟销量"`
	Views       int               `json:"views" gorm:"default:0;comment:浏览量"`
	SpecType    uint8             `json:"spec_type" gorm:"not null;default:0;comment:规格:0单,1多"`
	Deadline    datetime.Datetime `json:"deadline"`
	CreatedAt   datetime.Datetime `json:"created_at"`
	UpdatedAt   datetime.Datetime `json:"updated_at"`
	DeletedAt   gorm.DeletedAt    `json:"deleted_at"`
}

// Seeder
func (m *Item) Seeder() {

	// 如果菜单已存在，不执行Seeder操作
	if service.NewMenuService().IsExist(90) {
		return
	}

	// 创建菜单
	menuSeeders := []*appmodel.Menu{
		{Id: 90, Name: "商品管理", GuardName: "admin", Icon: "icon-shop", Type: 1, Pid: 0, Sort: 0, Path: "/item", Show: 1, IsEngine: 0, IsLink: 0, Status: 1},
		{Id: 91, Name: "商品列表", GuardName: "admin", Icon: "", Type: 2, Pid: 90, Sort: 0, Path: "/api/admin/item/index", Show: 1, IsEngine: 1, IsLink: 0, Status: 1},
		{Id: 92, Name: "商品分类", GuardName: "admin", Icon: "", Type: 2, Pid: 90, Sort: 0, Path: "/api/admin/itemCategory/index", Show: 1, IsEngine: 1, IsLink: 0, Status: 1},
	}
	db.Client.Create(&menuSeeders)
}
