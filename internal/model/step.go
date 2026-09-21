// File: step.go
// Purpose: describe one stage (one pivot column) of the forward elimination.
// Receives: nothing (data definition only).
// Previous stage: elimination (fills the steps while working).
// Next stage: solver (passes them on) and the CLI/web presentations.
// Restrictions: only data; contains no presentation text and no logic.

package model

// RowOperation means: row Row = row Row - Factor * pivot row.
type RowOperation struct {
	Row    int
	Factor float64
}

// Step is the work done on one column. Rows and columns are 0-based.
// Matrix is the augmented matrix after the step (nil when the column had no pivot
// or when the calculation failed during the step).
type Step struct {
	Column      int
	PivotRow    int
	NoPivot     bool
	SwappedWith int // row swapped into PivotRow, or -1 when no swap was needed
	Pivot       float64
	Operations  []RowOperation
	Matrix      [][]float64
}
