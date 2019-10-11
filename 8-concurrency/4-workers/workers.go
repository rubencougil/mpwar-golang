package main

import (
	"fmt"
	"time"
)

func Worker(worker int, c <-chan int) {
	for num := range c {
		fmt.Printf("Worker %d: %d\n", worker, num)
		time.Sleep(time.Second)
	}
}

func main() {

	c := make(chan int)

	for i:=0; i<3; i++  {
		go Worker(i,c)
	}

	for j:=0; j<100 ;j++  {
		c<-j
	}
}
