package wormhole

import (
	"crypto/tls"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type httphandler struct {
	to string
}

func (h *httphandler) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	request, err := http.NewRequest(req.Method, h.to+req.RequestURI, req.Body)
	if err != nil {
		return
	}

	to := h.to[strings.Index(h.to, "://")+3:]
	for key := range req.Header {
		request.Header.Set(key, strings.ReplaceAll(req.Header.Get(key), req.Host, to))
	}
	client := http.Client{
		Timeout:   time.Minute,
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
	}
	res, err := client.Do(request)
	if err != nil {
		return
	}
	defer res.Body.Close()

	for key := range res.Header {
		writer.Header().Set(key, res.Header.Get(key))
	}
	writer.WriteHeader(res.StatusCode)
	io.Copy(writer, res.Body)
}

func serveHTTP(cfg map[string]string) {
	for on := range cfg {
		log.Printf("listening http %s to %s\n", on, cfg[on])
		go http.ListenAndServe(on, &httphandler{cfg[on]})
	}
}
