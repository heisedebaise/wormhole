package wormhole

import (
	"io"
	"os"
	"strings"
)

type capture struct {
	reader io.Reader
	file   *os.File
}

func (c *capture) init(uri, contentType string) (err error) {
	if err = os.MkdirAll(contentType, os.ModePerm); err != nil {
		return
	}

	if c.file, err = os.OpenFile(contentType+uri[strings.LastIndex(uri, "/"):], os.O_CREATE|os.O_RDWR, os.ModePerm); err != nil {
		return
	}

	return
}

func (c *capture) Read(p []byte) (n int, err error) {
	n, err = c.reader.Read(p)
	c.file.Write(p[:n])

	return
}

func (c *capture) close() {
	if c.file != nil {
		c.file.Close()
	}
}
