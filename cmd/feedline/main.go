package main

import (
	"feedline/internal/server"
	"feedline/pkg/receiver"
	"flag"
	"time"
)

func main() {
	webDir := flag.String("webdir", "./web", "load templates from this directory")
	addr := flag.String("addr", ":8080", "listen to this address")
	interval := flag.Duration("interval", time.Duration(10 * float64(time.Minute)), "time between refreshes")
	flag.Parse()

	go func() {
		receiver.Refresh()
	}()
	go receiver.Listen(*interval)
	server.Serve(*webDir, *addr)
}
