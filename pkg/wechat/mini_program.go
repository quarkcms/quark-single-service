package wechat

import (
	"errors"

	"github.com/quarkcloudio/quark-smart/v2/pkg/utils"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram"
	"github.com/silenceper/wechat/v2/miniprogram/config"
	"github.com/silenceper/wechat/v2/miniprogram/encryptor"
)

type WechatMiniProgram struct {
	mini *miniprogram.MiniProgram
}

// 初始化微信小程序
func NewWechatMiniProgram() *WechatMiniProgram {
	// 首先在 configs 表中添加 name 为 WECHAT_APP_ID 和 WECHAT_APP_SECRET 的记录
	return &WechatMiniProgram{
		mini: wechat.NewWechat().GetMiniProgram(&config.Config{
			AppID:     utils.GetConfig("WECHAT_APP_ID"),
			AppSecret: utils.GetConfig("WECHAT_APP_SECRET"),
			Cache:     cache.NewMemcache(),
		}),
	}
}

// 获取微信授权用户信息
func (p *WechatMiniProgram) GetWechatUser(iv, code, encryptedData string) (encryptor.PlainData, error) {
	// 登录凭证校验
	authResponse, err := p.mini.GetAuth().Code2Session(code)
	if err != nil {
		return encryptor.PlainData{}, err
	}
	if authResponse.ErrCode != 0 {
		return encryptor.PlainData{}, errors.New(authResponse.ErrMsg)
	}
	// 解密数据，获取微信用户信息
	wechatUser, err := p.mini.GetEncryptor().Decrypt(authResponse.SessionKey, encryptedData, iv)
	if err != nil {
		return encryptor.PlainData{}, err
	}

	// 设置微信用户openid&unionid
	wechatUser.OpenID = authResponse.OpenID
	wechatUser.UnionID = authResponse.UnionID

	return *wechatUser, nil
}
