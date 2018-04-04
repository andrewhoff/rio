package main

import (
	"flag"
	"fmt"

	"github.com/andrewhoff/rio/client"
	"github.com/andrewhoff/rio/server"
)

var c = flag.Bool("c", false, "-c for client mode")
var s = flag.Bool("s", false, "-s for server mode")

func run() int {
	flag.Parse()

	if *c {
		client.Run()
		return 0
	} else if *s {
		server.Run()
		return 0
	} else {
		fmt.Print("\x1b[31;1m")
		fmt.Println("[ERROR] Must specify mode")
	}

	return 1
}
