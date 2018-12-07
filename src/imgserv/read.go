package imgserv

import (
	"httpserv"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"util"

	"github.com/nfnt/resize"
)

func read(writer http.ResponseWriter, request *http.Request, uri string) int {
	if indexOf := strings.Index(uri, "?"); indexOf > -1 {
		uri = uri[0:indexOf]
	}
	uri = uri[len(cfg.Root):]
	path := absolute(uri)
	info, err := os.Stat(path)
	if err == nil {
		if info.IsDir() {
			log.Printf("read dir %s\n", uri)

			return httpserv.Send404(writer)
		}

		return httpserv.ServeFile(writer, request, info, path)
	}

	index := strings.LastIndex(uri, "/") + 1
	names := strings.Split(uri[index:], ".")
	length := len(names)
	if length <= 2 {
		log.Printf("length<=2 %s\n", uri)

		return httpserv.Send404(writer)
	}

	suffix := names[length-1]
	if suffix != "jpg" && suffix != "jpeg" && suffix != "png" {
		log.Printf("not a jpeg|png file %s\n", uri)

		return httpserv.Send404(writer)
	}

	origin := absolute(uri[0:index] + names[0] + "." + suffix)
	if !util.Exists(origin) {
		log.Printf("origin jpeg|png file not exists %s\n", uri)

		return httpserv.Send404(writer)
	}

	scale, quality, err := getScaleQuality(names)
	if err != nil || (scale == 0 && quality == 0) {
		log.Printf("fail to get scale[%d] or quality[%d] %s %q\n", scale, quality, uri, err)

		return httpserv.Send404(writer)
	}

	file, err := os.Open(origin)
	if err != nil {
		log.Printf("fail to read origin jpeg|png file %s %q\n", uri, err)

		return httpserv.Send404(writer)
	}
	defer file.Close()

	image, err := decode(suffix, file)
	if err != nil {
		log.Printf("fail to decode origin jpeg file %s %q\n", uri, err)

		return httpserv.Send404(writer)
	}

	if scale > 0 {
		image = resize.Resize(uint(image.Bounds().Dx()*scale/100), 0, image, resize.Lanczos3)
	}

	out, err := os.Create(path)
	if err != nil {
		log.Printf("fail to create scale|quality jpeg file %s %q\n", uri, err)

		return httpserv.Send404(writer)
	}
	defer out.Close()

	if quality <= 0 || quality > 100 {
		quality = 100
	}

	if err := encode(suffix, out, image, quality); err != nil {
		log.Printf("fail to encode scale[%d]|quality[%d] jpeg file %s %q\n", scale, quality, uri, err)

		return httpserv.Send404(writer)
	}

	return httpserv.ServeFile(writer, request, nil, path)
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

func decode(suffix string, file *os.File) (image.Image, error) {
	if suffix == "png" {
		return png.Decode(file)
	}

	return jpeg.Decode(file)
}

func encode(suffix string, out *os.File, image image.Image, quality int) error {
	if suffix == "png" {
		return png.Encode(out, image)
	}

	return jpeg.Encode(out, image, &jpeg.Options{Quality: quality})
}
