package main

import (
	"encoding/json"
	"io"
	"log"
	"net"
	"os"
)

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

	ch := make(chan bool, 1)
	for addr := range m {
		if server, ok := m[addr].(string); ok {
			go serve(addr, server)
		}
	}
	<-ch
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

	ch := make(chan bool, 2)
	go copy(accept, dial, ch)
	go copy(dial, accept, ch)
	<-ch
}

func copy(reader io.Reader, wirter io.Writer, ch chan bool) {
	buffer := make([]byte, 1024)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}

		if n > 0 {
			wirter.Write(buffer[:n])
		}
	}
	ch <- true
}
