package request

// 验证码
type Captcha struct {
	Id    string `json:"id" form:"id"`
	Value string `json:"value" form:"value"`
}

// 登录
type LoginReq struct {
	Username string  `json:"username" form:"username"`
	Password string  `json:"password" form:"password"`
	Captcha  Captcha `json:"captcha" form:"captcha"`
}
