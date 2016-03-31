package main

import (
	"flag"

	"github.com/asciimoo/privacyscore/server"
)

var listen = flag.String("listen", "127.0.0.1:1080", "server listen address")

func main() {

	flag.Parse()
	if len(flag.Args()) == 0 {
		server.Run(listen)
	}
}
