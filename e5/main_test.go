package main

import (
	"reflect"
	"testing"
)

func TestDosTerminanSeguidos(t *testing.T) {
	p := &Planificador{}
	p.Agregar("P1", 2)
	p.Agregar("P2", 2)
	p.Agregar("P3", 5)
	obtenido := p.Ejecutar(2)
	esperado := []string{"P1@t=2", "P2@t=4", "P3@t=9"}
	if !reflect.DeepEqual(obtenido, esperado) {
		t.Errorf("obtenido %v; se esperaba %v", obtenido, esperado)
	}
}
