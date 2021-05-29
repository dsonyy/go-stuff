package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jacobsa/go-serial/serial"
)

// Calibration offsets for: ACCEL_FS_2, MPU6050_GYRO_FS_250
const (
	xAccOffset  = -844
	yAccOffset  = 78
	zAccOffset  = 1542
	xGyroOffset = 244
	yGyroOffset = -228
	zGyroOffset = 161
)

func main() {
	// Setting up the port
	options := serial.OpenOptions{
		PortName:        os.Args[1],
		BaudRate:        9600,
		DataBits:        8,
		MinimumReadSize: 1,
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
			xAccRaw := mergeBytes(buf[3], buf[4]) + xAccOffset
			yAccRaw := mergeBytes(buf[5], buf[6]) + yAccOffset
			zAccRaw := mergeBytes(buf[7], buf[8]) + zAccOffset
			xGyroRaw := mergeBytes(buf[9], buf[10]) + xGyroOffset
			yGyroRaw := mergeBytes(buf[11], buf[12]) + yGyroOffset
			zGyroRaw := mergeBytes(buf[13], buf[14]) + zGyroOffset

			fmt.Printf("%5d %5d %5d   %5d %5d %5d \n", xAccRaw, yAccRaw, zAccRaw, xGyroRaw, yGyroRaw, zGyroRaw)
		}
	}
}

// mergeBytes merges two uint8s to one uint16.
func mergeBytes(left8 byte, right8 byte) (v int) {
	// awesome conversion to signed int
	v = int((uint16(left8) << 8) | uint16(right8))
	if v >= 32768 {
		v = -65536 + v
	}
	return
}
