package main

import "fmt"

type Nodo struct {
	Valor     int
	Siguiente *Nodo
}

// Lista mantiene Cabeza y Cola para insertar al final sin recorrer la lista.
type Lista struct {
	Cabeza *Nodo
	Cola   *Nodo
	n      int
}

func (l *Lista) InsertarFinal(v int) {
	nuevo := &Nodo{Valor: v}
	if l.Cabeza == nil {
		l.Cabeza, l.Cola = nuevo, nuevo
	} else {
		l.Cola.Siguiente = nuevo
		l.Cola = nuevo
	}
	l.n++
}

func (l *Lista) Eliminar(v int) bool {
	if l.Cabeza == nil {
		return false
	}
	if l.Cabeza.Valor == v {
		l.Cabeza = l.Cabeza.Siguiente
		if l.Cabeza == nil {
			l.Cola = nil
		}
		l.n--
		return true
	}
	ant := l.Cabeza
	for ant.Siguiente != nil {
		if ant.Siguiente.Valor == v {
			ant.Siguiente = ant.Siguiente.Siguiente
			l.n--
			return true
		}
		ant = ant.Siguiente
	}
	return false
}

func (l *Lista) Longitud() int { return l.n }

func (l *Lista) String() string {
	s := ""
	for a := l.Cabeza; a != nil; a = a.Siguiente {
		s += fmt.Sprintf("%d -> ", a.Valor)
	}
	return s + "nil"
}

func main() {
	l := &Lista{}
	for _, v := range []int{10, 20, 30, 40} {
		l.InsertarFinal(v)
	}
	l.Eliminar(20)
	l.InsertarFinal(50)
	fmt.Println(l, "| n =", l.Longitud())
	l.Eliminar(50)
	l.InsertarFinal(60)
	fmt.Println(l, "| n =", l.Longitud())
}
