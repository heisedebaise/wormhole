package temporary

import (
	"httpserv"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"util"
)

func copy(writer http.ResponseWriter, request *http.Request) int {
	request.ParseForm()
	if !util.InWhiteList(httpserv.GetIP(request)) && !util.CheckSign(request.Form) {
		return httpserv.Send404(writer)
	}

	uri := httpserv.GetParam(request, "uri", "")
	if uri == "" {
		log.Printf("uri is empty\n")

		return httpserv.Send404(writer)
	}

	path, _ := filepath.Abs(uri[1:])
	if info, err := os.Stat(path); err != nil || info.IsDir() {
		log.Printf("not exists or read dir %s %q\n", path, err)

		return httpserv.Send404(writer)
	}

	source, err := os.Open(path)
	if err != nil {
		log.Printf("open %s fail %q\n", path, err)

		return httpserv.Send404(writer)
	}
	defer source.Close()

	name := util.RandomString(32) + uri[strings.LastIndex(uri, "."):]
	target, err := os.Create(absolute(name))
	if err != nil {
		log.Printf("create %s fail %q\n", absolute(name), err)

		return httpserv.Send404(writer)
	}
	defer target.Close()

	if _, err := io.Copy(target, source); err != nil {
		log.Printf("copy %s fail %q\n", absolute(name), err)

		return httpserv.Send404(writer)
	}

	writer.Write([]byte(cfg.Root + name))

	return httpserv.SendCode(writer, 200)
}
