package httpserv

import (
	"net/http"
	"os"
	"strconv"
)

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
