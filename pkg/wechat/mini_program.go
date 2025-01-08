package wechat

import (
	"github.com/quarkcloudio/quark-smart/v2/pkg/utils"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram"
	"github.com/silenceper/wechat/v2/miniprogram/config"
)

// 初始化微信小程序
func NewWechatMiniProgram() *miniprogram.MiniProgram {
	// 首先在 configs 表中添加 name 为 WECHAT_APP_ID 和 WECHAT_APP_SECRET 的记录
	return wechat.NewWechat().GetMiniProgram(&config.Config{
		AppID:     utils.GetConfig("WECHAT_APP_ID"),
		AppSecret: utils.GetConfig("WECHAT_APP_SECRET"),
		Cache:     cache.NewMemcache(),
	})
}
