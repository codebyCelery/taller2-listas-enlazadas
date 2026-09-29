# Guía de trabajo en equipo (Git + GitHub)

## Una sola vez (Celery)
1. En GitHub: New repository → `taller2-listas-enlazadas` → **Private** → sin README.
2. Settings → Collaborators → invitar a Madelein, Leticia y Yeongmi.
3. Subir esta carpeta:
```bash
git init -b main
git add .
git commit -m "Estructura inicial con codigo original de los 6 ejercicios"
git remote add origin https://github.com/USUARIO/taller2-listas-enlazadas.git
git push -u origin main
```
4. (Recomendado) Settings → Branches → proteger `main` exigiendo Pull Request.

## Cada integrante (una vez)
```bash
git clone https://github.com/USUARIO/taller2-listas-enlazadas.git
cd taller2-listas-enlazadas
git checkout -b nombre-eX      # ej: madelein-e1-e2
```
Instalar Go (https://go.dev/dl) y verificar con `go version`.

## Por cada ejercicio, EN ESTE ORDEN
1. **A:** escribe tu predicción en `eX/informe.md` → `git add` → `git commit -m "eX: prediccion (antes de ejecutar)"` → `git push`.
2. Recién ahora ejecuta `go run ./eX`. Pega la salida en el informe y compara.
3. **B:** hipótesis, corrección mínima, test que falla con el original.
   Commit del test primero (falla) y luego del arreglo (pasa): `eX: test que rompe el original` / `eX: correccion minima`.
4. **C:** implementa, prueba y justifica. Commits pequeños.
5. Antes de abrir el PR: `gofmt -l .`, `go vet ./...`, `go test ./...`.
6. Abre un Pull Request hacia `main`; otra compañera lo revisa y le hace las preguntas D.

## Para no perder puntos
- El commit de la predicción (A) debe ir **antes** de ejecutar.
- Anota tu uso de IA en `bitacora_ia.md` (todo, salvo el código de las estructuras).
- Todas deben poder defender los 6 ejercicios: se sortea a una persona y su nota vale para el equipo.
