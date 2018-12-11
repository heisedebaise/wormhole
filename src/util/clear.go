package util

import (
	"io/ioutil"
	"log"
	"os"
	"time"
)

type clear struct {
	Path     string
	Timeout  string
	nTimeout int64
}

var clears []clear

func init() {
	LoadConfig(&clears, "clear")
	for i, c := range clears {
		c.nTimeout = int64(ParseTime(c.Timeout))
		clears[i] = c
	}
	log.Printf("clear : %+v\n", clears)

	autoClear()
}

func autoClear() {
	go func() {
		for {
			for _, c := range clears {
				timeout := time.Now().Unix() - c.nTimeout
				scanToClear(c, c.Path, timeout)
			}

			time.Sleep(time.Minute)
		}
	}()
}

func scanToClear(c clear, path string, timeout int64) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return
	}

	if len(files) == 0 {
		os.Remove(path)

		return
	}

	for _, file := range files {
		p := path + string(os.PathSeparator) + file.Name()
		if file.IsDir() {
			scanToClear(c, p, timeout)

			continue
		}

		if file.ModTime().Unix() < timeout {
			os.Remove(p)
		}
	}
}
