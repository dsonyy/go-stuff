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
		MinimumReadSize: 17,
	}
	port, err := serial.Open(options)
	defer port.Close()
	if err != nil {
		log.Println(err)
		return
	}

	// Loop
	for {
		buf := make([]byte, 17)
		_, err = port.Read(buf)
		if err != nil {
			log.Println(err)
			break
		}

		if buf[0] == 'L' && buf[1] == 'D' && buf[2] == '-' {
			xAccRaw := mergeBytes(buf[3], buf[4])
			yAccRaw := mergeBytes(buf[5], buf[6])
			zAccRaw := mergeBytes(buf[7], buf[8])
			xGyroRaw := mergeBytes(buf[9], buf[10])
			yGyroRaw := mergeBytes(buf[11], buf[12])
			zGyroRaw := mergeBytes(buf[13], buf[14])

			fmt.Printf("%5d %5d %5d   %5d %5d %5d \n", xAccRaw, yAccRaw, zAccRaw, xGyroRaw, yGyroRaw, zGyroRaw)
		}
	}
}

// mergeBytes merges two uint8s to one uint16.
func mergeBytes(left8 byte, right8 byte) (v int) {
	v = int((uint16(left8) << 8) | uint16(right8))
	if v >= 32768 {
		v = -65536 + v
	}
	return
}
