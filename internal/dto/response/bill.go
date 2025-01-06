package response

import "github.com/quarkcloudio/quark-go/v3/utils/datetime"

type BillDetailResp struct {
	Id        int               `json:"id"`
	Uid       int               `json:"uid"`
	LinkId    string            `json:"link_id"`
	BillNo    string            `json:"bill_no"`
	PM        uint8             `json:"pm"`
	Title     string            `json:"title"`
	Category  string            `json:"category"`
	Type      string            `json:"type"`
	Number    float64           `json:"number"`
	Balance   float64           `json:"balance"`
	Mark      string            `json:"mark"`
	Status    int8              `json:"status"`
	CreatedAt datetime.Datetime `json:"created_at"`
}
