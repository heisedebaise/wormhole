package wormhole

import (
	"io"
	"log"
	"net"
	"sync/atomic"
)

func tcp(config map[string]string) {
	for on := range config {
		go listenTCP(on, config[on])
	}
}

func listenTCP(on, to string) error {
	log.Printf("listening tcp %s to %s\n", on, to)
	listener, err := net.Listen("tcp", on)
	if err != nil {
		return err
	}
	defer listener.Close()

	for {
		accept, err := listener.Accept()
		if err != nil {
			log.Printf("forward %v to %s err %v\n", on, to, err)

			continue
		}

		log.Printf("forward %v to %s\n", accept.LocalAddr(), to)
		go tcpAgent(accept, to)
	}
}

func tcpAgent(accept net.Conn, to string) {
	defer accept.Close()

	dial, err := net.Dial("tcp", to)
	if err != nil {
		log.Printf("dial to %s err %v\n", to, err)

		return
	}
	defer dial.Close()

	atomic.AddInt64(&count, 1)
	ch := make(chan bool, 2)
	go tcpCopy(accept, dial, &request, ch)
	go tcpCopy(dial, accept, &response, ch)
	<-ch
	atomic.AddInt64(&count, -1)
}

func tcpCopy(reader io.Reader, wirter io.Writer, sum *int64, ch chan bool) {
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
