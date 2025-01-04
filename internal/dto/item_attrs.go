package dto

type ItemAttrItem struct {
	Name string `json:"name"` // 属性值
}

type ItemAttr struct {
	Name  string         `json:"name"` // 属性名
	Items []ItemAttrItem `json:"itemAttrItem"`
}
