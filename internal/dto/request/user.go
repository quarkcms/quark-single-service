package request

// 更新用户信息
type UpdateUserReq struct {
	Nickname string `json:"nickname"`
	Phone    string `json:"phone"`
	Avatar   string `json:"avatar"`
}
