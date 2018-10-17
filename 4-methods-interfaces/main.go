package main

type Pelicula struct {
	Nombre string
	Director string
}

func (p *Pelicula) setNombre(nombre string) {
	p.Nombre = nombre
}

func (p *Pelicula) getNombre() string {
	return p.Nombre
}

func (p *Pelicula) getNombreAndDirector() (string, string) {
	return p.Nombre, p.Director
}

