package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Bye bye!")
		os.Exit(0)
	}()

	fmt.Println("My name is Bob.")
	reader := bufio.NewReader(os.Stdin)
	for {
		msg, _ := reader.ReadString('\n')
		fmt.Println("I heard", msg)
	}
}
