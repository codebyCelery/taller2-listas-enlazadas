package main

import "fmt"

type Nodo struct {
	Valor     int
	Siguiente *Nodo
}

func Desde(vals ...int) *Nodo {
	var cab, cola *Nodo
	for _, v := range vals {
		n := &Nodo{Valor: v}
		if cola == nil {
			cab = n
		} else {
			cola.Siguiente = n
		}
		cola = n
	}
	return cab
}

func Str(c *Nodo) string {
	s := ""
	for ; c != nil; c = c.Siguiente {
		s += fmt.Sprintf("%d ", c.Valor)
	}
	return s
}
func InvertirEnGrupos(cabeza *Nodo, k int) *Nodo {
	centinela := &Nodo{Siguiente: cabeza}
	colaPrevia := centinela
	actual := cabeza
	for actual != nil {
		 //corrección
		prueba := actual           
		n := 0                      
		for prueba != nil && n < k { 
			prueba = prueba.Siguiente 
			n++                       
		} 
		if n < k { 
			break
		} 
		// fin de la corrección
		inicioGrupo := actual
		var ant *Nodo
		cont := 0
		for actual != nil && cont < k {
			sig := actual.Siguiente
			actual.Siguiente = ant
			ant = actual
			actual = sig
			cont++
		}
		colaPrevia.Siguiente = ant
		inicioGrupo.Siguiente = actual
		colaPrevia = inicioGrupo
	}
	return centinela.Siguiente
}
//Implementacion de la parte C
func invertir(cabeza *Nodo) *Nodo {
	var ant *Nodo
	actual := cabeza
	for actual != nil {
		sig := actual.Siguiente
		actual.Siguiente = ant
		ant = actual
		actual = sig
	}
	return ant
}

func EsPalindromo(cabeza *Nodo) bool {
	if cabeza == nil || cabeza.Siguiente == nil {
		return true
	}
	lento, rapido := cabeza, cabeza
	for rapido.Siguiente != nil && rapido.Siguiente.Siguiente != nil {
		lento = lento.Siguiente
		rapido = rapido.Siguiente.Siguiente
	}
	segunda := invertir(lento.Siguiente)
	es := true
	a, b := cabeza, segunda
	for b != nil {
		if a.Valor != b.Valor {
			es = false
			break
		}
		a = a.Siguiente
		b = b.Siguiente
	}
	lento.Siguiente = invertir(segunda)
	return es
}
func main() {
	fmt.Println(Str(InvertirEnGrupos(Desde(1, 2, 3, 4, 5, 6), 3)))
	fmt.Println(Str(InvertirEnGrupos(Desde(1, 2, 3, 4, 5, 6, 7, 8), 3)))
}
