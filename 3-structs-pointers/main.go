package main

import "fmt"

type Persona struct {
	Nombre string
	Edad   int
}

func main() {

	pPointer := &Persona{Nombre: "Victoria"}
	pValue := Persona{Nombre: "Mario"}

	cambiarNombrePointer(pPointer)
	cambiarNombre(pValue)

	fmt.Println(pPointer.Nombre)
	fmt.Println(pValue.Nombre)
}

func cambiarNombrePointer(p *Persona) {
	p.Nombre = "Felipe"
}

func cambiarNombre(p Persona) {
	p.Nombre = "Felipe"
}
