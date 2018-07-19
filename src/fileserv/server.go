package fileserv

import (
	"net/http"
	"protocol"
)

// Root 获取URI前缀。
func Root() string {
	return cfg.Root
}

// Handler 处理HTTP(S)请求。
func Handler(writer http.ResponseWriter, request *http.Request, uri string) int {
	if uri == cfg.Save {
		_, _, code := protocol.Save(writer, request, maxSize, root, cfg.Root)

		return code
	}

	return read(writer, request, uri)
}
