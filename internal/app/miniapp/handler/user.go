package handler

import (
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-smart/v2/internal/dto"
	"github.com/quarkcloudio/quark-smart/v2/internal/dto/request"
	"github.com/quarkcloudio/quark-smart/v2/internal/dto/response"
	"github.com/quarkcloudio/quark-smart/v2/internal/service"
	"github.com/quarkcloudio/quark-smart/v2/pkg/utils"
)

// 结构体
type User struct{}

// 用户中心
func (p *User) Index(ctx *quark.Context) error {
	uid, _ := service.NewAuthService(ctx).GetUid()
	user, _ := service.NewUserService().GetInfoById(uid)
	userInfo := response.UserInfoResp{
		Id:       user.Id,
		Nickname: user.Nickname,
		Phone:    user.Phone,
		Avatar:   utils.GetImagePath(user.Avatar),
	}
	return ctx.JSONOk("ok", userInfo)
}

// 更新用户信息
func (p *User) Save(ctx *quark.Context) error {
	var param request.UpdateUserReq
	if err := ctx.Bind(&param); err != nil {
		return ctx.JSONError(err.Error())
	}

	// 参数校验
	if param.Phone != "" && !utils.CheckRegex("^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\\d{8}$", param.Phone) {
		return ctx.JSONError("手机号格式不正确")
	}

	uid, _ := service.NewAuthService(ctx).GetUid()
	if _, err := service.NewUserService().UpdateUser(dto.SaveUserDTO{
		Id:       uid,
		Nickname: param.Nickname,
		Phone:    param.Phone,
		Avatar:   param.Avatar,
	}); err != nil {
		return ctx.JSONError("更新用户信息失败")
	}
	return ctx.JSONOk("更新成功")
}

// 注销用户信息
func (p *User) Delete(ctx *quark.Context) error {
	uid, _ := service.NewAuthService(ctx).GetUid()
	if err := service.NewUserService().DeleteUser(uid); err != nil {
		return ctx.JSONError("注销失败")
	}
	return ctx.JSONOk("注销成功")
}
