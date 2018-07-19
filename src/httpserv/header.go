package httpserv

import (
	"mime/multipart"
	"net/http"
	"strings"
)

// GetIP 获取请求方IP地址。
func GetIP(request *http.Request) string {
	if cfg.RealIP != "" {
		if realIP := GetHeader(request, cfg.RealIP); realIP != "" {
			return realIP
		}
	}

	ip := request.RemoteAddr
	ip = ip[0:strings.LastIndex(ip, ":")]
	if ip == "[::1]" {
		ip = "127.0.0.1"
	}

	return ip
}

// GetHeader 获取请求头信息。
func GetHeader(request *http.Request, name string) string {
	return request.Header.Get(name)
}

// AppendSuffix 添加文件名后缀。
func AppendSuffix(name string, handler *multipart.FileHeader) string {
	lastIndex := strings.LastIndex(handler.Filename, ".")
	if lastIndex > -1 {
		return name + handler.Filename[lastIndex:]
	}

	contentType := handler.Header.Get("content-type")
	lastIndex = strings.LastIndex(contentType, "/")
	if lastIndex > -1 {
		return name + "." + contentType[lastIndex+1:]
	}

	return name
}

// SetHeader 设置返回头信息。
func SetHeader(writer http.ResponseWriter, name string, value string) {
	writer.Header().Set(name, value)
}
