# Ejercicio 6: Merge Sort estable en lista

**Responsable:** Celery

---

## A · Predice
> ⚠️ Escribe esta sección y haz **commit ANTES de ejecutar** el programa (`git log` es la evidencia).

**Predicción (antes de ejecutar):**
Beto(18) Dani(16) Ana(14) Caro(11) 
Eva(18) Beto(18) Fito(15) Caro(15) Ana(15) Dani(12)

**Salida real (después de ejecutar):**

```text
(pega aquí la salida de la terminal)
```

**Comparación:** en qué acerté y en qué no.

---

## B · Depura

**Hipótesis del defecto** 
Eva(18) se procesa antes que Beto(18) porque en la función Mezclar usas una condición de comparación estricta (if a.Nota > b.Nota). Esa condición resulta falsa cuando ambas notas son iguales ($18 > 18$). Por eso el programa entra en el bloque else y extrae primero el nodo de la sublista derecha b, donde se encuentra Eva. La lista derecha contiene los elementos que aparecían después en la secuencia original. Al dar prioridad a b cuando hay empate, se invierte el orden relativo de los alumnos con la misma nota, rompiendo la estabilidad del algoritmo. Para solucionarlo y mantener a Beto primero, la condición debe ser mayor o igual (a.Nota >= b.Nota). Así, ante un empate, el algoritmo siempre preserva el elemento de la sublista izquierda a.

**Corrección mínima (diff):**
-		if a.Nota > b.Nota {
+		if a.Nota >= b.Nota {

**Test que rompe el original**
Test que rompe el original: TestMergeSortIgualInsertionSort. Ordena los mismos datos con MergeSort e InsertionSort y exige que den exactamente el mismo orden, incluidos los empates. Con el código original falla en los casos 1 y 4, porque Mezclar usa > estricto y en un empate toma primero el nodo de la derecha, invirtiendo el orden original (por ejemplo, A B C sale como C B A). InsertionSort es estable y conserva ese orden, por eso los resultados difieren. Con >= el test pasa.
*Evidencia:
--- FAIL: TestMergeSortIgualInsertionSort (0.00s)
    main_test.go:18: caso 1: MergeSort="Eva(18) Beto(18) Fito(15) Caro(15) Ana(15) Dani(12) " InsertionSort="Beto(18) Eva(18) Ana(15) Caro(15) Fito(15) Dani(12) "
    main_test.go:18: caso 4: MergeSort="C(5) B(5) A(5) " InsertionSort="A(5) B(5) C(5) "
FAIL
FAIL    taller2/e6      0.011s
*Después de la corrección (> por >=):
ok      taller2/e6      0.003s
## C · Construye

- [ ] InsertionSort estable sobre la misma lista
- [ ] Test que verifique que MergeSort e InsertionSort dan EXACTAMENTE el mismo orden, incluidos los empates

**Decisiones de diseño y justificación:**

---

## D · Defiende (preparación, no se entrega)

1. **Define estabilidad y explica por qué importa en este problema.**
   - Mi respuesta (ensayo): 

2. **¿Qué pasaría con dos nodos si ObtenerMedio iniciara rapido := c?**
   - Mi respuesta (ensayo): 

3. **Si la lista ya estuviera ordenada por nombre, ¿qué garantiza la estabilidad al ordenar luego por nota?**
   - Mi respuesta (ensayo): 

