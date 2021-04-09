package main

import "fmt"

func main() {

	// --- Variables ---

	var foo string
	foo = "Foo"
	fmt.Println(foo)

	bar := "Bar"
	fmt.Println(bar)

	// --- Arrays ---

	var cheeses [2]string
	cheeses[0] = "Grullere"
	cheeses[1] = "Feta"

	// --- Slices ---

	var cars = make([]string, 2)
	cars[0] = "Honda"
	cars[1] = "Volvo"
	//cars[2] = "Bmw"
	cars = append(cars, "Seat", "Fiat", "Toyota")
	cars = append(cars[:2], cars[2+1:]...)

	// --- Maps ---

	var players = make(map[string]int)
	players["Mike"] = 8
	players["John"] = 10
	delete(players, "Mike")
}
