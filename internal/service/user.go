package service

import (
	"github.com/quarkcloudio/quark-go/v3/dto"
	"github.com/quarkcloudio/quark-go/v3/model"
	appservice "github.com/quarkcloudio/quark-go/v3/service"
	"github.com/quarkcloudio/quark-go/v3/utils/datetime"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

// 获取普通用户JWT信息
func (p *UserService) GetUserClaims(userInfo model.User) *dto.UserClaims {
	return appservice.NewUserService().GetUserClaims(userInfo)
}

// 获取当前认证的用户信息，默认参数为tokenString
func (p *UserService) GetAuthUser(appKey string, tokenString string) (userClaims *dto.UserClaims, err error) {
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
