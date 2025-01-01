package dashboard

import (
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/template/admin/dashboard"
	"github.com/quarkcloudio/quark-smart/v2/internal/app/admin/metric"
)

type Index struct {
	dashboard.Template
}

// 初始化
func (p *Index) Init(ctx *quark.Context) interface{} {
	p.Title = "仪表盘"

	return p
}

// 内容
func (p *Index) Cards(ctx *quark.Context) []interface{} {

	return []any{
		&metric.TotalAdmin{},
		&metric.TotalLog{},
		&metric.TotalImage{},
		&metric.TotalFile{},
		&metric.SystemInfo{},
		&metric.TeamInfo{},
	}
}
