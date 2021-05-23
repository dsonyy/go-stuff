package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Bye bye!")
		os.Exit(0)
	}()

	fmt.Println("I'm Alice.")
	for {
		time.Sleep(time.Duration(rand.Intn(2000)+100) * time.Millisecond)
		fmt.Println("Hello?")
		time.Sleep(time.Duration(rand.Intn(2000)+100) * time.Millisecond)
		fmt.Fprintf(os.Stderr, "Hi?")
	}
}
