# ARCHITECTURE.md

# Gauss — Resolutor de sistemas de ecuaciones lineales

Lenguaje: **Go** · Interfaces: **CLI** y **Web** sobre un mismo núcleo · Entrega: **zip**

---

## 1. Propósito

Resolver sistemas de n ecuaciones lineales con n incógnitas mediante el método de Gauss (eliminación hacia adelante con pivoteo parcial y sustitución hacia atrás), mostrando al usuario el proceso completo: matriz original, pasos de la eliminación, matriz escalonada y resultado.

## 2. Problema que resuelve

Resolver un sistema a mano es lento y propenso a errores, y muchas herramientas fallan o dan resultados engañosos con sistemas singulares, inconsistentes o numéricamente delicados.

Este programa:

1. Acepta sistemas de cualquier tamaño n ≥ 1.
2. Clasifica correctamente el sistema: solución única, infinitas soluciones o sin solución.
3. Trata esos tres casos como **resultados normales**, no como errores.
4. Detecta y reporta de forma controlada cualquier fallo numérico (valores no finitos), sin caerse.
5. Muestra el proceso para que el usuario entienda por qué se obtuvo cada resultado.

---

## 3. Decisiones cerradas

| Tema | Decisión |
|---|---|
| Tipo de sistema | Solo cuadrados: n ecuaciones, n incógnitas |
| Tamaño n | Sin límite máximo propio; debe ser entero ≥ 1 |
| Infinitas soluciones | Solo se informa; se muestra la matriz final para que se vea el motivo |
| Sin solución | Se informa y se señala la fila contradictoria |
| Pasos de la eliminación | Siempre se muestran |
| Precisión | Cálculo en doble precisión (64 bits); presentación con máximo 2 decimales |
| Tolerancia | Relativa: `n × ε × máx\|aᵢⱼ\|` |
| Entrada | Menú de forma de ingreso; también archivo y argumentos; ayuda en cada paso |
| CLI | No verbosa: solo pide datos y muestra resultados |
| Web | Mismo núcleo; validación previa con mensaje; datos inválidos nunca llegan al núcleo |
| Lenguaje | Go, solo biblioteca estándar |

---

## 4. Casos de resultado

Sea `r` el rango de la matriz de coeficientes y `r'` el de la matriz ampliada.

| Caso | Condición | Variables libres |
|---|---|---|
| Solución única | r = r' = n | 0 |
| Infinitas soluciones | r = r' < n | n − r |
| Sin solución (inconsistente) | r < r' | no aplica |
| Fallo numérico controlado | aparece un valor no finito durante el cálculo | no aplica |

Un sistema es inconsistente cuando la matriz escalonada contiene una fila `0 … 0 | c` con `|c|` mayor que la tolerancia. Esa fila es la **fila contradictoria** y debe señalarse.

---

## 5. Entradas y salidas

### 5.1 Entrada (estructura neutra)

```text
Sistema
├── n   : entero ≥ 1
├── A   : matriz n × n de coeficientes
└── b   : vector de n términos independientes
```

El núcleo recibe siempre esta estructura, sin importar si el dato vino de la CLI o de la web.

### 5.2 Salida (estructura neutra)

```text
Resultado
├── caso                  : única | infinitas | sinSolucion | falloNumerico
├── matrizOriginal        : matriz ampliada de entrada
├── pasos                 : un paso por columna de pivote (intercambio, operaciones entre filas y matriz resultante)
├── matrizEscalonada      : matriz ampliada tras la eliminación
├── rangoCoeficientes     : r
├── rangoAmpliada         : r'
├── variablesLibres       : n − r (solo si aplica)
├── solucion              : vector de n valores (solo si caso = única)
└── filaContradictoria    : índice de la fila (solo si caso = sinSolucion)
```

### 5.3 Presentación

Reglas comunes a CLI y web:

- Máximo 2 decimales; el redondeo ocurre solo al mostrar.
- Los valores de magnitud 1e15 o mayor se muestran en notación científica, solo para mantener la legibilidad.
- Un `-0.00` se muestra como `0.00`.
- Cuando el pivoteo reordena filas, la presentación lo indica en el paso correspondiente.

Ejemplo de salida CLI (solución única):

```text
Matriz original:
  [  1.00   1.00   1.00 |  6.00 ]
  [  2.00  -1.00   1.00 |  3.00 ]
  [  1.00   2.00  -1.00 |  2.00 ]

Paso 1 (columna 1): intercambiar fila 1 con fila 2 (pivote 2.00)
  fila 2 = fila 2 - 0.50 × fila 1
  fila 3 = fila 3 - 0.50 × fila 1
  [  2.00  -1.00   1.00 |  3.00 ]
  [  0.00   1.50   0.50 |  4.50 ]
  [  0.00   2.50  -1.50 |  0.50 ]

Paso 2 (columna 2): intercambiar fila 2 con fila 3 (pivote 2.50)
  fila 3 = fila 3 - 0.60 × fila 2
  [  2.00  -1.00   1.00 |  3.00 ]
  [  0.00   2.50  -1.50 |  0.50 ]
  [  0.00   0.00   1.40 |  4.20 ]

Paso 3 (columna 3): pivote 1.40 en fila 3
  [  2.00  -1.00   1.00 |  3.00 ]
  [  0.00   2.50  -1.50 |  0.50 ]
  [  0.00   0.00   1.40 |  4.20 ]

Matriz escalonada:
  [  2.00  -1.00   1.00 |  3.00 ]
  [  0.00   2.50  -1.50 |  0.50 ]
  [  0.00   0.00   1.40 |  4.20 ]

Resultado: SOLUCIÓN ÚNICA
Rango coeficientes: 3 | Rango ampliada: 3 | Variables libres: 0
  x1 = 1.00
  x2 = 2.00
  x3 = 3.00
```

Ejemplos de la línea de resultado en los demás casos:

```text
Resultado: INFINITAS SOLUCIONES
Rango coeficientes: 2 | Rango ampliada: 2 | Variables libres: 1

Resultado: SIN SOLUCIÓN
Rango coeficientes: 2 | Rango ampliada: 3
Fila contradictoria: 3  (0 0 0 | -1.00)
```

---

## 6. Validación en capas

Ningún dato inválido debe iniciar el cálculo por ningún camino.

```text
Dato del usuario
      ↓
[1] Interfaz (CLI: al leer / Web: en la página)
      mensaje inmediato de valor no válido
      ↓
[2] Adaptador (CLI: parseo / Web: servidor, formcheck)
      rechaza antes de llamar al núcleo
      ↓
[3] Núcleo (validation)
      última defensa, obligatoria en toda entrada
      ↓
    Cálculo
```

- La capa [1] mejora la experiencia, pero **no es de confianza**: en web se puede saltar.
- La capa [2] garantiza que un dato inválido no llega al núcleo.
- La capa [3] siempre se ejecuta; en la CLI es la barrera final.

Reglas de validación:

- n entero ≥ 1.
- A es n × n y b tiene n elementos.
- Todos los valores son números reales finitos (sin letras, símbolos, NaN ni infinito).
- El mensaje de error indica qué se esperaba y qué dato falló.

---

## 7. Precisión, tolerancia y estabilidad numérica

- Los cálculos usan doble precisión (64 bits); no se redondea nada antes de presentar.
- **Pivoteo parcial:** en cada columna se elige la fila con mayor valor absoluto.
- **Tolerancia relativa:**

```text
tolerancia = n × ε × máx|aᵢⱼ|
```

  - `ε` es el épsilon de la máquina en doble precisión (≈ 2.2e-16).
  - `máx|aᵢⱼ|` es el mayor valor absoluto de la matriz ampliada original.
  - Un pivote o término independiente con valor absoluto ≤ tolerancia se considera cero.
  - Se calcula en un único lugar (`tolerance`) y no se repite en otras funciones.
- **Vigilancia de valores no finitos:** tras cada operación se comprueba que el resultado sea finito; si no lo es, el proceso se detiene y devuelve `falloNumerico`.
- Limitación conocida: ninguna tolerancia es infalible; un sistema casi singular puede clasificarse de forma imprecisa. Por eso los pasos y la matriz final se muestran siempre.

### Prototipo del algoritmo (pseudocódigo)

```text
resolverSistema(sistema):
    validarSistema(sistema)
    tol = calcularTolerancia(sistema)
    escalonada, pasos, rango = eliminarAdelante(sistema, tol)
    caso = clasificarSistema(escalonada, rango, tol)
    si caso == solucionUnica:
        solucion = sustituirAtras(escalonada)
    devolver armarResultado(...)

eliminarAdelante(matrizAmpliada, tol):
    fila = 0
    para columna en 0..n-1:
        pivote = filaConMayorValorAbsoluto(columna, desdeFila = fila)
        si |matriz[pivote][columna]| <= tol: continuar   // columna sin pivote
        intercambiar(fila, pivote) y registrarPaso
        para cada filaInferior:
            restarMultiplo(filaInferior, fila) y registrarPaso
        fila = fila + 1
    devolver matriz, pasos, rango = fila
```

---

## 8. Arquitectura: núcleo y adaptadores

```text
CLI (menú, teclado, archivo, argumentos) ─┐
                                          ├→ model.Sistema → solver → model.Resultado ─┬→ cli/present
Web (formulario) ─────────────────────────┘                                            └→ web/present
                                                      (ambas presentaciones usan format)
```

### Regla de dependencia

```text
cmd/* → cli | web → solver → (validation, tolerance, elimination, classification, substitution) → model
                     numeric (hoja): lo usan el núcleo y los adaptadores
                     format → model: lo usan solo las presentaciones
```

- El núcleo (`model` a `solver`) **no importa** `cli`, `web` ni `format`.
- Las interfaces solo dependen hacia el núcleo, nunca al revés; `cli` y `web` no se importan entre sí.
- El núcleo no lee, no imprime, no conoce HTTP ni la consola.
- `numeric` y `format` son compartidos y no contienen lógica de resolución.
- Prohibido duplicar lógica de cálculo o validación de dominio en los adaptadores.

---

## 9. Estructura de carpetas

```text
gauss/
├── ARCHITECTURE.md
├── go.mod
├── cmd/
│   ├── cli/main.go
│   └── web/main.go
├── internal/
│   ├── numeric/        finite.go, parse.go
│   ├── model/          system.go, step.go, result.go
│   ├── validation/     validator.go
│   ├── tolerance/      tolerance.go
│   ├── elimination/    elimination.go
│   ├── classification/ classifier.go
│   ├── substitution/   substitution.go
│   ├── solver/         solver.go
│   ├── format/         number.go, text.go
│   ├── cli/            run.go, args.go, menu.go, prompt.go, interactive.go,
│   │                   file.go, parse.go, help.go, present.go
│   └── web/            server.go, handlers.go, formcheck.go, present.go
│       └── templates/  index.html
└── docs/               analisis.md, capacidades.md
```

Las pruebas (`*_test.go`) se ubican junto al código, un archivo de pruebas por paquete, según la convención de Go.

---

## 10. Responsabilidad por archivo

Cada archivo lleva una cabecera con estos mismos campos (ver sección 12).

### 10.1 Puntos de entrada

| Archivo | Responsabilidad | Recibe de | Envía a | Restricciones |
|---|---|---|---|---|
| `cmd/cli/main.go` | Arrancar la CLI | Sistema operativo (argumentos) | `internal/cli` | Sin lógica de cálculo, formato ni entrada de datos |
| `cmd/web/main.go` | Arrancar el servidor web | Sistema operativo | `internal/web` | Sin lógica de cálculo, formato ni HTML |

### 10.2 Núcleo

| Archivo | Responsabilidad | Recibe de | Envía a | Restricciones |
|---|---|---|---|---|
| `model/system.go` | Definir `System` y construirlo desde filas ampliadas | Adaptadores CLI y web | `solver` | Solo datos y su construcción; sin validación ni E/S |
| `model/step.go` | Definir `Step` (una columna de pivote) y `RowOperation` | — | `elimination`, presentaciones | Solo datos; sin texto de presentación |
| `model/result.go` | Definir `Result` y los cuatro casos | — | `solver`, presentaciones | Solo datos; sin lógica |
| `validation/validator.go` | Validar un `System` | `solver` | `solver` | No modifica datos; no imprime; no depende de adaptadores |
| `tolerance/tolerance.go` | Calcular la tolerancia relativa | `solver` | `elimination`, `classification` | Único lugar donde se calcula |
| `elimination/elimination.go` | Eliminación hacia adelante con pivoteo; un paso por columna | `solver`, `tolerance` | `classification`, `solver` | No clasifica ni resuelve; no modifica la entrada; se detiene ante valores no finitos |
| `classification/classifier.go` | Decidir el caso, el rango ampliado y la fila contradictoria | `elimination`, `tolerance` | `solver` | No calcula la solución; no modifica la matriz |
| `substitution/substitution.go` | Sustitución hacia atrás | `solver` | `solver` | Solo para solución única; informa fallo en vez de devolver valores no finitos |
| `solver/solver.go` | Orquestar el núcleo y armar el `Result` | Adaptadores CLI y web | Adaptadores CLI y web | Sin fórmulas propias; sin entrada/salida |

### 10.3 Compartido

| Archivo | Responsabilidad | Recibe de | Envía a | Restricciones |
|---|---|---|---|---|
| `numeric/finite.go` | Comprobar que un valor es finito | `validation`, `elimination`, `substitution` | Quien lo llama | Sin estado ni E/S |
| `numeric/parse.go` | Convertir texto en números (sintaxis decimal estricta) y separar filas | `cli`, `web/formcheck.go` | Los mismos adaptadores | Solo punto decimal; sin letras, hexadecimales, NaN ni infinito; no conoce CLI ni HTTP |
| `format/number.go` | Convertir números y matrices a texto con máximo 2 decimales | Presentaciones CLI y web | Presentaciones | No altera valores internos; nunca `-0.00`; notación científica desde 1e15 |
| `format/text.go` | Textos comunes: título de paso, operación, caso, rangos, contradicción, solución | Presentaciones CLI y web | Presentaciones | Solo presentación; sin cálculos ni E/S; misma redacción en ambas interfaces |

### 10.4 Adaptador CLI

| Archivo | Responsabilidad | Recibe de | Envía a | Restricciones |
|---|---|---|---|---|
| `cli/run.go` | Encadenar argumentos, origen de datos, solver y presentación; devolver el código de salida | `cmd/cli` | `solver`, `present` | Sin cálculos ni reglas de formato |
| `cli/args.go` | Interpretar `-f` y `-m` | `run` | `run` | No lee archivos ni imprime; `-f` y `-m` son excluyentes |
| `cli/menu.go` | Mostrar el menú de forma de ingreso y encaminar | `run` (sin datos por argumentos) | `interactive`, `file` | No lee el sistema por sí mismo |
| `cli/prompt.go` | Preguntar, leer una línea y repetir hasta obtener respuesta válida | `menu`, `interactive`, `file` | Los mismos | No conoce sistemas ni el solver |
| `cli/interactive.go` | Pedir n y las ecuaciones por teclado, con ayuda y reintento | `menu` | `run` | No calcula; no reserva memoria según n antes de recibir los datos |
| `cli/file.go` | Leer y parsear el sistema desde un archivo (pide la ruta si hace falta) | `menu` o `run` | `run` | No calcula; comprueba el número de líneas contra n |
| `cli/parse.go` | Pasar texto de la CLI (n, ecuaciones, `-m`) a `System` | `interactive`, `file`, `run` | Los mismos | Solo formato de texto; sin E/S |
| `cli/help.go` | Textos de ayuda y preguntas breves por paso | `menu`, `interactive`, `file`, `args` | Consola | Solo texto; sin lógica |
| `cli/present.go` | Mostrar matriz original, pasos, matriz escalonada y resultado | `run` (`Result`) | Consola | Sin cálculos; usa `format` |

### 10.5 Adaptador web

| Archivo | Responsabilidad | Recibe de | Envía a | Restricciones |
|---|---|---|---|---|
| `web/server.go` | Configurar rutas y arrancar el servidor HTTP | `cmd/web` | `handlers` | Sin lógica de dominio |
| `web/handlers.go` | Recibir el formulario, delegar y devolver la respuesta | Navegador | `formcheck`, `solver`, `present` | No calcula; limita el tamaño de la petición (1 MiB) |
| `web/formcheck.go` | Validar el texto del formulario en el servidor antes del núcleo | `handlers` | `handlers` | Un dato inválido nunca continúa hacia `solver`; devuelve mensaje claro |
| `web/present.go` | Construir el modelo de vista y renderizar la plantilla | `handlers` (`Result`) | Navegador | Sin cálculos; usa `format` |
| `web/templates/index.html` | Formulario, validación de página y vista del resultado | `present` | Navegador | Sin lógica de dominio; la validación de página es solo un mensaje inmediato |

### 10.6 Documentación

| Archivo | Responsabilidad |
|---|---|
| `docs/analisis.md` | Registro del análisis realizado |
| `docs/capacidades.md` | Capacidades actuales del programa y su uso |

---

## 11. Flujos

### 11.1 CLI

```text
cmd/cli/main.go
      ↓
¿argumentos con datos?
   ┌───┴───┐
  NO      SÍ
   │       │
   ▼       ▼
 menú    args → file (o valores directos)
   │       │
   ▼       │
interactive / file
   └───┬───┘
       ▼
  Sistema mostrado como matriz ampliada
       ▼
  solver.Resolver
       ▼
  present: matriz original, pasos, matriz escalonada, resultado
```

Menú y ayuda (breves):

```text
Forma de ingreso:
  1. Escribir los datos
  2. Leer desde archivo
Opción:

Número de ecuaciones (entero mayor que 0):
Ecuación 1 (n coeficientes y término independiente, separados por espacio):
```

Un dato inválido produce un mensaje corto con lo esperado y se vuelve a pedir.

### 11.2 Web

```text
Navegador → formulario (validación de página)
      ↓
handlers → formcheck (validación del servidor)
      ↓ (solo si es válido)
solver.Resolver
      ↓
present → plantilla de resultado
```

Si `formcheck` detecta un dato inválido, se responde con un mensaje de error y no se invoca al núcleo.

---

## 12. Convenciones de código

- Diseño **top-down**: el programa se divide en partes pequeñas; cada función hace una sola cosa. No hay límite fijo de líneas por función.
- Código lo más limpio posible, sin nada innecesario.
- Sin dependencias externas: solo la biblioteca estándar de Go. Cualquier dependencia nueva se consulta antes.
- Nombres siguiendo la sintaxis de Go: `PascalCase` para lo exportado y `camelCase` para lo interno. El camelCase sin más se usa solo en prototipos de algoritmos.
- **Cabecera obligatoria en cada archivo**, sin excepción:

```go
// File: elimination.go
// Purpose: forward elimination with partial pivoting, recording each step.
// Receives: model.System (augmented matrix) and the tolerance.
// Previous stage: solver (calls it after validation and tolerance).
// Next stage: classification (echelon matrix) and solver (steps).
// Restrictions: does not classify, does not solve, does not print;
//               stops and reports if a non-finite value appears.
```

---

## 13. Pruebas

Pruebas unitarias junto al código (`*_test.go`, un archivo por paquete). Casos mínimos:

- Solución única (incluido n = 1 y un caso que requiere intercambio de filas).
- Infinitas soluciones y sin solución (con fila contradictoria correcta).
- Sistemas casi singulares alrededor de la tolerancia.
- Valores no finitos durante el cálculo (`falloNumerico`).
- Validación: n ≤ 0, matriz no cuadrada, letras y caracteres extraños, NaN e infinito.
- Formato: dos decimales y ausencia de `-0.00`.
- Web: un dato inválido no llega al núcleo.

---

## 14. Entrega y documentación posterior

- El paquete completo del proyecto se entrega en **zip**.
- Después de generar el proyecto se documentan, en `docs/`, el análisis realizado y las capacidades actuales del programa.

---

## 15. Decisiones tomadas al aprobar el documento

Estos puntos no salieron del análisis inicial; se propusieron y quedaron aprobados con el documento.

1. **Formato del archivo de entrada:** texto plano; primera línea con `n`; luego `n` líneas con `n + 1` números separados por espacio (coeficientes y término independiente). Las filas también pueden separarse con `;`.
2. **Argumentos de ejecución:** `-f <archivo>` para leer de archivo y `-m "<filas separadas por ;>"` para pasar la matriz ampliada. El menú se omite cuando se usa cualquiera de los dos.
3. **Separador decimal:** punto (`2.5`). La coma se rechaza con un mensaje claro.
4. **Web:** el núcleo no limita n; el adaptador web limita el tamaño de la petición a 1 MiB para no agotar memoria. La entrada web es un cuadro de texto con una ecuación por línea.
5. **Fallo numérico controlado:** cuarto resultado posible cuando aparece un valor no finito.

## 16. Limitaciones conocidas

- Se guarda la matriz tras cada columna de pivote para mostrar los pasos: la memoria crece como n³. Con valores de n muy grandes el límite lo pone la máquina.
- Ninguna tolerancia es infalible: un sistema casi singular puede clasificarse de forma imprecisa.
