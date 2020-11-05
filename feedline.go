package main

import (
	"feedline/lib/receiver"
	"feedline/lib/server"
)

func main() {
	go receiver.Refresh()
	go receiver.Listen()
	server.Serve("localhost:8579")
}
