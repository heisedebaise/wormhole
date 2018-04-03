package main

import (
	"fmt"
	"image/jpeg"
	"io"
	"io/ioutil"
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
	Listen string
	Root   string
}

var cfg config = config{":8192", "image"}

func handler(writer http.ResponseWriter, request *http.Request) {
	uri := request.RequestURI
	if uri == "/save" {
		save(writer, request, uri)
	} else {
		read(writer, request, uri)
	}
}

func save(writer http.ResponseWriter, request *http.Request, uri string) {
	filename := request.Header.Get("filename")
	path := request.Header.Get("path")
	file, _, err := request.FormFile("image")
	if err != nil {
		writer.WriteHeader(404)

		return
	}
	defer file.Close()

	empty := filename == ""
	if empty {
		f, _, _ := request.FormFile("image")
		if filename, err = util.Md5(f); err != nil {
			writer.WriteHeader(404)

			return
		}
		filename = filename + ".jpg"
	}

	if path != "" {
		os.MkdirAll(absolute(path), os.ModePerm)
		filename = path + "/" + filename
	}
	if empty && util.Exists(filename) {
		fmt.Fprintf(writer, "%s", filename)

		return
	}

	out, err := os.OpenFile(absolute(filename), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		writer.WriteHeader(404)

		return
	}
	defer out.Close()
	io.Copy(out, file)
	clean(path, filename)

	fmt.Fprintf(writer, "%s", filename)
}

func clean(path string, filename string) {
	files, err := ioutil.ReadDir(absolute(path))
	if err != nil {
		return
	}

	names := strings.Split(filename, ".")
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
		writer.WriteHeader(404)

		return
	}

	origin := absolute(uri[0:index+1] + names[0] + "." + names[3])
	if !util.Exists(origin) {
		writer.WriteHeader(404)

		return
	}

	scale, err := strconv.Atoi(names[1])
	if err != nil {
		writer.WriteHeader(404)

		return
	}

	quality, err := strconv.Atoi(names[2])
	if err != nil {
		writer.WriteHeader(404)

		return
	}

	file, err := os.Open(origin)
	if err != nil {
		writer.WriteHeader(404)

		return
	}
	defer file.Close()

	img, err := jpeg.Decode(file)
	if err != nil {
		writer.WriteHeader(404)

		return
	}

	m := resize.Resize(uint(img.Bounds().Dx()*scale/100), 0, img, resize.Lanczos3)
	out, err := os.Create(path)
	if err != nil {
		writer.WriteHeader(404)

		return
	}
	defer out.Close()

	jpeg.Encode(out, m, &jpeg.Options{Quality: quality})
	http.ServeFile(writer, request, path)
}

func absolute(uri string) string {
	return cfg.Root + uri
}

func main() {
	if err := util.LoadConfig(&cfg, "image"); err != nil {
		fmt.Println("Fail to load conf/image.json")
		fmt.Println(err)

		return
	}

	cfg.Root, _ = filepath.Abs(cfg.Root)
	os.MkdirAll(cfg.Root, os.ModePerm)
	cfg.Root = cfg.Root + "/"

	protocol.Http(cfg.Listen, "/", handler)
}
