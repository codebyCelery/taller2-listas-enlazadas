package main

import "testing"

func TestMergeSortIgualInsertionSort(t *testing.T) {
	casos := [][]Alumno{
		{A("Ana", 14), A("Beto", 18), A("Caro", 11), A("Dani", 16)},
		{A("Ana", 15), A("Beto", 18), A("Caro", 15), A("Dani", 12), A("Eva", 18), A("Fito", 15)},
		{},
		{A("Solo", 10)},
		{A("A", 5), A("B", 5), A("C", 5)},
		{A("A", 1), A("B", 2), A("C", 3), A("D", 4)},
	}
	for i, datos := range casos {
		m := Str(MergeSort(Desde(datos...)))
		in := Str(InsertionSort(Desde(datos...)))
		if m != in {
			t.Errorf("caso %d: MergeSort=%q InsertionSort=%q", i, m, in)
		}
	}
}
