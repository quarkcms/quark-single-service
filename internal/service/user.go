package service

import (
	"context"
	"errors"

	"github.com/golang-jwt/jwt/v4"
	"github.com/quarkcloudio/quark-go/v3/dal/db"
	appDTO "github.com/quarkcloudio/quark-go/v3/dto"
	"github.com/quarkcloudio/quark-go/v3/model"
	appservice "github.com/quarkcloudio/quark-go/v3/service"
	"github.com/quarkcloudio/quark-go/v3/utils/datetime"
	"github.com/quarkcloudio/quark-smart/v2/config"
	"github.com/quarkcloudio/quark-smart/v2/internal/dto"
	"github.com/quarkcloudio/quark-smart/v2/pkg/wechat"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

// 获取普通用户JWT信息
func (p *UserService) GetUserClaims(userInfo model.User) *appDTO.UserClaims {
	return appservice.NewUserService().GetUserClaims(userInfo)
}

// 获取当前认证的用户信息，默认参数为tokenString
func (p *UserService) GetAuthUser(appKey string, tokenString string) (userClaims *appDTO.UserClaims, err error) {
	return appservice.NewUserService().GetAuthUser(appKey, tokenString)
}

// 通过ID获取用户信息
func (p *UserService) GetInfoById(id interface{}) (user model.User, err error) {
	return appservice.NewUserService().GetInfoById(id)
}

// 通过用户名获取用户信息
func (p *UserService) GetInfoByUsername(username string) (user model.User, err error) {
	return appservice.NewUserService().GetInfoByUsername(username)
}

// 更新最后一次登录数据
func (p *UserService) UpdateLastLogin(uid int, lastLoginIp string, lastLoginTime datetime.Datetime) error {
	return appservice.NewUserService().UpdateLastLogin(uid, lastLoginIp, lastLoginTime)
}

// 通过wxopenid获取用户信息
func (p *UserService) GetInfoByWxOpenid(wxOpenid string) (user model.User) {
	db.Client.Model(model.User{}).Where("wx_openid = ?", wxOpenid).Last(&user)
	return user
}

// 微信小程序授权
func (p *UserService) AuthByWechatMiniProgram(param dto.WechatMiniProgramDTO) (token string, err error) {
	// 初始化微信小程序
	mini := wechat.NewWechatMiniProgram()
	// 登录凭证校验
	authResponse, err := mini.GetAuth().Code2Session(param.Code)
	if err != nil {
		return token, err
	}
	if authResponse.ErrCode != 0 {
		return token, errors.New(authResponse.ErrMsg)
	}

	// 解密数据，获取微信用户信息
	plainData, err := mini.GetEncryptor().Decrypt(authResponse.SessionKey, param.EncryptedData, param.Iv)
	if err != nil {
		return token, err
	}

	user := p.GetInfoByWxOpenid(authResponse.OpenID)
	if user.Id > 0 {
		if err = db.Client.Model(model.User{}).Where("id = ?", user.Id).Updates(&model.User{
			LastLoginIp:   param.ClientIp,
			LastLoginTime: datetime.Now(),
		}).Error; err != nil {
			return token, err
		}
	} else {
		user = model.User{
			Nickname:      plainData.NickName,
			Sex:           plainData.Gender,
			Avatar:        plainData.AvatarURL,
			WxOpenid:      authResponse.OpenID,
			WxUnionid:     authResponse.UnionID,
			LastLoginIp:   param.ClientIp,
			LastLoginTime: datetime.Now(),
		}
		if err = db.Client.Model(model.User{}).Create(&user).Error; err != nil {
			return token, err
		}
	}

	// 生成token
	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, p.GetUserClaims(user)).SignedString(config.App.Key)

	return token, err
}

// 微信网页授权
func (p *UserService) AuthByWechatOfficialAccount(param dto.WechatOfficialAccountDTO) (token string, err error) {
	// 初始化网页授权
	auth := wechat.NewWechatOfficialAccount().GetOauth()
	authResponse, err := auth.GetUserInfoByCodeContext(context.Background(), param.Code)
	if err != nil {
		return token, err
	}
	if authResponse.ErrCode != 0 {
		return token, errors.New(authResponse.ErrMsg)
	}

	user := p.GetInfoByWxOpenid(authResponse.OpenID)
	if user.Id > 0 {
		if err = db.Client.Model(model.User{}).Where("id = ?", user.Id).Updates(&model.User{
			LastLoginIp:   param.ClientIp,
			LastLoginTime: datetime.Now(),
		}).Error; err != nil {
			return token, err
		}
	} else {
		user = model.User{
			Nickname:      authResponse.Nickname,
			Avatar:        authResponse.HeadImgURL,
			Sex:           int(authResponse.Sex),
			WxOpenid:      authResponse.OpenID,
			WxUnionid:     authResponse.Unionid,
			LastLoginIp:   param.ClientIp,
			LastLoginTime: datetime.Now(),
		}
		if err = db.Client.Model(model.User{}).Create(&user).Error; err != nil {
			return token, err
		}
	}

	// 生成token
	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, p.GetUserClaims(user)).SignedString(config.App.Key)

	return token, err
}
