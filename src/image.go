package main

import (
	"fmt"
	"image/jpeg"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"protocol"
	"strconv"
	"strings"
	"util"

	"github.com/nfnt/resize"
)

type config struct {
	Listen  string
	Root    string
	SaveUri string
}

var cfg config = config{":8192", "image", "/save"}

func handler(writer http.ResponseWriter, request *http.Request) {
	uri := request.RequestURI
	if uri == cfg.SaveUri {
		save(writer, request, uri)
	} else {
		read(writer, request, uri)
	}
}

func save(writer http.ResponseWriter, request *http.Request, uri string) {
	request.ParseMultipartForm(1024 * 1024 * 1024)
	if !util.CheckSign(request.Form) {
		log.Println("fail to check sign !")
		protocol.Send404(writer)

		return
	}

	name := request.Form["name"][0]
	empty := name == ""
	if empty {
		file, _, err := request.FormFile("file")
		if err != nil {
			protocol.Send404(writer)

			return
		}
		defer file.Close()
		if name, err = util.Md5FromReader(file); err != nil {
			protocol.Send404(writer)

			return
		}
		name = name + ".jpg"
	}

	path := request.Form["path"][0]
	if path != "" {
		os.MkdirAll(absolute(path), os.ModePerm)
		name = path + "/" + name
	}
	if empty && util.Exists(name) {
		fmt.Fprintf(writer, "%s", name)

		return
	}

	file, _, err := request.FormFile("file")
	if err != nil {
		protocol.Send404(writer)

		return
	}
	defer file.Close()

	out, err := os.OpenFile(absolute(name), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		protocol.Send404(writer)

		return
	}
	defer out.Close()
	io.Copy(out, file)
	clean(path, name)

	fmt.Fprintf(writer, "%s", name)
}

func clean(path string, name string) {
	files, err := ioutil.ReadDir(absolute(path))
	if err != nil {
		return
	}

	names := strings.Split(name, ".")
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		ns := strings.Split(file.Name(), ".")
		if len(ns) == 4 && ns[0] == names[0] && ns[3] == names[1] {
			os.Remove(absolute(path + "/" + file.Name()))
		}
	}
}

func read(writer http.ResponseWriter, request *http.Request, uri string) {
	path := absolute(uri)
	if util.Exists(path) {
		http.ServeFile(writer, request, path)

		return
	}

	index := strings.LastIndex(uri, "/")
	names := strings.Split(uri[index+1:], ".")
	if len(names) != 4 || (names[3] != "jpg" && names[3] != "jpeg") {
		protocol.Send404(writer)

		return
	}

	origin := absolute(uri[0:index+1] + names[0] + "." + names[3])
	if !util.Exists(origin) {
		protocol.Send404(writer)

		return
	}

	scale, err := strconv.Atoi(names[1])
	if err != nil {
		protocol.Send404(writer)

		return
	}

	quality, err := strconv.Atoi(names[2])
	if err != nil {
		protocol.Send404(writer)

		return
	}

	file, err := os.Open(origin)
	if err != nil {
		protocol.Send404(writer)

		return
	}
	defer file.Close()

	image, err := jpeg.Decode(file)
	if err != nil {
		protocol.Send404(writer)

		return
	}

	img := resize.Resize(uint(image.Bounds().Dx()*scale/100), 0, image, resize.Lanczos3)
	out, err := os.Create(path)
	if err != nil {
		protocol.Send404(writer)

		return
	}
	defer out.Close()

	jpeg.Encode(out, img, &jpeg.Options{Quality: quality})
	http.ServeFile(writer, request, path)
}

func absolute(uri string) string {
	return cfg.Root + uri
}

func main() {
	if err := util.LoadConfig(&cfg, "image"); err != nil {
		return
	}

	cfg.Root, _ = filepath.Abs(cfg.Root)
	os.MkdirAll(cfg.Root, os.ModePerm)
	cfg.Root = cfg.Root + "/"

	protocol.Http(cfg.Listen, "/", handler)
}
