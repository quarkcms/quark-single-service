package upload

import (
	"reflect"

	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/model"
	"github.com/quarkcloudio/quark-go/v3/service"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/message"
	"github.com/quarkcloudio/quark-go/v3/template/admin/upload"
	"github.com/quarkcloudio/quark-smart/v2/config"
	"github.com/quarkcloudio/quark-smart/v2/internal/dto/response"
)

type File struct {
	upload.Template
}

// 初始化
func (p *File) Init(ctx *quark.Context) interface{} {

	// 限制文件大小
	p.LimitSize = config.App.UploadFileSize

	// 限制文件类型
	p.LimitType = config.App.UploadFileType

	// 设置文件上传路径
	p.SavePath = config.App.UploadFileSavePath

	return p
}

// 上传前回调
func (p *File) BeforeHandle(ctx *quark.Context, fileSystem *quark.FileSystem) (*quark.FileSystem, *quark.FileInfo, error) {
	fileHash, err := fileSystem.GetFileHash()
	if err != nil {
		return fileSystem, nil, err
	}

	getFileInfo, _ := service.NewFileService().GetInfoByHash(fileHash)
	if getFileInfo.Id != 0 {
		fileInfo := &quark.FileInfo{
			Name: getFileInfo.Name,
			Size: getFileInfo.Size,
			Ext:  getFileInfo.Ext,
			Path: getFileInfo.Path,
			Url:  getFileInfo.Url,
			Hash: getFileInfo.Hash,
		}

		return fileSystem, fileInfo, err
	}

	return fileSystem, nil, err
}

// 上传完成后回调
func (p *File) AfterHandle(ctx *quark.Context, result *quark.FileInfo) error {
	driver := reflect.
		ValueOf(ctx.Template).
		Elem().
		FieldByName("Driver").
		String()

	// 重写url
	if driver == quark.LocalStorage {
		result.Url = service.NewFileService().GetPath(result.Url)
	}

	adminInfo, err := service.NewUserService().GetAuthUser(ctx.Engine.GetConfig().AppKey, ctx.Token())
	if err != nil {
		return ctx.JSON(200, message.Error(err.Error()))
	}

	// 插入数据库
	id, err := service.NewFileService().InsertGetId(model.File{
		ObjType: "ADMIN",
		ObjId:   adminInfo.Id,
		Name:    result.Name,
		Size:    result.Size,
		Ext:     result.Ext,
		Path:    result.Path,
		Url:     result.Url,
		Hash:    result.Hash,
		Status:  1,
	})
	if err != nil {
		return ctx.JSON(200, message.Error(err.Error()))
	}

	return ctx.JSON(200, message.Success("上传成功", "", response.UploadFileResp{
		Id:          id,
		ContentType: result.ContentType,
		Ext:         result.Ext,
		Hash:        result.Hash,
		Name:        result.Name,
		Path:        result.Path,
		Size:        result.Size,
		Url:         result.Url,
	}))
}
