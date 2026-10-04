# Ejercicio 6: Merge Sort estable en lista

**Responsable:** Celery

---

## A · Predice

> ⚠️ Escribe esta sección y haz **commit ANTES de ejecutar** el programa (`git log` es la evidencia).

**Predicción (antes de ejecutar):**

```text
Beto(18) Dani(16) Ana(14) Caro(11) 
Eva(18) Beto(18) Fito(15) Caro(15) Ana(15) Dani(12) 
```

**Salida real (después de ejecutar):**

```text
Beto(18) Dani(16) Ana(14) Caro(11) 
Eva(18) Beto(18) Fito(15) Caro(15) Ana(15) Dani(12) 
```

**Comparación:**

Mi predicción coincidió con la salida real en las dos líneas, porque seguí el código tal cual. La primera línea es correcta según la especificación, ya que las notas son todas distintas. La segunda coincide con lo que hace el programa, pero no es correcta: `Eva(18)` aparece antes que `Beto(18)` aunque Beto estaba antes en la lista original. Eso indica que `Mezclar` no es estable cuando hay empates, y es lo que corregiré en B.

---

## B · Depura

**Hipótesis del defecto** (cita la evidencia: salida observada vs. esperada):

- Observada: `Eva(18) Beto(18) Fito(15) Caro(15) Ana(15) Dani(12)`
- Esperada: `Beto(18) Eva(18) Ana(15) Caro(15) Fito(15) Dani(12)`

El defecto está en la condición `if a.Nota > b.Nota` de `Mezclar`. Cuando las notas son iguales (por ejemplo `18 > 18`) la condición es falsa, el programa entra en el `else` y toma primero el nodo de la sublista derecha `b`, que contiene los elementos que venían después en la lista original. Así se invierte el orden de los empates y se rompe la estabilidad. La primera línea de `main` sale bien porque todas sus notas son distintas y no hay empates. La corrección es usar `>=`, de modo que ante un empate se conserve el elemento de la sublista izquierda `a`.

**Corrección mínima (diff):**

```diff
-		if a.Nota > b.Nota {
+		if a.Nota >= b.Nota {
```

**Test que rompe el original:** `TestMergeSortIgualInsertionSort`.

Ordena los mismos datos con `MergeSort` e `InsertionSort` y exige que den exactamente el mismo orden, incluidos los empates. Con el código original falla en los casos 1 y 4, porque `Mezclar` usa `>` estricto y en un empate toma primero el nodo de la derecha, invirtiendo el orden original (por ejemplo, `A B C` sale como `C B A`). `InsertionSort` es estable y conserva ese orden, por eso los resultados difieren. Con `>=` el test pasa.

Evidencia con el código original:

```text
--- FAIL: TestMergeSortIgualInsertionSort (0.00s)
    main_test.go:18: caso 1: MergeSort="Eva(18) Beto(18) Fito(15) Caro(15) Ana(15) Dani(12) " InsertionSort="Beto(18) Eva(18) Ana(15) Caro(15) Fito(15) Dani(12) "
    main_test.go:18: caso 4: MergeSort="C(5) B(5) A(5) " InsertionSort="A(5) B(5) C(5) "
FAIL
FAIL    taller2/e6      0.011s
```

Después de la corrección (`>` por `>=`):

```text
ok      taller2/e6      0.003s
```

---

## C · Construye

- [x] `InsertionSort` estable sobre la misma lista
- [x] Test que verifica que `MergeSort` e `InsertionSort` dan exactamente el mismo orden, incluidos los empates

**Decisiones de diseño y justificación:**

- **`>=` en el bucle interno:** cada nodo `c` se inserta después de todos los nodos ya ordenados cuya nota es mayor o **igual** a la suya. Los nodos con la misma nota que ya estaban en la sublista ordenada venían antes en la lista original, así que `c` queda detrás de ellos y se respeta el orden original. Con `>` quedaría delante de sus empates y se perdería la estabilidad.
- **Insertar en una sublista ordenada, solo con punteros:** se recorre la lista original en orden y se guarda `sig` antes de cambiar `c.Siguiente`, para no perder el resto de la lista. No se intercambian valores, solo se reenlazan nodos. Es más lento que Merge Sort (O(n²)), pero simple y estable, y sirve como referencia para comparar.
- **Test contra `InsertionSort`:** como `InsertionSort` es estable por construcción, sirve de referencia. El test cubre notas distintas, empates mezclados, lista vacía, un solo elemento, todas las notas iguales y una lista ya ordenada.
