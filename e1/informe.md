# Ejercicio 1: Lista simple con puntero Cola

**Responsable:** Madelein · **Revisora (PR):** _______

---

## A · Predice
> Escribe esta sección y haz **commit ANTES de ejecutar** el programa (`git log` es la evidencia).

**Predicción (antes de ejecutar):** Escribe las dos líneas que imprime main (lista y valor de n).

```text
10 -> 30 -> 40 -> 50 -> nil  n = 4
10 -> 30 -> 40 -> 60 -> nil  n = 4
```

**Salida real (después de ejecutar):**

```text
10 -> 30 -> 40 -> 50 -> nil | n = 4
10 -> 30 -> 40 -> nil | n = 4
```

**Comparación:** en qué acerté y en qué no.

---
La primera línea sale correctamente en la segunda también se ve correcta la lista y n, pero internamente  "cola" queda apuntando al nodo eliminado 50, causando el defecto

## B · Depura

**Hipótesis del defecto** (cita la evidencia: salida observada vs. esperada):

El error está en "eliminar()", cuando se elimina el último nodo se actualiza "siguiente" pero no se actualiza "Cola"

**Corrección mínima (diff):**

```diff
en esta parte:
			ant.Siguiente = ant.Siguiente.Siguiente
			l.n--
			return true
se agrega:
			if ant.Siguiente == nil {
			 l.Cola = ant
}
el resultado seria:
			ant.Siguiente = ant.Siguiente.Siguiente
			if ant.Siguiente == nil {
			 l.Cola = ant
			}
			l.n--
			return true
```

**Test que rompe el original** (nombre del test en `main_test.go` y por qué falla con el original):
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

---

## C · Construye

- [ A ] Validar() error (Cola es el último nodo alcanzable, n coincide, Cabeza == nil ⇔ Cola == nil)
- [ B ] EliminarUltimo() (int, error) — error si la lista está vacía
- [ C ] Llamar a Validar() después de cada operación en los tests

**Decisiones de diseño y justificación:**

A.  Rpta.
---
	if (l.Cabeza == nil) != (l.Cola == nil) {
        return fmt.Errorf("Cabeza y Cola no coinciden")
    }

    contador := 0
    var ultimo *Nodo

    for p := l.Cabeza; p != nil; p = p.Siguiente {
        contador++
        ultimo = p
    }

    if contador != l.n {
        return fmt.Errorf("n no coincide con la cantidad de nodos")
    }

    if ultimo != l.Cola {
        return fmt.Errorf("Cola no es el último nodo")
    }

    return nil
}
---

B. Rpta.
---
	func (l *Lista) EliminarUltimo() (int, error) {
    if l.Cabeza == nil {
        return 0, fmt.Errorf("lista vacía")
    }

    if l.Cabeza == l.Cola {
        valor := l.Cabeza.Valor
        l.Cabeza = nil
        l.Cola = nil
        l.n--
        return valor, nil
    }

    actual := l.Cabeza

    for actual.Siguiente != l.Cola {
        actual = actual.Siguiente
    }

    valor := l.Cola.Valor
    actual.Siguiente = nil
    l.Cola = actual
    l.n--

    return valor, nil
}
---
C.  Rpta. 
--- 
	func TestLista(t *testing.T) {
    l := &Lista{}

    l.InsertarFinal(10)
    if err := l.Validar(); err != nil {
        t.Fatal(err)
    }

    l.InsertarFinal(20)
    if err := l.Validar(); err != nil {
        t.Fatal(err)
    }

    l.InsertarFinal(30)
    if err := l.Validar(); err != nil {
        t.Fatal(err)
    }

    l.Eliminar(20)
    if err := l.Validar(); err != nil {
        t.Fatal(err)
    }

    l.EliminarUltimo()
    if err := l.Validar(); err != nil {
        t.Fatal(err)
    }

---
## D · Defiende (preparación, no se entrega)

1. **¿Qué invariante se rompía y en qué línea exacta?**
   - Se rompía el invariante de que "cola" debe apuntar al último nodo de la lista
   Se rompe en: ant.Siguiente = ant.Siguiente.Siguiente

2. **¿Por qué eliminar un nodo intermedio nunca revelaba el defecto?**
   - Porque "cola" sigue siendo el mismo último nodo, solo aparece el problema cuando se elimina el último nodo

3. **Después de EliminarUltimo, ¿qué nodo debe quedar como Cola y cómo lo encuentras?**
   - Queda como Cola el penúltimo nodo, es decir, el último nodo que queda en la lista, se encuentra recorriendo la lista hasta que:
     actual.Siguiente == l.Cola
   Entonces
     l.Cola = actual
