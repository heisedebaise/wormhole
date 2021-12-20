package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"sync/atomic"
	"time"
)

var units = []string{"B", "KB", "MB", "GB", "TB"}
var count, request, response int64

func main() {
	bs, err := os.ReadFile("map.json")
	if err != nil {
		log.Printf("load map.json fail %v\n", err)

		return
	}

	m := make(map[string]interface{})
	if err = json.Unmarshal(bs, &m); err != nil {
		log.Printf("unmarshal map.json fail %v\n", err)

		return
	}

	if len(m) == 0 {
		log.Printf("map.json %v is empty\n", m)

		return
	}

	go stat()

	ch := make(chan bool, 1)
	for addr := range m {
		if server, ok := m[addr].(string); ok {
			log.Printf("listening %s to %s\n", addr, server)
			if strings.HasPrefix(server, "https://") {
				go https(addr, server)
			} else {
				go serve(addr, server)
			}
		}
	}
	<-ch
}

func stat() {
	for {
		time.Sleep(time.Minute)
		n1, u1 := flow(request)
		n2, u2 := flow(response)
		atomic.StoreInt64(&request, 0)
		atomic.StoreInt64(&response, 0)
		log.Printf("count=%d;request=%d%s/s;response=%d%s/s\n", count, n1, units[u1], n2, units[u2])
	}
}

func flow(n int64) (int64, int) {
	n /= 60
	unit := 0
	for n > 1024 {
		n >>= 10
		unit += 1
	}

	return n, unit
}

func serve(addr, server string) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer listener.Close()

	for {
		listen(listener, server)
	}
}

func listen(listener net.Listener, server string) {
	accept, err := listener.Accept()
	if err != nil {
		return
	}

	log.Printf("forward %v to %s\n", accept.LocalAddr(), server)
	go agent(accept, server)
}

func agent(accept net.Conn, server string) {
	defer accept.Close()

	dial, err := net.Dial("tcp", server)
	if err != nil {
		log.Printf("dial to %s fail %v\n", server, err)

		return
	}
	defer dial.Close()

	atomic.AddInt64(&count, 1)
	ch := make(chan bool, 2)
	go copy(accept, dial, &request, ch)
	go copy(dial, accept, &response, ch)
	<-ch
	atomic.AddInt64(&count, -1)
}

func copy(reader io.Reader, wirter io.Writer, sum *int64, ch chan bool) {
	buffer := make([]byte, 1024)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}

		if n > 0 {
			wirter.Write(buffer[:n])
			atomic.AddInt64(sum, int64(n))
		}
	}
	ch <- true
}

func https(addr, server string) error {
	http.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
		atomic.AddInt64(&count, 1)
		atomic.AddInt64(&request, req.ContentLength)
		defer func() {
			req.Body.Close()
			atomic.AddInt64(&count, -1)
		}()

		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Printf("read request body %s%s to %s:%s fail %v\n", addr, req.RequestURI, server, req.RequestURI, err)

			return
		}

		r, err := http.NewRequest(req.Method, server+req.RequestURI, bytes.NewBuffer(body))
		if err != nil {
			log.Printf("new request %s%s to %s:%s fail %v\n", addr, req.RequestURI, server, req.RequestURI, err)

			return
		}

		client := http.Client{
			Timeout:   time.Minute,
			Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
		}
		r.Header = req.Header
		res, err := client.Do(r)
		if err != nil {
			log.Printf("do request %s%s to %s:%s fail %v\n", addr, req.RequestURI, server, req.RequestURI, err)

			return
		}
		defer res.Body.Close()

		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Printf("read response %s%s from %s:%s fail %v\n", addr, req.RequestURI, server, req.RequestURI, err)

			return
		}

		n, err := writer.Write(b)
		if err != nil {
			log.Printf("write response %s%s from %s:%s fail %v\n", addr, req.RequestURI, server, req.RequestURI, err)

			return
		}

		atomic.AddInt64(&response, int64(n))
	})

	return http.ListenAndServeTLS(addr, "cert.pem", "key.pem", nil)
}
