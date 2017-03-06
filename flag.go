package main

import (
	"flag"
)

type config struct {
	port string
}

func parseArgs() config {
	c := config{}
	flag.StringVar(&c.port, "port", "3000", "open port number")
	flag.Parse()
	return c
}
