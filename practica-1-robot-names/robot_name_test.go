package robotname

import (
	"regexp"
	"testing"
)

var namePat = regexp.MustCompile(`^[A-Z]{2}\d{3}$`)

func New() *Robot { return new(Robot) }

func TestNameValid(t *testing.T) {
	n := New().Name()
	if !namePat.MatchString(n) {
		t.Errorf(`Nombre de Robot no válido %q, el patrón es "AA###".`, n)
	}
}

func TestNameSticks(t *testing.T) {
	r := New()
	n1 := r.Name()
	n2 := r.Name()
	if n2 != n1 {
		t.Errorf(`El nombre del robot ha cambiado.  Ahora %s, era %s.`, n2, n1)
	}
}

func TestSuccessiveRobotsHaveDifferentNames(t *testing.T) {
	n1 := New().Name()
	if New().Name() == n1 {
		t.Errorf(`Robots con el mismo nombre detectados.  Dos %s's.`, n1)
	}
}

func TestResetName(t *testing.T) {
	r := New()
	n1 := r.Name()
	r.Reset()
	if r.Name() == n1 {
		t.Errorf(`El nombre del Robot no ha sido reseteado.  Sigue siendo %s.`, n1)
	}
}
