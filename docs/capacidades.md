# Capacidades actuales

## Qué hace

Resuelve sistemas de n ecuaciones lineales con n incógnitas (n ≥ 1, sin límite propio) por el método de Gauss con pivoteo parcial, y muestra el proceso completo.

Resultados posibles:

| Resultado | Qué se muestra |
|---|---|
| Solución única | Rangos, variables libres (0) y el valor de cada incógnita |
| Infinitas soluciones | Rangos, variables libres y la matriz escalonada |
| Sin solución | Rangos, la fila contradictoria y la matriz escalonada |
| Fallo numérico | Aviso de que apareció un valor no finito; sin caída del programa |

Siempre se muestran la matriz original y cada paso de la eliminación. Los números se presentan con máximo dos decimales (cálculo interno en doble precisión).

## Requisitos

Go 1.21 o superior. Sin dependencias externas.

## Uso de la CLI

```text
go run ./cmd/cli                       # menú de ingreso
go run ./cmd/cli -f sistema.txt        # datos desde archivo
go run ./cmd/cli -m "2 1 5; 1 3 10"    # matriz ampliada por argumento
```

Formato del archivo: primera línea con n; luego n líneas con n + 1 números separados por espacio (coeficientes y término independiente). Punto decimal, sin comas.

```text
3
1 1 1 6
2 -1 1 3
1 2 -1 2
```

Sin argumentos se muestra un menú: escribir los datos o leer desde archivo. Cada pregunta lleva una ayuda breve y, ante un dato inválido, se explica qué se esperaba y se vuelve a preguntar.

Códigos de salida: 0 correcto, 1 error de datos, 2 uso incorrecto de argumentos.

## Uso de la web

```text
go run ./cmd/web                # http://localhost:8080
go run ./cmd/web -addr :9000    # otro puerto
```

Se escribe la matriz ampliada, una ecuación por línea. La página avisa de inmediato si hay letras o caracteres extraños; el servidor comprueba de nuevo y un dato inválido nunca llega al núcleo. Las peticiones se limitan a 1 MiB.

## Compilar y probar

```text
go build ./...
go vet ./...
go test ./...
```

## Estado de verificación

El proyecto se generó en un entorno sin Go instalado, por lo que **no se compiló ni se ejecutaron las pruebas allí**. La lógica numérica se comprobó con un port equivalente en Python sobre los casos de ejemplo (solución única, infinitas soluciones, sin solución, n = 1 y matriz de ceros). Antes de usarlo conviene ejecutar `go build ./...` y `go test ./...`.

## Límites conocidos

- La memoria de los pasos crece como n³; con n muy grande el límite lo pone la máquina.
- Ninguna tolerancia es infalible: un sistema casi singular puede clasificarse de forma imprecisa.
- Solo sistemas cuadrados.
