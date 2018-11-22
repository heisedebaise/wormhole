package rsync

import (
	"io"
	"log"
	"net"
	"time"
	"util"
)

func read(conn net.Conn, callback func(message []byte) bool, eof func(conn net.Conn)) {
	var messages []byte
	var message []byte
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				eof(conn)
			} else {
				log.Println(err)
			}

			break
		}

		messages = append(messages, buffer[:n]...)
		message, messages = readMessage(messages)
		if message == nil {
			continue
		}

		if !callback(message) {
			break
		}
	}
}

func readMessage(messages []byte) ([]byte, []byte) {
	length := uint32(len(messages))
	if length < 4 {
		return nil, messages
	}

	size := util.BytesToUint32(messages[:4]) + 4
	if length < size {
		return nil, messages
	}

	return messages[4:size], messages[size:]
}

func alive(conn net.Conn) bool {
	buffer := make([]byte, 1024)
	conn.SetReadDeadline(time.Now().Add(time.Millisecond))
	if _, err := conn.Read(buffer); err == io.EOF {
		return false
	}

	return true
}
