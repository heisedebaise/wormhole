package util

import (
	"log"
	"os"
	"time"
)

var file *os.File

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	os.MkdirAll("logs", os.ModePerm)
	logFile()
	timer()
}

func timer() {
	go func() {
		for {
			timer := time.NewTimer(time.Second)
			<-timer.C
			now := time.Now()
			if now.Hour() == 0 && now.Minute() == 0 && now.Second() == 0 {
				logFile()
			}
			timer.Stop()
		}
	}()
}

func logFile() {
	if f, err := os.OpenFile("logs/"+time.Now().Format("2006-01-02"), os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666); err == nil {
		if file != nil {
			file.Close()
		}
		file = f
		log.SetOutput(file)
	}
}
