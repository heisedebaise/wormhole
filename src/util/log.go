package util

import (
	"log"
	"os"
	"time"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	os.MkdirAll("logs", os.ModePerm)
	if file, err := os.OpenFile("logs/"+time.Now().Format("20060102"), os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666); err == nil {
		log.SetOutput(file)
	}
}
