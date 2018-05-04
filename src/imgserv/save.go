package imgserv

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"protocol"
	"strings"
	"util"
)

func save(writer http.ResponseWriter, request *http.Request) {
	request.ParseMultipartForm(maxSize)
	if !util.CheckSign(request.Form) {
		protocol.Send404(writer)

		return
	}

	name := protocol.GetParam(request, "name", "")
	empty := name == ""
	if empty {
		file, handler, err := request.FormFile("file")
		if err != nil {
			log.Printf("fail to load file: %q\n", err)
			protocol.Send404(writer)

			return
		}
		defer file.Close()
		if name, err = util.Md5FromReader(file); err != nil {
			log.Printf("fail to sum md5: %q\n", err)
			protocol.Send404(writer)

			return
		}

		name = suffix(name, handler)
	}

	path := protocol.GetParam(request, "path", "")
	if path != "" {
		os.MkdirAll(absolute(path), os.ModePerm)
		name = path + "/" + name
	}
	if empty && util.Exists(absolute(name)) {
		fmt.Fprintf(writer, "%s", cfg.Root+name)

		return
	}

	file, handler, err := request.FormFile("file")
	if err != nil {
		log.Printf("fail to load file: %q\n", err)
		protocol.Send404(writer)

		return
	}
	defer file.Close()

	lastIndex := strings.LastIndex(name, ".")
	if lastIndex == -1 {
		name = suffix(name, handler)
	}

	out, err := os.OpenFile(absolute(name), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Printf("fail to open file: %q\n", err)
		protocol.Send404(writer)

		return
	}
	defer out.Close()
	io.Copy(out, file)
	clean(path, name)

	fmt.Fprintf(writer, "%s", cfg.Root+name)
}

func suffix(name string, handler *multipart.FileHeader) string {
	return name + handler.Filename[strings.LastIndex(handler.Filename, "."):len(handler.Filename)]
}
