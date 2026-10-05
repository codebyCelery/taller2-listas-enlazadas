package main

import "testing"

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

func construir(vals ...int) *ListaDoble {
	l := &ListaDoble{}
	for _, v := range vals {
		l.InsertarOrdenado(v)
	}
	return l
}

func TestValidarDespuesDeInsertar(t *testing.T) {
	l := &ListaDoble{}
	for _, v := range []int{40, 10, 30, 20, 50} {
		l.InsertarOrdenado(v)
		if err := l.Validar(); err != nil {
			t.Fatalf("después de insertar %d: %v", v, err)
		}
	}
}

func TestEliminarNodo(t *testing.T) {
	casos := []struct {
		nombre   string
		borrar   func(l *ListaDoble) *NodoDoble
		adelante string
		atras    string
	}{
		{"cabeza", func(l *ListaDoble) *NodoDoble { return l.Cabeza }, "20 30 ", "30 20 "},
		{"intermedio", func(l *ListaDoble) *NodoDoble { return l.Cabeza.Siguiente }, "10 30 ", "30 10 "},
		{"cola", func(l *ListaDoble) *NodoDoble { return l.Cola }, "10 20 ", "20 10 "},
	}
	for _, c := range casos {
		l := construir(10, 20, 30)
		l.EliminarNodo(c.borrar(l))
		if err := l.Validar(); err != nil {
			t.Errorf("%s: %v", c.nombre, err)
		}
		if l.Adelante() != c.adelante || l.Atras() != c.atras {
			t.Errorf("%s: %q | %q", c.nombre, l.Adelante(), l.Atras())
		}
	}
	unico := construir(7)
	unico.EliminarNodo(unico.Cabeza)
	if unico.Cabeza != nil || unico.Cola != nil {
		t.Errorf("nodo único: Cabeza y Cola deberían ser nil")
	}
	if err := unico.Validar(); err != nil {
		t.Errorf("nodo único: %v", err)
	}
}
