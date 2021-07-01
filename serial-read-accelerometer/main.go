package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/tarm/serial"
)

// Calibration offsets for: ACCEL_FS_2, MPU6050_GYRO_FS_250
const (
	xAccOffset  = 812.0
	yAccOffset  = 118.0
	zAccOffset  = -14750.0 + 16384.0
	xGyroOffset = 55.0
	yGyroOffset = -56.0
	zGyroOffset = 39.0
)

func main() {
	// Logger
	log.SetFlags(0)

	// Setting up the port
	config := &serial.Config{Name: os.Args[1], Baud: 19200}
	port, err := serial.OpenPort(config)
	if err != nil {
		log.Println(err)
		return
	}
	defer port.Close()

	// Loop
	data := make([]byte, 18)
	t := time.Now()
	for {
		if err := readFrame(port, data, 'L'); err != nil {
			log.Printf("error: %s\n", err)
		}
		log.Println(data)
		if data[0] == 'L' && data[1] == 'D' && data[2] == 12 && data[3] == '+' && data[16] == '#' {
			xAccRaw := mergeBytes(data[4], data[5]) + xAccOffset
			yAccRaw := mergeBytes(data[6], data[7]) + yAccOffset
			zAccRaw := mergeBytes(data[8], data[9]) + zAccOffset
			xGyroRaw := mergeBytes(data[10], data[11]) + xGyroOffset
			yGyroRaw := mergeBytes(data[12], data[13]) + yGyroOffset
			zGyroRaw := mergeBytes(data[14], data[15]) + zGyroOffset
			timeDiff := time.Now().Sub(t)
			t = time.Now()

			log.Printf("OK %5d %5d %5d   %5d %5d %5d   %d\n", xAccRaw, yAccRaw, zAccRaw, xGyroRaw, yGyroRaw, zGyroRaw, timeDiff.Nanoseconds())
			fmt.Printf("%5d %5d %5d   %5d %5d %5d\n", xAccRaw, yAccRaw, zAccRaw, xGyroRaw, yGyroRaw, zGyroRaw)
		} else {
			log.Println("BAD FRAME")
			continue
		}
	}
}

func readFrame(port *serial.Port, data []byte, header rune) (err error) {
	scan := false
	for i := 0; i < 18; i++ {
		buf := make([]byte, 1)
		_, err := port.Read(buf)
		if err != nil {
			return errors.New("cannot read from port")
		}

		if scan {
			data[i] = buf[0]
		} else {
			if buf[0] == byte(header) {
				scan = true
				data[i] = buf[0]
			} else {
				return errors.New("lost some data")
			}
		}
	}
	return nil
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
