package speech

import (
	"httpserv"
	"net/http"
	"util"
)

func save(writer http.ResponseWriter, request *http.Request) int {
	if !util.InWhiteList(httpserv.GetIP(request)) && !util.CheckSign(request.Form) {
		return httpserv.Send404(writer)
	}

	id := httpserv.GetParam(request, "id", "")
	data := httpserv.GetParam(request, "data", "")
	if id == "" || data == "" {
		return httpserv.Send404(writer)
	}

	writer.Write([]byte("success"))

	return 200
}

func getData(auth string) string {
	return getPath(auth, "") + "data"
}
