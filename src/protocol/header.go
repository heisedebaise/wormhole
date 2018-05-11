package protocol

import (
	"mime/multipart"
	"net/http"
	"strings"
)

func GetIp(request *http.Request) string {
	ip := request.RemoteAddr
	ip = ip[0:strings.LastIndex(ip, ":")]
	if ip == "[::1]" {
		ip = "127.0.0.1"
	}

	return ip
}

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
