package service

import (
	"errors"
	"strconv"

	"github.com/quarkcloudio/quark-go/v3/dal/db"
	"github.com/quarkcloudio/quark-go/v3/utils/rand"
	"github.com/quarkcloudio/quark-smart/v2/internal/dto"
	"github.com/quarkcloudio/quark-smart/v2/internal/dto/request"
	"github.com/quarkcloudio/quark-smart/v2/internal/model"
	"github.com/quarkcloudio/quark-smart/v2/pkg/utils"
	"github.com/sony/sonyflake"
	"gorm.io/gorm"
)

type OrderService struct{}

func NewOrderService() *OrderService {
	return &OrderService{}
}

// 根据订单id获取订单信息
func (p *OrderService) GetOrderById(orderId interface{}) (order model.Order, err error) {
	err = db.Client.Where("id = ?", orderId).Find(&order).Error
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
		attrValueInfo, err := NewItemService().GetItemAttrValueWithDeleteById(v.ItemId, v.AttrValueId)
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

func (p *OrderService) Submit(uid int, submitOrderReq request.SubmitOrderReq) (orderNo string, err error) {
	realname := submitOrderReq.Realname
	userPhone := submitOrderReq.UserPhone
	userAddress := submitOrderReq.UserAddress
	orderDetails := submitOrderReq.OrderDetails
	if len(orderDetails) < 1 {
		return "", errors.New("参数错误")
	}

	// 开始事务
	tx := db.Client.Begin()

	// 雪花算法生成订单号
	sf := sonyflake.NewSonyflake(sonyflake.Settings{})
	getOrderNo, err := sf.NextID()
	if err != nil {
		return
	}

	orderNo = strconv.FormatUint(getOrderNo, 10)
	order := model.Order{
		OrderNo:     orderNo,
		Uid:         uid,
		Realname:    realname,
		UserPhone:   userPhone,
		UserAddress: userAddress,
		VerifyCode:  rand.MakeNumeric(8),
	}

	// 创建订单
	err = tx.Create(&order).Error
	if err != nil {
		tx.Rollback()
		return "", err
	}

	totalNum := 0
	totalPrice := 0.00
	payPrice := 0.00
	cost := 0.00
	for _, orderDetail := range orderDetails {
		var item model.Item
		err = tx.Where("id = ?", orderDetail.ItemId).First(&item).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return "", err
		}
		if item.Id == 0 {
			tx.Rollback()
			return "", errors.New("商品不存在")
		}
		if item.Status != 1 {
			tx.Rollback()
			return "", errors.New("商品已下架")
		}
		if orderDetail.PayNum <= 0 {
			tx.Rollback()
			return "", errors.New("请选择购买商品")
		}

		image := ""
		sku := ""
		price := 0.00
		// 计算总量
		totalNum = totalNum + orderDetail.PayNum
		// 单规格
		if item.SpecType == 0 {
			if item.Stock <= 0 {
				tx.Rollback()
				return "", errors.New("商品已售完")
			}
			if orderDetail.PayNum > item.Stock {
				tx.Rollback()
				return "", errors.New("库存不足")
			}
			result := tx.Model(&item).Where("version = ?", item.Version).Update("stock", item.Stock-orderDetail.PayNum)
			if result.RowsAffected == 0 {
				// 如果没有行被更新，说明库存已经被其他事务修改，回滚事务
				tx.Rollback()
				return "", errors.New("购买失败，请重试")
			}
			// 更新版本号
			tx.Model(&item).Update("version", item.Version+1)
			price = item.Price
			image = item.Image
			totalPrice = totalPrice + float64(orderDetail.PayNum)*item.Price
			payPrice = payPrice + float64(orderDetail.PayNum)*item.Price
			cost = cost + float64(orderDetail.PayNum)*item.Cost
		}
		// 多规格
		if item.SpecType == 1 {
			var attrValue model.ItemAttrValue
			err = tx.Where("id = ?", orderDetail.AttrValueId).Where("item_id = ?", item.Id).First(&attrValue).Error
			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				tx.Rollback()
				return "", err
			}
			if attrValue.Id == 0 {
				tx.Rollback()
				return "", errors.New("商品不存在")
			}
			if attrValue.Status != 1 {
				tx.Rollback()
				return "", errors.New("商品已下架")
			}
			if attrValue.Stock <= 0 {
				tx.Rollback()
				return "", errors.New("商品已售完")
			}
			if orderDetail.PayNum > attrValue.Stock {
				tx.Rollback()
				return "", errors.New("库存不足")
			}
			result := tx.Model(&attrValue).Where("version = ?", attrValue.Version).Update("stock", attrValue.Stock-orderDetail.PayNum)
			if result.RowsAffected == 0 {
				// 如果没有行被更新，说明库存已经被其他事务修改，回滚事务
				tx.Rollback()
				return "", errors.New("购买失败，请重试")
			}
			// 更新版本号
			tx.Model(&attrValue).Update("version", attrValue.Version+1)
			sku = attrValue.Suk
			price = attrValue.Price
			if attrValue.Image != "" {
				image = attrValue.Image
			}
			totalPrice = totalPrice + float64(orderDetail.PayNum)*attrValue.Price
			payPrice = payPrice + float64(orderDetail.PayNum)*attrValue.Price
			cost = cost + float64(orderDetail.PayNum)*attrValue.Cost
		}

		orderDetail := model.OrderDetail{
			OrderId:     order.Id,
			ItemId:      item.Id,
			OrderNo:     orderNo,
			Name:        item.Name,
			Content:     item.Content,
			AttrValueId: orderDetail.AttrValueId,
			Image:       image,
			SKU:         sku,
			Price:       price,
			PayNum:      orderDetail.PayNum,
		}

		// 创建订单详情
		err = tx.Create(&orderDetail).Error
		if err != nil {
			tx.Rollback()
			return "", err
		}
	}

	orderStatus := model.OrderStatus{
		OrderId:       order.Id,
		ChangeType:    "create_order",
		ChangeMessage: "订单生成",
	}

	// 创建订单详情
	err = tx.Create(&orderStatus).Error
	if err != nil {
		tx.Rollback()
		return "", err
	}

	// 更新订单金额等信息
	err = tx.Model(&order).Where("id = ?", order.Id).Updates(&model.Order{
		TotalNum:   totalNum,
		TotalPrice: totalPrice,
		Cost:       cost,
	}).Error
	if err != nil {
		tx.Rollback()
		return "", err
	}

	tx.Commit()
	return
}

func (p *OrderService) Refund() {
	db.Client.Create(&model.Order{})
}

func (p *OrderService) Verify() {
	db.Client.Create(&model.Order{})
}
