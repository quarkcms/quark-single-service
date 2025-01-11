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

// 获取订单数量
// all:全部
// pendingPayment:待支付
// pendingShipment:待发货（预留）
// pendingVerify:待核销
// pendingReceipt:待收货（预留）
// pendingReview:待评价（预留）
// completed:已完成
// refund:退款申请中
// refunded:已退款
// deleted:已删除
func (p *OrderService) GetNum(uid interface{}, status string) int64 {
	var num int64
	query := db.Client.Model(&model.Order{})
	if uid != nil {
		query.Where("uid = ?", uid)
	}
	switch status {
	case "all":
		// 全部
	case "pendingPayment":
		// 待支付
		query.Where("paid", 0)
	case "pendingShipment":
		// 待发货（预留）
		query.Where("paid", 1).Where("status = ?", 0).Where("shipping_type = ?", 1)
	case "pendingVerify":
		// 待核销，到店自提订单需要核销
		query.Where("paid", 1).Where("status = ?", 0).Where("shipping_type = ?", 2)
	case "pendingReceipt":
		// 待收货（预留）
		query.Where("paid", 1).Where("status = ?", 1).Where("shipping_type = ?", 1)
	case "pendingReview":
		// 待评价（预留）
		query.Where("paid", 1).Where("status = ?", 2)
	case "completed":
		// 已完成
		query.Where("paid", 1).Where("status = ?", 3)
	case "refund":
		// 退款申请中
		query.Where("paid", 1).Where("status = ?", -1)
	case "refunded":
		// 已退款
		query.Where("paid", 1).Where("status = ?", -2)
	case "deleted":
		// 已删除
		query.Unscoped().Where("deleted_at IS NOT NULL")
	}
	query.Count(&num)
	return num
}

// 根据订单状态获取订单数量
// all:全部
// pendingPayment:待支付
// pendingShipment:待发货（预留）
// pendingVerify:待核销
// pendingReceipt:待收货（预留）
// pendingReview:待评价（预留）
// completed:已完成
// refund:退款申请中
// refunded:已退款
// deleted:已删除
func (p *OrderService) GetNumByStatus(status string) int64 {
	return p.GetNum(nil, status)
}

// 根据用户ID、订单状态获取订单数量
// all:全部
// pendingPayment:待支付
// pendingShipment:待发货（预留）
// pendingVerify:待核销
// pendingReceipt:待收货（预留）
// pendingReview:待评价（预留）
// completed:已完成
// refund:退款申请中
// refunded:已退款
// deleted:已删除
func (p *OrderService) GetNumByUidAndStatus(uid interface{}, status string) (num int64, err error) {
	if uid == nil {
		return 0, errors.New("参数错误")
	}
	return p.GetNum(uid, status), nil
}

// 获取订单信息
func (p *OrderService) GetOrder(uid, orderId interface{}) (orderDto dto.OrderDTO, err error) {
	order := model.Order{}
	query := db.Client
	if uid != nil {
		query.Where("uid = ?", uid)
	}
	err = query.Where("id = ?", orderId).First(&order).Error
	if err != nil {
		return
	}
	orderDetails, err := p.GetOrderDetailsByOrderId(orderId)
	if err != nil {
		return
	}
	orderDto = dto.OrderDTO{
		Id:                    order.Id,
		OrderNo:               order.OrderNo,
		Uid:                   order.Uid,
		Realname:              order.Realname,
		UserPhone:             order.UserPhone,
		UserAddress:           order.UserAddress,
		TotalNum:              order.TotalNum,
		TotalPrice:            order.TotalPrice,
		PayPrice:              order.PayPrice,
		Paid:                  order.Paid,
		PayTime:               order.PayTime,
		PayType:               order.PayType,
		OrderDetails:          orderDetails,
		Status:                order.Status,
		RefundStatus:          order.RefundStatus,
		RefundReasonImg:       order.RefundReasonImg,
		RefundReasonExplain:   order.RefundReasonExplain,
		RefundReason:          order.RefundReason,
		RefundRejectionReason: order.RefundRejectionReason,
		RefundReasonTime:      order.RefundReasonTime,
		RefundPrice:           order.RefundPrice,
		Remark:                order.Remark,
		MerchantId:            order.MerchantId,
		IsMerchantCheck:       order.IsMerchantCheck,
		Cost:                  order.Cost,
		VerifyCode:            order.VerifyCode,
		ShippingType:          order.ShippingType,
		ClerkId:               order.Id,
		CreatedAt:             order.CreatedAt,
		UpdatedAt:             order.UpdatedAt,
	}
	return
}

// 根据ID获取订单信息
func (p *OrderService) GetOrderById(orderId interface{}) (orderDto dto.OrderDTO, err error) {
	if orderId == nil {
		return orderDto, errors.New("参数错误")
	}
	return p.GetOrder(nil, orderId)
}

// 获取用户订单信息
func (p *OrderService) GetUserOrder(uid, orderId interface{}) (orderDto dto.OrderDTO, err error) {
	if uid == nil || orderId == nil {
		return orderDto, errors.New("参数错误")
	}
	return p.GetOrder(uid, orderId)
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
		PayPrice:   payPrice,
		Cost:       cost,
	}).Error
	if err != nil {
		tx.Rollback()
		return "", err
	}

	err = tx.Commit().Error
	if err != nil {
		return
	}

	// 重建items表attr_values字段值
	for _, orderDetail := range orderDetails {
		NewItemService().RebuildItemAttrValues(orderDetail.ItemId)
	}

	return
}

// 删除订单
func (p *OrderService) Delete(uid interface{}, id interface{}) (err error) {
	order, err := p.GetOrderById(id)
	if err != nil {
		return err
	}

	// 后台可删除未付款订单
	if uid == nil && order.Paid == 1 {
		return errors.New("已付款订单无法删除")
	}

	// 用户可删除未付款、已完成订单
	if order.Paid == 1 && order.Status != 3 {
		return errors.New("已付款未完成订单无法删除")
	}

	tx := db.Client.Begin()
	if uid != nil {
		err = tx.Where("uid = ?", uid).Where("id = ?", id).Delete(&model.Order{}).Error
	} else {
		err = tx.Model(&model.Order{}).Where("id = ?", id).Update("is_system_del", 1).Error
		if err != nil {
			tx.Rollback()
			return err
		}
		err = tx.Where("id = ?", id).Delete(&model.Order{}).Error
	}

	if err != nil {
		tx.Rollback()
		return err
	}

	// 删除未付款订单，将归还库存
	if order.Paid == 0 {
		for _, orderDetail := range order.OrderDetails {
			item, err := NewItemService().GetItem(orderDetail.ItemId, nil, false)
			if err != nil {
				tx.Rollback()
				return err
			}
			// 单规格归还库存
			if item.SpecType == 0 {
				tx.Model(&model.Item{}).Where("id = ?", item.Id).Update("stock", item.Stock+orderDetail.PayNum)
			}
			// 多规格归还库存
			if item.SpecType == 1 {
				var attrValue model.ItemAttrValue
				tx.Where("id = ?", orderDetail.AttrValueId).Where("item_id = ?", item.Id).First(&attrValue)
				if attrValue.Id != 0 {
					tx.Model(&model.ItemAttrValue{}).Where("id = ?", orderDetail.AttrValueId).Update("stock", attrValue.Stock-orderDetail.PayNum)
				}
			}
		}
	}

	err = tx.Commit().Error
	if err != nil {
		return
	}

	// 重建items表attr_values字段值
	for _, orderDetail := range order.OrderDetails {
		NewItemService().RebuildItemAttrValues(orderDetail.ItemId)
	}

	return nil
}

// 后台管理员根据订单ID删除订单
func (p *OrderService) DeleteBySystem(id interface{}) (err error) {
	if id == nil {
		return errors.New("参数错误")
	}
	return p.Delete(nil, id)
}

// 前台用户删除订单
func (p *OrderService) DeleteByUser(uid interface{}, id interface{}) (err error) {
	if uid == nil || id == nil {
		return errors.New("参数错误")
	}
	return p.Delete(uid, id)
}

// 订单退款
func (p *OrderService) Refund(orderId interface{}, refundPrice float64) (err error) {
	return
}

// 订单核销
func (p *OrderService) Verify(orderId interface{}, verifyCode interface{}) (err error) {
	return
}

// 管理后台根据订单id核销
func (p *OrderService) VerifyById(orderId interface{}) (err error) {
	return p.Verify(orderId, nil)
}
