package main

import "fmt"

type Robot interface {
	PowerOn() error
}

type Terminator struct {}

func (t Terminator) PowerOn() error {
	return nil
}

func sayHello(r Robot) {
	r.PowerOn()
	fmt.Println("Hello!")
}

func main() {
	t1000 := &Terminator{}
	sayHello(t1000)
}


