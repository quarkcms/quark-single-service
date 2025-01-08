package wechat

import (
	"context"
	"errors"

	"github.com/quarkcloudio/quark-smart/v2/pkg/utils"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/officialaccount"
	"github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/oauth"
)

type WechatOfficialAccount struct {
	officialaccount *officialaccount.OfficialAccount
}

// 初始化微信公众号
func NewWechatOfficialAccount() *WechatOfficialAccount {
	// 首先在 configs 表中添加 name 为 WECHAT_APP_ID 和 WECHAT_APP_SECRET 的记录
	return &WechatOfficialAccount{
		officialaccount: wechat.NewWechat().GetOfficialAccount(&config.Config{
			AppID:     utils.GetConfig("WECHAT_APP_ID"),
			AppSecret: utils.GetConfig("WECHAT_APP_SECRET"),
			Cache:     cache.NewMemcache(),
		}),
	}
}

// 获取微信授权用户信息
func (p *WechatOfficialAccount) GetWechatUser(ctx context.Context, code string) (oauth.UserInfo, error) {
	// 登录凭证校验
	wechatUser, err := p.officialaccount.GetOauth().GetUserInfoByCodeContext(ctx, code)
	if err != nil {
		return oauth.UserInfo{}, err
	}
	if wechatUser.ErrCode != 0 {
		return oauth.UserInfo{}, errors.New(wechatUser.ErrMsg)
	}

	return wechatUser, nil
}
