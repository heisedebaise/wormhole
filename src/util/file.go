package util

import (
	"log"
	"os"
	"strconv"
)

func Exists(path string) bool {
	_, err := os.Stat(path)

	return err == nil
}

func ExistsFile(path string) bool {
	info, err := os.Stat(path)

	return err == nil && !info.IsDir()
}

func ByteSize(size string) int64 {
	length := len(size)
	if length == 0 {
		return 0
	}

	suffix := size[length-1:]
	number, err := strconv.ParseInt(size[0:length-1], 10, 64)
	if err != nil {
		log.Printf("failure to convert %s to byte size %q\n", size, err)

		return 0
	}

	if suffix == "K" || suffix == "k" {
		return number << 10
	}

	if suffix == "M" || suffix == "m" {
		return number << 20
	}

	if suffix == "G" || suffix == "g" {
		return number << 30
	}

	if suffix == "T" || suffix == "t" {
		return number << 40
	}

	number, err = strconv.ParseInt(size, 10, 64)
	if err != nil {
		log.Printf("failure to convert %s to byte size %q\n", size, err)

		return 0
	}

	return number
}
