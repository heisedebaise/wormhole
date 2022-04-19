package wormhole

import (
	"compress/gzip"
	"crypto/tls"
	"io"
	"log"
	"net/http"
	"strings"
	"sync/atomic"
	"time"
)

type httpcfg struct {
	On  string `json:"on"`
	TLS struct {
		Key  string `json:"key"`
		Cert string `json:"cert"`
	} `json:"tls"`
	OnV3 bool   `json:"onv3"`
	To   string `json:"to"`
	ToV3 bool   `json:"tov3"`
}

type handler struct {
	on string
	to string
}

func (h *handler) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	atomic.AddInt64(&count, 1)
	atomic.AddInt64(&request, req.ContentLength)
	defer func() {
		req.Body.Close()
		atomic.AddInt64(&count, -1)
	}()

	r, err := http.NewRequest(req.Method, h.to+req.RequestURI, req.Body)
	if err != nil {
		log.Printf("new request %s%s to %s:%s err %v\n", h.on, req.RequestURI, h.to, req.RequestURI, err)

		return
	}

	client := http.Client{Timeout: time.Minute}
	if strings.HasPrefix(h.to, "https://") {
		client.Transport = &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	}
	r.Header = req.Header
	res, err := client.Do(r)
	if err != nil {
		log.Printf("do request %s%s to %s:%s err %v\n", h.on, req.RequestURI, h.to, req.RequestURI, err)

		return
	}
	defer res.Body.Close()

	writer.WriteHeader(res.StatusCode)
	gz := false
	for key := range res.Header {
		if key == "Content-Encoding" && res.Header.Get(key) == "gzip" {
			gz = true
		}
		writer.Header().Set(key, res.Header.Get(key))
	}

	var reader io.Reader
	if gz {
		reader, _ = gzip.NewReader(res.Body)
	} else {
		reader = res.Body
	}
	n, err := io.Copy(writer, reader)
	if err != nil {
		log.Printf("write response %s%s from %s:%s err %v\n", h.on, req.RequestURI, h.to, req.RequestURI, err)

		return
	}

	atomic.AddInt64(&response, int64(n))
}

func https(cfgs []httpcfg) {
	for _, cfg := range cfgs {
		if cfg.On == "" || cfg.To == "" {
			continue
		}

		h := &handler{
			on: cfg.On,
			to: cfg.To,
		}
		if cfg.TLS.Key == "" || cfg.TLS.Cert == "" {
			go http.ListenAndServe(cfg.On, h)
		} else {
			go http.ListenAndServeTLS(cfg.On, cfg.TLS.Cert, cfg.TLS.Key, h)
		}
		log.Printf("listening http/s %s to %s\n", cfg.On, cfg.To)
	}
}
