package rsync

import (
	"net"
	"util"
)

func write(conn net.Conn, bytes []byte) error {
	if _, err := conn.Write(util.Uint32ToBytes(uint32(len(bytes)))); err != nil {
		return err
	}

	if _, err := conn.Write(bytes); err != nil {
		return err
	}

	return nil
}
