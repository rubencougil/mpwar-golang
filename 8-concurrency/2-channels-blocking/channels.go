package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int)

	go func() {
		for {
			fmt.Printf("%d", <-c)
			time.Sleep(5 * time.Second)
		}
	}()

	c <- 1
	fmt.Printf("... waiting for channel to be free ...")
	c <- 2
}
