package wechat

import (
	"github.com/quarkcloudio/quark-smart/v2/pkg/utils"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/officialaccount"
	"github.com/silenceper/wechat/v2/officialaccount/config"
)

// 初始化微信公众号
func NewWechatOfficialAccount() *officialaccount.OfficialAccount {
	// 首先在 configs 表中添加 name 为 WECHAT_APP_ID 和 WECHAT_APP_SECRET 的记录
	return wechat.NewWechat().GetOfficialAccount(&config.Config{
		AppID:     utils.GetConfig("WECHAT_APP_ID"),
		AppSecret: utils.GetConfig("WECHAT_APP_SECRET"),
		Cache:     cache.NewMemcache(),
	})
}
