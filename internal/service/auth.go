package service

import (
	"errors"

	"github.com/golang-jwt/jwt"
	"github.com/quarkcloudio/quark-go/v3"
	appdto "github.com/quarkcloudio/quark-go/v3/dto"
	"github.com/quarkcloudio/quark-go/v3/utils/datetime"
	"github.com/quarkcloudio/quark-go/v3/utils/hash"
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

// 获取当前登录用户JWT信息
func (p *AuthService) GetUserClaims() (userClaims appdto.UserClaims, err error) {
	userClaims = appdto.UserClaims{}
	err = p.ctx.JwtAuthUser(&userClaims)
	return userClaims, err
}

// 获取当前登录用户信息
func (p *AuthService) GetUser() (userId int, err error) {
	userClaims := appdto.UserClaims{}
	err = p.ctx.JwtAuthUser(&userClaims)
	return userClaims.Id, err
}

// 获取当前登录用户ID
func (p *AuthService) GetUid() (userId int, err error) {
	userClaims, err := p.GetUserClaims()
	return userClaims.Id, err
}

// 账号密码授权
func (p *AuthService) GetTokenByUsername(username, password string) (token string, err error) {
	// 初始化用户服务层
	userService := NewUserService()
	user, err := userService.GetInfoByUsername(username)
	if err != nil {
		return token, err
	}
	if !hash.Check(user.Password, password) {
		return token, errors.New("账号或密码错误")
	}

	// 生成token
	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, userService.GetUserClaims(user)).SignedString(config.App.Key)

	return token, err
}

// 微信小程序授权
func (p *AuthService) GetTokenByWechatMiniProgram(param dto.WechatAuthDTO) (token string, err error) {
	// 获取微信授权用户信息
	wechatUser, err := wechat.NewWechatMiniProgram().GetWechatUser(param.Iv, param.Code, param.EncryptedData)
	if err != nil {
		return token, err
	}

	// 初始化用户服务层
	userService := NewUserService()
	user := userService.GetInfoByWxOpenid(wechatUser.OpenID)
	if user.Id > 0 {
		user, err = userService.Update(dto.SaveUserDTO{
			Id:            user.Id,
			LastLoginIp:   p.ctx.ClientIP(),
			LastLoginTime: datetime.Now(),
		})
	} else {
		user, err = userService.Create(dto.SaveUserDTO{
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

	// 生成token
	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, userService.GetUserClaims(user)).SignedString(config.App.Key)

	return token, err
}

// 微信网页授权
func (p *AuthService) GetTokenByWechatOfficialAccount(param dto.WechatAuthDTO) (token string, err error) {
	// 获取微信授权用户信息
	wechatUser, err := wechat.NewWechatOfficialAccount().GetWechatUser(p.ctx.Request.Context(), param.Code)
	if err != nil {
		return token, err
	}

	// 初始化用户服务层
	userService := NewUserService()
	user := userService.GetInfoByWxOpenid(wechatUser.OpenID)
	if user.Id > 0 {
		user, err = userService.Update(dto.SaveUserDTO{
			Id:            user.Id,
			LastLoginIp:   p.ctx.ClientIP(),
			LastLoginTime: datetime.Now(),
		})
	} else {
		user, err = userService.Create(dto.SaveUserDTO{
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

	// 生成token
	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, userService.GetUserClaims(user)).SignedString(config.App.Key)

	return token, err
}
