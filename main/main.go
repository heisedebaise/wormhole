package main

import (
	"flag"
	"wormhole"
)

func main() {
	var c = flag.String("c", "map.json", "config file")
	flag.Parse()

	wormhole.Serve(*c)
}
