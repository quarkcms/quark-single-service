package service

import (
	"github.com/quarkcloudio/quark-go/v3/dal/db"
	"github.com/quarkcloudio/quark-smart/v2/internal/model"
)

type OrderService struct{}

func NewOrderService() *OrderService {
	return &OrderService{}
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
