package service

import (
	"github.com/quarkcloudio/quark-go/v3/dal/db"
	"github.com/quarkcloudio/quark-go/v3/utils/datetime"
	"github.com/quarkcloudio/quark-smart/v2/internal/dto/response"
	"github.com/quarkcloudio/quark-smart/v2/internal/model"
)

type BillService struct{}

func NewBillService() *BillService {
	return &BillService{}
}

// 获取详情
func (p *BillService) GetInfoById(id int) (bill response.BillDetailResp) {
	db.Client.Model(model.Bill{}).Where("id = ?", id).Last(&bill)
	return bill
}

// 更新备注
func (p *BillService) UpdateMarkById(id int, mark string) error {
	return db.Client.Model(model.Bill{}).Where("id = ?", id).Update("mark", mark).Error
}

// 获取周期账单列表
func (p *BillService) GetListByPeriod(startTime, endTime datetime.Datetime) []response.BillDetailResp {
	bills := make([]response.BillDetailResp, 0)
	db.Client.Model(model.Bill{}).
		Where("status = ? AND created_at BETWEEN ? AND ?", 1, startTime, endTime).
		Order("id desc").
		Find(&bills)
	return bills
}
