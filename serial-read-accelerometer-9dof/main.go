package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"time"

	"github.com/tarm/serial"
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
	data := make([]byte, 42)
	t := time.Now()
	for {
		if err := readFrame(port, data, 'L'); err != nil {
			log.Printf("error: %s\n", err)
		}
		log.Println(data)
		if data[0] == 'L' && data[1] == 'F' && data[2] == 36 && data[3] == '+' && data[40] == '#' {
			xAccRaw := Float32frombytes(data[4:8])
			yAccRaw := Float32frombytes(data[8:12])
			zAccRaw := Float32frombytes(data[12:16])
			xGyroRaw := Float32frombytes(data[16:20])
			yGyroRaw := Float32frombytes(data[20:24])
			zGyroRaw := Float32frombytes(data[24:28])
			xMagRaw := Float32frombytes(data[28:32])
			yMagRaw := Float32frombytes(data[32:36])
			zMagRaw := Float32frombytes(data[36:40])
			timeDiff := time.Now().Sub(t)
			t = time.Now()

			log.Printf("OK %f %f %f   %f %f %f   %f %f %f   %d\n",
				xAccRaw, yAccRaw, zAccRaw,
				xGyroRaw, yGyroRaw, zGyroRaw,
				xMagRaw, yMagRaw, zMagRaw, timeDiff.Nanoseconds())
			fmt.Printf("%f %f %f   %f %f %f   %f %f %f\n", xAccRaw, yAccRaw, zAccRaw, xGyroRaw, yGyroRaw, zGyroRaw, xMagRaw, yMagRaw, zMagRaw)
		} else {
			log.Println("BAD FRAME")
			continue
		}
	}
}

func readFrame(port *serial.Port, data []byte, header rune) (err error) {
	scan := false
	for i := 0; i < 42; i++ {
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

// convert 4 bytes to float32
func Float32frombytes(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)
	float := math.Float32frombits(bits)
	return float
}
