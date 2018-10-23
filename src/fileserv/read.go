package fileserv

import (
	"httpserv"
	"log"
	"net/http"
	"os"
	"strings"
)

func read(writer http.ResponseWriter, request *http.Request, uri string) int {
	uri = uri[len(cfg.Root):]
	path := absolute(uri)
	info, err := os.Stat(path)
	if err != nil || info.IsDir() {
		log.Printf("not exists or read dir %s %q\n", path, err)

		return httpserv.Send404(writer)
	}

	httpserv.SetHeader(writer, "Content-Disposition", "attachment;filename="+uri[strings.LastIndex(uri, "/")+1:])

	return httpserv.ServeFile(writer, request, info, path)
}
