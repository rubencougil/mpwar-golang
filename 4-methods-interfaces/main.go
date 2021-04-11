package main

import "fmt"

type Pelicula struct {
	Nombre   string
	Director string
}

func NewPelicula(nombre, director string) *Pelicula {
	return &Pelicula{
		Nombre:   nombre,
		Director: director,
	}
}

func (p *Pelicula) getNombre() string {
	return p.Nombre
}

func (p *Pelicula) getNombreAndDirector() (string, string) {
	return p.Nombre, p.Director
}

func main() {
	p := NewPelicula("Rambo", "Spielberg")
	nombre, director := p.getNombreAndDirector()
	fmt.Printf("Nombre de Película: %s, Director %s", nombre, director)
}
