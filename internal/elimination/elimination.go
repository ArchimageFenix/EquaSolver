// File: elimination.go
// Purpose: forward elimination with partial pivoting, recording one Step per column.
// Receives: a validated model.System and the tolerance.
// Previous stage: solver (after validation and tolerance).
// Next stage: solver, which hands the echelon matrix to classification.
// Restrictions: does not classify, does not solve, does not print; never mutates the
//               input system; stops and reports Failed if a non-finite value appears.

package elimination

import (
	"math"

	"gauss/internal/model"
	"gauss/internal/numeric"
)

// Outcome is what the forward elimination hands to the next stage.
type Outcome struct {
	Original [][]float64
	Echelon  [][]float64
	Steps    []model.Step
	Rank     int
	Failed   bool
}

type columnStatus int

const (
	columnPivoted columnStatus = iota
	columnSkipped
	columnFailed
)

// Run triangulates the augmented matrix of the system.
func Run(system model.System, tolerance float64) Outcome {
	matrix := augment(system)
	outcome := Outcome{Original: cloneMatrix(matrix)}
	row := 0
	for column := 0; column < system.N && row < system.N; column++ {
		step, status := processColumn(matrix, column, row, tolerance)
		outcome.Steps = append(outcome.Steps, step)
		if status == columnFailed {
			outcome.Failed = true
			break
		}
		if status == columnPivoted {
			row++
		}
	}
	outcome.Echelon = matrix
	outcome.Rank = row
	return outcome
}

func processColumn(matrix [][]float64, column, row int, tolerance float64) (model.Step, columnStatus) {
	pivotIndex := findPivot(matrix, column, row)
	pivot := matrix[pivotIndex][column]
	step := model.Step{Column: column, PivotRow: row, SwappedWith: -1, Pivot: pivot}
	if math.Abs(pivot) <= tolerance {
		step.NoPivot = true
		return step, columnSkipped
	}
	if pivotIndex != row {
		matrix[row], matrix[pivotIndex] = matrix[pivotIndex], matrix[row]
		step.SwappedWith = pivotIndex
	}
	operations, finite := clearBelow(matrix, column, row)
	step.Operations = operations
	if !finite {
		return step, columnFailed
	}
	step.Matrix = cloneMatrix(matrix)
	return step, columnPivoted
}

// findPivot returns the row (from fromRow down) with the largest magnitude in column.
func findPivot(matrix [][]float64, column, fromRow int) int {
	best := fromRow
	for i := fromRow + 1; i < len(matrix); i++ {
		if math.Abs(matrix[i][column]) > math.Abs(matrix[best][column]) {
			best = i
		}
	}
	return best
}

// clearBelow zeroes the column under the pivot row; it reports false on a non-finite value.
func clearBelow(matrix [][]float64, column, row int) ([]model.RowOperation, bool) {
	var operations []model.RowOperation
	for below := row + 1; below < len(matrix); below++ {
		if matrix[below][column] == 0 {
			continue
		}
		factor := matrix[below][column] / matrix[row][column]
		operations = append(operations, model.RowOperation{Row: below, Factor: factor})
		if !numeric.IsFinite(factor) || !subtractMultiple(matrix, below, row, factor, column) {
			return operations, false
		}
	}
	return operations, true
}

// subtractMultiple does target = target - factor*source from fromColumn on and
// leaves an exact zero in fromColumn.
func subtractMultiple(matrix [][]float64, target, source int, factor float64, fromColumn int) bool {
	for j := fromColumn; j < len(matrix[target]); j++ {
		matrix[target][j] -= factor * matrix[source][j]
		if !numeric.IsFinite(matrix[target][j]) {
			return false
		}
	}
	matrix[target][fromColumn] = 0
	return true
}

func augment(system model.System) [][]float64 {
	matrix := make([][]float64, system.N)
	for i := range matrix {
		matrix[i] = make([]float64, system.N+1)
		copy(matrix[i], system.A[i])
		matrix[i][system.N] = system.B[i]
	}
	return matrix
}

func cloneMatrix(matrix [][]float64) [][]float64 {
	clone := make([][]float64, len(matrix))
	for i, row := range matrix {
		clone[i] = append([]float64(nil), row...)
	}
	return clone
}
