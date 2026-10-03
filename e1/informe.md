# Ejercicio 1: Lista simple con puntero Cola

**Responsable:** Madelein · **Revisora (PR):** _______

---

## A · Predice
> ⚠️ Escribe esta sección y haz **commit ANTES de ejecutar** el programa (`git log` es la evidencia).

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

## B · Depura

**Hipótesis del defecto** (cita la evidencia: salida observada vs. esperada):

**Corrección mínima (diff):**

```diff
(pega aquí el diff)
```

**Test que rompe el original** (nombre del test en `main_test.go` y por qué falla con el original):

---

## C · Construye

- [ ] Validar() error (Cola es el último nodo alcanzable, n coincide, Cabeza == nil ⇔ Cola == nil)
- [ ] EliminarUltimo() (int, error) — error si la lista está vacía
- [ ] Llamar a Validar() después de cada operación en los tests

**Decisiones de diseño y justificación:**

---

## D · Defiende (preparación, no se entrega)

1. **¿Qué invariante se rompía y en qué línea exacta?**
   - Mi respuesta (ensayo): 

2. **¿Por qué eliminar un nodo intermedio nunca revelaba el defecto?**
   - Mi respuesta (ensayo): 

3. **Después de EliminarUltimo, ¿qué nodo debe quedar como Cola y cómo lo encuentras?**
   - Mi respuesta (ensayo): 

