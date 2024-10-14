package middleware

import (
	"github.com/quarkcloudio/quark-go/v3/pkg/builder"
	"github.com/quarkcloudio/quark-smart/v2/config"
)

// 结构体
type AppMiddleware struct{}

// 中间件
func (am *AppMiddleware) Handle(ctx *builder.Context) error {
	if config.App.Env == "demo" {
		isForbiddenRoute := false
		forbiddenRoutes := []string{
			"/api/admin/admin/store",
			"/api/admin/admin/save",
			"/api/admin/admin/delete",
			"/api/admin/admin/editable",
			"/api/admin/admin/action/delete",
			"/api/admin/admin/action/change-status",
			"/api/admin/menu/store",
			"/api/admin/menu/save",
			"/api/admin/menu/delete",
			"/api/admin/menu/editable",
			"/api/admin/menu/action/delete",
			"/api/admin/menu/action/change-status",
			"/api/admin/account/action/change-account",
		}
		for _, forbiddenRoute := range forbiddenRoutes {
			if ctx.Path() == forbiddenRoute {
				isForbiddenRoute = true
			}
		}
		if isForbiddenRoute {
			return ctx.JSON(200, builder.Error("演示站点禁止了此操作！"))
		}
	}

	return ctx.Next()
}
