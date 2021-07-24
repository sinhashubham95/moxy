package main

import (
	"github.com/sinhashubham95/bleep"
	"os"
	"syscall"

	"github.com/sinhashubham95/moxy/persistence"
)

func init() {
	listener := bleep.New()
	listener.Add(func(os.Signal) {
		persistence.Close()
	})
	go listener.Listen(syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
}
