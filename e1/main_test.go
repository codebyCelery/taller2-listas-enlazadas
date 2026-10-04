package main

package main
	import "testing"
	func TestEliminarUltimoActualizaCola(t *testing.T) {
	    l := &Lista{}
	    l.InsertarFinal(10)
	    l.InsertarFinal(20)
	    l.InsertarFinal(30)
	    l.Eliminar(30)
	    if l.Cola.Valor != 20 {
	        t.Errorf("esperaba Cola = 20, obtuvo %d", l.Cola.Valor)
}
