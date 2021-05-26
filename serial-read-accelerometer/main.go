package main

import (
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
			xGyro := mergeBytes(buf[3], buf[4])
			yGyro := mergeBytes(buf[5], buf[6])
			zGyro := mergeBytes(buf[7], buf[8])
			xAcc := mergeBytes(buf[9], buf[10])
			yAcc := mergeBytes(buf[11], buf[12])
			zAcc := mergeBytes(buf[13], buf[14])

			log.Printf("xg: %5d, yg: %5d: zg: %5d, xa: %5d, ya: %5d, za: %5d\n", xGyro, yGyro, zGyro, xAcc, yAcc, zAcc)
		}
	}
}

/// mergeBytes merges two uint8s to one uint16.
func mergeBytes(left8 byte, right8 byte) uint16 {
	return uint16(left8) + uint16(right8)<<8
}
