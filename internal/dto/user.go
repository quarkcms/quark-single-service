package dto

import "github.com/quarkcloudio/quark-go/v3/utils/datetime"

// 保存用户
type SaveUserDTO struct {
	Id            int
	Username      string
	Nickname      string
	Sex           int
	Email         string
	Phone         string
	Password      string
	Avatar        string
	DepartmentId  int
	PositionIds   string
	LastLoginIp   string
	LastLoginTime datetime.Datetime
	WxOpenid      string
	WxUnionid     string
	Status        int
}
