package metric

import (
	"github.com/quarkcloudio/quark-go/v3/dal/db"
	"github.com/quarkcloudio/quark-go/v3/model"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/statistic"
	"github.com/quarkcloudio/quark-go/v3/template/admin/dashboard/metrics"
)

type TotalAdmin struct {
	metrics.Value
}

// 初始化
func (p *TotalAdmin) Init() *TotalAdmin {
	p.Title = "用户数量"
	p.Col = 6

	return p
}

// 计算数值
func (p *TotalAdmin) Calculate() *statistic.Component {

	return p.
		Init().
		Count(db.Client.Model(&model.User{})).
		SetValueStyle(map[string]string{"color": "#3f8600"})
}
