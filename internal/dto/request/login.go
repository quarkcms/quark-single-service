package request

// 验证码
type Captcha struct {
	Id    string `json:"id"`
	Value string `json:"value"`
}

// 登录
type LoginReq struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
	Captcha  Captcha `json:"captcha"`
}
