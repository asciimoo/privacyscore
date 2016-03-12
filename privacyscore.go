package main

import (
	"flag"

	"github.com/asciimoo/privacyscore/server"
)

var listen = flag.String("listen", "localhost:1080", "listen on address")

func main() {
	server.Run(listen)
}
