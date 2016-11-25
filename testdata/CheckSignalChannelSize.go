package main

import (
	"os"
	"os/signal"
	"syscall"
)

func fn() {
	c0 := make(chan os.Signal)
	signal.Notify(c0, os.Interrupt) // MATCH /channel buffer size 0 is too small to catch 1 signal/

	c1 := make(chan os.Signal, 1)
	signal.Notify(c1, os.Interrupt, syscall.SIGHUP) // MATCH /channel buffer size 1 is too small to catch 2 signal/

	c2 := make(chan os.Signal, 1)
	signal.Notify(c2, os.Interrupt)
	signal.Notify(c2, syscall.SIGHUP) // MATCH /channel buffer size 1 is too small to catch 2 signal/
}
