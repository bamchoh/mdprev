package main

import (
	"flag"
)

type config struct {
	port      string
	storePath string
}

func parseArgs() config {
	c := config{}
	flag.StringVar(&c.port, "port", "3000", "open port number")
	flag.StringVar(&c.storePath, "path", findUserConfigPath(), "json file store path")
	flag.Parse()

	return c
}
