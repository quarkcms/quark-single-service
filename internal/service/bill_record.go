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
func (p *BillRecordService) GetDetailById(id int) response.BillRecordDetailResp {
	var billRecord response.BillRecordDetailResp
	db.Client.Model(model.BillRecord{}).Where("id = ?", id).Last(&billRecord)
	return billRecord
}