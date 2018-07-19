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
			if trustable = trust(message); trustable {
				log.Printf("同步认证[%s]成功\n", conn.RemoteAddr().String())
				write(conn, []byte(id))

				return true
			}

			return false
		}

		log.Printf("receive: %s\n", string(message))

		return true
	})
}

func trust(message []byte) bool {
	argot := string(message)
	if argot == cfg.Argot {
		return true
	}

	log.Printf("同步认证失败[%s!=%s]！\n", argot, cfg.Argot)

	return false
}
