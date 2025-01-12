package service

import (
	"github.com/quarkcloudio/quark-go/v3/dal/db"
	"github.com/quarkcloudio/quark-go/v3/model"
	appservice "github.com/quarkcloudio/quark-go/v3/service"
	"github.com/quarkcloudio/quark-go/v3/utils/datetime"
	"github.com/quarkcloudio/quark-smart/v2/internal/dto"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
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

// 新增用户
func (p *UserService) CreateUser(param dto.SaveUserDTO) (model.User, error) {
	user := model.User{
		Username:      param.Username,
		Nickname:      param.Nickname,
		Sex:           param.Sex,
		Email:         param.Email,
		Phone:         param.Phone,
		Password:      param.Password,
		Avatar:        param.Avatar,
		DepartmentId:  param.DepartmentId,
		PositionIds:   param.PositionIds,
		LastLoginIp:   param.LastLoginIp,
		LastLoginTime: param.LastLoginTime,
		WxOpenid:      param.WxOpenid,
		WxUnionid:     param.WxUnionid,
		Status:        param.Status,
	}
	if err := db.Client.Model(model.User{}).Create(&user).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

// 更新用户
func (p *UserService) UpdateUser(param dto.SaveUserDTO) (model.User, error) {
	user := model.User{
		Id:            param.Id,
		Username:      param.Username,
		Nickname:      param.Nickname,
		Sex:           param.Sex,
		Email:         param.Email,
		Phone:         param.Phone,
		Password:      param.Password,
		Avatar:        param.Avatar,
		DepartmentId:  param.DepartmentId,
		PositionIds:   param.PositionIds,
		LastLoginIp:   param.LastLoginIp,
		LastLoginTime: param.LastLoginTime,
		WxOpenid:      param.WxOpenid,
		WxUnionid:     param.WxUnionid,
		Status:        param.Status,
	}
	if err := db.Client.Model(model.User{}).Where("id = ?", user.Id).Updates(&user).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

// 删除用户
func (p *UserService) DeleteUser(id int) error {
	return db.Client.Model(model.User{}).Where("id = ?", id).Delete(&model.User{}).Error
}
