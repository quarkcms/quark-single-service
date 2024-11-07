package metric

import (
	"github.com/quarkcloudio/quark-go/v3/dal/db"
	"github.com/quarkcloudio/quark-go/v3/model"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/statistic"
	"github.com/quarkcloudio/quark-go/v3/template/admin/dashboard/metrics"
)

type TotalPicture struct {
	metrics.Value
}

// 初始化
func (p *TotalPicture) Init() *TotalPicture {
	p.Title = "图片数量"
	p.Col = 6

	return p
}

// 计算数值
func (p *TotalPicture) Calculate() *statistic.Component {

	return p.
		Init().
		Count(db.Client.Model(&model.Picture{})).
		SetValueStyle(map[string]string{"color": "#cf1322"})
}
