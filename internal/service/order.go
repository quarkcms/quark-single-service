package service

import (
	"github.com/quarkcloudio/quark-go/v3/dal/db"
	"github.com/quarkcloudio/quark-smart/v2/internal/dto"
	"github.com/quarkcloudio/quark-smart/v2/internal/dto/request"
	"github.com/quarkcloudio/quark-smart/v2/internal/model"
	"github.com/quarkcloudio/quark-smart/v2/pkg/utils"
)

type OrderService struct{}

func NewOrderService() *OrderService {
	return &OrderService{}
}

// 根据订单id获取订单信息
func (p *OrderService) GetOrderByOrderId(orderId interface{}) (order model.Order, err error) {
	err = db.Client.Where("order_id = ?", orderId).Find(&order).Error
	return
}

// 根据订单id获取订单详细信息
func (p *OrderService) GetOrderDetailsByOrderId(orderId interface{}) (orderDetails []dto.OrderDetailDTO, err error) {
	list := []model.OrderDetail{}
	err = db.Client.Where("order_id = ?", orderId).Find(&list).Error
	for _, v := range list {

		// 获取购买商品信息
		itemInfo, err := NewItemService().GetItemWithDeleteById(v.ItemId)
		if err != nil {
			return nil, err
		}

		// 获取购买商品规格信息
		attrValueInfo, err := NewItemService().GetItemAttrValueById(v.AttrValueId)
		if err != nil {
			return nil, err
		}

		orderDetail := dto.OrderDetailDTO{
			Id:            v.Id,
			OrderId:       v.OrderId,
			ItemId:        v.ItemId,
			ItemInfo:      itemInfo,
			OrderNo:       v.OrderNo,
			Name:          v.Name,
			AttrValueId:   v.AttrValueId,
			AttrValueInfo: attrValueInfo,
			Image:         utils.GetImagePath(v.Image),
			SKU:           v.SKU,
			Price:         v.Price,
			PayNum:        v.PayNum,
		}
		orderDetails = append(orderDetails, orderDetail)
	}
	return
}

func (p *OrderService) Submit(uid interface{}, submitOrderReq request.SubmitOrderReq) (orderNo string, err error) {
	order := model.Order{}
	err = db.Client.Create(&order).Error
	return
}

func (p *OrderService) Refund() {
	db.Client.Create(&model.Order{})
}

func (p *OrderService) Verify() {
	db.Client.Create(&model.Order{})
}
