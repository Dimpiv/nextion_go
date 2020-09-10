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
		Port: "COM3",
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
				fmt.Printf("Response from Nextion: %s\n", r)
			}
		}
	}()

	dp.Start()
}
