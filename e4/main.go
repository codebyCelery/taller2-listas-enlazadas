package main

import "fmt"

type Nodo struct {
	Valor     int
	Siguiente *Nodo
}

func Desde(vals ...int) *Nodo {
	var cab, cola *Nodo
	for _, v := range vals {
		n := &Nodo{Valor: v}
		if cola == nil {
			cab = n
		} else {
			cola.Siguiente = n
		}
		cola = n
	}
	return cab
}

func Str(c *Nodo) string {
	s := ""
	for ; c != nil; c = c.Siguiente {
		s += fmt.Sprintf("%d ", c.Valor)
	}
	return s
}

// InvertirEnGrupos invierte bloques consecutivos de k nodos.
// Especificación: si al final quedan menos de k nodos, ese bloque
// final se deja en su orden original.
func InvertirEnGrupos(cabeza *Nodo, k int) *Nodo {
	centinela := &Nodo{Siguiente: cabeza}
	colaPrevia := centinela
	actual := cabeza
	for actual != nil {
		inicioGrupo := actual
		var ant *Nodo
		cont := 0
		for actual != nil && cont < k {
			sig := actual.Siguiente
			actual.Siguiente = ant
			ant = actual
			actual = sig
			cont++
		}
		colaPrevia.Siguiente = ant
		inicioGrupo.Siguiente = actual
		colaPrevia = inicioGrupo
	}
	return centinela.Siguiente
}

func main() {
	fmt.Println(Str(InvertirEnGrupos(Desde(1, 2, 3, 4, 5, 6), 3)))
	fmt.Println(Str(InvertirEnGrupos(Desde(1, 2, 3, 4, 5, 6, 7, 8), 3)))
}
