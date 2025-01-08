package wechat

import (
	"log"

	"github.com/silenceper/wechat/v2/officialaccount"
	"github.com/silenceper/wechat/v2/officialaccount/message"
)

// 微信模板消息
type WechatTemplateMessage struct {
	officialaccount *officialaccount.OfficialAccount
}

// 消息结构体
type Message struct {
	TemplateID string
	ToUser     string
	Data       map[string]*message.TemplateDataItem
}

// 初始化微信模板消息
func NewWechatTemplateMessage() *WechatTemplateMessage {
	return &WechatTemplateMessage{
		officialaccount: NewWechatOfficialAccount().officialaccount,
	}
}

// 发送
func (p *WechatTemplateMessage) Send(msg *Message) {
	if _, err := p.officialaccount.GetTemplate().Send(&message.TemplateMessage{
		ToUser:     msg.ToUser,
		TemplateID: msg.TemplateID,
		Data:       msg.Data,
		// URL:        utils.GetDomain() + "/pages/student/registration/detail/detail?id=", // 点击模板消息跳转的页面
	}); err != nil {
		log.Println("发送模板消息失败，错误信息：", err)
	}
}
