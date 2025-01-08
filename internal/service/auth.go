package service

import (
	"errors"

	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/dto"
	"github.com/quarkcloudio/quark-go/v3/model"
)

type AuthService struct {
	ctx *quark.Context
}

func NewAuthService(ctx *quark.Context) *AuthService {
	return &AuthService{ctx}
}

// 获取当前登录用户信息
func (p *AuthService) GetUser() (user model.User, err error) {
	userClaims := dto.UserClaims{}
	err = p.ctx.JwtAuthUser(&userClaims)
	if err != nil {
		return user, err
	}
	userInfo, err := NewUserService().GetInfoById(userClaims.Id)
	if user.Status == 0 {
		return user, errors.New("用户被禁用")
	}
	return userInfo, err
}

// 获取当前登录用户ID
func (p *AuthService) GetUid() (userId int, err error) {
	userClaims, err := p.GetUser()
	return userClaims.Id, err
}
