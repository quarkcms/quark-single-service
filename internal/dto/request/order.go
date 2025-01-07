package request

// 订单详情
type OrderDetail struct {
	ItemId      int `json:"item_id"`
	AttrValueId int `json:"attr_value_id"`
	PayNum      int `json:"pay_num"`
}

// 提交订单
type SubmitOrderReq struct {
	Realname     string        `json:"realname"`
	UserPhone    string        `json:"user_phone"`
	UserAddress  string        `json:"user_address"`
	OrderDetails []OrderDetail `json:"order_details"`
}
