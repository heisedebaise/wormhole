package rsync

import (
	"io/ioutil"
	"memory"
	"net"
	"os"
	"path/filepath"
	"strings"
)

func write(conn net.Conn, bytes []byte) error {
	nSize := len(bytes) + 8
	bSize := make([]byte, 8)
	for i := uint(0); i < 8; i++ {
		bSize[i] = byte(nSize>>((7-i)<<3)) & 0xff
	}
	if _, err := conn.Write(bSize); err != nil {
		return err
	}

	if _, err := conn.Write(bytes); err != nil {
		return err
	}

	return nil
}

func saveFile(message []byte) {
	length, uri := readLengghUnique(message)
	path, err := filepath.Abs(uri[1:])
	if err != nil {
		return
	}

	if err = os.MkdirAll(path[:strings.LastIndex(path, "/")], 0755); err != nil {
		return
	}

	ioutil.WriteFile(path, message[length:], 0755)
}

func putMemory(message []byte) {
	length, unique := readLengghUnique(message)
	memory.Put(unique, message[length:])
}

func readLengghUnique(message []byte) (int, string) {
	length := (int(message[1]) << 8) + int(message[2]) + 3
	unique := string(message[3:length])

	return length, unique
}
