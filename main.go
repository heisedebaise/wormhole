package main

import (
	"encoding/json"
	"io"
	"log"
	"net"
	"os"
	"sync/atomic"
	"time"
)

var units = []string{"B", "KB", "MB", "GB", "TB"}
var count, request, response int32

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
			go serve(addr, server)
		}
	}
	<-ch
}

func stat() {
	for {
		time.Sleep(time.Minute)
		n1, u1 := flow(request)
		n2, u2 := flow(response)
		atomic.StoreInt32(&request, 0)
		atomic.StoreInt32(&response, 0)
		log.Printf("count=%d;request=%d%s/s;response=%d%s/s\n", count, n1, units[u1], n2, units[u2])
	}
}

func flow(n int32) (int32, int) {
	n /= 60
	unit := 0
	for n > 1024 {
		n >>= 10
		unit += 1
	}

	return n, unit
}

func serve(addr, server string) error {
	log.Printf("listening %s to %s\n", addr, server)
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

	atomic.AddInt32(&count, 1)
	ch := make(chan bool, 2)
	go copy(accept, dial, &request, ch)
	go copy(dial, accept, &response, ch)
	<-ch
	atomic.AddInt32(&count, -1)
}

func copy(reader io.Reader, wirter io.Writer, sum *int32, ch chan bool) {
	buffer := make([]byte, 1024)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}

		if n > 0 {
			wirter.Write(buffer[:n])
			atomic.AddInt32(sum, int32(n))
		}
	}
	ch <- true
}
