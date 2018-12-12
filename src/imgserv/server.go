package imgserv

import (
	"httpserv"
	"net/http"
)

func handle(writer http.ResponseWriter, request *http.Request, uri string) int {
	if uri == cfg.Save {
		return save(writer, request)
	}

	return read(writer, request, uri)
}

// Serve 服务。
func Serve() {
	httpserv.Handler(cfg.Root, handle)
}
