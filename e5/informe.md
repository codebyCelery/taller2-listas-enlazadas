# Ejercicio 5: Planificador Round Robin

**Responsable:** Yeongmi · **Revisora (PR):** _______

---

## A · Predice

> ⚠️ Escribe esta sección y haz **commit ANTES de ejecutar** el programa (`git log` es la evidencia).

**Predicción (antes de ejecutar):** Escribe las dos líneas de main (nombre@tiempo de cada proceso).

```text
[P2@t=10 P1@t=14 P3@t=18]
[P1@t=2 P2@t=4 P3@t=9]
```

**Salida real (después de ejecutar):**

```text
[P2@t=10 P1@t=14 P3@t=18]
[P1@t=2 P2@t=4 P2@t=6]
```

**Comparación:** en qué acerté y en qué no.
Acerté en la primera línea y en el comienzo de la segunda (P1@t=2 y P2@t=4), pero no en el final: yo esperaba P3@t=9 y el programa imprimió P2@t=6, o sea P2 repetido y P3 nunca terminó. Como mi predicción seguía lo que debería hacer un Round Robin correcto, la diferencia revela un defecto en la lista circular: cuando P1 y P2 terminan seguidos, ant avanza hacia un nodo ya desenlazado y el siguiente borrado se hace en el lugar equivocado.

---

## B · Depura

**Hipótesis del defecto** (cita la evidencia: salida observada vs. esperada):

* ¿En qué nodo queda ant justo después de retirar un proceso?
  Observada: [P1@t=2 P2@t=4 P2@t=6]. Esperada: [P1@t=2 P2@t=4 P3@t=9]. Al terminar un proceso, ant avanza al nodo ya desenlazado (nodo fantasma); si el siguiente también termina, se desenlaza desde el fantasma y la lista real no cambia: P2 se repite y P3 nunca termina. En el primer caso no pasa porque entre dos finalizaciones corre un proceso que no termina y reubica ant.

**Corrección mínima (diff):**

```diff
-        ant = act
+        if act.Tiempo != 0 {
+            ant = act
+        }
```

**Test que rompe el original** (nombre del test en `main_test.go` y por qué falla con el original):
TestDosTerminanSeguidos: con P1:2, P2:2, P3:5 y quantum 2, exige [P1@t=2 P2@t=4 P3@t=9]. El original falla porque da [P1@t=2 P2@t=4 P2@t=6]: P1 y P2 terminan seguidos y ant queda en un nodo fantasma. No usa el primer caso de main porque ahí pasa aunque exista el defecto.

---

## C · Construye

* [ ] Soportar llegadas: (nombre, llegada, ráfaga); entra a la cola cuando reloj ≥ llegada
* [ ] Decidir y justificar dónde se inserta un proceso que llega respecto a act
* [ ] Retornar también el tiempo de espera promedio

**Decisiones de diseño y justificación:**

---

## D · Defiende (preparación, no se entrega)

1. **¿Qué es un “nodo fantasma” en el código original y cómo vuelve a ejecutarse?**

   * Mi respuesta (ensayo):
2. **¿Por qué el primer caso da el resultado correcto?**

   * Mi respuesta (ensayo):
3. **¿Cómo cambiaría la salida con quantum = 1? ¿Y con un quantum mayor que todas las ráfagas?**

   * Mi respuesta (ensayo):

