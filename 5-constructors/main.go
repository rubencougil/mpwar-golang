package main

import (
	"fmt"
	"github.com/pkg/errors"
)

type Robot struct {
	name   string
	config *Config
}

func (r *Robot) Name() string {
	return r.name
}

func (r *Robot) Config() *Config {
	if r.config == nil {
		return &Config{age: -1}
	}
	return r.config
}

type Config struct {
	age int
}

func NewRobot(name string, config *Config) *Robot {
	return &Robot{name: name, config: config}
}

func NewRobotVariadic(name string, config ...*Config) (*Robot, error) {
	if len(config) > 1 {
		return nil, errors.New("too many parameters")
	}
	if len(config) == 0 {
		return &Robot{name: name}, nil
	}

	return &Robot{name: name, config: config[0]}, nil
}

func NewRobotFuncOptions(name string, config ...func(*Robot)) (*Robot, error) {
	robot := &Robot{name: name, config: &Config{}}
	for _, c := range config {
		c(robot)
	}
	return robot, nil
}

func main() {
	robotOne := NewRobot("one", nil)
	fmt.Printf("Robot name: %s \n", robotOne.Name())
	robotTwo := NewRobot("two", &Config{age: 15})
	fmt.Printf("Robot name: %s, age: %d \n", robotTwo.Name(), robotTwo.Config().age)
	robotThree, _ := NewRobotVariadic("three")
	fmt.Printf("Robot name: %s, age: %d \n", robotThree.name, robotThree.Config().age)
	robotFour, _ := NewRobotVariadic("four", &Config{age: 5})
	fmt.Printf("Robot name: %s, age: %d \n", robotFour.name, robotFour.Config().age)

	ageF := func(r *Robot) {
		r.config.age = 17
	}
	robotFive, _ := NewRobotFuncOptions("five", ageF)
	fmt.Printf("Robot name: %s, age: %d \n", robotFive.name, robotFive.Config().age)
}
