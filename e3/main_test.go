package main

import "testing"

// TestInsertarOrdenadoAtras usa datos desordenados para forzar inserciones intermedias.
// Con el código original falla: el Anterior del nodo siguiente al nuevo no se actualiza.
func TestInsertarOrdenadoAtras(t *testing.T) {
	l := &ListaDoble{}
	for _, v := range []int{40, 10, 30, 20, 50} {
		l.InsertarOrdenado(v)
	}
	esperado := "50 40 30 20 10 "
	obtenido := l.Atras()
	if obtenido != esperado {
		t.Errorf("Atras()=%q; se esperaba %q", obtenido, esperado)
	}
}

// TestInsertarOrdenadoAdelante comprueba que el recorrido hacia adelante sigue ordenado.
func TestInsertarOrdenadoAdelante(t *testing.T) {
	l := &ListaDoble{}
	for _, v := range []int{40, 10, 30, 20, 50} {
		l.InsertarOrdenado(v)
	}
	esperado := "10 20 30 40 50 "
	obtenido := l.Adelante()
	if obtenido != esperado {
		t.Errorf("Adelante()=%q; se esperaba %q", obtenido, esperado)
	}
}
