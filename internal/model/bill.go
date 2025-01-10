package model

import (
	"github.com/quarkcloudio/quark-go/v3/dal/db"
	appmodel "github.com/quarkcloudio/quark-go/v3/model"
	"github.com/quarkcloudio/quark-go/v3/service"
	"github.com/quarkcloudio/quark-go/v3/utils/datetime"
)

// Bill 定义用户账单表的结构体
type Bill struct {
	Id        int               `json:"id" gorm:"primaryKey;autoIncrement;comment:用户账单id"`         // 用户账单id
	Uid       int               `json:"uid" gorm:"not null;default:0;comment:用户uid;index:openid"`  // 用户uid
	LinkId    string            `json:"link_id" gorm:"not null;default:'0';comment:关联id"`          // 关联id
	BillNo    string            `json:"bill_no" gorm:"default:null;comment:交易单号"`                  // 交易单号
	PM        uint8             `json:"pm" gorm:"not null;default:0;comment:0=支出,1=获得"`            // 0 = 支出 1 = 获得
	Title     string            `json:"title" gorm:"not null;default:'';comment:账单标题"`             // 账单标题
	Category  string            `json:"category" gorm:"not null;default:'';comment:明细种类"`          // 明细种类
	Type      string            `json:"type" gorm:"not null;default:'';comment:明细类型"`              // 明细类型
	Number    float64           `json:"number" gorm:"unsigned;not null;default:0.00;comment:明细数字"` // 明细数字
	Balance   float64           `json:"balance" gorm:"unsigned;not null;default:0.00;comment:剩余"`  // 剩余
	Mark      string            `json:"mark" gorm:"not null;default:'';comment:备注"`                // 备注
	Status    int8              `json:"status" gorm:"not null;default:1;comment:0=待确定,1=有效,-1=无效"` // 0 = 待确定 1 = 有效 -1 = 无效
	CreatedAt datetime.Datetime `json:"created_at" gorm:"type:datetime(0);"`
	UpdatedAt datetime.Datetime `json:"updated_at" gorm:"type:datetime(0);"` // 记录更新时间
}

// Seeder
func (m *Bill) Seeder() {

	// 如果菜单已存在，不执行Seeder操作
	if service.NewMenuService().IsExist(97) {
		return
	}

	// 创建菜单
	menuSeeders := []*appmodel.Menu{
		{Id: 97, Name: "财务管理", GuardName: "admin", Icon: "icon-moneycollect", Type: 1, Pid: 0, Sort: 0, Path: "/bill", Show: 1, IsEngine: 0, IsLink: 0, Status: 1},
		{Id: 98, Name: "资金流水", GuardName: "admin", Icon: "", Type: 2, Pid: 97, Sort: 0, Path: "/api/admin/bill/index", Show: 1, IsEngine: 1, IsLink: 0, Status: 1},
		{Id: 99, Name: "账单记录", GuardName: "admin", Icon: "", Type: 2, Pid: 97, Sort: 0, Path: "/api/admin/billRecord/index", Show: 1, IsEngine: 1, IsLink: 0, Status: 1},
	}
	db.Client.Create(&menuSeeders)
}
