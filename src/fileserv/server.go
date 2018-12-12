package fileserv

import (
	"httpserv"
	"net/http"
)

func handle(writer http.ResponseWriter, request *http.Request, uri string) int {
	if uri == cfg.Save {
		_, _, code := httpserv.Save(writer, request, maxSize, root, cfg.Root)

		return code
	}

	return read(writer, request, uri)
}

// Serve 服务。
func Serve() {
	httpserv.Handler(cfg.Root, handle)
}
