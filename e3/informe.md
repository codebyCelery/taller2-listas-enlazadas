# Ejercicio 3: Lista doble ordenada

**Responsable:** Leticia · **Revisora (PR):** _______

---

## A · Predice
> ⚠️ Escribe esta sección y haz **commit ANTES de ejecutar** el programa (`git log` es la evidencia).

**Predicción (antes de ejecutar):** Escribe las dos líneas de main (recorrido adelante | atrás).

```text
(escribe aquí tu predicción)
```

**Salida real (después de ejecutar):**

```text
(pega aquí la salida de la terminal)
```

**Comparación:** en qué acerté y en qué no.

---

## B · Depura

**Hipótesis del defecto** (cita la evidencia: salida observada vs. esperada):
- ¿Qué puntero queda desactualizado y en qué caso de inserción?

**Corrección mínima (diff):**

```diff
(pega aquí el diff)
```

**Test que rompe el original** (nombre del test en `main_test.go` y por qué falla con el original):

---

## C · Construye

- [ ] Validar() error (n.Siguiente == nil o n.Siguiente.Anterior == n; Cola es el último nodo)
- [ ] EliminarNodo(n *NodoDoble): cabeza, cola y nodo único

**Decisiones de diseño y justificación:**

---

## D · Defiende (preparación, no se entrega)

1. **¿Por qué el primer caso (10, 20, 30, 40, 50) funciona?**
   - Mi respuesta (ensayo): 

2. **En una inserción intermedia, ¿qué cuatro punteros cambian?**
   - Mi respuesta (ensayo): 

3. **¿Qué punteros cambian al eliminar la cabeza, un nodo intermedio y la cola?**
   - Mi respuesta (ensayo): 

