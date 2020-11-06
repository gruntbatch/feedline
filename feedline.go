package main

import (
	"feedline/lib/receiver"
	"feedline/lib/server"
	"flag"
)

func main() {
	templateDir := flag.String("templatedir", "./", "load templates from this directory")
	addr := flag.String("addr", ":8080", "listen to this address")
	flag.Parse()

	go receiver.Refresh()
	go receiver.Listen()
	server.Serve(*templateDir, *addr)
}
