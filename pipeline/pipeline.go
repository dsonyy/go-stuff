package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	pipeString = "->"
)

// readPipe reads data from streamSrc and prints it on this program's stderr.
func readPipe(streamSrc *io.ReadCloser) {
	scanner := bufio.NewScanner(*streamSrc)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		log.Println(line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

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

func main() {
	log.SetFlags(0)
	cmds := []exec.Cmd{}

	os.Args = []string{"", "python", "->"}

	// Get a path to the running executable
	exe, err := os.Executable()
	if err != nil {
		panic(err)
	}
	cwd := filepath.Dir(exe)

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
	cmds[0].Stdin = os.Stdin
	cmds[len(cmds)-1].Stdout = os.Stdout
	cmds[len(cmds)-1].Stderr = os.Stderr
	cmds[0].Dir = cwd

	cmds[0].Start()

	cmds[0].Wait()

	// // Get stdout and stdin of each process
	// stdouts := []io.ReadCloser{}
	// stdins := []io.WriteCloser{}
	// for _, cmd := range cmds {
	// 	stdout, err := cmd.StdoutPipe()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	stdin, err := cmd.StdinPipe()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	stdouts = append(stdouts, stdout)
	// 	stdins = append(stdins, stdin)
	// }

	// // Start reading/writing goroutines
	// // stdins[0] = os.Stdin
	// for i := range cmds[:len(cmds)-1] {
	// 	go readWritePipe(&stdouts[i], &stdins[i+1])
	// }
	// cmds[len(cmds)-1].Stdout = os.Stdout

	// // Start every process
	// for _, cmd := range cmds {
	// 	if err := cmd.Start(); err != nil {
	// 		panic(err)
	// 	}
	// }

	// for _, cmd := range cmds {
	// 	cmd.Wait()
	// }
}
