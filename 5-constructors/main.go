package main

import (
	"fmt"
)

type Config struct {
	age int
}

type Robot struct {
	name   string
	config *Config
}

func NewRobot(name string, config *Config) *Robot {
	return &Robot{
		name:   name,
		config: config,
	}
}

func (r *Robot) getName() string {
	return r.name
}

func (r *Robot) getConfig() *Config {
	return r.config
}

func main() {

	robotOne := NewRobot("one", nil)
	fmt.Printf("Robot name: %s \n", robotOne.getName())

	robotTwo := NewRobot("two", &Config{age: 15})
	fmt.Printf("Robot name: %s, age: %d \n", robotTwo.getName(), robotTwo.getConfig().age)
}
