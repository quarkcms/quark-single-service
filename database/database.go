package database

import (
	"github.com/quarkcloudio/quark-go/v3/dal/db"
	"github.com/quarkcloudio/quark-smart/v2/internal/model"
)

// 执行数据库操作
func Handle() {

	// 迁移数据
	db.Client.AutoMigrate(
		&model.Post{},
		&model.Category{},
		&model.Banner{},
		&model.BannerCategory{},
		&model.Navigation{},
	)

	// 数据填充
	(&model.Post{}).Seeder()
	(&model.Category{}).Seeder()
	(&model.Banner{}).Seeder()
	(&model.BannerCategory{}).Seeder()
	(&model.Navigation{}).Seeder()
}

// 高级功能
func MiniAppHandle() {
	// 迁移数据
	db.Client.AutoMigrate(
		&model.Item{},
		&model.ItemAttr{},
		&model.ItemAttrValue{},
		&model.ItemCategory{},
		&model.Order{},
		&model.OrderDetail{},
		&model.OrderStatus{},
		&model.Bill{},
		&model.BillRecord{},
	)

	// 数据填充
	(&model.Item{}).Seeder()
	(&model.Order{}).Seeder()
	(&model.Bill{}).Seeder()
}
