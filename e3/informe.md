# Ejercicio 3: Lista doble ordenada

---

## A · Predice

**Predicción (antes de ejecutar):** Escribe las dos líneas de main (recorrido adelante | atrás).

```text
Predicción (antes de ejecutar):

A: 10 20 30 40 50  | B: 50 40 30 20 10
A: 10 20 30 40 50  | B: 50 40 30 20 10
```

**Salida real (después de ejecutar):**

```text
Salida real (después de ejecutar):

A: 10 20 30 40 50  | B: 50 40 30 20 10
A: 10 20 30 40 50  | B: 50 40 10
```

**Comparación:** en qué acerté y en qué no.
```text
Mi predicción coincidió con la salida real en la primera línea, pero no en la segunda. Predije que en los dos casos la lista
quedaría ordenada y que el recorrido hacia atrás sería el inverso exacto del recorrido hacia adelante, porque la función se llama
InsertarOrdenado. En la segunda línea, el recorrido hacia adelante (A) sí salió bien, pero hacia atrás (B) salió 50 40 10: se
saltó el 30 y el 20. Como A solo usa los punteros Siguiente y B solo usa los punteros Anterior, esto indica que los Siguiente
son correctos y que el defecto está en algún puntero Anterior que no se actualiza. La primera línea sale bien porque, al insertar
en orden, cada número va al final, donde no hay un nodo siguiente que deba apuntar de vuelta al nuevo. Esto es lo que corregiré en B.
```
## B · Depura

**Hipótesis del defecto** (cita la evidencia: salida observada vs. esperada):
- ¿Qué puntero queda desactualizado y en qué caso de inserción?
 El defecto está en la inserción intermedia de InsertarOrdenado. Cuando el nuevo nodo queda entre dos nodos, el código actualiza nuevo.Siguiente, nuevo.Anterior y actual.Siguiente, pero no actualiza el puntero Anterior del nodo que queda después del nuevo. Ese nodo sigue apuntando hacia atrás a actual, como si el nuevo no existiera. Por ejemplo, al insertar el 30 entre el 10 y el 40, el 40 sigue apuntando hacia atrás al 10. Por eso el recorrido hacia adelante (Siguiente) es correcto, pero el recorrido hacia atrás (Anterior) se salta los nodos insertados en el medio. La primera línea de main sale bien porque, al insertar 10, 20, 30, 40, 50 en orden, cada nodo se agrega al final, donde no hay un nodo siguiente cuyo Anterior haya que actualizar. La corrección es agregar ese cuarto puntero cuando el nuevo nodo tiene un siguiente.

**Corrección mínima (diff):**

```diff
 	nuevo.Siguiente = actual.Siguiente
 	nuevo.Anterior = actual
 	actual.Siguiente = nuevo
 	if nuevo.Siguiente == nil {
 		l.Cola = nuevo
+	} else {
+		nuevo.Siguiente.Anterior = nuevo
 	}
```

**Test que rompe el original:** ```TestInsertarOrdenadoAtras```

Inserta 40, 10, 30, 20, 50, que obliga a hacer inserciones intermedias, y exige que Atras() devuelva "50 40 30 20 10 ". Con el código original falla, porque el Anterior del nodo siguiente al insertado no se actualiza y Atras() devuelve "50 40 10 ". No se usa el caso 10, 20, 30, 40, 50, porque ahí todas las inserciones son al final y el test pasaría incluso con el defecto. Con la corrección, el test pasa.

Evidencia con el código original:
```text
--- FAIL: TestInsertarOrdenadoAtras (0.00s)
    main_test.go:13: Atras()="50 40 10 "; se esperaba "50 40 30 20 10 "
FAIL
FAIL    taller2/e3      0.392s
```

## C · Construye

- [x] Validar() error (n.Siguiente == nil o n.Siguiente.Anterior == n; Cola es el último nodo)
- [x] EliminarNodo(n *NodoDoble): cabeza, cola y nodo único

**Decisiones de diseño y justificación:**

Validar: recorre la lista una vez (O(n)) y comprueba que cada nodo y su siguiente se apunten de ida y de vuelta. Al final verifica que Cola sea el último nodo. Si algo falla, devuelve un error.

EliminarNodo: solo actualiza los vecinos del nodo (O(1)). Si era la cabeza, Cabeza pasa al siguiente. Si era la cola, Cola pasa al anterior. Si era el único, las dos quedan en nil.

Tests: se llama a Validar() después de cada operación para asegurar que ningún puntero quedó roto.

---



