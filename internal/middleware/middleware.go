package middleware

import (
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-smart/v2/config"
	"github.com/quarkcloudio/quark-smart/v2/internal/service"
)

// 中间件
func AppMiddleware(ctx *quark.Context) error {
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
			return ctx.JSON(200, quark.Error("演示站点禁止了此操作！"))
		}
	}
	return ctx.Next()
}

// MiniApp中间件
func MiniAppMiddleware(ctx *quark.Context) error {
	_, err := service.NewAuthService(ctx).GetUser()
	if err != nil {
		return ctx.JSON(401, quark.Error(err.Error()))
	}
	return ctx.Next()
}
