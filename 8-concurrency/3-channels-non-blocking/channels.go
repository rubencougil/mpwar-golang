package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int)

	go func(){
		for {
			select {
				case num := <-c:
					fmt.Printf("%d\n", num)
				default:
					time.Sleep(time.Second)
					fmt.Printf(".")
			}
		}
	}()

	for i:=0; i<10; i++  {
		c<-i
		time.Sleep(5 * time.Second)
	}
}

