package rsync

import (
	"log"
	"net"
)

// Listen 启动同步监听。
func Listen() {
	go listen()
}

func listen() {
	listener, err := net.Listen("tcp", cfg.Listen)
	if err != nil {
		log.Fatalln(err)

		return
	}

	log.Printf("rsync listening on %s\n", cfg.Listen)
	defer listener.Close()

	go connect()

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go receive(conn)
	}
}

func receive(conn net.Conn) {
	defer conn.Close()
	trustable := false
	read(conn, func(message []byte) bool {
		if !trustable {
			argot := string(message)
			if trustable = argot == cfg.Argot; trustable {
				log.Printf("remote [%s] argot is right\n", conn.RemoteAddr().String())
				write(conn, []byte(id))

				return true
			}

			log.Printf("illegal remote [%s] argot [%s]\n", conn.RemoteAddr().String(), argot)

			return false
		}

		log.Printf("receive: %s\n", string(message))

		return true
	})
}
