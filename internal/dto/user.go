package dto

// 微信小程序授权
type WechatMiniProgramDTO struct {
	Iv            string // 小程序授权所需的参数，由前端传递
	Code          string // 小程序授权所需的参数，由前端传递
	EncryptedData string // 小程序授权所需的参数，由前端传递
	ClientIp      string // 要授权用户的ip地址，用于更新最后一次登录ip
}

// 微信公众号/网页授权
type WechatOfficialAccountDTO struct {
	Code     string // 小程序授权所需的参数，由前端传递
	ClientIp string // 要授权用户的ip地址，用于更新最后一次登录ip
}
