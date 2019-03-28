package main

import (
	"auth"
	"fileserv"
	"httpserv"
	"imgserv"
	"rsync"
	"speech"
	"temporary"
	"wserv"
)

func main() {
	rsync.Listen()

	wserv.Serve()
	auth.Serve()
	imgserv.Serve()
	fileserv.Serve()
	temporary.Serve()
	speech.Serve()
	httpserv.HTTP("/")
}
