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
			xAccRaw := mergeBytes(buf[3], buf[4])
			yAccRaw := mergeBytes(buf[5], buf[6])
			zAccRaw := mergeBytes(buf[7], buf[8])
			xGyroRaw := mergeBytes(buf[9], buf[10])
			yGyroRaw := mergeBytes(buf[11], buf[12])
			zGyroRaw := mergeBytes(buf[13], buf[14])

			// xAcc := float64(xAccRaw-0) / 16384.0
			// yAcc := float64(yAccRaw-0) / 16384.0
			// zAcc := float64(zAccRaw-0) / 16384.0
			// xGyro := float64(xGyroRaw+42) / 16.4
			// yGyro := float64(yGyroRaw-9) / 16.4
			// zGyro := float64(zGyroRaw+29) / 16.4

			xAcc := float64(xAccRaw) * 9.8 * 2 / 32768
			yAcc := float64(yAccRaw) * 9.8 * 2 / 32768
			zAcc := float64(zAccRaw) * 9.8 * 2 / 32768
			xGyro := float64(xGyroRaw) * (3.14159 / 180) / 1000 / 32768
			yGyro := float64(yGyroRaw) * (3.14159 / 180) / 1000 / 32768
			zGyro := float64(zGyroRaw) * (3.14159 / 180) / 1000 / 32768

			log.Printf("LD-   %5.1f %5.1f %5.1f   %5.1f %5.1f %5.1f   S\n", xGyro, yGyro, zGyro, xAcc, yAcc, zAcc)
		}
	}
}

// mergeBytes merges two uint8s to one uint16.
func mergeBytes(left8 byte, right8 byte) uint16 {
	return (uint16(left8) << 8) | uint16(right8)
}
