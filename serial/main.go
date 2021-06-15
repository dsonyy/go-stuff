package main

import (
	"fmt"
	"log"

	"github.com/jacobsa/go-serial/serial"
)

func main() {
	// Setting up the port
	options := serial.OpenOptions{
		PortName:        "COM8",
		BaudRate:        9600,
		DataBits:        8,
		MinimumReadSize: 4,
	}
	port, err := serial.Open(options)
	defer port.Close()
	if err != nil {
		log.Println(err)
		return
	}

	// Loop
	for {
		var line string
		_, err := fmt.Scanln(&line)
		if err != nil {
			log.Println(err)
			break
		}

		n, err := port.Write([]byte(line))
		if err != nil {
			log.Println(err)
			break
		}
		log.Println("Wrote", n, "bytes:", []byte(line))

		buf := make([]byte, 128)
		n, err = port.Read(buf)
		if err != nil {
			log.Println(err)
			break
		}
		log.Println("Received:", buf)
	}
}
