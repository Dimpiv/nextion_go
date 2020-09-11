package main

import (
	"Nextion/nextion"
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	dp := nextion.DisplayNextion{
		Port: "/dev/ttyUSB0",
		Baud: 9600,
	}

	consoleInput := make(chan string)

	go func() {
		for {
			fmt.Print("Enter command: ")
			reader := bufio.NewReader(os.Stdin)
			t, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			consoleInput <- t[:len(t)-1]
		}
	}()

	go func() {
		for {
			select {
			case command := <-consoleInput:
				dp.Input <- command
			case r := <-dp.Output:
				fmt.Printf("\rResponse from Nextion: %s\n", r)
				fmt.Print("Enter command: ")
			}
		}
	}()

	dp.Start()
}
