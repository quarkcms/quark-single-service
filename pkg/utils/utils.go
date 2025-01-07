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
	return service.NewAttachmentService().GetFilePath(id)
}

// 获取多文件路径
func GetFilePaths(id interface{}) []string {
	if id == nil {
		return nil
	}
	return service.NewAttachmentService().GetPaths(id)
}

// 获取图片路径
func GetImagePath(id interface{}) string {
	if id == nil {
		return ""
	}
	return service.NewAttachmentService().GetImagePath(id)
}

// 获取多图片路径
func GetImagePaths(id interface{}) []string {
	if id == nil {
		return nil
	}
	return service.NewAttachmentService().GetPaths(id)
}

// 设置配置
func SetConfig(key string, value string) {
	service.NewConfigService().SetValue(key, value)
}

// 获取配置
func GetConfig(key string) string {
	return service.NewConfigService().GetValue(key)
}

// 获取域名
func GetDomain() string {
	domain := service.NewConfigService().GetValue("WEB_SITE_DOMAIN")
	ssl := service.NewConfigService().GetValue("SSL_OPEN")
	http := ""
	if domain != "" {
		http = "http://"
		if ssl == "1" {
			http = "https://"
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

// 比较工具
// 检查元素item是否存在于切片slice中
// 如果存在，返回true；如果不存在，返回false
func Contains[T comparable](slice []T, item T) bool {
	for _, value := range slice {
		if value == item {
			return true
		}
	}
	return false
}

// 过滤器
// 条件函数返回true，元素会被包含在结果中
func Filter[T interface{}](slice []T, condition func(T) bool) (result []T) {
	for _, value := range slice {
		if condition(value) {
			result = append(result, value)
		}
	}
	return result
}

// 脱敏工具
func Desensitize(content string, start, end int) string {
	if start < 0 || end < 0 || start > end {
		return content
	}
	var contentRune []rune
	for key, value := range content {
		if key >= start && key <= end {
			contentRune = append(contentRune, '*')
		} else {
			contentRune = append(contentRune, value)
		}
	}
	return string(contentRune)
}
