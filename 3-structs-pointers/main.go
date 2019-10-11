package main

import "fmt"

type Persona struct {
	Nombre string
	Edad   int
}

type Cuadro struct {
	Color   string
	Medidas Medidas
}

type Medidas struct {
	Ancho float32
	Alto  float32
}

func main() {

	// --- Struct ---

	var p Persona
	p.Nombre = "Mike"
	p.Edad = 32

	var p2 = Persona{Nombre: "John", Edad: 29}

	var p3 = new(Persona)
	p3.Nombre = "Alex"
	p3.Edad = 31

	p4 := Persona{
		Nombre: "Berta",
		Edad:   22,
	}

	// --- Struct con Struct anidado ---

	cuadro := Cuadro{
		Color: "blanco",
		Medidas: Medidas{
			Alto:  1.23,
			Ancho: 12.3,
		},
	}

	// --- Punteros ----

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
