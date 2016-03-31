package main

import (
	"flag"
	"fmt"

	"github.com/asciimoo/privacyscore/checker"
	"github.com/asciimoo/privacyscore/server"
)

var listen = flag.String("listen", "127.0.0.1:1080", "server listen address")

func main() {

	flag.Parse()
	if len(flag.Args()) == 0 {
		server.Run(listen)
	} else {
		c, err := checker.Run(flag.Args()[0])
		if err == nil {
			fmt.Println(c.Result.Penalties.GetScore())
		} else {
			fmt.Println("Something went wrong:", err)
		}
	}
}
