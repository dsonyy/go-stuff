package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

const (
	pipeString = "->"
)

// readWritePipe reads data from streamSrc and passes it to streamTgt.
func readWritePipe(streamSrc *io.ReadCloser, streamTgt *io.WriteCloser) {
	scanner := bufio.NewScanner(*streamSrc)
	writer := bufio.NewWriter(*streamTgt)

	for scanner.Scan() {
		line := scanner.Text()
		if _, err := writer.WriteString(line + "\n"); err != nil {
			panic(err)
		}
		writer.Flush()
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

// Help prints usage of this program.
func help() {

}

func main() {
	log.SetFlags(0)
	cmds := []exec.Cmd{}

	// Handling Ctrl+C - killing every process
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		for _, cmd := range cmds {
			if err := cmd.Process.Kill(); err != nil {
				panic(err)
			}
		}
		os.Exit(0)
	}()

	// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// Process command line arguments
	cmdPath := ""
	cmdArgs := []string{}
	for _, word := range os.Args[1:] {
		if word == pipeString {
			cmds = append(cmds, *exec.Command(cmdPath, cmdArgs...))
			cmdPath = ""
			cmdArgs = []string{}
		} else if cmdPath == "" {
			cmdPath = word
		} else {
			cmdArgs = append(cmdArgs, word)
		}
	}
	if cmdPath != "" {
		cmds = append(cmds, *exec.Command(cmdPath, cmdArgs...))
	}
	if len(cmds) == 0 {
		help()
		return
	}

	// Get stdout and stdin of each process
	stdouts := []*io.ReadCloser{}
	stdins := []*io.WriteCloser{}
	for i := range cmds {
		stdout, err := cmds[i].StdoutPipe()
		if err != nil {
			panic(err)
		}
		stdin, err := cmds[i].StdinPipe()
		if err != nil {
			panic(err)
		}
		stdouts = append(stdouts, &stdout)
		stdins = append(stdins, &stdin)
	}

	// Redirect stdout of the last process
	cmds[len(cmds)-1].Stdout = os.Stdout

	// Start reading/writing goroutines
	for i := range cmds[:len(cmds)-1] {
		go readWritePipe(stdouts[i], stdins[i+1])
	}

	// Prepare and start each process
	for i := range cmds {
		cmds[i].Stderr = os.Stderr
		cmds[i].Dir = cwd
		cmds[i].Start()
	}

	// Wait for each process
	for i := range cmds {
		cmds[i].Wait()
	}

}
