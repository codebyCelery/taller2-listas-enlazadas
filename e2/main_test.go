package main

import (
	"testing"
)

func TestDesdeYString(t *testing.T) {
	l := Desde(1, 2, 3)

	esperado := "1 -> 2 -> 3 -> nil"
	if l.String() != esperado {
		t.Errorf("Esperaba %q, obtuve %q", esperado, l.String())
	}
}

func TestEliminarTodosCasosPrincipales(t *testing.T) {
	// Caso 1: Elementos repetidos al inicio, medio y final
	l := Desde(3, 7, 1, 7, 5, 7)
	borrados := l.EliminarTodos(7)

	if borrados != 3 {
		t.Errorf("Esperaba borrar 3 elementos, borró %d", borrados)
	}

	esperado := "3 -> 1 -> 5 -> nil"
	if l.String() != esperado {
		t.Errorf("Esperaba %q, obtuve %q", esperado, l.String())
	}
}

func TestEliminarTodosConsecutivosYEliminarTodo(t *testing.T) {
	// Caso con muchos 7 al principio y en el medio
	l := Desde(7, 7, 2, 7, 7, 7, 4)
	borrados := l.EliminarTodos(7)

	if borrados != 5 {
		t.Errorf("Esperaba borrar 5 elementos, borró %d", borrados)
	}

	esperado := "2 -> 4 -> nil"
	if l.String() != esperado {
		t.Errorf("Esperaba %q, obtuve %q", esperado, l.String())
	}
}

func TestEliminarTodosNoExiste(t *testing.T) {
	l := Desde(1, 2, 3)
	borrados := l.EliminarTodos(99)

	if borrados != 0 {
		t.Errorf("Esperaba borrar 0 elementos, borró %d", borrados)
	}

	esperado := "1 -> 2 -> 3 -> nil"
	if l.String() != esperado {
		t.Errorf("La lista no debería haber cambiado, obtuve %q", l.String())
	}
}

func TestEliminarTodosListaVacia(t *testing.T) {
	l := &Lista{}
	borrados := l.EliminarTodos(5)

	if borrados != 0 {
		t.Errorf("Esperaba 0 borrados en lista vacía, obtuve %d", borrados)
	}

	if l.Cabeza != nil {
		t.Error("La cabeza debería seguir siendo nil")
	}
}
