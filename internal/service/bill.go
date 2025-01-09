package service

import (
	"strconv"

	"github.com/quarkcloudio/quark-go/v3/dal/db"
	"github.com/quarkcloudio/quark-go/v3/utils/datetime"
	"github.com/quarkcloudio/quark-go/v3/utils/rand"
	"github.com/quarkcloudio/quark-smart/v2/internal/dto/response"
	"github.com/quarkcloudio/quark-smart/v2/internal/model"
)

type BillService struct{}

func NewBillService() *BillService {
	return &BillService{}
}

// 生成账单号
func (p *BillService) GenerateBillNo() string {
	return strconv.FormatInt(datetime.Now().Unix(), 10) + rand.MakeNumeric(6) + datetime.DateNow().Format("20060102")
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
