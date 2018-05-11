package protocol

import "net/http"

func GetParam(request *http.Request, name string, defaultValue string) string {
	if _, has := request.Form[name]; has {
		return request.Form[name][0]
	}

	return defaultValue
}
