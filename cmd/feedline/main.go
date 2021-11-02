package main

import (
	"feedline/internal/server"
	"feedline/pkg/receiver"
	"flag"
)

func main() {
	webDir := flag.String("webdir", "./web", "load templates from this directory")
	addr := flag.String("addr", ":8080", "listen to this address")
	flag.Parse()

	go func() {
		receiver.Refresh()
		receiver.TidyDismissed()
	}()
	go receiver.Listen()
	server.Serve(*webDir, *addr)
}
