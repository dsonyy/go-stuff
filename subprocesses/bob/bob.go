package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	stdout := bufio.NewWriter(os.Stdout)
	stdin := bufio.NewReader(os.Stdin)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		stdout.WriteString("Bye bye!\n")
		stdout.Flush()
		os.Exit(0)
	}()

	stdout.WriteString("My name is Bob.\n")
	stdout.Flush()
	for {
		// msg, _ := stdin.ReadByte()
		msg, _ := stdin.ReadString('\n')
		stdout.WriteString(fmt.Sprintf("I heard: %s", msg))
		stdout.Flush()
	}
}
