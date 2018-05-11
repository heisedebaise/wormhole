package protocol

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"util"
)

func Upload(writer http.ResponseWriter, request *http.Request, maxSize int64, absRoot string, root string) (string, string, bool) {
	request.ParseMultipartForm(maxSize)
	if !util.InWhiteList(GetIp(request)) && !util.CheckSign(request.Form) {
		Send404(writer)

		return "", "", false
	}

	name := GetParam(request, "name", "")
	empty := name == ""
	if empty {
		file, handler, err := request.FormFile("file")
		if err != nil {
			log.Printf("fail to load file: %q\n", err)
			Send404(writer)

			return "", "", false
		}
		defer file.Close()
		if name, err = util.Md5FromReader(file); err != nil {
			log.Printf("fail to sum md5: %q\n", err)
			Send404(writer)

			return "", "", false
		}

		name = AppendSuffix(name, handler)
	}

	path := GetParam(request, "path", "")
	if path != "" {
		os.MkdirAll(util.FormatPath(absRoot+path), os.ModePerm)
	}
	if empty && util.Exists(util.FormatPath(absRoot+path+"/"+name)) {
		fmt.Fprintf(writer, "%s", util.FormatPath(root+path+"/"+name))

		return path, name, false
	}

	file, handler, err := request.FormFile("file")
	if err != nil {
		log.Printf("fail to load file: %q\n", err)
		Send404(writer)

		return "", "", false
	}
	defer file.Close()

	if strings.Index(name, ".") == -1 {
		name = AppendSuffix(name, handler)
	}

	out, err := os.OpenFile(util.FormatPath(absRoot+path+"/"+name), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Printf("fail to open file: %q\n", err)
		Send404(writer)

		return "", "", false
	}
	defer out.Close()
	io.Copy(out, file)

	fmt.Fprintf(writer, "%s", util.FormatPath(root+path+"/"+name))

	return path, name, true
}
