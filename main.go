package main

import (
	"flag"
	"wormhole/serve"
)

func main() {
	var c = flag.String("c", "map.json", "config file")
	flag.Parse()

	serve.Serve(*c)
}
