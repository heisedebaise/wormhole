package httpserv

import "net/http"

// GetParam 获取请求参数。
func GetParam(request *http.Request, name string, defaultValue string) string {
	if _, has := request.Form[name]; has {
		return request.Form[name][0]
	}

	return defaultValue
}
