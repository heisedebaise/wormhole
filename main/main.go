package main

import (
	"flag"
	"wormhole"
)

func main() {
	var c = flag.String("c", "config.json", "config file")
	flag.Parse()

	wormhole.Serve(*c)
}
