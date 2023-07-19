package service

import (
	"github.com/quarkcms/quark-smart/internal/admin/service/dashboard"
	"github.com/quarkcms/quark-smart/internal/admin/service/login"
	"github.com/quarkcms/quark-smart/internal/admin/service/resource"
)

// 注册服务
var Provider = []interface{}{
	&login.Index{},
	&dashboard.Index{},
	&resource.Article{},
	&resource.Page{},
	&resource.Category{},
	&resource.Banner{},
	&resource.BannerCategory{},
	&resource.Navigation{},
}