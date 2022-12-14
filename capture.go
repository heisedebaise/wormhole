package wormhole

import (
	"bytes"
	"compress/gzip"
	"io"
	"os"
	"strings"
)

type capture struct {
	reader io.Reader
	file   *os.File
	buffer *bytes.Buffer
	gzip   bool
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

	if c.gzip {
		c.buffer = &bytes.Buffer{}
	}

	return
}

func (c *capture) Read(p []byte) (n int, err error) {
	n, err = c.reader.Read(p)
	if c.gzip {
		c.buffer.Write(p[:n])
	} else {
		c.file.Write(p[:n])
	}

	return
}

func (c *capture) close() {
	if c.file == nil {
		return
	}
	defer c.file.Close()

	if c.gzip {
		reader, err := gzip.NewReader(c.buffer)
		if err != nil {
			Log("read capture from gzip err %v", err)

			return
		}
		defer reader.Close()

		io.Copy(c.file, reader)
	}
}
