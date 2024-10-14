package service

import "github.com/quarkcloudio/quark-smart/v2/internal/tool/service/upload"

// 注册服务
var Providers = []interface{}{
	&upload.File{},
	&upload.Image{},
}
