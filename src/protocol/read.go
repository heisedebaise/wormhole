package protocol

import (
	"net/http"
	"os"
	"strconv"
)

// ServeFile 读取文件并返回。
func ServeFile(writer http.ResponseWriter, request *http.Request, info os.FileInfo, path string) {
	if info == nil {
		info, _ = os.Stat(path)
	}
	etag := strconv.FormatInt(info.ModTime().UnixNano(), 16)
	if GetHeader(request, "If-None-Match") == etag {
		SendCode(writer, 304)

		return
	}

	SetHeader(writer, "ETag", etag)
	http.ServeFile(writer, request, path)
}
