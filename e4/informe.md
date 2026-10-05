# Ejercicio 4: Invertir en grupos de k
---

## A · Predice

**Predicción (antes de ejecutar):** Escribe las dos líneas de main.

```text
3 2 1 6 5 4
3 2 1 6 5 4 7 8
```

**Salida real (después de ejecutar):**

```text
3 2 1 6 5 4
3 2 1 6 5 4 8 7
```

**Comparación:** en qué acerté y en qué no.

La primera línea coincidió. La segunda no: el último bloque [7 8] tiene menos de 3 nodos y aun así se invirtió (8 7), lo que contradice la especificación.

## B · Depura

**Hipótesis del defecto** (cita la evidencia: salida observada vs. esperada):
Observada: 3 2 1 6 5 4 8 7
Esperada: 3 2 1 6 5 4 7 8

El defecto está en InvertirEnGrupos. El for interno invierte nodos "mientras haya nodos y cont < k", pero el código nunca comprueba, antes de empezar, si quedan al menos k nodos. Si quedan menos, igual los invierte, y el último bloque incompleto queda al revés, lo que contradice la especificación. Con la lista del 1 al 8 y k = 3, el bloque [7 8] tiene solo 2 nodos y aun así sale 8 7. El defecto se activa cuando n no es múltiplo de k (y también cuando k > n). La primera línea de main sale bien porque 6 es múltiplo de 3: todos los bloques tienen exactamente k nodos y nunca aparece un bloque incompleto. La corrección es contar hacia adelante si quedan k nodos antes de invertir y, si no, salir del bucle sin tocarlos.

**Corrección mínima (diff):**

```diff
 	for actual != nil {
+		prueba := actual
+		n := 0
+		for prueba != nil && n < k {
+			prueba = prueba.Siguiente
+			n++
+		}
+		if n < k {
+			break
+		}
 		inicioGrupo := actual
 		var ant *Nodo
 		cont := 0
```

**Test que rompe el original:** ```text TestInvertirEnGruposBloqueIncompleto ```

Invierte la lista 1, 2, 3, 4, 5, 6, 7, 8 con k = 3 y exige que el resultado sea "3 2 1 6 5 4 7 8 ". Con el código original falla, porque el bloque incompleto [7 8] también se invierte y el resultado es "3 2 1 6 5 4 8 7 ". No se usa el caso de 6 nodos, porque 6 es múltiplo de 3 y el test pasaría incluso con el defecto. Con la corrección, el bloque incompleto no se toca y el test pasa.

Evidencia con el código original:
 ```text
--- FAIL: TestInvertirEnGruposBloqueIncompleto (0.00s)
    main_test.go:9: InvertirEnGrupos(1..8, k=3)="3 2 1 6 5 4 8 7 "; se esperaba "3 2 1 6 5 4 7 8 "
FAIL
FAIL    taller2/e4      0.390s

ok      taller2/e4      0.400s
```
---

## C · Construye

- [x] EsPalindromo(cabeza *Nodo) bool sin copiar a un slice: medio → invertir segunda mitad → comparar
- [x] La lista debe quedar exactamente como estaba (probarlo en un test)

**Decisiones de diseño y justificación:**

Medio con dos punteros: uno avanza de uno en uno y otro de dos en dos. Cuando el rápido llega al final, el lento está en el medio.

Invertir y comparar: se invierte la segunda mitad con los tres punteros (ant, actual, sig) y se compara con la primera. Es O(n) en tiempo y O(1) en espacio, sin copiar a un slice.
Restaurar: al final se vuelve a invertir la segunda mitad, para que la lista quede exactamente como estaba, como pide el enunciado.

Tests: lista vacía, un nodo, par (1 2 2 1), impar (1 2 1) y no palíndromo (1 2 3).

---
