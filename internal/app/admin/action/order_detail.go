package action

import (
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/component"
	"github.com/quarkcloudio/quark-go/v3/template/admin/resource/actions"
)

type OrderDetailAction struct {
	actions.Drawer
}

// 创建-抽屉类型
func OrderDetail() *OrderDetailAction {
	return &OrderDetailAction{}
}

// 初始化
func (p *OrderDetailAction) Init(ctx *quark.Context) interface{} {

	// 文字
	p.Name = "订单详情"

	// 类型
	p.Type = "link"

	// 设置按钮大小,large | middle | small | default
	p.Size = "small"

	// 关闭时销毁 Drawer 里的子元素
	p.DestroyOnClose = true

	// 执行成功后刷新的组件
	p.Reload = "table"

	// 宽度
	p.Width = 1000

	// 设置展示位置
	p.SetOnlyOnIndexTableRow(true)

	return p
}

// 内容
func (p *OrderDetailAction) GetBody(ctx *quark.Context) interface{} {
	// 返回数据
	return component.New().SetComponent("order")
}
