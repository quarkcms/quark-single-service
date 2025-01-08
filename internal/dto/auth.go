package dto

// 微信授权
type WechatAuthDTO struct {
	Code          string // 小程序、网页公众号授权所需的参数，由前端传递
	Iv            string // 小程序授权所需的参数，由前端传递
	EncryptedData string // 小程序授权所需的参数，由前端传递
}
