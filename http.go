package wormhole

import (
    "crypto/tls"
    "io"
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
	uri := req.RequestURI
	request, err := http.NewRequest(req.Method, h.to+uri, req.Body)
	if err != nil {
		return
	}

	to := h.to[strings.Index(h.to, "://")+3:]
	for key := range req.Header {
		request.Header.Set(key, strings.ReplaceAll(req.Header.Get(key), req.Host, to))
	}
	request.Header.Set("Accept-Encoding", "")
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

	var reader io.Reader
	reader = res.Body
	length := 0
	if len(h.replace) > 0 {
		if reader, length, err = replace(res.Body, h.replace); err != nil {
			return
		}
	}

    if length > 0 {
        writer.Header().Set("content-length", strconv.Itoa(length))
    }
    
	if len(h.capture) > 0 {
		if mirror, ok := h.capture["mirror"]; ok && mirror == "yes" {
			h.copy(writer, reader, uri)
		} else if ct, ok := h.capture["content-type"]; ok {
			rct := res.Header.Get("content-type")
			if ok, _ = regexp.MatchString(ct, rct); ok {
				h.copy(writer, reader, uri)
			}
		}
	} else {
		io.Copy(writer, reader)
	}
}

func (h *httphandler) copy(writer http.ResponseWriter, reader io.Reader, uri string) {
	c := capture{reader: reader}
	if err := c.init(uri); err != nil {
		Log("init capture err %v", err)

		return
	}
	defer c.close()

	io.Copy(writer, &c)
}

func serveHTTP(cfg map[string]string, replace, capture map[string]string) {
	for on := range cfg {
		Log("listening http %s to %s", on, cfg[on])
		go http.ListenAndServe(on, &httphandler{cfg[on], replace, capture})
	}
}
