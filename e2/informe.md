# Ejercicio 2: Eliminar todas las apariciones

**Responsable:** Madelein · **Revisora (PR):** _______

---

## A · Predice
>  Escribe esta sección y haz **commit ANTES de ejecutar** el programa (`git log` es la evidencia).

**Predicción (antes de ejecutar):** Escribe las dos líneas que imprime main (cantidad borrada y lista resultante).

```text
3 3 -> 1 -> 5 -> nil
5 2 -> 4 -> nil
```

**Salida real (después de ejecutar):**

```text
3 3 -> 1 -> 5 -> nil
4 2 -> 7 -> 4 -> nil
```

**Comparación:** en qué acerté y en qué no.

La primera sale correcta la segunda también parece correcta, porque el código actual funciona para esos casos
---

## B · Depura

**Hipótesis del defecto** (cita la evidencia: salida observada vs. esperada):

El defecto aparece cuando hay valores repetidos consecutivamente después de la cabeza porque después de eliminar un nodo ant avanza inmediatamente y puede saltarse otro nodo con el mismo valor

**Corrección mínima (diff):**

```diff
en esta parte:

         if ant.Siguiente.Valor == v {
             ant.Siguiente = ant.Siguiente.Siguiente
             borrados++
         }
         
         ant = ant.Siguiente
se cambia por:
         if ant.Siguiente.Valor == v {
             ant.Siguiente = ant.Siguiente.Siguiente
             borrados++
         } else {
             ant = ant.Siguiente
         }
```

**Test que rompe el original** (nombre del test en `main_test.go` y por qué falla con el original):
---
      func TestEliminarTodosConsecutivos(t *testing.T) {
          l := Desde(1, 7, 7, 7, 2)
      
          borrados := l.EliminarTodos(7)
      
          if borrados != 3 {
              t.Errorf("esperaba 3 borrados, obtuvo %d", borrados)
          }
      
          esperado := "1 -> 2 -> nil"
          if l.String() != esperado {
              t.Errorf("esperaba %s, obtuvo %s", esperado, l.String())
          }
      }
en el primer caso de main no sirve porque los 7 estan separados, no hay dos 7 juntos 
---

## C · Construye

- [ A] EliminarDuplicados(): deja solo la primera aparición de cada valor, conservando el orden

  ---
        func (l *Lista) EliminarDuplicados() {
          vistos := make(map[int]bool)
      
          for l.Cabeza != nil && vistos[l.Cabeza.Valor] {
              l.Cabeza = l.Cabeza.Siguiente
          }
      
          if l.Cabeza == nil {
              return
          }
      
          vistos[l.Cabeza.Valor] = true
          ant := l.Cabeza
      
          for ant.Siguiente != nil {
              if vistos[ant.Siguiente.Valor] {
                  ant.Siguiente = ant.Siguiente.Siguiente
              } else {
                  vistos[ant.Siguiente.Valor] = true
                  ant = ant.Siguiente
              }
          }
---
- [ B ] Pruebas: lista vacía, todos iguales, duplicados consecutivos, duplicados no consecutivos
- lista vacia 
--- 
      l := Desde()
      l.EliminarDuplicados()
      // nil
---
- todos iguales
---
      l := Desde(5, 5, 5, 5)
      l.EliminarDuplicados()
      // 5 -> nil
---
- duplicados consecutivos
---
      l := Desde(1, 2, 2, 2, 3)
      l.EliminarDuplicados()
      // 1 -> 2 -> 3 -> nil
---
- duplicados no consecutivos
---
      l := Desde(1, 2, 3, 2, 4, 1)
      l.EliminarDuplicados()
      // 1 -> 2 -> 3 -> 4 -> nil
    
---

## D · Defiende (preparación, no se entrega)

1. **¿Qué representa ant y en qué momento exacto debe avanzar?**
   - "ant" representa el nodo anterior al que estamos revisando solo avanza cuando no elimina un nodo

2. **¿Por qué el primer caso da la respuesta correcta?**
   - Porque los 7 no están juntos después de eliminar un 7, el siguiente nodo no es otro 7 así que avanzar "ant" no se salta ninguno

3. **¿Qué casos de prueba mínimos necesitas para confiar en EliminarTodos? Justifica cada uno.**
- Lista vacía: comprueba que no falle
- Valor inexistente: comprueba que no cambie la lista
- Un solo valor: comprueba una eliminación simple
- Valor al inicio: comprueba la eliminación de la cabeza
- Valor al final: comprueba la eliminación del último
- Varios consecutivos: comprueba que no se salte nodos
- Varios separados: comprueba eliminaciones en diferentes posiciones

