package config

import (
	"time"

	"github.com/quarkcloudio/quark-smart/v2/pkg/env"
)

type AppConfig struct {
	Version             string   // 应用版本
	Name                string   // 应用名称
	Debug               bool     // 开启Debug模式
	Recover             bool     // 崩溃后自动恢复
	Pro                 bool     // 开启高级功能
	Env                 string   // 项目环境
	Host                string   // 服务地址
	Key                 string   // 令牌加密key，如果设置绝对不可泄漏
	RootPath            string   // Web根目录
	StaticPath          string   // 静态文件路径
	TemplatePath        string   // 模版文件路径
	UploadFileSize      int64    // 上传文件大小限制
	UploadFileType      []string // 上传文件类型限制
	UploadFileSavePath  string   // 上传文件保存路径
	UploadImageSize     int64    // 上传图片大小限制
	UploadImageType     []string // 上传图片类型限制
	UploadImageSavePath string   // 上传图片保存路径
	Logger              bool     // 是否开启日志
	LoggerFilePath      string   // 日志文件路径
}

// APP配置信息
var App = &AppConfig{

	// 应用版本
	Version: "2.0.1",

	// 应用名称
	Name: env.Get("APP_NAME", "QuarkSmart").(string),

	// 开启Debug模式
	Debug: false,

	// 崩溃后自动恢复
	Recover: true,

	// 开启高级功能
	Pro: env.Get("APP_PRO").(string) == "true",

	// 项目环境
	Env: env.Get("APP_ENV").(string),

	// 服务地址
	Host: env.Get("APP_HOST", "127.0.0.1:3000").(string),

	// 令牌加密key，如果设置绝对不可泄漏
	Key: env.Get("APP_KEY").(string),

	// Web根目录
	RootPath: env.Get("APP_ROOT_PATH", "./web/app").(string),

	// 静态文件路径
	StaticPath: env.Get("APP_STATIC_PATH", "./web/static").(string),

	// 模版文件路径
	TemplatePath: env.Get("APP_TEMPLATE_PATH", "web/template/*.html").(string),

	// 上传文件大小限制
	UploadFileSize: 1024 * 1024 * 1024 * 2,

	// 上传文件类型限制，尽量使用文件的MIME名称
	UploadFileType: []string{
		"image/png",
		"image/gif",
		"image/jpeg",
		"video/mp4",
		"video/mpeg",
		"application/x-xls",
		"application/x-ppt",
		"application/msword",
		"application/zip",
		"application/pdf",
		"application/vnd.ms-excel",
		"application/vnd.ms-powerpoint",
		"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
		"application/vnd.openxmlformats-officedocument.wordprocessingml.document",
		"application/vnd.openxmlformats-officedocument.presentationml.presentation",
	},

	// 上传文件保存路径
	UploadFileSavePath: "./web/app/storage/files/" + time.Now().Format("20060102") + "/",

	// 上传图片大小限制
	UploadImageSize: 1024 * 1024 * 1024 * 2,

	// 上传图片类型限制，尽量使用文件的MIME名称
	UploadImageType: []string{
		"image/png",
		"image/gif",
		"image/jpeg",
	},

	// 上传图片保存路径
	UploadImageSavePath: "./web/app/storage/images/" + time.Now().Format("20060102") + "/",

	// 是否开启日志
	Logger: false,

	// 日志文件路径
	LoggerFilePath: "./app.log",
}
