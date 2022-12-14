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

func (c *capture) init(uri string) (err error) {
	path := "capture"
	name := ""
	index := strings.LastIndex(uri, "/")
	if index == -1 {
		path += "/"
	} else {
		path += uri[:index+1]
		name = uri[index+1:]
	}
	if name == "" {
		name = "_"
	}

	if err = os.MkdirAll(path, os.ModePerm); err != nil {
		return
	}

	if c.file, err = os.OpenFile(path+name, os.O_CREATE|os.O_RDWR, os.ModePerm); err != nil {
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
