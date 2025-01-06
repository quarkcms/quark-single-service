package service

import (
	"github.com/quarkcloudio/quark-go/v3/dal/db"
	"github.com/quarkcloudio/quark-smart/v2/internal/model"
)

type OrderService struct{}

func NewOrderService() *OrderService {
	return &OrderService{}
}

// 根据订单id获取订单信息
func (p *OrderService) GetOrderDetailsByOrderId(orderId interface{}) (list []model.OrderDetail, err error) {
	err = db.Client.Where("order_id = ?", orderId).Find(&list).Error
	return
}

func (p *OrderService) Submit() {
	db.Client.Create(&model.Order{})
}

func (p *OrderService) Refund() {
	db.Client.Create(&model.Order{})
}

func (p *OrderService) Verify() {
	db.Client.Create(&model.Order{})
}
