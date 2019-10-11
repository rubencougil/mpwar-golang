package main

import (
	"fmt"
)

type Robot struct {
	Name   string
	Config *Config
}

type Config struct {
	Age int
}

func NewRobot(name string, config *Config) *Robot {
	return &Robot{Name: name, Config: config}
}

func NewRobotVariadic(name string, config ...*Config) *Robot {
	if len(config) == 0 {
		return &Robot{Name: name, Config: &Config{Age: 0}}
	}

	return &Robot{Name: name, Config: config[0]}
}

func main() {
	robotOne := NewRobot("one", nil)
	fmt.Printf("Robot name: %s \n", robotOne.Name)
	robotTwo := NewRobot("two", &Config{Age: 15})
	fmt.Printf("Robot name: %s, age: %d \n", robotTwo.Name, robotTwo.Config.Age)

	robotThree := NewRobotVariadic("three")
	fmt.Printf("Robot name: %s, age: %d \n", robotThree.Name, robotThree.Config.Age)
	robotFour := NewRobotVariadic("four", &Config{Age: 5})
	fmt.Printf("Robot name: %s, age: %d \n", robotFour.Name, robotFour.Config.Age)
}
