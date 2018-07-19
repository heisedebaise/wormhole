package rsync

import "net"

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
