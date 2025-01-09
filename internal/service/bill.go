package service

import (
	"github.com/quarkcloudio/quark-go/v3/dal/db"
	"github.com/quarkcloudio/quark-smart/v2/internal/dto/response"
	"github.com/quarkcloudio/quark-smart/v2/internal/model"
)

type BillService struct{}

func NewBillService() *BillService {
	return &BillService{}
}

// 获取详情
func (p *BillService) GetDetailById(id int) response.BillDetailResp {
	var bill response.BillDetailResp
	db.Client.Model(model.Bill{}).Where("id = ?", id).Last(&bill)
	return bill
}

// 更新备注
func (p *BillService) UpdateMarkById(id int, mark string) error {
	return db.Client.Model(model.Bill{}).Where("id = ?", id).Update("mark", mark).Error
}
