package httpserv

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"rsync"
	"strings"
	"util"
)

// Save 处理文件上传请求。
func Save(writer http.ResponseWriter, request *http.Request, maxSize int64, absRoot string, root string) (string, string, int) {
	request.ParseMultipartForm(maxSize)
	if !util.InWhiteList(GetIP(request)) && !util.CheckSign(request.Form) {
		return "", "", Send404(writer)
	}

	name := GetParam(request, "name", "")
	empty := name == ""
	if empty {
		file, handler, err := request.FormFile("file")
		if err != nil {
			log.Printf("fail to load file: %q\n", err)

			return "", "", Send404(writer)
		}
		defer file.Close()
		if name, err = util.Md5FromReader(file); err != nil {
			log.Printf("fail to sum md5: %q\n", err)

			return "", "", Send404(writer)
		}

		name = AppendSuffix(name, handler)
	}

	path := GetParam(request, "path", "")
	if path != "" {
		os.MkdirAll(util.FormatPath(absRoot+path), os.ModePerm)
	}
	log.Println(absolute(absRoot, path, name))
	if empty && util.Exists(util.FormatPath(absRoot+path+"/"+name)) {
		fmt.Fprintf(writer, "%s", util.FormatPath(root+path+"/"+name))

		return path, name, 200
	}

	file, handler, err := request.FormFile("file")
	if err != nil {
		log.Printf("fail to load file: %q\n", err)

		return "", "", Send404(writer)
	}
	defer file.Close()

	if strings.Index(name, ".") == -1 {
		name = AppendSuffix(name, handler)
	}

	absPath := util.FormatPath(absRoot + path + "/" + name)
	out, err := os.OpenFile(absPath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Printf("fail to open file: %q\n", err)

		return "", "", Send404(writer)
	}
	defer out.Close()
	io.Copy(out, file)

	uri := util.FormatPath(root + path + "/" + name)
	fmt.Fprintf(writer, "%s", util.FormatPath(root+path+"/"+name))

	rsync.SendFile(uri, absPath)

	return path, name, 200
}
