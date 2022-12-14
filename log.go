package wormhole

import "log"

func Log(format string, v ...any) {
	log.Printf(format+"\n", v...)
}
