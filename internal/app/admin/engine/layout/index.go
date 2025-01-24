package layout

import (
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/template/admin/layout"
	"github.com/quarkcloudio/quark-smart/v2/config"
)

type Index struct {
	layout.Template
}

// 初始化
func (p *Index) Init(ctx *quark.Context) interface{} {

	// layout 的左上角 的 title
	p.Title = config.Admin.Title

	// layout 的左上角 的 logo
	p.Logo = config.Admin.Logo

	// layout 的菜单模式,side：右侧导航，top：顶部导航，mix：混合模式
	p.Layout = config.Admin.Layout

	// layout 的菜单模式为mix时，是否自动分割菜单
	p.SplitMenus = config.Admin.SplitMenus

	// layout 的内容模式,Fluid：定宽 1200px，Fixed：自适应
	p.ContentWidth = config.Admin.ContentWidth

	// 主题色,"#1890ff"
	p.PrimaryColor = config.Admin.PrimaryColor

	// 是否固定 header 到顶部
	p.FixedHeader = config.Admin.FixedHeader

	// 是否固定导航
	p.FixSiderbar = config.Admin.FixSiderbar

	// 使用 IconFont 的图标配置
	p.IconfontUrl = config.Admin.IconfontUrl

	// 当前 layout 的语言设置，'zh-CN' | 'zh-TW' | 'en-US'
	p.Locale = config.Admin.Locale

	// 侧边菜单宽度
	p.SiderWidth = config.Admin.SiderWidth

	// 网站版权 time.Now().Format("2006") + " QuarkGo"
	p.Copyright = config.Admin.Copyright

	// 友情链接
	p.Links = config.Admin.Links

	// 右上角菜单
	p.RightMenus = config.Admin.RightMenus

	return p
}
