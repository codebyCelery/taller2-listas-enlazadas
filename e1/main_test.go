package main

import (
	"testing"
)

func TestInsertarFinalYLongitud(t *testing.T) {
	l := &Lista{}

	if l.Longitud() != 0 {
		t.Errorf("Esperaba longitud 0, obtuve %d", l.Longitud())
	}

	l.InsertarFinal(10)
	l.InsertarFinal(20)
	l.InsertarFinal(30)

	if l.Longitud() != 3 {
		t.Errorf("Esperaba longitud 3, obtuve %d", l.Longitud())
	}

	esperado := "10 -> 20 -> 30 -> nil"
	if l.String() != esperado {
		t.Errorf("Esperaba %q, obtuve %q", esperado, l.String())
	}
}

func TestEliminar(t *testing.T) {
	l := &Lista{}
	for _, v := range []int{10, 20, 30, 40} {
		l.InsertarFinal(v)
	}

	
	if !l.Eliminar(20) {
		t.Error("Debería haber eliminado el 20")
	}
	if l.Longitud() != 3 {
		t.Errorf("Esperaba longitud 3 tras eliminar, obtuve %d", l.Longitud())
	}

	
	if !l.Eliminar(10) {
		t.Error("Debería haber eliminado el 10 (cabeza)")
	}

	
	if l.Eliminar(999) {
		t.Error("No debería poder eliminar un valor que no existe")
	}

	
	if !l.Eliminar(40) {
		t.Error("Debería haber eliminado el 40")
	}
	if !l.Eliminar(30) {
		t.Error("Debería haber eliminado el 30")
	}

	if l.Longitud() != 0 {
		t.Errorf("Esperaba lista vacía (longitud 0), obtuve %d", l.Longitud())
	}

	if l.Cabeza != nil || l.Cola != nil {
		t.Error("Cabeza y Cola deberían ser nil cuando la lista queda vacía")
	}
}


func TestEliminarListaVacia(t *testing.T) {
	l := &Lista{}
	if l.Eliminar(10) {
		t.Error("No se puede eliminar de una lista vacía")
	}
}


func TestColaActualizadaTrasEliminarCola(t *testing.T) {
	l := &Lista{}
	l.InsertarFinal(100)
	l.InsertarFinal(200)

	
	l.Eliminar(200)

	if l.Cola == nil || l.Cola.Valor != 100 {
		t.Errorf("La cola debería ser 100, pero es %v", l.Cola)
	}
}
