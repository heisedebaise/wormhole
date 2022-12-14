package wormhole

import (
	"bytes"
	"io"
)

func replace(reader io.Reader, m map[string]string) (io.Reader, int, error) {
	bs, err := io.ReadAll(reader)
	if err != nil {
		return nil, 0, err
	}

	for key := range m {
		bs = bytes.ReplaceAll(bs, []byte(key), []byte(m[key]))
	}

	return bytes.NewReader(bs), len(bs), nil
}
