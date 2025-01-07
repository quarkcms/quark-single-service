package service

import (
	"github.com/quarkcloudio/quark-smart/v2/internal/app/admin/dashboard"
	"github.com/quarkcloudio/quark-smart/v2/internal/app/admin/layout"
	"github.com/quarkcloudio/quark-smart/v2/internal/app/admin/login"
	"github.com/quarkcloudio/quark-smart/v2/internal/app/admin/resource"
	"github.com/quarkcloudio/quark-smart/v2/internal/app/admin/upload"
)

// 注册服务
var Providers = []interface{}{
	&login.Index{},
	&dashboard.Index{},
	&layout.Index{},
	&resource.Article{},
	&resource.Page{},
	&resource.Category{},
	&resource.Banner{},
	&resource.BannerCategory{},
	&resource.Navigation{},
	&upload.File{},
	&upload.Image{},
}

// 注册高级功能服务
var ProProviders = []interface{}{
	&resource.Item{},
	&resource.ItemCategory{},
	&resource.Order{},
	&resource.RefundOrder{},
	&resource.VerifyOrder{},
	&resource.Bill{},
	&resource.BillRecord{},
	&resource.BillDetail{},
}
