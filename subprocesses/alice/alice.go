package main

import (
	"bufio"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	stdout, stderr := bufio.NewWriter(os.Stdout), bufio.NewWriter(os.Stderr)
	stdout.Buffered()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		stdout.WriteString("Hello?\n")
		stdout.Flush()
		os.Exit(0)
	}()

	stdout.WriteString("My name is alice?\n")
	stdout.Flush()
	for {
		time.Sleep(1 * time.Second)
		stdout.WriteString("Hello?\n")
		stdout.Flush()
		stderr.WriteString("Hi?\n")
		stderr.Flush()
	}
}
