package upload

import (
	"reflect"

	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/model"
	"github.com/quarkcloudio/quark-go/v3/service"
	"github.com/quarkcloudio/quark-go/v3/template/tool/upload"
	"github.com/quarkcloudio/quark-smart/v2/config"
)

type Image struct {
	upload.Template
}

// 初始化
func (p *Image) Init(ctx *quark.Context) interface{} {

	// 限制文件大小
	p.LimitSize = config.App.UploadImageSize

	// 限制文件类型
	p.LimitType = config.App.UploadImageType

	// 设置文件上传路径
	p.SavePath = config.App.UploadImageSavePath

	return p
}

// 上传前回调
func (p *Image) BeforeHandle(ctx *quark.Context, fileSystem *quark.FileSystem) (*quark.FileSystem, *quark.FileInfo, error) {
	fileHash, err := fileSystem.GetFileHash()
	if err != nil {
		return fileSystem, nil, err
	}

	pictureInfo, _ := service.NewPictureService().GetInfoByHash(fileHash)
	if pictureInfo.Id != 0 {
		fileInfo := &quark.FileInfo{
			Name:   pictureInfo.Name,
			Size:   pictureInfo.Size,
			Width:  pictureInfo.Width,
			Height: pictureInfo.Height,
			Ext:    pictureInfo.Ext,
			Path:   pictureInfo.Path,
			Url:    pictureInfo.Url,
			Hash:   pictureInfo.Hash,
		}

		return fileSystem, fileInfo, err
	}

	return fileSystem, nil, err
}

// 上传完成后回调
func (p *Image) AfterHandle(ctx *quark.Context, result *quark.FileInfo) error {
	driver := reflect.
		ValueOf(ctx.Template).
		Elem().
		FieldByName("Driver").String()

	// 重写url
	if driver == quark.LocalStorage {
		result.Url = service.NewPictureService().GetPath(result.Url)
	}

	// 插入数据库
	id, err := service.NewPictureService().InsertGetId(model.Picture{
		Name:   result.Name,
		Size:   result.Size,
		Width:  result.Width,
		Height: result.Height,
		Ext:    result.Ext,
		Path:   result.Path,
		Url:    result.Url,
		Hash:   result.Hash,
		Status: 1,
	})

	if err != nil {
		return ctx.JSONError(err.Error())
	}

	return ctx.JSONOk("上传成功", map[string]interface{}{
		"id":          id,
		"contentType": result.ContentType,
		"ext":         result.Ext,
		"hash":        result.Hash,
		"height":      result.Height,
		"width":       result.Width,
		"name":        result.Name,
		"path":        result.Path,
		"size":        result.Size,
		"url":         result.Url,
	})
}
