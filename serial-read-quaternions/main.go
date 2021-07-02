package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"math"
	"os"

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
	data := make([]byte, 22)
	for {
		if err := readFrame(port, data, 'L'); err != nil {
			log.Printf("error: %s\n", err)
		}
		log.Println(data)
		if data[0] == 'L' && data[1] == 'Q' && data[2] == 16 && data[3] == '+' && data[20] == '#' {
			qw := math.Acos(float64(Float32frombytes(data[4:8]))) * 2 * 57.2957795
			qx := Float32frombytes(data[8:12])
			qy := Float32frombytes(data[12:16])
			qz := Float32frombytes(data[16:20])

			fmt.Printf("%5f %5f %5f %5f\n", qw, qx, qy, qz)
		} else {
			log.Println("BAD FRAME")
			continue
		}
	}
}

func readFrame(port *serial.Port, data []byte, header rune) (err error) {
	scan := false
	for i := 0; i < 22; i++ {
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

// convert 4 bytes to float32
func Float32frombytes(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)
	float := math.Float32frombits(bits)
	return float
}
