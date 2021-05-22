package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	fmt.Println("I'm Alice?")
	for {
		time.Sleep(time.Duration(rand.Intn(2000)+100) * time.Millisecond)
		fmt.Println("Hello?")
		time.Sleep(time.Duration(rand.Intn(2000)+100) * time.Millisecond)
		fmt.Fprintf(os.Stderr, "Hi?")
	}
}
