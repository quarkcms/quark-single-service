package utils

import (
	"regexp"
	"strings"

	"github.com/quarkcloudio/quark-go/v3/service"
)

// 获取文件路径
func GetFilePath(id interface{}) string {
	if id == nil {
		return ""
	}

	return service.NewFileService().GetPath(id)
}

// 获取多文件路径
func GetFilePaths(id interface{}) []string {
	if id == nil {
		return nil
	}

	return service.NewFileService().GetPaths(id)
}

// 获取图片路径
func GetPicturePath(id interface{}) string {
	if id == nil {
		return ""
	}

	return service.NewPictureService().GetPath(id)
}

// 获取多图片路径
func GetPicturePaths(id interface{}) []string {
	if id == nil {
		return nil
	}

	return service.NewPictureService().GetPaths(id)
}

// 获取配置
func GetConfig(key string) string {
	return service.NewConfigService().GetValue(key)
}

// 获取域名
func GetDomain() string {

	http := ""
	domain := service.NewConfigService().GetValue("WEB_SITE_DOMAIN")
	ssl := service.NewConfigService().GetValue("SSL_OPEN")
	if domain != "" {
		if ssl == "1" {
			http = "https://"
		} else {
			http = "http://"
		}
	}

	return http + domain
}

// 内容中的地址替换
func ReplaceContentSrc(content string) string {

	reg := regexp.MustCompile(`src="(/[^"]*)"`)

	return reg.ReplaceAllStringFunc(content, func(src string) string {
		return "src= \"" + GetDomain() + src[strings.Index(src, "\"")+1:] + "\""
	})
}

// 正则验证
// expr 正则表达式
// content 要验证的内容
func CheckRegex(expr, content string) bool {

	r, err := regexp.Compile(expr)
	if err != nil {
		return false
	}

	return r.MatchString(content)
}
