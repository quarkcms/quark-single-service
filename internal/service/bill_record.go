package service

import (
	"github.com/quarkcloudio/quark-go/v3/dal/db"
	"github.com/quarkcloudio/quark-smart/v2/internal/dto/response"
	"github.com/quarkcloudio/quark-smart/v2/internal/model"
)

type BillRecordService struct{}

func NewBillRecordService() *BillRecordService {
	return &BillRecordService{}
}

// 获取详情
func (p *BillRecordService) GetInfoById(id int) (billRecord response.BillRecordDetailResp) {
	db.Client.Model(model.BillRecord{}).Where("id = ?", id).Last(&billRecord)
	return billRecord
}

// 创建账单
func (p *BillRecordService) CreateBillRecord(billRecord model.BillRecord) error {
	return db.Client.Model(model.BillRecord{}).Create(&billRecord).Error
}
