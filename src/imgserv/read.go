package imgserv

import (
	"image/jpeg"
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
		} else {
			http.ServeFile(writer, request, path)
		}

		return
	}

	index := strings.LastIndex(uri, "/")
	names := strings.Split(uri[index+1:], ".")
	length := len(names)
	if length <= 2 {
		protocol.Send404(writer)

		return
	}

	suffix := names[length-1]
	if suffix != "jpg" && suffix != "jpeg" {
		protocol.Send404(writer)

		return
	}

	origin := absolute(uri[0:index+1] + names[0] + "." + suffix)
	if !util.Exists(origin) {
		protocol.Send404(writer)

		return
	}

	scale, quality, err := getScaleQuality(names)
	if err != nil || (scale == 0 && quality == 0) {
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

	if scale > 0 {
		image = resize.Resize(uint(image.Bounds().Dx()*scale/100), 0, image, resize.Lanczos3)
	}

	out, err := os.Create(path)
	if err != nil {
		protocol.Send404(writer)

		return
	}
	defer out.Close()

	if quality <= 0 || quality > 100 {
		quality = 100
	}

	jpeg.Encode(out, image, &jpeg.Options{Quality: quality})
	http.ServeFile(writer, request, path)
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
