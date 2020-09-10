package main

import (
	"Nextion/nextion"
	"fmt"
	"time"
)

func main() {
	dp := nextion.DisplayNextion{
		Port: "COM3",
		Baud: 9600,
	}

	go func() {
		t1 := time.NewTicker(time.Second * 5)
		for {
			select {
			case <-t1.C:
				dp.Input <- "page page0"
			case r := <-dp.Output:
				fmt.Printf("Response from Nextion: %s\n", r)
			}

		}
	}()

	dp.Start()
}
