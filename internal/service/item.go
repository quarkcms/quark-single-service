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

func (p *ItemService) GetNumByStatus(status interface{}) int64 {
	var num int64
	query := db.Client.Model(&model.Item{})
	if status != nil {
		query.Where("status = ?", status)
	}
	query.Count(&num)
	return num
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
	items := []dto.ItemAttrDTO{}
	mapstructure.Decode(attrs, &items)
	for _, item := range items {
		p.StoreOrUpdateItemAttr(itemId, item.Name, item.Items)
	}
}

func (p *ItemService) StoreOrUpdateItemAttr(itemId int, attrName string, attrItems []dto.ItemAttrItemDTO) {
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

func (p *ItemService) DeleteItemAttrValuesByItemId(itemId interface{}, notInSuks []interface{}) {
	db.Client.
		Where("item_id = ?", itemId).
		Where("suk NOT IN ?", notInSuks).
		Delete(&model.ItemAttrValue{})
}

func (p *ItemService) StoreOrUpdateItemAttrValues(itemId int, attrValues interface{}) {
	if attrValues == nil {
		p.DeleteItemAttrValuesByItemId(itemId, []interface{}{})
		return
	}

	items := []dto.ItemAttrValueDTO{}
	mapstructure.Decode(attrValues, &items)
	suks := []interface{}{}
	for _, item := range attrValues.([]interface{}) {
		i := dto.ItemAttrValueDTO{}
		mapstructure.Decode(item, &i)
		suks = append(suks, i.Suk)
		i.AttrValue = item.(map[string]interface{})["attr_value"]
		p.StoreOrUpdateItemAttrValue(itemId, i.Suk, i)
	}

	// 清理原数据
	p.DeleteItemAttrValuesByItemId(itemId, suks)
}

func (p *ItemService) StoreOrUpdateItemAttrValue(itemId int, suk string, attrValue dto.ItemAttrValueDTO) {
	getItemAttrValue := model.ItemAttrValue{}
	db.Client.Where("item_id = ?", itemId).Where("suk = ?", suk).First(&getItemAttrValue)
	imageJsonData, _ := json.Marshal(attrValue.Image)
	attrValueJsonData, _ := json.Marshal(attrValue.AttrValue)
	isDefault := 0
	status := 0
	if attrValue.IsDefault {
		isDefault = 1
	}
	if attrValue.Status {
		status = 1
	}
	itemAttrValue := model.ItemAttrValue{
		ItemId:    itemId,
		Suk:       attrValue.Suk,
		Stock:     attrValue.Stock,
		Sales:     attrValue.Sales,
		Price:     attrValue.Price,
		Image:     string(imageJsonData),
		Cost:      attrValue.Cost,
		OtPrice:   attrValue.OtPrice,
		AttrValue: string(attrValueJsonData),
		IsDefault: isDefault,
		Status:    status,
	}
	if getItemAttrValue.Id != 0 {
		db.Client.
			Model(&model.ItemAttrValue{}).
			Where("item_id = ?", itemId).
			Where("suk = ?", suk).
			Updates(&itemAttrValue)
	} else {
		db.Client.Create(&itemAttrValue)
	}
}
