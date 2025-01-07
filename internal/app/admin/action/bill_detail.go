package action

import (
	"strings"

	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/template/admin/resource/actions"
)

type BillDetailAction struct {
	actions.Link
}

// 账单详情
func BillDetail(name string) *BillDetailAction {

	action := &BillDetailAction{}

	action.Name = "账单详情"
	if name != "" {
		action.Name = name
	}

	return action
}

// 初始化
func (p *BillDetailAction) Init(ctx *quark.Context) interface{} {

	// 设置按钮类型,primary | ghost | dashed | link | text | default
	p.Type = "link"

	// 设置按钮大小,large | middle | small | default
	p.Size = "small"

	// 设置展示位置
	p.SetOnlyOnIndexTableRow(true)

	return p
}

// 跳转链接
func (p *BillDetailAction) GetHref(ctx *quark.Context) string {
	return "#/layout/index?api=" + strings.Replace(ctx.Path(), "/billRecord/index", "/billDetail/index&id=${id}", -1)
}
