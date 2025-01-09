package pay

import (
	"context"
	"errors"
	"io/ioutil"
	"log"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"github.com/quarkcloudio/quark-smart/v2/pkg/utils"
)

// gopay 文档：https://github.com/go-pay/gopay/blob/main/doc/alipay_v3.md

type AliPay struct {
	Config *AliPayConfig
	Client *alipay.Client
}

type AliPayConfig struct {
	AppId                string // 应用ID
	PrivateKeyPath       string // 应用私钥文件路径，支持PKCS1和PKCS8
	IsProd               bool   // 是否是正式环境，沙箱环境请选择新版沙箱应用
	AppPublicCertPath    string // appPublicCert.crt 文件路径
	AlipayRootCertPath   string // alipayRootCert 文件路径
	AlipayPublicCertPath string // alipayPublicCert.crt 文件路径
}

// 初始化支付宝支付客户端
//
// 如果使用默认配置首先需要在 configs 表中添加 name 为 ALI_PAY_APP_ID、ALI_PAY_PRIVATE_KEY_PATH、ALI_PAY_IS_PROD、ALI_PAY_APP_PUBLIC_CERT_PATH、ALI_PAY_ROOT_CERT_PATH、ALI_PAY_PUBLIC_CERT_PATH 的记录
func NewAliPay(param ...AliPayConfig) *AliPay {
	var config AliPayConfig
	if len(param) <= 0 {
		config = AliPayConfig{
			AppId:                utils.GetConfig("ALI_PAY_APP_ID"),
			PrivateKeyPath:       utils.GetConfig("ALI_PAY_PRIVATE_KEY_PATH"),
			IsProd:               utils.GetConfig("ALI_PAY_IS_PROD") == "true",
			AppPublicCertPath:    utils.GetConfig("ALI_PAY_APP_PUBLIC_CERT_PATH"),
			AlipayRootCertPath:   utils.GetConfig("ALI_PAY_ROOT_CERT_PATH"),
			AlipayPublicCertPath: utils.GetConfig("ALI_PAY_PUBLIC_CERT_PATH"),
		}
	} else {
		config = param[0]
	}

	// 读取私钥内容
	privateKeyBytes, err := ioutil.ReadFile(config.PrivateKeyPath)
	if err != nil {
		log.Println("读取私钥文件失败：", err)
		return nil
	}

	// 初始化支付宝支付客户端
	client, err := alipay.NewClient(config.AppId, string(privateKeyBytes), config.IsProd)
	if err != nil {
		log.Println("初始化支付宝支付客户端失败：", err)
		return nil
	}

	// 读取证书内容
	alipayPublicCertBytes, err := ioutil.ReadFile(config.AlipayPublicCertPath)
	if err != nil {
		log.Println("读取支付宝支付公钥失败：", err)
		return nil
	}

	// 自动同步验签
	client.AutoVerifySign(alipayPublicCertBytes)
	if err = client.SetCertSnByPath(config.AppPublicCertPath, config.AlipayRootCertPath, config.AlipayPublicCertPath); err != nil {
		log.Println("设置证书失败：", err)
		return nil
	}

	return &AliPay{
		Config: &config,
		Client: client,
	}
}

// 支付宝电脑网站支付
//
// 具体传参请参考官方文档：https://opendocs.alipay.com/open/028r8t
func (p *AliPay) TradePagePay(param map[string]interface{}) (string, error) {
	var bodyMap gopay.BodyMap
	for key, value := range param {
		bodyMap.Set(key, value)
	}

	// 获取支付链接
	return p.Client.TradePagePay(context.Background(), bodyMap)
}

// 支付宝手机网站支付
//
// 具体传参请参考官方文档：https://opendocs.alipay.com/open/02ivbs
func (p *AliPay) TradeWapPay(param map[string]interface{}) (string, error) {
	var bodyMap gopay.BodyMap
	for key, value := range param {
		bodyMap.Set(key, value)
	}

	// 获取支付链接
	return p.Client.TradeWapPay(context.Background(), bodyMap)
}

// 支付宝 APP 支付
//
// 具体传参请参考官方文档：https://opendocs.alipay.com/open/02e7gq
func (p *AliPay) TradeAppPay(param map[string]interface{}) (string, error) {
	var bodyMap gopay.BodyMap
	for key, value := range param {
		bodyMap.Set(key, value)
	}

	// 获取拉起 APP 支付
	return p.Client.TradeAppPay(context.Background(), bodyMap)
}

// 支付宝订单退款
//
// 具体传参请参考官方文档：https://opendocs.alipay.com/open/02e7go
func (p *AliPay) TradeRefund(param map[string]interface{}) (*alipay.TradeRefund, error) {
	var bodyMap gopay.BodyMap
	for key, value := range param {
		bodyMap.Set(key, value)
	}

	tradeRefundResponse, err := p.Client.TradeRefund(context.Background(), bodyMap)
	if err != nil {
		return nil, errors.New("支付宝订单退款错误：" + err.Error())
	}
	if tradeRefundResponse.Response.Code != "10000" {
		return nil, errors.New("支付宝订单退款请求错误，错误码：" + tradeRefundResponse.Response.Code + "，错误消息：" + tradeRefundResponse.Response.Msg)
	}

	return tradeRefundResponse.Response, nil
}
