// File: text.go
// Purpose: build the texts shared by CLI and web (step titles, row operations,
//          case names, rank summary, contradiction, solution lines).
// Receives: model.Step, model.RowOperation, model.Case and model.Result values.
// Previous stage: cli/present.go and web/present.go.
// Next stage: the same presentations, which receive plain strings.
// Restrictions: presentation only; no calculations; no I/O; the same wording for
//               both interfaces. Rows and columns are shown 1-based.

package format

import (
	"fmt"
	"strconv"
	"strings"

	"gauss/internal/model"
)

// NumericFailureMessage explains a numeric failure to the user.
const NumericFailureMessage = "Se encontró un valor no finito durante el cálculo: los datos son demasiado extremos para procesarlos."

// StepTitle describes what happened in a step; number is the 1-based position.
func StepTitle(number int, step model.Step) string {
	head := fmt.Sprintf("Paso %d (columna %d): ", number, step.Column+1)
	switch {
	case step.NoPivot:
		return head + "sin pivote, se omite la columna"
	case step.SwappedWith >= 0:
		return head + fmt.Sprintf("intercambiar fila %d con fila %d (pivote %s)", step.PivotRow+1, step.SwappedWith+1, Number(step.Pivot))
	default:
		return head + fmt.Sprintf("pivote %s en fila %d", Number(step.Pivot), step.PivotRow+1)
	}
}

// OperationText describes one row operation against the given pivot row (0-based).
func OperationText(operation model.RowOperation, pivotRow int) string {
	sign, factor := "-", operation.Factor
	if factor < 0 {
		sign, factor = "+", -factor
	}
	return fmt.Sprintf("fila %d = fila %d %s %s × fila %d", operation.Row+1, operation.Row+1, sign, Number(factor), pivotRow+1)
}

// CaseTitle returns the name of a case in capital letters.
func CaseTitle(c model.Case) string {
	switch c {
	case model.UniqueSolution:
		return "SOLUCIÓN ÚNICA"
	case model.InfiniteSolutions:
		return "INFINITAS SOLUCIONES"
	case model.NoSolution:
		return "SIN SOLUCIÓN"
	default:
		return "FALLO NUMÉRICO"
	}
}

// VerdictLines returns the detail lines that follow the case title.
func VerdictLines(result model.Result) []string {
	switch result.Case {
	case model.NumericFailure:
		return []string{NumericFailureMessage}
	case model.NoSolution:
		return []string{rankSummary(result), contradictionText(result)}
	default:
		return []string{rankSummary(result)}
	}
}

// SolutionLines returns one "xi = value" line per unknown.
func SolutionLines(solution []float64) []string {
	lines := make([]string, len(solution))
	for i, value := range solution {
		lines[i] = variableName(i) + " = " + Number(value)
	}
	return lines
}

func variableName(index int) string {
	return "x" + strconv.Itoa(index+1)
}

func rankSummary(result model.Result) string {
	summary := fmt.Sprintf("Rango coeficientes: %d | Rango ampliada: %d", result.RankCoefficients, result.RankAugmented)
	if result.Case == model.NoSolution {
		return summary
	}
	return summary + fmt.Sprintf(" | Variables libres: %d", result.FreeVariables)
}

func contradictionText(result model.Result) string {
	row := result.Echelon[result.ContradictoryRow]
	n := len(row) - 1
	coefficients := Matrix([][]float64{row[:n]})[0]
	return fmt.Sprintf("Fila contradictoria: %d  (%s | %s)", result.ContradictoryRow+1, strings.Join(coefficients, " "), Number(row[n]))
}
