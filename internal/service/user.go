package service

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/quarkcloudio/quark-go/v3/dto"
	"github.com/quarkcloudio/quark-go/v3/model"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

// 获取普通用户JWT信息
func (p *UserService) GetUserClaims(userInfo model.User) *dto.UserClaims {
	return &dto.UserClaims{
		Id:        userInfo.Id,
		Username:  userInfo.Username,
		Nickname:  userInfo.Nickname,
		Sex:       userInfo.Sex,
		Email:     userInfo.Email,
		Phone:     userInfo.Phone,
		Avatar:    userInfo.Avatar,
		GuardName: "user",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 过期时间，默认24小时
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     // 颁发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                     // 不早于时间
			Issuer:    "QuarkGo",                                          // 颁发人
			Subject:   "User Token",                                       // 主题信息
		},
	}
}
