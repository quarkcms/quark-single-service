package upload

import (
	"encoding/json"
	"reflect"

	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/model"
	"github.com/quarkcloudio/quark-go/v3/service"
	"github.com/quarkcloudio/quark-go/v3/template/tool/upload"
	"github.com/quarkcloudio/quark-smart/v2/config"
	"github.com/quarkcloudio/quark-smart/v2/internal/dto/response"
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

	imageInfo, err := service.NewAttachmentService().GetInfoByHash(fileHash)
	if err != nil {
		return fileSystem, nil, err
	}

	if imageInfo.Id != 0 {
		var extra map[string]interface{}
		if imageInfo.Extra != "" {
			_ = json.Unmarshal([]byte(imageInfo.Extra), &extra)
		}

		fileInfo := &quark.FileInfo{
			Name:  imageInfo.Name,
			Size:  imageInfo.Size,
			Ext:   imageInfo.Ext,
			Path:  imageInfo.Path,
			Url:   imageInfo.Url,
			Hash:  imageInfo.Hash,
			Extra: extra,
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
		result.Url = service.NewAttachmentService().GetPath(result.Url)
	}

	extra := ""
	if result.Extra != nil {
		extraData, err := json.Marshal(result.Extra)
		if err == nil {
			extra = string(extraData)
		}
	}

	// 插入数据库
	id, err := service.NewAttachmentService().InsertGetId(model.Attachment{
		Name:   result.Name,
		Type:   "IMAGE",
		Size:   result.Size,
		Ext:    result.Ext,
		Path:   result.Path,
		Url:    result.Url,
		Hash:   result.Hash,
		Extra:  extra,
		Status: 1,
	})

	if err != nil {
		return ctx.JSONError(err.Error())
	}

	return ctx.JSONOk("上传成功", response.UploadResp{
		Id:          id,
		ContentType: result.ContentType,
		Ext:         result.Ext,
		Hash:        result.Hash,
		Name:        result.Name,
		Path:        result.Path,
		Size:        result.Size,
		Url:         result.Url,
		Extra:       result.Extra,
	})
}
