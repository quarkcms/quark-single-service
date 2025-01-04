package service

import "fmt"

type ItemService struct{}

func NewItemService() *ItemService {
	return &ItemService{}
}

func (p *ItemService) StoreOrUpdateItemAttrs(itemId interface{}, attrs interface{}) {
	fmt.Println(attrs)
}

func (p *ItemService) StoreOrUpdateItemAttr(itemId interface{}, attrName string, attrItems interface{}) {
}

func (p *ItemService) StoreOrUpdateItemAttrValues(itemId interface{}, attrValues interface{}) {
}

func (p *ItemService) StoreOrUpdateItemAttrValue(itemId interface{}, suk string, attrvalue interface{}) {
}
