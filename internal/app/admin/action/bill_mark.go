package action

import (
	"strconv"

	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/form/rule"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/message"
	"github.com/quarkcloudio/quark-go/v3/template/admin/resource"
	"github.com/quarkcloudio/quark-go/v3/template/admin/resource/actions"
	"github.com/quarkcloudio/quark-smart/v2/internal/service"
	"gorm.io/gorm"
)

type BillMarkAction struct {
	actions.ModalForm
}

// 资金流水备注
func BillMark(name string) *BillMarkAction {

	action := &BillMarkAction{}

	action.Name = "备注"
	if name != "" {
		action.Name = name
	}

	return action
}

// 初始化
func (p *BillMarkAction) Init(ctx *quark.Context) interface{} {

	// 类型
	p.Type = "link"

	// 设置按钮大小,large | middle | small | default
	p.Size = "small"

	// 关闭时销毁 Modal 里的元素
	p.DestroyOnClose = true

	//  执行成功后刷新的组件
	p.Reload = "table"

	// 在表格行内展示
	p.SetOnlyOnIndexTableRow(true)

	// 行为接口接收的参数
	p.SetApiParams([]string{
		"id",
	})

	return p
}

// 字段
func (p *BillMarkAction) Fields(ctx *quark.Context) []interface{} {

	field := &resource.Field{}

	return []interface{}{
		field.Hidden("id", "ID"),

		field.TextArea("mark", "").
			SetRules([]rule.Rule{
				rule.Required("备注必须填写"),
				rule.Max(100, "备注不能超过100个字符"),
			}).
			SetPlaceholder("请输入备注"),
	}
}

// 表单数据（异步获取）
func (p *BillMarkAction) Data(ctx *quark.Context) map[string]interface{} {

	id, _ := strconv.Atoi(ctx.Query("id").(string))

	bill := service.NewBillService().GetInfoById(id)

	return map[string]interface{}{
		"id":   id,
		"mark": bill.Mark,
	}
}

// 执行行为句柄
func (p *BillMarkAction) Handle(ctx *quark.Context, query *gorm.DB) error {

	var param struct {
		Id   int    `json:"id"`
		Mark string `json:"mark"`
	}

	if err := ctx.Bind(&param); err != nil {
		return ctx.JSON(200, message.Error(err.Error()))
	}

	if err := service.NewBillService().UpdateMarkById(param.Id, param.Mark); err != nil {
		return ctx.JSON(200, message.Error("操作失败"))
	}

	return ctx.JSON(200, message.Success("操作成功"))
}
