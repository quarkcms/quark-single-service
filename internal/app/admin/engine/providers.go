package service

import (
	"github.com/quarkcloudio/quark-smart/v2/internal/app/admin/engine/dashboard"
	"github.com/quarkcloudio/quark-smart/v2/internal/app/admin/engine/layout"
	"github.com/quarkcloudio/quark-smart/v2/internal/app/admin/engine/login"
	"github.com/quarkcloudio/quark-smart/v2/internal/app/admin/engine/resource"
	"github.com/quarkcloudio/quark-smart/v2/internal/app/admin/engine/upload"
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
