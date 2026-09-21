// File: help.go
// Purpose: hold the short help texts and questions shown at each data-entry step.
// Receives: nothing (texts and a small formatting helper).
// Previous stage: menu.go, interactive.go, file.go and args.go (they ask for the texts).
// Next stage: the console, through prompt.go.
// Restrictions: text only; no logic, no reading and no calculations. Keep every
//               message short: the CLI must not be verbose.

package cli

import "fmt"

const usageText = "Uso: gauss [-f archivo | -m \"fila; fila; ...\"]\n" +
	"  -f  archivo de texto: primera línea n; luego n líneas con n+1 números\n" +
	"  -m  matriz ampliada, filas separadas por ';' (ejemplo: -m \"2 1 5; 1 3 10\")\n" +
	"  sin argumentos: se muestra el menú de ingreso"

const (
	modeQuestion  = "Opción (1 o 2):"
	countQuestion = "Número de ecuaciones (entero mayor que 0):"
	pathQuestion  = "Ruta del archivo (1.ª línea: n; luego n líneas de n+1 números):"
)

var menuLines = []string{
	"Forma de ingreso:",
	"  1. Escribir los datos",
	"  2. Leer desde archivo",
}

func rowQuestion(index, n int) string {
	return fmt.Sprintf("Ecuación %d de %d (%d coeficientes y el término independiente, separados por espacio):", index, n, n)
}
