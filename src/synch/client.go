package synch

import (
	"log"
	"net"
)

var servers map[string]*net.TCPConn

func connect() {
	servers = make(map[string]*net.TCPConn, len(cfg.Nodes))
	for _, node := range cfg.Nodes {
		addr, err := net.ResolveTCPAddr("tcp4", node)
		if err != nil {
			log.Println(err)

			continue
		}

		conn, err := net.DialTCP("tcp", nil, addr)
		if err != nil {
			log.Println(err)

			continue
		}

		if err := write(conn, []byte(cfg.Argot)); err == nil {
			read(conn, func(message []byte) bool {
				if string(message) != id {
					servers[node] = conn
				}

				return false
			})
		}
	}
}

// Send 发送
func Send(bytes []byte) {
	for _, conn := range servers {
		if conn == nil {
			continue
		}

		if err := write(conn, bytes); err != nil {
			log.Println(err)
		}
	}
}
