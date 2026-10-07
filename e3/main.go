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
	} else {
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

// Implementaciones de la parte C
func (l *ListaDoble) Validar() error {
	if l.Cabeza == nil {
		if l.Cola != nil {
			return fmt.Errorf("lista vacía pero Cola no es nil")
		}
		return nil
	}
	if l.Cabeza.Anterior != nil {
		return fmt.Errorf("la cabeza (%d) tiene un Anterior", l.Cabeza.Valor)
	}
	var ultimo *NodoDoble
	for n := l.Cabeza; n != nil; n = n.Siguiente {
		if n.Siguiente != nil && n.Siguiente.Anterior != n {
			return fmt.Errorf("el nodo %d no es el Anterior de su siguiente (%d)", n.Valor, n.Siguiente.Valor)
		}
		ultimo = n
	}
	if l.Cola != ultimo {
		return fmt.Errorf("Cola no apunta al último nodo (%d)", ultimo.Valor)
	}
	return nil
}

func (l *ListaDoble) EliminarNodo(n *NodoDoble) {
	if n.Anterior != nil {
		n.Anterior.Siguiente = n.Siguiente
	} else {
		l.Cabeza = n.Siguiente
	}
	if n.Siguiente != nil {
		n.Siguiente.Anterior = n.Anterior
	} else {
		l.Cola = n.Anterior
	}
	n.Anterior, n.Siguiente = nil, nil
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
