package imgserv

import (
	"net/http"
)

// Root 获取URI前缀。
func Root() string {
	return cfg.Root
}

// Handler 处理HTTP(S)请求。
func Handler(writer http.ResponseWriter, request *http.Request, uri string) {
	if uri == cfg.Save {
		save(writer, request)
	} else {
		read(writer, request, uri)
	}
}
