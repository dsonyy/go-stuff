package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func handle_interrupt() {
	fmt.Println("Bye bye!")
}

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		handle_interrupt()
		os.Exit(0)
	}()

	fmt.Println("My name is Bob.")
	reader := bufio.NewReader(os.Stdin)
	for {
		msg, _ := reader.ReadString(' ')
		fmt.Println("I heard", msg)
	}
}
