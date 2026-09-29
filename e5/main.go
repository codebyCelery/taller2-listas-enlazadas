package main

import "fmt"

type Proceso struct {
	Nombre    string
	Tiempo    int // ráfagas de CPU restantes
	Siguiente *Proceso
}

// Planificador: lista circular; Cola apunta al último, Cola.Siguiente es el primero.
type Planificador struct {
	Cola    *Proceso
	Tamanio int
}

func (p *Planificador) Agregar(nombre string, tiempo int) {
	nuevo := &Proceso{Nombre: nombre, Tiempo: tiempo}
	if p.Cola == nil {
		nuevo.Siguiente = nuevo
	} else {
		nuevo.Siguiente = p.Cola.Siguiente
		p.Cola.Siguiente = nuevo
	}
	p.Cola = nuevo
	p.Tamanio++
}

// Ejecutar simula Round Robin y retorna el orden de finalización.
func (p *Planificador) Ejecutar(quantum int) []string {
	var orden []string
	reloj := 0
	ant := p.Cola
	act := p.Cola.Siguiente
	for p.Tamanio > 0 {
		usa := min(quantum, act.Tiempo)
		act.Tiempo -= usa
		reloj += usa
		if act.Tiempo == 0 {
			orden = append(orden, fmt.Sprintf("%s@t=%d", act.Nombre, reloj))
			ant.Siguiente = act.Siguiente // desenlaza el proceso terminado
			p.Tamanio--
		}
		ant = act
		act = act.Siguiente
	}
	p.Cola = nil
	return orden
}

func main() {
	a := &Planificador{}
	a.Agregar("P1", 6)
	a.Agregar("P2", 4)
	a.Agregar("P3", 8)
	fmt.Println(a.Ejecutar(2))

	b := &Planificador{}
	b.Agregar("P1", 2)
	b.Agregar("P2", 2)
	b.Agregar("P3", 5)
	fmt.Println(b.Ejecutar(2))
}
