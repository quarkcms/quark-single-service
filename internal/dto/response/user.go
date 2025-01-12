package response

// 用户信息
type UserInfoResp struct {
	Id       int    `json:"id"`
	Nickname string `json:"nickname"`
	Phone    string `json:"phone"`
	Avatar   string `json:"avatar"`
}
