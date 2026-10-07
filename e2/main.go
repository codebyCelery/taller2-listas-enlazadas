package main

import "fmt"

type Nodo struct {
	Valor     int
	Siguiente *Nodo
}

type Lista struct{ Cabeza *Nodo }

func Desde(vals ...int) *Lista {
	l := &Lista{}
	var cola *Nodo
	for _, v := range vals {
		n := &Nodo{Valor: v}
		if cola == nil {
			l.Cabeza = n
		} else {
			cola.Siguiente = n
		}
		cola = n
	}
	return l
}

func (l *Lista) String() string {
	s := ""
	for a := l.Cabeza; a != nil; a = a.Siguiente {
		s += fmt.Sprintf("%d -> ", a.Valor)
	}
	return s + "nil"
}

// EliminarTodos borra TODAS las apariciones de v y retorna cuántas borró.
func (l *Lista) EliminarTodos(v int) int {
	borrados := 0
	for l.Cabeza != nil && l.Cabeza.Valor == v {
		l.Cabeza = l.Cabeza.Siguiente
		borrados++
	}
	ant := l.Cabeza
	for ant != nil && ant.Siguiente != nil {
		if ant.Siguiente.Valor == v {
			ant.Siguiente = ant.Siguiente.Siguiente
			borrados++
		} else {
			ant = ant.Siguiente
		}
	}
	return borrados
}

func main() {
	a := Desde(3, 7, 1, 7, 5, 7)
	fmt.Println(a.EliminarTodos(7), a)
	b := Desde(7, 7, 2, 7, 7, 7, 4)
	fmt.Println(b.EliminarTodos(7), b)
}
