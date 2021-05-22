package main

import (
	"bufio"
	"fmt"
	"os/exec"
)

func main() {
	alice := exec.Command("python", "test.py")
	alice_stdout, _ := alice.StdoutPipe()
	alice.Start()
	r := bufio.NewReader(alice_stdout)
	for {
		line, _, _ := r.ReadLine()
		fmt.Printf(string(line))
	}
}
