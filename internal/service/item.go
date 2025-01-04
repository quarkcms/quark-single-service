package service

import (
	"encoding/json"

	"github.com/mitchellh/mapstructure"
	"github.com/quarkcloudio/quark-go/v3/dal/db"
	"github.com/quarkcloudio/quark-smart/v2/internal/dto"
	"github.com/quarkcloudio/quark-smart/v2/internal/model"
)

type ItemService struct{}

func NewItemService() *ItemService {
	return &ItemService{}
}

func (p *ItemService) DeleteItemAttrsByItemId(itemId interface{}) {
	db.Client.
		Where("item_id = ?", itemId).
		Delete(&model.ItemAttr{})
}

func (p *ItemService) StoreOrUpdateItemAttrs(itemId int, attrs interface{}) {
	p.DeleteItemAttrsByItemId(itemId)
	if attrs == nil {
		return
	}
	items := []dto.ItemAttr{}
	mapstructure.Decode(attrs, &items)
	for _, item := range items {
		p.StoreOrUpdateItemAttr(itemId, item.Name, item.Items)
	}
}

func (p *ItemService) StoreOrUpdateItemAttr(itemId int, attrName string, attrItems []dto.ItemAttrItem) {
	attrValues := ""
	for index, attrItem := range attrItems {
		if index == 0 {
			attrValues = attrItem.Name
		} else {
			attrValues = attrValues + "," + attrItem.Name
		}
	}
	attrItemsJsonData, _ := json.Marshal(attrItems)
	itemAttr := model.ItemAttr{
		ItemId:     itemId,
		AttrName:   attrName,
		AttrValues: attrValues,
		AttrItems:  string(attrItemsJsonData),
	}
	db.Client.Create(&itemAttr)
}

func (p *ItemService) StoreOrUpdateItemAttrValues(itemId interface{}, attrValues interface{}) {
}

func (p *ItemService) StoreOrUpdateItemAttrValue(itemId interface{}, suk string, attrvalue interface{}) {
}
