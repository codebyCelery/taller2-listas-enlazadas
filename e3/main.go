package main

import "fmt"

type NodoDoble struct {
	Valor     int
	Anterior  *NodoDoble
	Siguiente *NodoDoble
}

type ListaDoble struct {
	Cabeza *NodoDoble
	Cola   *NodoDoble
}

// InsertarOrdenado mantiene la lista en orden ascendente.
func (l *ListaDoble) InsertarOrdenado(v int) {
	nuevo := &NodoDoble{Valor: v}
	if l.Cabeza == nil {
		l.Cabeza, l.Cola = nuevo, nuevo
		return
	}
	if v < l.Cabeza.Valor {
		nuevo.Siguiente = l.Cabeza
		l.Cabeza.Anterior = nuevo
		l.Cabeza = nuevo
		return
	}
	actual := l.Cabeza
	for actual.Siguiente != nil && actual.Siguiente.Valor < v {
		actual = actual.Siguiente
	}
	nuevo.Siguiente = actual.Siguiente
	nuevo.Anterior = actual
	actual.Siguiente = nuevo
	if nuevo.Siguiente == nil {
		l.Cola = nuevo
	}else{
		nuevo.Siguiente.Anterior = nuevo
		}
	//este else es la correción
}

func (l *ListaDoble) Adelante() string {
	s := ""
	for a := l.Cabeza; a != nil; a = a.Siguiente {
		s += fmt.Sprintf("%d ", a.Valor)
	}
	return s
}

func (l *ListaDoble) Atras() string {
	s := ""
	for a := l.Cola; a != nil; a = a.Anterior {
		s += fmt.Sprintf("%d ", a.Valor)
	}
	return s
}

func main() {
	l := &ListaDoble{}
	for _, v := range []int{10, 20, 30, 40, 50} {
		l.InsertarOrdenado(v)
	}
	fmt.Println("A:", l.Adelante(), "| B:", l.Atras())
	m := &ListaDoble{}
	for _, v := range []int{40, 10, 30, 20, 50} {
		m.InsertarOrdenado(v)
	}
	fmt.Println("A:", m.Adelante(), "| B:", m.Atras())
}
