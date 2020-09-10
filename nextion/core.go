package nextion

import (
	"bytes"
	"log"

	"go.bug.st/serial"
)

type DisplayNextion struct {
	Port       string
	Baud       int
	Input      chan string
	Output     chan string
	serialPort serial.Port
	err        error
}

func (d *DisplayNextion) Start() {
	mode := &serial.Mode{
		BaudRate: d.Baud,
	}

	d.Input = make(chan string)
	d.Output = make(chan string)

	d.serialPort, d.err = serial.Open(d.Port, mode)
	if d.err != nil {
		log.Fatal(d.err)
	}
	defer d.serialPort.Close()

	buf := make([]byte, 1)
	var result []byte

	go d.SendStringToNextion()

	for {
		_, err := d.serialPort.Read(buf)
		if err != nil {
			log.Print(err)
		}

		result = append(result, buf[0])

		if i := bytes.Index(result, END); i >= 0 {
			//log.Printf("Response hex: %v\n", CheckReturnedCode(result[0:i]))
			d.Output <- CheckReturnedCode(result[0:i])
			result = nil
		}
	}
}

func (d *DisplayNextion) SendStringToNextion() {
	for {
		select {
		case s := <-d.Input:
			comm := stringToHexBytes(s)
			log.Printf("Send command: %s\n", comm[:len(comm)-3])
			_, err := d.serialPort.Write(comm)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
