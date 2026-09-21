<p align="center">
  <img src="assets/logo.png" alt="EquaSolver Logo" width="500">
</p>

# EquaSolver

**Resolutor de sistemas de ecuaciones lineales mediante eliminación de Gauss**

EquaSolver es una aplicación desarrollada en **Go** para resolver y analizar sistemas cuadrados de **n ecuaciones lineales con n incógnitas**.

Utiliza **eliminación de Gauss con pivoteo parcial** para transformar la matriz aumentada del sistema y determinar su naturaleza. Cuando existe una solución única, aplica **sustitución hacia atrás** para calcular las incógnitas.

El programa no se limita a entregar el resultado final: muestra la matriz original, las operaciones realizadas durante la eliminación, la matriz escalonada y el resultado del análisis.

Dispone de interfaces **CLI y Web**, ambas construidas sobre el mismo núcleo matemático.

---

## ¿Qué puede resolver?

EquaSolver trabaja con sistemas cuadrados:

```text
n ecuaciones
n incógnitas
n ≥ 1
```

De forma general:

```text
a₁₁x₁ + a₁₂x₂ + ... + a₁ₙxₙ = b₁
a₂₁x₁ + a₂₂x₂ + ... + a₂ₙxₙ = b₂
                  ...
aₙ₁x₁ + aₙ₂x₂ + ... + aₙₙxₙ = bₙ
```

Por ejemplo:

```text
 2x + 3y -  z + 2w =  12
 4x -  y + 2z +  w =  11
-2x + 5y + 3z -  w =   7
 3x + 2y + 4z + 3w =  25
```

se representa mediante la matriz aumentada:

```text
[  2   3  -1   2 | 12 ]
[  4  -1   2   1 | 11 ]
[ -2   5   3  -1 |  7 ]
[  3   2   4   3 | 25 ]
```

No existe un algoritmo diferente para `2×2`, `3×3`, `4×4`, etc. El núcleo trabaja de forma general con la dimensión `n`.

---

## Clasificación del sistema

Que un sistema sea cuadrado no significa necesariamente que posea una única solución.

EquaSolver analiza los rangos de la matriz de coeficientes y de la matriz ampliada para distinguir los siguientes casos:

| Resultado                | Condición                                        |
| ------------------------ | ------------------------------------------------ |
| **Solución única**       | `r = r' = n`                                     |
| **Infinitas soluciones** | `r = r' < n`                                     |
| **Sin solución**         | `r < r'`                                         |
| **Fallo numérico**       | Se detecta un valor no finito durante el cálculo |

donde:

* `r` es el rango de la matriz de coeficientes.
* `r'` es el rango de la matriz ampliada.
* `n` es el número de incógnitas.

Los casos de infinitas soluciones o ausencia de solución se consideran **resultados matemáticos del análisis**, no errores de ejecución.

Cuando existe una contradicción, una fila equivalente a:

```text
0  0  0  ...  0 | c
```

con `c ≠ 0` permite identificar un sistema sin solución.

---

## Método de resolución

El proceso general seguido por EquaSolver es:

```text
Sistema n × n
      │
      ▼
Validación
      │
      ▼
Matriz aumentada
      │
      ▼
Cálculo de tolerancia
      │
      ▼
Eliminación de Gauss
      │
      ├── Selección de pivote
      ├── Pivoteo parcial
      └── Eliminación hacia adelante
      │
      ▼
Matriz escalonada
      │
      ▼
Clasificación
      │
      ├── Solución única
      │        │
      │        ▼
      │  Sustitución hacia atrás
      │        │
      │        ▼
      │     Solución
      │
      ├── Infinitas soluciones
      │
      └── Sin solución
```

### Pivoteo parcial

En cada columna se selecciona como pivote la fila disponible cuyo elemento tenga el **mayor valor absoluto**.

Si es necesario, las filas son intercambiadas antes de continuar con la eliminación.

Por ejemplo:

```text
2
4   ← mayor valor absoluto
-2
3
```

provoca que la fila que contiene `4` sea utilizada como primera fila pivote.

Esta estrategia mejora el comportamiento numérico respecto a utilizar siempre el elemento que originalmente ocupa la posición del pivote.

### Eliminación hacia adelante

Una vez seleccionado el pivote, se aplican operaciones elementales para eliminar los coeficientes situados debajo de él.

El proceso produce progresivamente una matriz escalonada de forma similar a:

```text
[a  b  c  d | e]
[0  f  g  h | i]
[0  0  j  k | l]
[0  0  0  m | n]
```

### Sustitución hacia atrás

Cuando el sistema posee solución única, EquaSolver comienza por la última ecuación y calcula las incógnitas en sentido inverso:

```text
xₙ → xₙ₋₁ → ... → x₂ → x₁
```

La sustitución hacia atrás solo se ejecuta cuando la clasificación determina que existe una solución única.

---

## El procedimiento es visible

EquaSolver está diseñado para mostrar **cómo se obtiene el resultado**, no únicamente el resultado final.

Una resolución incluye:

```text
Matriz original
      │
      ▼
Paso 1
  operaciones entre filas
  matriz resultante
      │
      ▼
Paso 2
  operaciones entre filas
  matriz resultante
      │
      ▼
     ...
      │
      ▼
Matriz escalonada
      │
      ▼
Clasificación
      │
      ▼
Resultado
```

Por ejemplo, durante un paso puede mostrarse:

```text
Paso 1 (columna 1):
intercambiar fila 1 con fila 2 (pivote 4.00)

fila 2 = fila 2 - 0.50 × fila 1
fila 3 = fila 3 + 0.50 × fila 1
fila 4 = fila 4 - 0.75 × fila 1
```

Esto permite utilizar EquaSolver tanto como herramienta de cálculo como para observar y estudiar el procedimiento de eliminación de Gauss.

---

## CLI y Web

EquaSolver dispone de **dos interfaces independientes sobre un mismo núcleo**.

```text
             ┌─────────────┐
             │     CLI     │
             └──────┬──────┘
                    │
                    ▼
             model.System
                    │
                    ▼
       ┌────────────────────────┐
       │    Núcleo EquaSolver   │
       │                        │
       │ validation             │
       │ tolerance              │
       │ elimination            │
       │ classification         │
       │ substitution           │
       │ solver                 │
       └────────────┬───────────┘
                    │
                    ▼
             model.Result
                    │
             ┌──────┴──────┐
             ▼             ▼
            CLI           Web
```

CLI y Web **no implementan sus propios algoritmos de Gauss**.

Ambas entregan los datos al mismo núcleo y reciben de él una estructura de resultado independiente de la interfaz.

Esto evita duplicar la lógica matemática y permite que una corrección o mejora del núcleo sea utilizada automáticamente por ambas interfaces.

---

## Interfaz CLI

La interfaz de línea de comandos permite utilizar EquaSolver directamente desde una terminal.

Admite diferentes formas de proporcionar un sistema:

### Entrada interactiva

El programa solicita el tamaño y posteriormente cada ecuación:

```text
Forma de ingreso:
  1. Escribir los datos
  2. Leer desde archivo

Opción:

Número de ecuaciones (entero mayor que 0):

Ecuación 1
(n coeficientes y término independiente, separados por espacio):
```

### Desde archivo

Puede utilizarse:

```text
-f <archivo>
```

para cargar un sistema almacenado en un archivo de texto.

### Desde argumentos

También puede proporcionarse directamente una matriz mediante:

```text
-m "<filas separadas por ;>"
```

Cuando se utiliza `-f` o `-m`, el menú interactivo puede omitirse.

---

## Interfaz Web

La interfaz Web proporciona acceso al mismo resolutor desde un navegador.

El flujo es:

```text
Navegador
    │
    ▼
Formulario
    │
    ▼
Validación Web
    │
    ▼
Servidor
    │
    ▼
Núcleo EquaSolver
    │
    ▼
Resultado
    │
    ▼
Presentación Web
```

Los datos son validados antes de llegar al núcleo. Si la entrada es incorrecta, el usuario recibe un mensaje y el cálculo no se inicia.

La Web no contiene lógica matemática propia: utiliza exactamente el mismo `solver` que la CLI.

---

## Precisión y estabilidad numérica

Los cálculos internos utilizan **doble precisión de 64 bits**.

Los valores no se redondean durante la resolución. El redondeo se realiza únicamente al presentar los resultados.

EquaSolver utiliza una tolerancia relativa basada en:

```text
tolerancia = n × ε × máx|aᵢⱼ|
```

donde `ε` representa el épsilon de máquina para doble precisión.

Un valor suficientemente pequeño respecto a esta tolerancia puede considerarse cero durante el análisis.

Además, después de las operaciones numéricas se controla que los valores permanezcan finitos. La aparición de `NaN` o infinito detiene el cálculo de forma controlada y produce un resultado de **fallo numérico**.

---

## Validación

Los datos pasan por diferentes niveles de validación antes del cálculo.

Entre otras condiciones, se verifica que:

* `n ≥ 1`.
* La matriz de coeficientes sea `n × n`.
* Existan exactamente `n` términos independientes.
* Los valores sean números reales finitos.
* No se introduzcan `NaN` ni infinito.
* La entrada tenga una estructura válida.

El núcleo mantiene además su propia validación como última barrera, independientemente de que los datos procedan de CLI o Web.

---

## Características principales

* Desarrollado en **Go**.
* Solo utiliza la **biblioteca estándar de Go**.
* Sistemas cuadrados de `n` ecuaciones y `n` incógnitas.
* Sin límite máximo artificial para `n` en el núcleo.
* Eliminación de Gauss.
* Pivoteo parcial.
* Sustitución hacia atrás.
* Cálculo de rangos.
* Detección de solución única.
* Detección de infinitas soluciones.
* Detección de sistemas inconsistentes.
* Control de valores numéricos no finitos.
* Visualización de la matriz original.
* Visualización paso a paso de la eliminación.
* Visualización de la matriz escalonada.
* Interfaz **CLI**.
* Interfaz **Web**.
* Entrada CLI interactiva, mediante archivo o argumentos.
* Núcleo matemático compartido entre interfaces.

---

## Arquitectura

El principio fundamental de EquaSolver es mantener el **núcleo matemático independiente de la forma en que el usuario interactúa con él**.

```text
Interfaces
    │
    ▼
Adaptadores
    │
    ▼
   Model
    │
    ▼
  Solver
    │
    ├── Validation
    ├── Tolerance
    ├── Elimination
    ├── Classification
    └── Substitution
    │
    ▼
  Result
    │
    ▼
Presentación
```

El núcleo:

* No conoce HTTP.
* No conoce la consola.
* No imprime resultados.
* No lee archivos.
* No contiene HTML.
* No depende de CLI ni Web.

Las interfaces dependen del núcleo, nunca al contrario.

La arquitectura completa, responsabilidades de cada módulo, flujo de datos y reglas de ampliación están documentadas en:

**[`ARCHITECTURE.md`](ARCHITECTURE.md)**

---

## Estructura general

```text
EquaSolver/
├── ARCHITECTURE.md
├── README.md
├── go.mod
├── assets/
│   └── logo.png
├── cmd/
│   ├── cli/
│   └── web/
├── internal/
│   ├── numeric/
│   ├── model/
│   ├── validation/
│   ├── tolerance/
│   ├── elimination/
│   ├── classification/
│   ├── substitution/
│   ├── solver/
│   ├── format/
│   ├── cli/
│   └── web/
└── docs/
```

Las pruebas unitarias se mantienen junto a los paquetes correspondientes siguiendo las convenciones habituales de Go.

---

## Limitaciones conocidas

EquaSolver trabaja actualmente con **sistemas cuadrados**: el número de ecuaciones debe coincidir con el número de incógnitas.

Para poder mostrar el procedimiento completo, el programa conserva la matriz correspondiente a cada columna de pivote. Por esta razón, el consumo de memoria aumenta significativamente para valores muy grandes de `n`.

Como ocurre con los métodos numéricos que utilizan aritmética de punto flotante, la tolerancia tampoco puede resolver de forma perfecta todos los sistemas casi singulares. Por ello EquaSolver conserva y muestra los pasos y la matriz final del proceso.

---

## Documentación

La documentación técnica del proyecto se encuentra en:

**[`ARCHITECTURE.md`](ARCHITECTURE.md)** — arquitectura, responsabilidades, dependencias y reglas del proyecto.

**`docs/analisis.md`** — análisis realizado durante el desarrollo.

**`docs/capacidades.md`** — capacidades actuales y formas de utilización.

---

## Estado

**En desarrollo.**

El proyecto continúa evolucionando manteniendo una separación estricta entre el núcleo matemático y sus interfaces.

---

## Autor

Desarrollado por **ArchimageFenix**.

**EquaSolver — Solve any system.**
