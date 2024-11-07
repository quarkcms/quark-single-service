package login

import (
	"github.com/dchest/captcha"
	"github.com/golang-jwt/jwt/v4"
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/service"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/form/rule"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/icon"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/message"
	"github.com/quarkcloudio/quark-go/v3/template/admin/login"
	"github.com/quarkcloudio/quark-go/v3/template/admin/resource"
	"github.com/quarkcloudio/quark-go/v3/utils/datetime"
	"github.com/quarkcloudio/quark-go/v3/utils/hash"
	"github.com/quarkcloudio/quark-smart/v2/config"
	"github.com/quarkcloudio/quark-smart/v2/internal/dto/request"
	"gorm.io/gorm"
)

type Index struct {
	login.Template
}

// 初始化
func (p *Index) Init(ctx *quark.Context) interface{} {

	// 登录页面Logo
	p.Logo = false

	// 登录页面标题
	p.Title = config.Admin.Title

	// 登录页面子标题
	p.SubTitle = config.Admin.SubTitle

	// 登录后跳转地址
	p.Redirect = "/layout/index?api=/api/admin/dashboard/index/index"

	return p
}

// 字段
func (p *Index) Fields(ctx *quark.Context) []interface{} {
	field := &resource.Field{}

	// 获取验证码ID链接
	captchaIdUrl := ctx.RouterPathToUrl("/api/admin/login/index/captchaId")

	// 验证码链接
	captchaUrl := ctx.RouterPathToUrl("/api/admin/login/index/captcha/:id")

	return []interface{}{
		field.Text("username").
			SetRules([]rule.Rule{
				rule.Required("请输入用户名"),
			}).
			SetPlaceholder("用户名").
			SetWidth("100%").
			SetSize("large").
			SetPrefix(icon.New().SetType("icon-user")),

		field.Password("password").
			SetRules([]rule.Rule{
				rule.Required("请输入密码"),
			}).
			SetPlaceholder("密码").
			SetWidth("100%").
			SetSize("large").
			SetPrefix(icon.New().SetType("icon-lock")),

		field.ImageCaptcha("captcha").
			SetCaptchaIdUrl(captchaIdUrl).
			SetCaptchaUrl(captchaUrl).
			SetRules([]rule.Rule{
				rule.Required("请输入验证码"),
			}).
			SetPlaceholder("验证码").
			SetWidth("100%").
			SetSize("large").
			SetPrefix(icon.New().SetType("icon-safetycertificate")),
	}
}

// 登录方法
func (p *Index) Handle(ctx *quark.Context) error {
	loginRequest := &request.LoginReq{}
	if err := ctx.Bind(loginRequest); err != nil {
		return ctx.JSON(200, message.Error(err.Error()))
	}
	if loginRequest.Captcha.Id == "" || loginRequest.Captcha.Value == "" {
		return ctx.JSON(200, message.Error("验证码不能为空"))
	}

	verifyResult := captcha.VerifyString(loginRequest.Captcha.Id, loginRequest.Captcha.Value)
	if !verifyResult {
		return ctx.JSON(200, message.Error("验证码错误"))
	}
	captcha.Reload(loginRequest.Captcha.Id)

	if loginRequest.Username == "" || loginRequest.Password == "" {
		return ctx.JSON(200, message.Error("用户名或密码不能为空"))
	}

	adminInfo, err := service.NewUserService().GetInfoByUsername(loginRequest.Username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.JSON(200, message.Error("用户不存在"))
		}
		return ctx.JSON(200, message.Error(err.Error()))
	}

	// 检验账号和密码
	if !hash.Check(adminInfo.Password, loginRequest.Password) {
		return ctx.JSON(200, message.Error("用户名或密码错误"))
	}

	config := ctx.Engine.GetConfig()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, service.NewUserService().GetAdminClaims(adminInfo))

	// 更新登录信息
	service.NewUserService().UpdateLastLogin(adminInfo.Id, ctx.ClientIP(), datetime.Now())

	// 获取token字符串
	tokenString, err := token.SignedString([]byte(config.AppKey))
	if err != nil {
		return ctx.JSON(200, message.Error(err.Error()))
	}

	return ctx.JSON(200, message.Success("登录成功", "", map[string]string{
		"token": tokenString,
	}))
}
