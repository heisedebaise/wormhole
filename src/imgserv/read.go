package imgserv

import (
	"image/jpeg"
	"log"
	"net/http"
	"os"
	"protocol"
	"strconv"
	"strings"
	"util"

	"github.com/nfnt/resize"
)

func read(writer http.ResponseWriter, request *http.Request, uri string) {
	uri = uri[len(cfg.Root):]
	path := absolute(uri)
	info, err := os.Stat(path)
	if err == nil {
		if info.IsDir() {
			protocol.Send404(writer)
			log.Printf("read dir %s\n", uri)
		} else {
			serveFile(writer, request, path)
		}

		return
	}

	index := strings.LastIndex(uri, "/") + 1
	names := strings.Split(uri[index:], ".")
	length := len(names)
	if length <= 2 {
		protocol.Send404(writer)
		log.Printf("length<=2 %s\n", uri)

		return
	}

	suffix := names[length-1]
	if suffix != "jpg" && suffix != "jpeg" {
		protocol.Send404(writer)
		log.Printf("not a jpeg file %s\n", uri)

		return
	}

	origin := absolute(uri[0:index] + names[0] + "." + suffix)
	if !util.Exists(origin) {
		protocol.Send404(writer)
		log.Printf("origin jpeg file not exists %s\n", uri)

		return
	}

	scale, quality, err := getScaleQuality(names)
	if err != nil || (scale == 0 && quality == 0) {
		protocol.Send404(writer)
		log.Printf("fail to get scale[%d] or quality[%d] %s %q\n", scale, quality, uri, err)

		return
	}

	file, err := os.Open(origin)
	if err != nil {
		protocol.Send404(writer)
		log.Printf("fail to read origin jpeg file %s %q\n", uri, err)

		return
	}
	defer file.Close()

	image, err := jpeg.Decode(file)
	if err != nil {
		protocol.Send404(writer)
		log.Printf("fail to decode origin jpeg file %s %q\n", uri, err)

		return
	}

	if scale > 0 {
		image = resize.Resize(uint(image.Bounds().Dx()*scale/100), 0, image, resize.Lanczos3)
	}

	out, err := os.Create(path)
	if err != nil {
		protocol.Send404(writer)
		log.Printf("fail to create scale|quality jpeg file %s %q\n", uri, err)

		return
	}
	defer out.Close()

	if quality <= 0 || quality > 100 {
		quality = 100
	}

	if err := jpeg.Encode(out, image, &jpeg.Options{Quality: quality}); err != nil {
		protocol.Send404(writer)
		log.Printf("fail to encode scale[%d]|quality[%d] jpeg file %s %q\n", scale, quality, uri, err)

		return
	}

	serveFile(writer, request, path)
}

func getScaleQuality(names []string) (scale int, quality int, err error) {
	scale = 0
	quality = 0
	for i := 1; i < len(names)-1; i++ {
		separator := len(names[i]) - 1
		number, err := strconv.Atoi(names[i][0:separator])
		if err != nil {
			return scale, quality, err
		}

		suffix := names[i][separator:]
		if suffix == "s" {
			scale = number
		} else if suffix == "q" {
			quality = number
		}
	}

	return scale, quality, nil
}

func serveFile(writer http.ResponseWriter, request *http.Request, path string) {
	log.Println("##########################")
	for key, value := range request.Header {
		log.Printf("%s:%s=%s\n", path, key, value)
	}

	info, _ := os.Stat(path)
	etag := strconv.FormatInt(info.ModTime().UnixNano(), 16)
	protocol.SetHeader(writer, "ETag", etag)
	http.ServeFile(writer, request, path)
}
