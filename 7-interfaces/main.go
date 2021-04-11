package main

import "fmt"

type RobotInterface interface {
	PowerOn() error
}

type Robot struct{}

func (t Robot) PowerOn() error {
	return nil
}

func main() {
	t := &Robot{}
	sayHello(t)
}

func sayHello(r RobotInterface) {
	_ = r.PowerOn()
	fmt.Println("Hello!")
}
