# EquaSolver

**Resolutor de sistemas cuadrados de ecuaciones lineales desarrollado en Go**

EquaSolver es una aplicación desarrollada en **Go** para resolver sistemas cuadrados de **n ecuaciones lineales con n incógnitas** mediante el método de eliminación de Gauss.

El programa transforma el sistema en una matriz aumentada, ejecuta el proceso de eliminación hacia adelante con **pivoteo parcial**, obtiene una matriz escalonada y utiliza **sustitución hacia atrás** para calcular las incógnitas.

EquaSolver dispone de dos interfaces —**CLI y Web**— construidas sobre el mismo núcleo matemático.


<p align="center">
  <img src="assets/logo.png" alt="EquaSolver Logo" width="500">
</p>

---

## ¿Qué resuelve?

EquaSolver trabaja exclusivamente con sistemas cuadrados:

```text id="s1wwh0"
n ecuaciones
n incógnitas
```

con:

```text id="3oq8v2"
n ≥ 1
```

De forma general:

```text id="jq2vb9"
a₁₁x₁ + a₁₂x₂ + ... + a₁ₙxₙ = b₁
a₂₁x₁ + a₂₂x₂ + ... + a₂ₙxₙ = b₂
                  ...
aₙ₁x₁ + aₙ₂x₂ + ... + aₙₙxₙ = bₙ
```

Por ejemplo, un sistema de `3 × 3`:

```text id="rqrwn9"
 2x +  y -  z =   8
-3x -  y + 2z = -11
-2x +  y + 2z =  -3
```

se representa mediante su matriz aumentada:

```text id="1jv16m"
[  2   1  -1 |   8 ]
[ -3  -1   2 | -11 ]
[ -2   1   2 |  -3 ]
```

El núcleo utiliza esta representación para realizar el proceso de resolución.

---

## Alcance actual

EquaSolver está diseñado para resolver **sistemas cuadrados completos que admiten una solución mediante el procedimiento implementado**.

Actualmente el programa no está orientado a resolver o clasificar:

* Sistemas incompletos.
* Sistemas inconsistentes.
* Sistemas con infinitas soluciones.
* Sistemas rectangulares de `m × n`.

Estas situaciones se encuentran fuera del alcance actual del resolutor.

El objetivo es mantener claramente delimitado el dominio matemático que el programa acepta antes de ampliar sus capacidades.

---

## Método de resolución

El núcleo utiliza **eliminación de Gauss con pivoteo parcial**.

El flujo matemático es:

```text id="zjokhg"
Sistema n × n
      │
      ▼
Matriz aumentada
      │
      ▼
Selección del pivote
      │
      ▼
Pivoteo parcial
      │
      ▼
Eliminación hacia adelante
      │
      ▼
Matriz escalonada
      │
      ▼
Sustitución hacia atrás
      │
      ▼
Valores de las incógnitas
```

### Eliminación hacia adelante

Las operaciones elementales transforman progresivamente la matriz aumentada hasta obtener una estructura escalonada.

### Pivoteo parcial

Durante la eliminación, el algoritmo puede intercambiar filas para seleccionar un pivote adecuado.

Esto evita depender únicamente del elemento que originalmente ocupa la posición del pivote y mejora el comportamiento numérico del proceso.

### Sustitución hacia atrás

Una vez obtenida la matriz escalonada, el sistema se recorre desde la última ecuación hacia la primera para calcular progresivamente las incógnitas.

---

## Sistemas de tamaño n × n

El algoritmo no está limitado específicamente a sistemas de `2 × 2` o `3 × 3`.

La estructura general permite trabajar con sistemas cuadrados de diferentes dimensiones:

```text id="79njpj"
1 × 1
2 × 2
3 × 3
4 × 4
...
n × n
```

Esto significa que no existe un algoritmo independiente para cada tamaño.

El mismo núcleo de Gauss procesa el sistema utilizando su dimensión `n`.

---

## El proceso es visible

EquaSolver no está pensado únicamente para recibir un sistema y devolver los valores finales.

También permite observar el procedimiento utilizado para llegar a ellos:

```text id="q6pcmx"
Sistema original
       │
       ▼
Matriz aumentada
       │
       ▼
Pasos de eliminación
       │
       ▼
Matriz escalonada
       │
       ▼
Sustitución hacia atrás
       │
       ▼
Solución
```

Esto permite utilizar EquaSolver no solamente como herramienta de cálculo, sino también como apoyo para comprender el funcionamiento del método de Gauss.

---

## CLI y Web

EquaSolver ofrece dos formas de utilizar el resolutor.

### CLI

La interfaz de línea de comandos permite introducir y resolver sistemas directamente desde una terminal.

```text id="d2ixcf"
Usuario
  │
  ▼
 CLI
  │
  ▼
Núcleo Gauss
  │
  ▼
Resultado
```

### Web

La interfaz Web permite utilizar el mismo resolutor desde un navegador mediante una presentación gráfica.

```text id="sq1yt5"
Usuario
  │
  ▼
 Web
  │
  ▼
Núcleo Gauss
  │
  ▼
Resultado
```

Las interfaces **no contienen implementaciones independientes del algoritmo matemático**.

Ambas utilizan el mismo núcleo:

```text id="mv7uzd"
        ┌─────────┐
        │   CLI   │
        └────┬────┘
             │
             ▼
      ┌──────────────┐
      │    Núcleo    │
      │    Gauss     │
      └──────────────┘
             ▲
             │
        ┌────┴────┐
        │   Web   │
        └─────────┘
```

De esta manera, una modificación o corrección en el algoritmo puede beneficiar a ambas interfaces sin duplicar la lógica matemática.

---

## Características

* Desarrollado en **Go**.
* Resolución de sistemas cuadrados `n × n`.
* Número de ecuaciones igual al número de incógnitas.
* Eliminación de Gauss.
* Pivoteo parcial.
* Sustitución hacia atrás.
* Representación mediante matriz aumentada.
* Visualización de la matriz original.
* Visualización del proceso de eliminación.
* Visualización de la matriz escalonada.
* Presentación de las incógnitas calculadas.
* Interfaz **CLI**.
* Interfaz **Web**.
* Núcleo matemático compartido.

---

## Tecnología

| Elemento            | Implementación                  |
| ------------------- | ------------------------------- |
| Lenguaje            | Go                              |
| Dominio             | Álgebra lineal                  |
| Problema            | Sistemas de ecuaciones lineales |
| Sistemas admitidos  | Cuadrados                       |
| Dimensión           | n × n                           |
| Algoritmo           | Eliminación de Gauss            |
| Selección de pivote | Pivoteo parcial                 |
| Resolución final    | Sustitución hacia atrás         |
| Interfaces          | CLI + Web                       |
| Núcleo              | Compartido                      |

---

## Arquitectura

EquaSolver mantiene una separación entre:

```text id="4kmqjt"
Entrada / Interfaces
        │
        ▼
Aplicación
        │
        ▼
Núcleo matemático
        │
        ▼
Resultados
        │
        ▼
Presentación
```

Esta separación permite mantener el algoritmo matemático como la pieza central del programa mientras CLI y Web funcionan como formas diferentes de interactuar con él.

La documentación técnica completa del diseño se encuentra en:

**[`ARCHITECTURE.md`](ARCHITECTURE.md)**

Allí se documentan la estructura de carpetas y archivos, responsabilidades de los módulos, flujo de datos, dependencias y reglas de ampliación del proyecto.

---

## Estado del proyecto

**En desarrollo.**

El alcance actual se concentra en resolver correctamente sistemas cuadrados mediante eliminación de Gauss antes de incorporar otros tipos de sistemas o métodos de resolución.

---

**EquaSolver — Solve linear systems with Gauss.**
