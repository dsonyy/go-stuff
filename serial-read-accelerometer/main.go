package main

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/tarm/serial"
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
	// Logger
	log.SetFlags(0)

	// Setting up the port
	config := &serial.Config{Name: os.Args[1], Baud: 9600}
	port, err := serial.OpenPort(config)
	if err != nil {
		log.Println(err)
		return
	}
	defer port.Close()

	// Loop
	data := make([]byte, 17)
	t := time.Now()
	for {
		if err := readFrame(port, data, 'L'); err != nil {
			log.Printf("error: %s\n", err)
		}

		if data[0] == 'L' && data[1] == 'D' && data[2] == 12 && data[15] == '#' {
			xAccRaw := mergeBytes(data[3], data[4]) + xAccOffset
			yAccRaw := mergeBytes(data[5], data[6]) + yAccOffset
			zAccRaw := mergeBytes(data[7], data[8]) + zAccOffset
			xGyroRaw := mergeBytes(data[9], data[10]) + xGyroOffset
			yGyroRaw := mergeBytes(data[11], data[12]) + yGyroOffset
			zGyroRaw := mergeBytes(data[13], data[14]) + zGyroOffset
			timeDiff := time.Now().Sub(t)
			t = time.Now()

			log.Printf("OK %5d %5d %5d   %5d %5d %5d   %d\n", xAccRaw, yAccRaw, zAccRaw, xGyroRaw, yGyroRaw, zGyroRaw, timeDiff.Nanoseconds())
		} else {
			log.Println("BAD FRAME")
			continue
		}
	}
}

func readFrame(port *serial.Port, data []byte, header rune) (err error) {
	scan := false
	for i := 0; i < 17; i++ {
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
