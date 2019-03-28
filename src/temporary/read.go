package temporary

import (
	"httpserv"
	"net/http"
	"strings"
)

func read(writer http.ResponseWriter, request *http.Request, uri string) int {
	if indexOf := strings.Index(uri, "?"); indexOf > -1 {
		uri = uri[0:indexOf]
	}

	return httpserv.Read(writer, request, absolute(uri[len(cfg.Root):]))
}
