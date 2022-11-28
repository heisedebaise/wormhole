package wormhole

import (
	"bytes"
	"crypto/tls"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type httphandler struct {
	to      string
	replace map[string]string
	capture map[string]string
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
	if len(h.capture) > 0 {
		if ct, ok := h.capture["content-type"]; ok {
			rct := res.Header.Get("content-type")
			if ok, _ = regexp.MatchString(ct, rct); ok {
				h.copy(writer, res, req.RequestURI, rct)

				return
			}
		}
	}

	if len(h.replace) == 0 {
		io.Copy(writer, res.Body)

		return
	}

	bs, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	for key := range h.replace {
		bs = bytes.ReplaceAll(bs, []byte(key), []byte(h.replace[key]))
	}
	writer.Header().Set("content-length", strconv.Itoa(len(bs)))
	writer.Write(bs)
}

func (h *httphandler) copy(writer http.ResponseWriter, res *http.Response, uri, ct string) {
	c := capture{reader: res.Body}
	if err := c.init(uri, ct); err != nil {
		log.Printf("init capture err %v\n", err)

		return
	}
	defer c.close()

	io.Copy(writer, &c)
}

func serveHTTP(cfg map[string]string, replace, capture map[string]string) {
	for on := range cfg {
		log.Printf("listening http %s to %s\n", on, cfg[on])
		go http.ListenAndServe(on, &httphandler{cfg[on], replace, capture})
	}
}
