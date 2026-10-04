package main

import "fmt"

type Alumno struct {
	Nombre    string
	Nota      int
	Siguiente *Alumno
}

func Desde(datos ...Alumno) *Alumno {
	var cab, cola *Alumno
	for i := range datos {
		n := &Alumno{Nombre: datos[i].Nombre, Nota: datos[i].Nota}
		if cola == nil {
			cab = n
		} else {
			cola.Siguiente = n
		}
		cola = n
	}
	return cab
}

func A(nombre string, nota int) Alumno { return Alumno{Nombre: nombre, Nota: nota} }

func Str(c *Alumno) string {
	s := ""
	for ; c != nil; c = c.Siguiente {
		s += fmt.Sprintf("%s(%d) ", c.Nombre, c.Nota)
	}
	return s
}

func ObtenerMedio(c *Alumno) *Alumno {
	lento, rapido := c, c.Siguiente
	for rapido != nil && rapido.Siguiente != nil {
		lento = lento.Siguiente
		rapido = rapido.Siguiente.Siguiente
	}
	return lento
}

// Mezclar combina dos listas ordenadas por Nota DESCENDENTE.
func Mezclar(a, b *Alumno) *Alumno {
	centinela := &Alumno{}
	cola := centinela
	for a != nil && b != nil {
		if a.Nota >= b.Nota {
			cola.Siguiente, a = a, a.Siguiente
		} else {
			cola.Siguiente, b = b, b.Siguiente
		}
		cola = cola.Siguiente
	}
	if a != nil {
		cola.Siguiente = a
	} else {
		cola.Siguiente = b
	}
	return centinela.Siguiente
}

// MergeSort ordena por Nota descendente. Debe ser ESTABLE: a igual nota,
// se respeta el orden de inscripción (orden original de la lista).
func MergeSort(c *Alumno) *Alumno {
	if c == nil || c.Siguiente == nil {
		return c
	}
	medio := ObtenerMedio(c)
	der := medio.Siguiente
	medio.Siguiente = nil
	return Mezclar(MergeSort(c), MergeSort(der))
}

func main() {
	l := Desde(A("Ana", 14), A("Beto", 18), A("Caro", 11), A("Dani", 16))
	fmt.Println(Str(MergeSort(l)))
	m := Desde(A("Ana", 15), A("Beto", 18), A("Caro", 15), A("Dani", 12), A("Eva", 18), A("Fito", 15))
	fmt.Println(Str(MergeSort(m)))
}
func InsertionSort(c *Alumno) *Alumno {
	var ordenada *Alumno
	for c != nil {
		sig := c.Siguiente
		if ordenada == nil || c.Nota > ordenada.Nota {
			c.Siguiente = ordenada
			ordenada = c
		} else {
			p := ordenada
			for p.Siguiente != nil && p.Siguiente.Nota >= c.Nota {
				p = p.Siguiente
			}
			c.Siguiente = p.Siguiente
			p.Siguiente = c
		}
		c = sig
	}
	return ordenada
}
