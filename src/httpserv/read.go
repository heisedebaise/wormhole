package httpserv

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// Read 读取/下载文件。
func Read(writer http.ResponseWriter, request *http.Request, path string) int {
	info, err := os.Stat(path)
	if err != nil || info.IsDir() {
		log.Printf("not exists or read dir %s %q\n", path, err)

		return Send404(writer)
	}

	SetHeader(writer, "Content-Disposition", "attachment;filename="+path[strings.LastIndex(path, "/")+1:])

	return ServeFile(writer, request, info, path)
}

// ServeFile 读取文件并返回。
func ServeFile(writer http.ResponseWriter, request *http.Request, info os.FileInfo, path string) int {
	if info == nil {
		var err error
		if info, err = os.Stat(path); err != nil {
			return Send404(writer)
		}
	}

	etag := strconv.FormatInt(info.ModTime().UnixNano(), 16)
	if GetHeader(request, "If-None-Match") == etag {
		return SendCode(writer, 304)
	}

	SetHeader(writer, "ETag", etag)
	http.ServeFile(writer, request, path)

	return 200
}
