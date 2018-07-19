package rsync

import (
	"log"
	"net"
)

func read(conn net.Conn, callback func(message []byte) bool) {
	var messages []byte
	var message []byte
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			log.Println(err)

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
	length := len(messages)
	if length < 8 {
		return nil, messages
	}

	size := 0
	for i := 0; i < 8; i++ {
		size = (size << 8) + int(messages[i])
	}

	if length < size {
		return nil, messages
	}

	return messages[8:size], messages[size:]
}
