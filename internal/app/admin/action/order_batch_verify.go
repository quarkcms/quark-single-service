package action

import (
	"strconv"
	"strings"

	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/message"
	"github.com/quarkcloudio/quark-go/v3/template/admin/resource/actions"
	"github.com/quarkcloudio/quark-smart/v2/internal/service"
	"gorm.io/gorm"
)

type OrderBatchVerifyAction struct {
	actions.Action
}

// 订单批量核销
func OrderBatchVerify() *OrderBatchVerifyAction {
	return &OrderBatchVerifyAction{}
}

// 初始化
func (p *OrderBatchVerifyAction) Init(ctx *quark.Context) interface{} {

	// 设置按钮文字
	p.Name = "批量核销"

	// 设置按钮类型,primary | ghost | dashed | link | text | default
	p.Type = "link"

	// 设置按钮大小,large | middle | small | default
	p.Size = "small"

	//  执行成功后刷新的组件
	p.Reload = "table"

	// 当行为在表格行展示时，支持js表达式
	p.WithConfirm("确定要核销吗？", "核销后数据将无法恢复，请谨慎操作！", "modal")

	// 在表格多选弹出层展示
	p.SetOnlyOnIndexTableAlert(true)

	return p
}

// 行为接口接收的参数，当行为在表格行展示的时候，可以配置当前行的任意字段
func (p *OrderBatchVerifyAction) GetApiParams() []string {
	return []string{
		"id",
	}
}

// 执行行为句柄
func (p *OrderBatchVerifyAction) Handle(ctx *quark.Context, query *gorm.DB) error {
	id := ctx.Query("id")
	if id == "" {
		return ctx.JSON(200, message.Error("参数错误！"))
	}

	ids := strings.Split(id.(string), ",")
	if len(ids) > 0 {
		for _, v := range ids {
			idInt, err := strconv.Atoi(v)
			if err != nil {
				return ctx.JSON(200, message.Error(err.Error()))
			}
			service.NewOrderService().VerifyById(idInt)
		}
	} else {
		idInt, err := strconv.Atoi(id.(string))
		if err != nil {
			return ctx.JSON(200, message.Error(err.Error()))
		}
		service.NewOrderService().VerifyById(idInt)
	}

	return ctx.JSON(200, message.Success("操作成功"))
}
