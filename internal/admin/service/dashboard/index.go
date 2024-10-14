package dashboard

import (
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/template/dashboard"
	"github.com/quarkcloudio/quark-go/v3/pkg/builder"
	"github.com/quarkcloudio/quark-smart/v2/internal/admin/service/metric"
)

type Index struct {
	dashboard.Template
}

// 初始化
func (p *Index) Init(ctx *builder.Context) interface{} {
	p.Title = "仪表盘"

	return p
}

// 内容
func (p *Index) Cards(ctx *builder.Context) []interface{} {

	return []any{
		&metric.TotalAdmin{},
		&metric.TotalLog{},
		&metric.TotalPicture{},
		&metric.TotalFile{},
		&metric.SystemInfo{},
		&metric.TeamInfo{},
	}
}
