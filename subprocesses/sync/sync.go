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

	pyPath := filepath.Join(cwd, "test.py")
	fmt.Println("test.py path:", pyPath)

	// pyStdout, err := os.Create(filepath.Join(cwd, "stdout.txt"))
	// if err != nil {
	//  	panic(err)
	// }

	pyCmd := exec.Command("python", pyPath)
	pyCmd.Dir = cwd
	// pyCmd.Stdout = pyStdout
	pyStdout, err := pyCmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	if err = pyCmd.Start(); err != nil {
		panic(err)
	}
	fmt.Println("Foo")

	scanner := bufio.NewScanner(pyStdout)
	fmt.Println("Foo")
	for scanner.Scan() {
		fmt.Println("Foo")

		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return

	// cmd := exec.Command("python", "test.py")
	// cwd, err := filepath.Abs()
	// cmd.Dir =
	// 	cmd.Start()

	// alice := exec.Command("python", "test.py")
	// alice_stdout, _ := alice.StdoutPipe()
	// alice.Start()
	// r := bufio.NewReader(alice_stdout)
	// for {
	// 	line, _, _ := r.ReadLine()
	// 	fmt.Printf(string(line))
	// }
}
