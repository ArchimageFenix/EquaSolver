# Análisis realizado

Registro del análisis previo al código, hecho en conjunto antes de escribir ARCHITECTURE.md.

## 1. Punto de partida

Programa que resuelve sistemas de n ecuaciones con el método de Gauss. Debía ser CLI, con diseño top-down y adaptable a web. En esa etapa aún no se había elegido lenguaje.

## 2. Funciones identificadas (top-down)

1. Leer el sistema (n, coeficientes, términos independientes).
2. Validar la entrada.
3. Eliminación hacia adelante con pivoteo parcial.
4. Clasificar el resultado.
5. Sustitución hacia atrás.
6. Presentar el resultado.

## 3. Entradas y salidas

- Entrada: n, matriz de coeficientes (n × n) y vector de términos independientes.
- Salida: matriz original, pasos de la eliminación, matriz escalonada, caso y solución o informe.

## 4. Decisiones tomadas

| Tema | Decisión |
|---|---|
| Sistemas admitidos | Solo cuadrados (n ecuaciones, n incógnitas) |
| Casos | Solución única, infinitas soluciones y sin solución son resultados normales, no errores |
| Infinitas soluciones | Solo se informa (rangos y variables libres); se muestra la matriz final |
| Sin solución | Se informa y se señala la fila contradictoria |
| Pasos | Se muestran siempre |
| Tamaño de n | Sin límite máximo propio; entero mayor o igual que 1 |
| Precisión | Cálculo en doble precisión; presentación con máximo dos decimales |
| Tolerancia | Relativa: n × ε × máx\|aᵢⱼ\| |
| Validación | En capas: página web, adaptador (servidor / CLI) y núcleo; ningún dato inválido llega al núcleo |
| CLI | No verbosa: solo pide y muestra; menú de forma de ingreso; ayuda breve en cada paso |
| Formas de entrada | Interactiva, archivo y argumentos, todas admitidas |
| Web | Mismo núcleo; mensaje de valor no válido; datos inválidos nunca llegan al núcleo |
| Lenguaje | Go, solo biblioteca estándar |

## 5. Alternativas descartadas

- **Java:** exige la JVM y, para la web, un servidor o framework adicional.
- **C / C++:** no traen servidor web (necesitarían una biblioteca externa) y la memoria dinámica para n arbitrario se gestiona a mano.

## 6. Ajustes durante la construcción

- Los pasos se registran **por columna de pivote** (intercambio, operaciones y matriz resultante) y no por cada operación entre filas, para que la memoria crezca como n³ y no como n⁴.
- Se añadieron `numeric` (comprobación de finitud y lectura estricta de números), `format/text.go` (textos comunes a CLI y web), `cli/run.go`, `cli/prompt.go` y `cli/parse.go`, para evitar duplicación y mantener una sola responsabilidad por archivo.
- Los valores de magnitud 1e15 o mayor se muestran en notación científica, solo por legibilidad.
- Se incluyó un cuarto resultado, fallo numérico controlado, para valores no finitos.
