package main

import "fmt"

func main() {

	// --- Variables ---

	var foo string
	foo = "Foo"
	foo := "Foo"
	fmt.Println(foo)

	bar := "Bar"
	fmt.Println(bar)

	// --- Arrays ---

	var cheeses [2]string
	cheeses[0] = "Gruyere"
	cheeses[1] = "Feta"

	// --- Slices ---

	var cars = make([]string, 2)
	cars[0] = "Audi"
	cars = append(cars, "Seat", "Fiat", "Toyota")
	fmt.Printf("%v", cars)

	// --- Maps ---

	var players = make(map[string]int)
	players["Mike"] = 8
	players["John"] = 10
	delete(players, "Mike")
}
