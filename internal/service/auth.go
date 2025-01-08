package service

import (
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/dto"
)

type AuthService struct {
	ctx *quark.Context
}

func NewAuthService(ctx *quark.Context) *AuthService {
	return &AuthService{ctx}
}

// 获取当前登录用户JWT信息
func (p *AuthService) GetUserClaims() (userClaims dto.UserClaims, err error) {
	userClaims = dto.UserClaims{}
	err = p.ctx.JwtAuthUser(&userClaims)
	return userClaims, err
}

// 获取当前登录用户信息
func (p *AuthService) GetUser() (userId int, err error) {
	userClaims := dto.UserClaims{}
	err = p.ctx.JwtAuthUser(&userClaims)
	return userClaims.Id, err
}

// 获取当前登录用户ID
func (p *AuthService) GetUid() (userId int, err error) {
	userClaims, err := p.GetUserClaims()
	return userClaims.Id, err
}
