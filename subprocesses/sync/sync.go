package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	exe, err := os.Executable()
	if err != nil {
		panic(err)
	}
	cwd := filepath.Dir(exe)
	fmt.Println("Executable path:", cwd)

	alicePath := filepath.Join(cwd, "../alice/alice.exe")
	fmt.Println("alice.go path:", alicePath)

	alice := exec.Command(alicePath)
	alice.Dir = cwd
	aliceStdout, err := alice.StdoutPipe()
	if err != nil {
		panic(err)
	}
	aliceStderr, err := alice.StderrPipe()
	if err != nil {
		panic(err)
	}
	if err := alice.Start(); err != nil {
		panic(err)
	}

	go func() {
		scanner := bufio.NewScanner(aliceStderr)
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println("Alice2:", line)
		}
		if err := scanner.Err(); err != nil {
			panic(err)
		}
	}()

	go func() {
		scanner := bufio.NewScanner(aliceStdout)
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println("Alice1:", line)
		}
		if err := scanner.Err(); err != nil {
			panic(err)
		}
	}()

	alice.Wait()
}
