package service

import (
	"encoding/json"
	"errors"

	"github.com/mitchellh/mapstructure"
	"github.com/quarkcloudio/quark-go/v3/dal/db"
	"github.com/quarkcloudio/quark-smart/v2/internal/dto"
	"github.com/quarkcloudio/quark-smart/v2/internal/dto/response"
	"github.com/quarkcloudio/quark-smart/v2/internal/model"
	"github.com/quarkcloudio/quark-smart/v2/pkg/utils"
	"gorm.io/gorm"
)

type ItemService struct{}

func NewItemService() *ItemService {
	return &ItemService{}
}

// 根据状态获取商品数量
func (p *ItemService) GetNumByStatus(status interface{}) int64 {
	var num int64
	query := db.Client.Model(&model.Item{})
	if status != nil {
		query.Where("status = ?", status)
	}
	query.Count(&num)
	return num
}

// 根据商品Id上传商品属性
func (p *ItemService) DeleteItemAttrsByItemId(itemId interface{}) {
	db.Client.
		Where("item_id = ?", itemId).
		Delete(&model.ItemAttr{})
}

// 跟进商品id创建或更新商品属性
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

// 根据商品id删除商品规格
func (p *ItemService) DeleteItemAttrValuesByItemId(itemId interface{}, notInSuks []interface{}) {
	db.Client.
		Where("item_id = ?", itemId).
		Where("suk NOT IN ?", notInSuks).
		Delete(&model.ItemAttrValue{})
}

// 创建或更新商品规格
func (p *ItemService) StoreOrUpdateItemAttrValues(itemId int, attrValues interface{}) {
	if attrValues == nil {
		p.DeleteItemAttrValuesByItemId(itemId, []interface{}{})
		return
	}

	items := []dto.AttrValueDTO{}
	mapstructure.Decode(attrValues, &items)
	suks := []interface{}{}
	for _, item := range attrValues.([]interface{}) {
		i := dto.AttrValueDTO{}
		mapstructure.Decode(item, &i)
		suks = append(suks, i.Suk)
		i.AttrValue = item.(map[string]interface{})["attr_value"]
		p.StoreOrUpdateItemAttrValue(itemId, i.Suk, i)
	}

	// 清理原数据
	p.DeleteItemAttrValuesByItemId(itemId, suks)
}

func (p *ItemService) StoreOrUpdateItemAttrValue(itemId int, suk string, attrValue dto.AttrValueDTO) {
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

// 获取商品属性
func (p *ItemService) GetItemAttrs(itemId int) (attrs []dto.AttrDTO, err error) {
	var list []model.ItemAttr
	err = db.Client.Where("item_id = ?", itemId).Find(&list).Error
	for _, v := range list {
		attrItems := []map[string]interface{}{}
		err = json.Unmarshal([]byte(v.AttrItems), &attrItems)
		attrs = append(attrs, dto.AttrDTO{
			Id:         v.Id,
			ItemId:     v.ItemId,
			AttrName:   v.AttrName,
			AttrValues: v.AttrValues,
			AttrItems:  attrItems,
		})
	}
	return
}

// 获取商品属性值
func (p *ItemService) GetItemAttrValues(itemId int) (attrValues []dto.AttrValueDTO, err error) {
	var list []model.ItemAttrValue
	err = db.Client.Where("item_id = ?", itemId).Find(&list).Error
	for _, v := range list {
		itemAttrValue := map[string]interface{}{}
		err = json.Unmarshal([]byte(v.AttrValue), &itemAttrValue)
		attrValues = append(attrValues, dto.AttrValueDTO{
			Id:        v.Id,
			ItemId:    v.ItemId,
			Suk:       v.Suk,
			Stock:     v.Stock,
			Sales:     v.Sales,
			Price:     v.Price,
			Image:     utils.GetImagePath(v.Image),
			Cost:      v.Cost,
			OtPrice:   v.OtPrice,
			AttrValue: itemAttrValue,
			IsDefault: v.IsDefault == 1,
			Status:    v.Status == 1,
		})
	}
	return
}

// 获取商品
func (p *ItemService) GetItem(itemId int, status interface{}, withDelete bool) (data dto.ItemDTO, err error) {
	item := model.Item{}
	query := db.Client
	if withDelete {
		query.Unscoped()
	}
	if status != nil {
		query.Where("status = ?", status)
	}
	err = query.Where("id = ?", itemId).First(&item).Error
	if err != nil {
		return
	}
	// 获取商品属性
	attrs, err := p.GetItemAttrs(itemId)
	if err != nil {
		return
	}
	// 获取商品属性值
	attrValues, err := p.GetItemAttrValues(itemId)
	if err != nil {
		return
	}
	data = dto.ItemDTO{
		Id:          item.Id,
		MerchantId:  item.MerchantId,
		Image:       utils.GetImagePath(item.Image),
		SliderImage: item.SliderImage,
		Name:        item.Name,
		Keyword:     item.Keyword,
		Description: item.Description,
		CategoryIds: item.CategoryIds,
		Price:       item.Price,
		OtPrice:     item.OtPrice,
		Sort:        item.Sort,
		Sales:       item.Sales,
		Stock:       item.Stock,
		Status:      item.Status,
		Cost:        item.Cost,
		FictiSales:  item.FictiSales,
		Views:       item.Views,
		FictiViews:  item.FictiViews,
		SpecType:    item.SpecType,
		Attrs:       attrs,
		AttrValues:  attrValues,
	}
	return
}

// 根据id获取商品
func (p *ItemService) GetItemById(itemId int, status int) (data dto.ItemDTO, err error) {
	return p.GetItem(itemId, 1, false)
}

// 根据id获取商品（包含已删除的）
func (p *ItemService) GetItemWithDeleteById(itemId int) (data dto.ItemDTO, err error) {
	return p.GetItem(itemId, nil, true)
}

// 获取商品规格值
func (p *ItemService) GetItemAttrValue(itemId int, attrValueId int, status interface{}, withDelete bool) (data dto.AttrValueDTO, err error) {
	itemAttrValue := model.ItemAttrValue{}
	query := db.Client
	if withDelete {
		query.Unscoped()
	}
	if status != nil {
		query.Where("status = ?", status)
	}
	err = query.Where("item_id = ?", itemId).Where("id = ?", attrValueId).First(&itemAttrValue).Error
	data = dto.AttrValueDTO{
		Id:        itemAttrValue.Id,
		ItemId:    itemAttrValue.ItemId,
		Suk:       itemAttrValue.Suk,
		Stock:     itemAttrValue.Stock,
		Sales:     itemAttrValue.Sales,
		Price:     itemAttrValue.Price,
		Image:     utils.GetImagePath(itemAttrValue.Image),
		Cost:      itemAttrValue.Cost,
		OtPrice:   itemAttrValue.OtPrice,
		AttrValue: itemAttrValue.AttrValue,
		IsDefault: itemAttrValue.IsDefault == 1,
		Status:    itemAttrValue.Status == 1,
	}
	return
}

// 根据id获取商品规格值
func (p *ItemService) GetItemAttrValueById(itemId int, attrValueId int) (data dto.AttrValueDTO, err error) {
	return p.GetItemAttrValue(itemId, attrValueId, 1, false)
}

// 根据id获取商品规格值（包含已删除的）
func (p *ItemService) GetItemAttrValueWithDeleteById(itemId int, attrValueId int) (data dto.AttrValueDTO, err error) {
	return p.GetItemAttrValue(itemId, attrValueId, nil, true)
}

// 检查商品状态
func (p *ItemService) CheckItemStatus(tx *gorm.DB, itemId int, attrValueId int, payNum int) (status bool, err error) {
	item, err := p.GetItem(itemId, 1, false)
	if err != nil {
		return false, err
	}
	if item.Id == 0 {
		return false, errors.New("商品已下架")
	}
	if payNum <= 0 {
		return false, errors.New("请选择购买商品")
	}
	// 单规格
	if item.SpecType == 0 {
		if item.Stock <= 0 {
			return false, errors.New("商品已售完")
		}
		if payNum > item.Stock {
			return false, errors.New("库存不足")
		}
	}

	// 多规格
	if item.SpecType == 1 {
		attrValue, err := p.GetItemAttrValueById(itemId, attrValueId)
		if err != nil {
			return false, err
		}
		if attrValue.Id == 0 {
			return false, errors.New("商品已下架")
		}
		if attrValue.Stock <= 0 {
			return false, errors.New("商品已售完")
		}
		if payNum > attrValue.Stock {
			return false, errors.New("库存不足")
		}
	}

	return true, nil
}

// 重建items表attr_values字段值
func (p *ItemService) RebuildItemAttrValues(itemId int) (err error) {
	item, err := NewItemService().GetItem(itemId, nil, false)
	if err != nil {
		return err
	}
	if item.SpecType != 1 {
		return
	}
	attrValues := item.AttrValues
	for k, attrValue := range attrValues {
		attrValue.Image = attrValue.ImageJson
		attrValues[k] = attrValue
	}
	attrValuesJsonData, err := json.Marshal(attrValues)
	if err != nil {
		return
	}
	return db.Client.Model(&model.Item{}).Where("id = ?", itemId).Update("attr_values", string(attrValuesJsonData)).Error
}

// 获取商品分类
func (p *ItemService) GetCategoriesByPid(pid int) []response.ItemCategoryResp {
	itemCategories := make([]response.ItemCategoryResp, 0)
	db.Client.Model(model.Category{}).
		Select("id", "pid", "title", "cover_id").
		Where("status = ? AND type = ?", 1, "ITEM").
		Where("pid = ?", pid).
		Order("sort, id").
		Find(&itemCategories)
	return itemCategories
}
