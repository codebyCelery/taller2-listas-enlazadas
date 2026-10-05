package main

import "testing"

func TestInvertirEnGruposBloqueIncompleto(t *testing.T) {
	obtenido := Str(InvertirEnGrupos(Desde(1, 2, 3, 4, 5, 6, 7, 8), 3))
	esperado := "3 2 1 6 5 4 7 8 "
	if obtenido != esperado {
		t.Errorf("InvertirEnGrupos(1..8, k=3)=%q; se esperaba %q", obtenido, esperado)
	}
}

func TestInvertirEnGruposCasos(t *testing.T) {
	casos := []struct {
		nombre   string
		lista    []int
		k        int
		esperado string
	}{
		{"n multiplo de k", []int{1, 2, 3, 4, 5, 6}, 3, "3 2 1 6 5 4 "},
		{"k = 1", []int{1, 2, 3}, 1, "1 2 3 "},
		{"k mayor que n", []int{1, 2}, 5, "1 2 "},
		{"k igual a n", []int{1, 2, 3}, 3, "3 2 1 "},
		{"lista vacia", []int{}, 3, ""},
	}
	for _, c := range casos {
		obtenido := Str(InvertirEnGrupos(Desde(c.lista...), c.k))
		if obtenido != c.esperado {
			t.Errorf("%s: obtenido %q; se esperaba %q", c.nombre, obtenido, c.esperado)
		}
	}
}

func TestEsPalindromo(t *testing.T) {
	casos := []struct {
		nombre   string
		lista    []int
		esperado bool
	}{
		{"vacia", []int{}, true},
		{"un nodo", []int{5}, true},
		{"par", []int{1, 2, 2, 1}, true},
		{"impar", []int{1, 2, 1}, true},
		{"no palindromo", []int{1, 2, 3}, false},
		{"casi", []int{1, 2, 3, 1}, false},
	}
	for _, c := range casos {
		cab := Desde(c.lista...)
		antes := Str(cab)
		if got := EsPalindromo(cab); got != c.esperado {
			t.Errorf("%s: EsPalindromo=%v; se esperaba %v", c.nombre, got, c.esperado)
		}
		if despues := Str(cab); despues != antes {
			t.Errorf("%s: la lista cambió: %q -> %q", c.nombre, antes, despues)
		}
	}
}
