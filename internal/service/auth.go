package service

import (
	"errors"

	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/model"
	appservice "github.com/quarkcloudio/quark-go/v3/service"
	"github.com/quarkcloudio/quark-go/v3/utils/datetime"
	"github.com/quarkcloudio/quark-smart/v2/config"
	"github.com/quarkcloudio/quark-smart/v2/internal/dto"
	"github.com/quarkcloudio/quark-smart/v2/pkg/wechat"
)

type AuthService struct {
	ctx *quark.Context
}

func NewAuthService(ctx *quark.Context) *AuthService {
	return &AuthService{ctx}
}

// 获取当前登录用户信息
func (p *AuthService) GetUser() (user model.User, err error) {
	return appservice.NewAuthService(p.ctx).GetUser()
}

// 获取当前登录用户ID
func (p *AuthService) GetUid() (userId int, err error) {
	return appservice.NewAuthService(p.ctx).GetUid()
}

// 模拟登录
func (p *AuthService) MockLogin() (token string, err error) {
	if !(config.App.Env == "develop" || config.App.Env == "dev" || config.App.Env == "development") {
		return "", errors.New("it must be a development environment")
	}
	uid := p.ctx.Query("uid", 1)
	user, err := NewUserService().GetInfoById(uid)
	if err != nil {
		return "", err
	}
	return appservice.NewAuthService(p.ctx).MakeToken(user, "user", 60*24*60*60)
}

// 账号密码授权
func (p *AuthService) Login(username, password string) (token string, err error) {
	return appservice.NewAuthService(p.ctx).UserLogin(username, password)
}

// 微信小程序授权
func (p *AuthService) WechatMPLogin(param dto.WechatAuthDTO) (token string, err error) {
	// 获取微信授权用户信息
	wechatUser, err := wechat.NewWechatMiniProgram().GetWechatUser(param.Iv, param.Code, param.EncryptedData)
	if err != nil {
		return token, err
	}

	// 初始化用户服务层
	userService := NewUserService()
	user := userService.GetInfoByWxOpenid(wechatUser.OpenID)
	if user.Id > 0 {
		user, err = userService.UpdateUser(dto.SaveUserDTO{
			Id:            user.Id,
			LastLoginIp:   p.ctx.ClientIP(),
			LastLoginTime: datetime.Now(),
		})
	} else {
		user, err = userService.CreateUser(dto.SaveUserDTO{
			Nickname:      wechatUser.NickName,
			Sex:           wechatUser.Gender,
			Avatar:        wechatUser.AvatarURL,
			WxOpenid:      wechatUser.OpenID,
			WxUnionid:     wechatUser.UnionID,
			LastLoginIp:   p.ctx.ClientIP(),
			LastLoginTime: datetime.Now(),
		})
	}
	if err != nil {
		return token, err
	}

	return appservice.NewAuthService(p.ctx).MakeToken(user, "user", 24*60*60)
}

// 微信网页授权
func (p *AuthService) WechatOALogin(param dto.WechatAuthDTO) (token string, err error) {
	// 获取微信授权用户信息
	wechatUser, err := wechat.NewWechatOfficialAccount().GetWechatUser(p.ctx.Request.Context(), param.Code)
	if err != nil {
		return token, err
	}

	// 初始化用户服务层
	userService := NewUserService()
	user := userService.GetInfoByWxOpenid(wechatUser.OpenID)
	if user.Id > 0 {
		user, err = userService.UpdateUser(dto.SaveUserDTO{
			Id:            user.Id,
			LastLoginIp:   p.ctx.ClientIP(),
			LastLoginTime: datetime.Now(),
		})
	} else {
		user, err = userService.CreateUser(dto.SaveUserDTO{
			Nickname:      wechatUser.Nickname,
			Sex:           int(wechatUser.Sex),
			Avatar:        wechatUser.HeadImgURL,
			WxOpenid:      wechatUser.OpenID,
			WxUnionid:     wechatUser.Unionid,
			LastLoginIp:   p.ctx.ClientIP(),
			LastLoginTime: datetime.Now(),
		})
	}
	if err != nil {
		return token, err
	}

	return appservice.NewAuthService(p.ctx).MakeToken(user, "user", 24*60*60)
}
