package response

import "github.com/quarkcloudio/quark-go/v3/utils/datetime"

type BillRecordDetailResp struct {
	Id            int               `json:"id"`
	Title         string            `json:"title"`
	Day           string            `json:"day"`
	Type          int8              `json:"type"`
	EntryPrice    float64           `json:"entry_price"`
	ExpPrice      float64           `json:"exp_price"`
	IncomePrice   float64           `json:"income_price"`
	StartDatetime datetime.Datetime `json:"start_datetime"`
	EndDatetime   datetime.Datetime `json:"end_datetime"`
	Status        int8              `json:"status"`
	CreatedAt     datetime.Datetime `json:"created_at"`
}
