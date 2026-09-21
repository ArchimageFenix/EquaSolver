// File: present.go
// Purpose: print the original matrix, the steps, the echelon matrix and the verdict.
// Receives: an output writer and a model.Result.
// Previous stage: run.go (after solver.Solve).
// Next stage: the console.
// Restrictions: no calculations; every number goes through package format; compact
//               output, without explanations beyond the data.

package cli

import (
	"fmt"
	"io"
	"strings"

	"gauss/internal/format"
	"gauss/internal/model"
)

func printResult(out io.Writer, result model.Result) {
	printMatrix(out, "Matriz original:", result.Original)
	printSteps(out, result.Steps)
	if result.Echelon != nil {
		printMatrix(out, "Matriz escalonada:", result.Echelon)
	}
	printVerdict(out, result)
}

func printMatrix(out io.Writer, title string, matrix [][]float64) {
	fmt.Fprintln(out, title)
	printRows(out, matrix)
	fmt.Fprintln(out)
}

func printSteps(out io.Writer, steps []model.Step) {
	for i, step := range steps {
		fmt.Fprintln(out, format.StepTitle(i+1, step))
		for _, operation := range step.Operations {
			fmt.Fprintln(out, "  "+format.OperationText(operation, step.PivotRow))
		}
		if step.Matrix != nil {
			printRows(out, step.Matrix)
		}
		fmt.Fprintln(out)
	}
}

func printVerdict(out io.Writer, result model.Result) {
	fmt.Fprintln(out, "Resultado: "+format.CaseTitle(result.Case))
	for _, line := range format.VerdictLines(result) {
		fmt.Fprintln(out, line)
	}
	for _, line := range format.SolutionLines(result.Solution) {
		fmt.Fprintln(out, "  "+line)
	}
}

func printRows(out io.Writer, matrix [][]float64) {
	cells := format.Matrix(matrix)
	width := widestCell(cells)
	for _, row := range cells {
		fmt.Fprintln(out, "  "+formatRow(row, width))
	}
}

func formatRow(row []string, width int) string {
	last := len(row) - 1
	parts := make([]string, 0, len(row))
	for _, cell := range row[:last] {
		parts = append(parts, fmt.Sprintf("%*s", width, cell))
	}
	return "[ " + strings.Join(parts, "  ") + " | " + fmt.Sprintf("%*s", width, row[last]) + " ]"
}

func widestCell(cells [][]string) int {
	width := 0
	for _, row := range cells {
		for _, cell := range row {
			if len(cell) > width {
				width = len(cell)
			}
		}
	}
	return width
}
