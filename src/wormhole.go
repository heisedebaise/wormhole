package main

import (
	"auth"
	"fileserv"
	"httpserv"
	"imgserv"
	"rsync"
	"speech"
	"wserv"
)

func main() {
	rsync.Listen()

	wserv.Serve()
	auth.Serve()
	imgserv.Serve()
	fileserv.Serve()
	speech.Serve()
	httpserv.HTTP("/")
}
