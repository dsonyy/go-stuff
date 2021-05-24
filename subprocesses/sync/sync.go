package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

func readPipe(prefix string, stream *io.ReadCloser) {
	scanner := bufio.NewScanner(*stream)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(prefix, line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func readWritePipe(prefix string, streamSrc *io.ReadCloser, streamTgt *io.WriteCloser) {
	scanner := bufio.NewScanner(*streamSrc)
	writer := bufio.NewWriter(*streamTgt)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(prefix, line)
		if _, err := writer.WriteString(line + "\n"); err != nil {
			panic(err)
		}
		writer.Flush()
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func main() {
	// Get a path to the running executable
	exe, err := os.Executable()
	if err != nil {
		panic(err)
	}
	cwd := filepath.Dir(exe)
	fmt.Println("Executable path:", cwd)

	// Create paths to alice.exe and bob.exe
	alicePath := filepath.Join(cwd, "../alice/alice.exe")
	bobPath := filepath.Join(cwd, "../bob/bob.exe")
	fmt.Println("alice.go path:", alicePath)
	fmt.Println("alice.go path:", bobPath)

	// Create Alice subprocess
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

	// Create Bob subprocess
	bob := exec.Command(bobPath)
	bobStdout, err := bob.StdoutPipe()
	if err != nil {
		panic(err)
	}
	bobStdin, err := bob.StdinPipe()
	if err != nil {
		panic(err)
	}

	// Start Alice and Bob
	if err := alice.Start(); err != nil {
		panic(err)
	}
	if err := bob.Start(); err != nil {
		panic(err)
	}

	// Start goroutines
	// go readPipe("Alice2:", &aliceStderr)                    // Read and print Alice's stderr
	// go readPipe("Alice1:", &aliceStdout)                    // Read and print Alice's stdout
	go readPipe("Bob1:", &bobStdout)                           // Read and print Bob's stdout
	go readWritePipe("Alice1->Bob0:", &aliceStdout, &bobStdin) // Read and print Alice's stdout and write it to Bob's stdin
	go readWritePipe("Alice2->Bob0:", &aliceStderr, &bobStdin) // Read and print Alice's stderr and write it to Bob's stdin

	// Wait for Alice and Bob (it never happens)
	alice.Wait()
	bob.Wait()
}
