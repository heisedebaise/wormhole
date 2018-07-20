package rsync

import (
	"bytes"
	"log"
	"net"
	"os"
	"time"
)

var servers map[string]net.Conn

func connect() {
	servers = make(map[string]net.Conn, len(cfg.Nodes))

	for {
		reconnect()
		time.Sleep(time.Duration(cfg.ReConnect) * time.Second)
	}
}

func reconnect() {
	for _, node := range cfg.Nodes {
		if conn, ok := servers[node]; ok {
			if conn == nil || alive(conn) {
				continue
			}

			conn.Close()
			delete(servers, node)
			log.Printf("server [%s] has closed\n", node)
		}

		conn, err := net.Dial("tcp", node)
		if err != nil {
			log.Println(err)

			continue
		}

		if err := write(conn, []byte(cfg.Argot)); err == nil {
			read(conn, func(message []byte) bool {
				if string(message) == id {
					conn.Close()
					servers[node] = nil
					log.Printf("close self server [%s] connection\n", node)
				} else {
					servers[node] = conn
					log.Printf("save server [%s] connection\n", node)
				}

				return false
			}, func(conn net.Conn) {
				conn.Close()
				delete(servers, node)
				log.Printf("server [%s] has closed\n", node)
			})
		}
	}
}

// SendFile 发送文件
func SendFile(uri string, path string) {
	var buffer bytes.Buffer
	buffer.WriteByte(byte(1))
	length := len(uri)
	buffer.WriteByte(byte((length >> 8) & 0xff))
	buffer.WriteByte(byte(length & 0xff))
	buffer.WriteString(uri)
	file, err := os.Open(path)
	if err != nil {
		return
	}

	buffer.ReadFrom(file)
	file.Close()
	go send(buffer.Bytes())
}

func send(bytes []byte) {
	for node, conn := range servers {
		if conn == nil {
			continue
		}

		if err := write(conn, bytes); err != nil {
			conn.Close()
			delete(servers, node)
			log.Println(err)
		}
	}
}
