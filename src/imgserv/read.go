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
	uri=uri[len(cfg.Root):len(uri)]
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
