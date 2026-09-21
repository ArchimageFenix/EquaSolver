// File: number.go
// Purpose: convert numbers and matrices to text with at most two decimals.
// Receives: float64 values and matrices from the presentations.
// Previous stage: cli/present.go and web/present.go (and format/text.go).
// Next stage: the same presentations, which receive plain strings.
// Restrictions: presentation only; never alters the values used by the core and
//               knows nothing about CLI or web. Never prints "-0.00".

package format

import (
	"math"
	"strconv"
)

// scientificThreshold is the magnitude from which values are shown in scientific
// notation, only to keep extreme values readable.
const scientificThreshold = 1e15

// Number formats a value with two decimals (scientific notation for extreme values).
func Number(value float64) string {
	if math.Abs(value) >= scientificThreshold {
		return strconv.FormatFloat(value, 'e', 2, 64)
	}
	text := strconv.FormatFloat(value, 'f', 2, 64)
	if text == "-0.00" {
		return "0.00"
	}
	return text
}

// Matrix formats every cell of a matrix.
func Matrix(matrix [][]float64) [][]string {
	cells := make([][]string, len(matrix))
	for i, row := range matrix {
		cells[i] = make([]string, len(row))
		for j, value := range row {
			cells[i][j] = Number(value)
		}
	}
	return cells
}
